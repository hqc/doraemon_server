package packager

import (
	//"fmt"
	//	"bufio"
	"bytes"
	"errors"
	//"github.com/lite_tool"
	"encoding/binary"
	"io"
	"reflect"
)

const (
	PackHead     uint8  = 0xff
	PackEnd      uint8  = 0xfe
	PBodyLenType uint16 = 0
)

type TypeMapValue struct {
	id    uint16
	type_ reflect.Type
}

type Packager struct {
	//TypeMap     map[uint16]reflect.Type
	typeMap   map[string]TypeMapValue
	writeBuff WriteBuffer
}

type WriteBuffer struct {
	data   []byte
	writer io.Writer
	seeker io.Seeker
}

func (sb *WriteBuffer) TryInit() {
	if sb.data != nil {
		return
	}
	sb.data = make([]byte, 256)
	sb.writer = bytes.NewBuffer(sb.data)
	sb.seeker = bytes.NewReader(sb.data)

	//See bytes.NewReader(Data)

}

//注册包
func (p *Packager) Reg(id uint16, obj interface{}) error {
	objType := reflect.TypeOf(obj)

	objTypeName := objType.Name()
	if _, exist := p.typeMap[objTypeName]; exist {
		return errors.New("reg package name duplicate")
	}

	p.typeMap[objTypeName] = TypeMapValue{id, objType}

	return nil
}

func (p *Packager) Write(writer io.Writer, obj interface{}) error {
	//寻找是否在typemap内
	objType := reflect.TypeOf(obj)
	var objData TypeMapValue
	var exist bool

	if objData, exist = p.typeMap[objType.Name()]; !exist {
		return errors.New("write package not register")
	}

	p.writeBuff.TryInit()
	p.writeBuff.seeker.Seek(0, 0)

	binary.Write(p.writeBuff.writer, binary.LittleEndian, PackHead) //写包头

	p.writeBuff.seeker.Seek(reflect.TypeOf(PBodyLenType).Len(), 1) //跳过 写包体长度 的位置

	binary.Write(p.writeBuff.writer, binary.LittleEndian, objData.id) //写入包id

	return nil
}
