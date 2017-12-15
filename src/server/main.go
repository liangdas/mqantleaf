package main

import (
	"github.com/liangdas/mqant"
	"server/gateleaf"
	"server/login"
	"github.com/liangdas/mqant/module/modules"
)
func main() {
	app := mqant.CreateApp()
	app.Run(true, //只有是在调试模式下才会在控制台打印日志, 非调试模式下只在日志文件中输出日志
		modules.MasterModule(),
		gateleaf.Module(),  //这是支持leaf网关模块
		login.Module(), //这是用户登录验证模块
	)
}

