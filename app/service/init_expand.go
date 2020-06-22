package service

import (
	"regin/app/cache"
	"regin/app/db"
)

// 初始化扩展包
func (a *Application) initExpand() {
	cache.Redis.Init(a.GetConfig("redis").ToAnyMap())       // 初始化Redis
	db.Query.LoadConfig(a.GetConfig("database").ToAnyMap()) // 初始化Database
}
