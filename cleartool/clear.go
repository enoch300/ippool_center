package cleartool

import (
	"context"
	"ippool_center/db/redis"
	. "ippool_center/utils/log"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var ctx = context.Background()

func clearExpire(key string) error {
	values, err := redis.RDB.HGetAll(ctx, key).Result()
	if err != nil {
		return err
	}

	for k, v := range values {
		t1 := time.Now().Unix()
		lt := strings.Split(v, "_")[2]
		t2, _ := strconv.ParseInt(lt, 10, 64)
		if t1-t2 > 50 {
			redis.RDB.HDel(ctx, key, k)
			GlobalLog.Infof("[cleartool] >>> clear expire hkey: %s, key: %s", key, k)
		}
	}
	return nil
}

func Run() {
	defer func() {
		if err := recover(); err != nil {
			GlobalLog.Errorf("[cleartool]: %v, stack: %v", err, string(debug.Stack()))
		}
	}()

	for {
		keys, err := redis.RDB.Keys(ctx, "*_*_*_*").Result()
		if err != nil {
			GlobalLog.Errorf("[cleartool] >>> %s", err)
			time.Sleep(5 * time.Second)
		}

		for _, key := range keys {
			if err := clearExpire(key); err != nil {
				GlobalLog.Errorf("[cleartool] >>> hdel expire key %s %s", key, err)
			}
		}
		time.Sleep(time.Minute)
	}
}
