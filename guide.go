package regin

import (
	"regin/origin"
	"regin/app"
)

type Guider struct{}

// 定义Guider
var Guide *Guider

func init() {
	Guide = &Guider{}
}

// 开启http服务
func (e *Guider) HttpService() {
	origin.Engine.HttpServer(app.NewHttp())
}


