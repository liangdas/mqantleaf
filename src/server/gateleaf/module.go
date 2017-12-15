/**
一定要记得在confin.json配置这个模块的参数,否则无法使用
*/
package gateleaf

import (
	"fmt"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate/base"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/gate"
)

var Module = func() module.Module {
	this := new(GateLeaf)
	return this
}

type GateLeaf struct {
	basegate.Gate //继承
}


func (this *GateLeaf) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "GateLeaf"
}
func (this *GateLeaf) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}

//与客户端通信的自定义粘包示例，需要mqant v1.6.4版本以上才能运行
//该示例只用于简单的演示，并没有实现具体的粘包协议
func (this *GateLeaf)CreateAgent() gate.Agent{
	agent:= NewAgent(this)
	return agent
}

func (this *GateLeaf) OnInit(app module.App, settings *conf.ModuleSettings) {
	//注意这里一定要用 gate.Gate 而不是 module.BaseModule
	this.Gate.OnInit(this, app, settings)


	//与客户端通信的自定义粘包示例，需要mqant v1.6.4版本以上才能运行
	//该示例只用于简单的演示，并没有实现具体的粘包协议
	this.Gate.SetCreateAgent(this.CreateAgent)

	this.Gate.SetSessionLearner(this)
	this.Gate.SetStorageHandler(this) //设置持久化处理器
}
//当连接建立
func (this *GateLeaf) Connect(session gate.Session)  {
	log.Info("客户端建立了链接")
}
//当连接关闭
func (this *GateLeaf) DisConnect(session gate.Session) {
	log.Info("客户端断开了链接")
}

/**
存储用户的Session信息
Session Bind Userid以后每次设置 settings都会调用一次Storage
*/
func (this *GateLeaf) Storage(Userid string, session gate.Session) (err error) {
	log.Info("需要处理对Session的持久化")
	return nil
}

/**
强制删除Session信息
*/
func (this *GateLeaf) Delete(Userid string) (err error) {
	log.Info("需要删除Session持久化数据")
	return nil
}

/**
获取用户Session信息
用户登录以后会调用Query获取最新信息
*/
func (this *GateLeaf) Query(Userid string) ([]byte,  error) {
	log.Info("查询Session持久化数据")
	return nil, fmt.Errorf("no redis")
}

/**
用户心跳,一般用户在线时60s发送一次
可以用来延长Session信息过期时间
*/
func (this *GateLeaf) Heartbeat(Userid string) {
	log.Info("用户在线的心跳包")
}
