/**
一定要记得在confin.json配置这个模块的参数,否则无法使用
*/
package login

import (
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
	"github.com/golang/protobuf/proto"
	"proto/user"
	"github.com/liangdas/mqant/log"
)

var Module = func() module.Module {
	gate := new(Login)
	return gate
}

type Login struct {
	basemodule.BaseModule
}

func (m *Login) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "Login"
}
func (m *Login) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}
func (m *Login) OnInit(app module.App, settings *conf.ModuleSettings) {
	m.BaseModule.OnInit(m, app, settings)

	m.GetServer().RegisterGO("HD_Login", m.login)  //我们约定所有对客户端的请求都以Handler_开头
}

func (m *Login) Run(closeSig chan bool) {
}

func (m *Login) OnDestroy() {
	//一定别忘了关闭RPC
	m.GetServer().OnDestroy()
}
/**
客户端请求处理函数
 */
func (m *Login) login(session gate.Session, msg []byte) ( []byte,  string) {
	//解析客户端发送过来的user.LoginRequest结构体
	request:=&user.LoginRequest{}
	proto.UnmarshalMerge(msg, request)
	/////


	//这里开始登陆处理等操作

	/////

	//组建处理结果数据包
	datamsg,err:=proto.Marshal(&user.LoginSuccessResponse{})
	if err!=nil{
		log.Error(err.Error())
		return nil, ""
	}

	//给客户端发送处理结果  在routemap.go函数中已绑定 Login/Success与 leaf msgid 的关系
	errstr:=session.Send("Login/Success",datamsg)
	if errstr!=""{
		log.Error(errstr)
		return nil, errstr
	}

	//leaf 的消息只能通过Send函数主动发送，rpc的返回值没有作用，但还是必须返回两个参数(第二个参数是string) 这是mqant rpc 约定的
	return nil, ""
}

