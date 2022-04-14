package udp

import (
	"fmt"
	"ippool_center/controller/redis"
	"ippool_center/peer"
	. "ippool_center/utils/log"
	"net"
	"strings"
	"time"
)

const PORT = 9981

func handleConn(conn *net.UDPAddr, msg []byte) {
	strArray := strings.Split(string(msg), "|") //"network|machine_id|app_id|inner_ip|inner_port|province|ipip"
	p := peer.Peer{
		T:         time.Now().Unix(),
		Network:   strArray[0],
		MachineId: strArray[1],
		AppId:     strArray[2],
		InnerIp:   strArray[3],
		InnerPort: strArray[4],
		Province:  strArray[5],
		Isp:       strArray[6],
		OuterIp:   conn.IP.String(),
		OuterPort: conn.Port,
	}

	err := redis.Store(p)
	if err != nil {
		GlobalLog.Errorf("Redis [INSERRT] hash: %s, key: %s, value: %s, %v",
			p.Format2NetAppIdProvinceIsp(),
			p.Format2MidInIpInPort(),
			p.Format2OutIpOutPort(),
			err)
		return
	}

	GlobalLog.Infof("Redis [INSERRT] hash: %s, key: %s, value: %s",
		p.Format2NetAppIdProvinceIsp(),
		p.Format2MidInIpInPort(),
		p.Format2OutIpOutPort(),
	)
}

func Listen() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: PORT})
	if err != nil {
		fmt.Printf("Listen UDP: %v\n", err)
		GlobalLog.Errorf("Listen UDP: %v", err)
		return
	}

	GlobalLog.Infof("Listen UDP: %s", listener.LocalAddr().String())
	for {
		data := make([]byte, 1024)
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			GlobalLog.Errorf("Read UDP: %s", err.Error())
			continue
		}
		GlobalLog.Infof("Recv from: %v, msg: %v", remoteAddr.String(), string(data[:n]))
		go handleConn(remoteAddr, data[:n])
	}
}
