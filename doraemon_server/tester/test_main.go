package main

import (
	"bytes"
	"encoding/binary"
	//"time"
	"fmt"
	//"math"
	//"reflect"
)

type S struct {
	F string `species:"gopher" color:"blue"`
}
type BodyHead struct {
	Id  uint32
	Len uint16
}

type Vertex struct {
	X int
	Y int
}

func main() {

	//var s S
	//st := reflect.TypeOf(S{})
	//field := st.Field(0)
	//fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

	//head := BodyHead{123, 55}

	buf := new(bytes.Buffer)
	//var pi float64 = math.Pi

	err := binary.Write(buf, binary.LittleEndian, BodyHead{1, 11})
	err = binary.Write(buf, binary.LittleEndian, BodyHead{2, 22})
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	fmt.Printf("% x\n", buf.Bytes())

	readHead := new(BodyHead)
	//readHead2 := new(BodyHead)
	reader := bytes.NewReader(buf.Bytes()[0:11])
	//reader := bytes.NewReader(buf.Bytes()[0:12])
	for {
		err = binary.Read(reader, binary.LittleEndian, readHead)

		if err != nil {
			fmt.Println("binary.Read failed:", err)
			break
		}
		fmt.Println(readHead)
	}
	u := new(uint32)
	reader.Seek(-5, 1)
	err = binary.Read(reader, binary.LittleEndian, u)
	if err != nil {
		fmt.Println("binary.Read uint  failed:", err)

	}

	fmt.Println(*u)
}
