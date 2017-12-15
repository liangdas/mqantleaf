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
//CustomAgent.go 的一些不常改动的函数就放在这里了
package gateleaf

import "github.com/liangdas/mqant/gate"
func (this *CustomAgent)Close(){
	this.conn.Close()
}
func (this *CustomAgent)OnClose() error{
	this.isclose = true
	//这个一定要调用，不然gate可能注销不了,造成内存溢出
	this.gate.GetAgentLearner().DisConnect(this) //发送连接断开的事件
	return nil
}
func (this *CustomAgent)Destroy(){
	this.conn.Destroy()
}
func (this *CustomAgent)RevNum() int64{
	return this.rev_num
}
func (this *CustomAgent)SendNum() int64{
	return this.send_num
}
func (this *CustomAgent)IsClosed() bool{
	return this.isclose
}
func (this *CustomAgent)GetSession() gate.Session{
	return this.session
}
