package redis

import (
	"context"
	"fmt"
	"ippool_center/db/redis"
	"ippool_center/peer"
	"time"
)

var ctx = context.Background()

func Store(p peer.Peer) (err error) {
	now := time.Now().Unix()
	err = redis.RDB.HSet(ctx, p.Format2NetAppidProvinceIsp(), p.Format2MidInIpInPort(), fmt.Sprintf("%s_%d", p.Format2OutIpOutPort(), now)).Err()
	if err != nil {
		return
	}

	err = redis.RDB.HSet(ctx, p.Format2Mid(), p.Format2AppidInIpInPort(), fmt.Sprintf("%s_%d", p.Format2OutIpOutPort(), now)).Err()
	if err != nil {
		return
	}

	err = redis.RDB.HSet(ctx, p.Format2NetAppidIsp(), p.Format2MidInIpInPort(), fmt.Sprintf("%s_%d", p.Format2OutIpOutPort(), now)).Err()
	if err != nil {
		return
	}

	isExist, err := redis.RDB.HExists(ctx, "machine_ipip", p.Format2Mid()).Result()
	if isExist {
		return err
	}

	err = redis.RDB.HSet(ctx, "machine_ipip", p.Format2Mid(), fmt.Sprintf("%s_%d", p.Format2ProvinceIsp(), now)).Err()
	if err != nil {
		return
	}

	return
}
