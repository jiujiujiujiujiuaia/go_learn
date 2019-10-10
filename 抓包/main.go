package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	"github.com/google/gopacket/pcap"
)

func main() {
	getAllService()
}

func zhuabao() {
	fmt.Println("packet start...")
	deviceName := "\\Device\\NPF_{80D0CD86-5229-4293-A7AD-3DC078D845CC}"
	//deviceName := "\\Device\\NPF_{3D211ED3-70EA-4012-8C25-581283A72B14}"
	snapLen := int32(65535)
	port := uint16(443)
	filter := getFilter(port)
	fmt.Printf("device:%v, snapLen:%v, port:%v\n", deviceName, snapLen, port)
	fmt.Println("filter:", filter)

	//打开网络接口，抓取在线数据
	handle, err := pcap.OpenLive(deviceName, snapLen, true, pcap.BlockForever)
	if err != nil {
		fmt.Printf("pcap open live failed: %v", err)
		return
	}

	// 设置过滤器
	if err := handle.SetBPFFilter(filter); err != nil {
		fmt.Printf("set bpf filter failed: %v", err)
		return
	}
	defer handle.Close()

	// 抓包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetSource.NoCopy = true
	for packet := range packetSource.Packets() {
		if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
			fmt.Println("unexpected packet")
			continue
		}
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Printf("packet:%v\n\n", packet)

		// tcp 层
		tcp := packet.TransportLayer().(*layers.TCP)

		//fmt.Printf("tcp byte:%v\n\n", tcp)
		// tcp payload，也即是tcp传输的数据
		if len(tcp.Payload) != 0 {
			continue
		}
		fmt.Printf("tcp payload:%v\n\n", string(tcp.Payload))

	}
}

//定义过滤器
func getFilter(port uint16) string {
	//filter := fmt.Sprintf("tcp and ((dst port %v)or (dst port 80)) and host 129.204.8.88", port)
	filter := "tcp and host 129.204.8.88"
	return filter
}

func getAllService() {
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Print device information
	fmt.Println("Devices found:")
	for _, d := range devices {
		fmt.Println("\nName: ", d.Name)
		fmt.Println("Description: ", d.Description)
		fmt.Println("Devices addresses: ", d.Description)

		for _, address := range d.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}
