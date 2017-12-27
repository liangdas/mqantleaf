// Copyright 2014 mqantleaf Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// 与leaf数据包解析的相关函数就放这里了
package gatewsleaf

import (
	"encoding/binary"
	"io"
	"errors"
)

/**
给客户端发送消息
 */
func (this *CustomAgent) WriteMarshal(_id uint16, body []byte) error{
	this.send_num++
	data,err:=this.Marshal(_id,body)
	if err!=nil{
		return err
	}
	//粘包完成后调下面的语句发送数据
	return this.Write(data...)
}

func (this *CustomAgent) Read() ([]byte, error) {
	msgData := make([]byte, 2048)
	if count, err := io.ReadAtLeast(this.r, msgData,2); err != nil {
		return nil, err
	}else{
		return msgData[:count], nil
	}
}
func (this *CustomAgent) Write(args ...[]byte) error {
	// get len
	var msgLen uint32
	for i := 0; i < len(args); i++ {
		msgLen += uint32(len(args[i]))
	}


	// don't copy
	if len(args) == 1 {
		this.w.Write(args[0])
		return nil
	}

	// merge the args
	msg := make([]byte, msgLen)
	l := 0
	for i := 0; i < len(args); i++ {
		copy(msg[l:], args[i])
		l += len(args[i])
	}

	//粘包完成后调下面的语句发送数据
	this.w.Write(msg)
	return this.w.Flush()
}

// goroutine safe
func (this *CustomAgent) Unmarshal(data []byte) (error) {
	if len(data) < 2 {
		return errors.New("protobuf data too short")
	}

	// id
	var id uint16
	if this.littleEndian {
		id = binary.LittleEndian.Uint16(data)
	} else {
		id = binary.BigEndian.Uint16(data)
	}
	return this.OnRecover(id,data[2:])
}

// goroutine safe
func (p *CustomAgent) Marshal(_id uint16,data []byte) ([][]byte, error) {
	id := make([]byte, 2)
	if p.littleEndian {
		binary.LittleEndian.PutUint16(id, _id)
	} else {
		binary.BigEndian.PutUint16(id, _id)
	}
	return [][]byte{id, data}, nil
}
