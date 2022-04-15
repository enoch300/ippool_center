package app

import (
	"fmt"
	"ippool_center/app/udp"
	"ippool_center/cleartool"
	"ippool_center/db/redis"
	"ippool_center/utils/log"
)

func Run() {
	if err := redis.Connect(); err != nil {
		fmt.Printf("Redis connect %v\n", err)
		log.GlobalLog.Errorf("Redis connect %v", err)
		return
	}
	go cleartool.Run()
	udp.Listen()
}
