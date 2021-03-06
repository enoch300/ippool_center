package peer

import "fmt"

type Peer struct {
	T         int64
	Network   string
	MachineId string
	AppId     string
	Province  string
	Isp       string
	InnerIp   string
	InnerPort string
	OuterIp   string
	OuterPort int
}

func (p *Peer) Format2NetAppidProvinceIsp() string {
	return fmt.Sprintf("%s_%s_%s_%s", p.Network, p.AppId, p.Province, p.Isp)
}

func (p *Peer) Format2Mid() string {
	return p.MachineId
}

func (p *Peer) Format2NetAppidIsp() string {
	return fmt.Sprintf("%s_%s_%s", p.Network, p.AppId, p.Isp)
}

func (p *Peer) Format2MidInIpInPort() string {
	return fmt.Sprintf("%s_%s_%s", p.MachineId, p.InnerIp, p.InnerPort)
}

func (p *Peer) Format2ProvinceIsp() string {
	return fmt.Sprintf("%s_%s", p.Province, p.Isp)
}

func (p *Peer) Format2AppidInIpInPort() string {
	return fmt.Sprintf("%s_%s_%s", p.AppId, p.InnerIp, p.InnerPort)
}

func (p *Peer) Format2OutIpOutPort() string {
	return fmt.Sprintf("%s_%d", p.OuterIp, p.OuterPort)
}
