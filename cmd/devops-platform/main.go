package main

import "devops-platform/pkg/beans"

// @title 运维系统
// @version 2.0
// @description 运维系统 api.
// @schemes http https
// @host 127.0.0.1
// @BasePath /
// @contact.name zhangxin
// @contact.email xin.zhang@hicom.com
func main() {
	beans.Start()
}
