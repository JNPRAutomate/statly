package main

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/Juniper/go-netconf/netconf"
)

type IntInfo struct {
	PhyInts []PhyInt `xml:"physical-interface"`
}

type PhyInt struct {
	Name         string      `xml:"name"`
	TrafficStats TrafficStat `xml:"traffic-statistics"`
	LogicalInt   LogicalInt  `xml:"logical-interface"`
}

type LogicalInt struct {
	Name         string      `xml:"name"`
	TrafficStats TrafficStat `xml:"traffic-statistics"`
}

type TrafficStat struct {
	InPackets  string `xml:"input-packets"`
	OutPackets string `xml:"output-packets"`
	InBytes    string `xml:"input-bytes"`
	OutBytes   string `xml:"output-bytes"`
}

func main() {
	username := "root"
	password := "Juniper"
	host := "10.0.1.150"

	s, err := netconf.DialSSH(host, netconf.SSHConfigPassword(username, password))
	if err != nil {
		panic(err)
	}

	defer s.Close()

	fmt.Printf("Session Id: %d\n\n", s.SessionID)

	reply, err := s.Exec(netconf.RawMethod("<get-interface-information><extensive/></get-interface-information>"))
	if err != nil {
		panic(err)
	}
	var ii IntInfo
	xml.Unmarshal([]byte(reply.Data), &ii)
	fmt.Printf("%#v\n", ii)
	for item := range ii.PhyInts {
		fmt.Println(strings.Trim(ii.PhyInts[item].Name, "\n"), strings.Trim(ii.PhyInts[item].TrafficStats.InBytes, "\n"), strings.Trim(ii.PhyInts[item].TrafficStats.OutBytes, "\n"))
	}
}
