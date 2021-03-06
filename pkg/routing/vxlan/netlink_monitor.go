package vxlan

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
	"net"
	"syscall"
	"time"
)

type NetlinkMonitor struct {
	socket    *nl.NetlinkSocket
	linkIndex int
}

func NewNetlinkMonitor(linkIndex int) (*NetlinkMonitor, error) {
	m := &NetlinkMonitor{
		linkIndex: linkIndex,
	}
	return m, nil
}

func (m *NetlinkMonitor) Start() error {
	socket, err := nl.Subscribe(syscall.NETLINK_ROUTE, syscall.RTNLGRP_NEIGH)
	if err != nil {
		return fmt.Errorf("Failed to subscribe to (NETLINK_ROUTE,RTNLGRP_NEIGH): %v", err)
	}

	m.socket = socket

	go m.watch()
	return nil
}

func (m *NetlinkMonitor) watch() {
	for {
		messages, err := m.socket.Receive()
		if err != nil {
			glog.Errorf("error reading from netlink monitor: %v ", err)
			time.Sleep(1 * time.Second)
			continue
		}

		for _, message := range messages {
			neigh, err := netlink.NeighDeserialize(message.Data)
			if err != nil {
				glog.Warningf("error deserializing netlink monitor message: %v", err)
				continue
			}

			if neigh.LinkIndex != m.linkIndex {
				// Ignore: not our device
				continue
			}

			if message.Header.Type == syscall.RTM_GETNEIGH {
			} else if message.Header.Type == syscall.RTM_NEWNEIGH {
			} else {
				// Ignore: We only care about the neighbour table, and don't care about deleted entries
				continue
			}

			if (neigh.State & netlink.NUD_REACHABLE) != 0 {
				// ignore: this does not affect ARP
				continue
			}

			if neigh.HardwareAddr != nil || neigh.IP == nil {
				// ignore: mac is already present, or ip is present
				continue
			}

			//glog.Infof("Got netlink monitor message: %v", neigh)
			//glog.Infof("\tLinkIndex: %v", neigh.LinkIndex)
			//glog.Infof("\tIP: %v", neigh.IP)
			//glog.Infof("\tHardwareAddr: %v", neigh.HardwareAddr)
			//
			//glog.Infof("\tState: %v", neigh.State)
			//if (neigh.State & netlink.NUD_INCOMPLETE) != 0 {
			//	glog.Infof("\t\t NUD_INCOMPLETE")
			//}
			//if (neigh.State & netlink.NUD_REACHABLE) != 0 {
			//	glog.Infof("\t\t NUD_REACHABLE")
			//}
			//if (neigh.State & netlink.NUD_STALE) != 0 {
			//	glog.Infof("\t\t NUD_STALE")
			//}
			//if (neigh.State & netlink.NUD_DELAY) != 0 {
			//	glog.Infof("\t\t NUD_DELAY")
			//}
			//if (neigh.State & netlink.NUD_PROBE) != 0 {
			//	glog.Infof("\t\t NUD_PROBE")
			//}
			//if (neigh.State & netlink.NUD_FAILED) != 0 {
			//	glog.Infof("\t\t NUD_FAILED")
			//}
			//if (neigh.State & netlink.NUD_NOARP) != 0 {
			//	glog.Infof("\t\t NUD_NOARP")
			//}
			//if (neigh.State & netlink.NUD_PERMANENT) != 0 {
			//	glog.Infof("\t\t NUD_PERMANENT")
			//}

			ip4 := neigh.IP.To4()
			if ip4 == nil {
				glog.Warningf("ignoring ipv6 request: %s", neigh.IP)
				continue
			}

			cidrIP := net.IPv4(ip4[0], ip4[1], ip4[2], 0)
			mac := mapToMAC(cidrIP)
			glog.Warningf("HACK mapping assuming /24 PodCIDR, 96 as second: %s -> %s", ip4, mac)

			// We inject directly for speed (i.e. we don't do a full sync)
			glog.V(2).Infof("NETLINK: ip neigh replace %s lladdr %s dev vxlanX", ip4, mac)
			err = netlink.NeighSet(&netlink.Neigh{
				LinkIndex:    m.linkIndex,
				State:        netlink.NUD_REACHABLE,
				Type:         syscall.RTN_UNICAST,
				IP:           ip4,
				HardwareAddr: mac,
			})
			if err != nil {
				glog.Warningf("error doing `ip neigh replace %s lladdr %s dev vxlanX`: %v", ip4, mac, err)
			}
		}
	}
}
