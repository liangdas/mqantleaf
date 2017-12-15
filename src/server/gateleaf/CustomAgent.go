// Copyright 2014 loolgame Author. All Rights Reserved.
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
//与客户端通信的自定义粘包示例，需要mqant v1.6.4版本以上才能运行
//该示例只用于简单的演示，并没有实现具体的粘包协议
package gateleaf

import (
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/network"
	"bufio"
	"github.com/liangdas/mqant/log"
	"fmt"
	"strings"
	"time"
)
func NewAgent(module module.RPCModule)*CustomAgent{
	a := &CustomAgent{
		module:module,
	}
	return a
}

type CustomAgent struct {
	gate.Agent
	module 		module.RPCModule
	session                          gate.Session
	conn                             network.Conn
	r                                *bufio.Reader
	w                                *bufio.Writer
	gate                             gate.Gate
	rev_num                          int64
	send_num                         int64
	last_storage_heartbeat_data_time int64 //上一次发送存储心跳时间
	isclose                          bool
	littleEndian			 bool //小端模式
	lenMsgLen			 int //数据头长度
	routeMap			 map[interface{}]interface{} //路由映射表 msgid--mqant rpc   mqant rpc--msgid
}
func (this *CustomAgent) OnInit(gate gate.Gate,conn network.Conn)error{
	log.Info("CustomAgent","OnInit")
	this.conn=conn
	this.gate=gate
	this.r=bufio.NewReader(conn)
	this.w=bufio.NewWriter(conn)
	this.isclose=false
	this.rev_num=0
	this.send_num=0
	this.littleEndian=false
	this.lenMsgLen=2
	this.routeMap=make(map[interface{}]interface{})
	this.RouteMap()

	return nil
}
/**
给客户端发送消息
 */
func (this *CustomAgent) WriteMsg(topic string, body []byte) error{
	if _id,ok:=this.routeMap[topic];ok{
		return this.WriteMarshal(_id.(uint16),body)
	}else{
		return fmt.Errorf("message id %d not registered", _id)
	}
	return nil
}


func (this *CustomAgent)Run() (err error){
	this.session, err = this.gate.NewSessionByMap( map[string]interface{}{
		"Sessionid": fmt.Sprintf("%d",time.Now().UnixNano()),
		"Network":   this.conn.RemoteAddr().Network(),
		"IP":        this.conn.RemoteAddr().String(),
		"Serverid":  this.module.GetServerId(),
		"Settings":  make(map[string]string),
	})
	this.gate.GetAgentLearner().Connect(this) //发送连接成功的事件 这是必须调用的

	//这里可以循环读取客户端的数据
	for{
		data, err := this.Read()
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}
		err=this.Unmarshal(data)
		if err!=nil{
			log.Error(err.Error())
		}
	}

	//这个函数返回后连接就会被关闭
	return err
}
/**
接收到一个数据包
 */
func (this *CustomAgent) OnRecover(_id uint16,msg []byte)error {
	if topic,ok:=this.routeMap[_id];ok{
		topics := strings.Split(topic.(string), "/")
		if len(topics) < 2 {
			errorstr := "Topic must be [moduleType@moduleID]/[handler]|[moduleType@moduleID]/[handler]/[msgid]"
			return fmt.Errorf(errorstr)
		} else if len(topics) == 3 {
			//msgid = topics[2]
		}
		startsWith := strings.HasPrefix(topics[1], "HD_")
		if !startsWith {
			return fmt.Errorf("Method(%s) must begin with 'HD_'", topics[1])
		}
		//下面调用后端模块
		moduleType:=topics[0]
		_func:=topics[1]
		e := this.module.RpcInvokeNR(moduleType,_func,this.GetSession(),msg)
		return e

	}else{
		return fmt.Errorf("message id %d not registered", _id)
	}
}


