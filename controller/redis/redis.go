package redis

import (
	"context"
	"ippool_center/db/redis"
	"ippool_center/peer"
)

var ctx = context.Background()

func Store(p peer.Peer) (err error) {
	err = redis.RDB.HSet(ctx, p.Format2NetAppIdProvinceIsp(), p.Format2MidInIpInPort(), p.Format2OutIpOutPort()).Err()
	if err != nil {
		return
	}

	err = redis.RDB.HSet(ctx, p.Format2Mid(), p.Format2AppIdInIpInPort(), p.Format2OutIpOutPort()).Err()
	if err != nil {
		return
	}

	return
}
