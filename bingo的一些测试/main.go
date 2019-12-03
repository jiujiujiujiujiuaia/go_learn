package main

import (
	"fmt"
)

func main() {
	//fmt.Println(^uint32(1))
	//fmt.Println(^uint32(0))
	//
	//a := byte(129)      //11111111 这是byte的极限， 因为 a := byte(256)//越界报错， 0~255正好256个数，不能再高了
	//b := uint8(255)     //11111111 这是uint8的极限，因为 c := uint8(256)//越界报错，0~255正好256个数，不能再高了
	//c := int8(127)      //01111111 这是int8的极限， 因为 b := int8(128)//越界报错， 0~127正好128个数，所以int8的极限只是256的一半
	//d := int8(a)        //11111111 打印出来则是-0000001，int8(128)、int8(255)、int8(byte(255))都报错越界，因为int极限是127，但是却可以写：int8(a)，第一位拿来当符号了
	//e := int16(127 + 3) //01111111 打印出来还是01111111
	//
	//fmt.Printf("%08b %d \n", a, a)
	//fmt.Printf("%08b %d \n", b, b)
	//fmt.Printf("%08b %d \n", c, c)
	//fmt.Printf("%08b %d \n", d, d)
	//fmt.Printf("%08b %d \n", e, e)
	b := int8(127)
	fmt.Println(b * 3)
	fmt.Printf("%08b %d \n", 381, 381)
	//overflow()
}

func overflow() {
	var a uint8 = 127
	for ; a < uint8(255); a++ {
		fmt.Printf("%08b %d \n", a, a)
		fmt.Printf("%08b %d \n", int8(a), int8(a))
		fmt.Println("-----------------")
	}
}
