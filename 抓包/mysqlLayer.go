package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

//	注册自定义的层
var layer = gopacket.LayerTypeMetadata{"MysqlLayer", gopacket.DecodeFunc(decodeMyLayer)}
var MyLayerType = gopacket.RegisterLayerType(12345, layer)

//	定义层的内容
type MyLayer struct {
	layers.BaseLayer
}

//	定义解析函数
func decodeMyLayer(data []byte, p gopacket.PacketBuilder) error {
	p.AddLayer(&MyLayer{layers.BaseLayer{data, data[5:]}})
	return p.NextDecoder(layers.LayerTypeEthernet)
}

//	定义一些解析包所需的接口
func (m MyLayer) LayerType() gopacket.LayerType {
	return MyLayerType
}
func (m MyLayer) LayerContents() []byte {
	return m.Contents
}
func (m MyLayer) LayerPayload() []byte {
	return m.Payload
}

func main() {
	fmt.Println("packet start...")
	//deviceName := "\\Device\\NPF_{80D0CD86-5229-4293-A7AD-3DC078D845CC}"
	deviceName := "\\Device\\NPF_{3D211ED3-70EA-4012-8C25-581283A72B14}"
	snapLen := int32(65535)
	port := uint16(3306)
	filter := getFilters(port)
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

		fmt.Printf("packet:%v\n", packet)

		// tcp 层
		tcp := packet.TransportLayer().(*layers.TCP)
		//fmt.Printf("tcp:%v\n", tcp)
		if len(tcp.Payload) == 0 {
			continue
		}
		mysqlPacket := gopacket.NewPacket(tcp.Payload, MyLayerType, gopacket.Default)

		log.Printf("mysql layer:%v\n", string(mysqlPacket.Layer(MyLayerType).LayerPayload()))
	}
}

func getFilters(port uint16) string {
	filter := fmt.Sprintf("tcp and dst port %v", port)
	return filter
}
