package service

import (
	"context"
	"time"
)

// 典型的 Option 设计模式
type Option func(*App)

// ShutdownCallback 采用 context.Context 来控制超时，而不是用 time.After 是因为
// - 超时本质上是使用这个回调的人控制的
// - 我们还希望用户知道，他的回调必须要在一定时间内处理完毕，而且他必须显式处理超时错误
type ShutdownCallback func(ctx context.Context)

// 这里我已经预先定义好了各种可配置字段
type App struct {
	servers []*Server

	// 优雅退出整个超时时间，默认30秒
	shutdownTimeout time.Duration

	// 优雅退出时候等待处理已有请求时间，默认10秒钟
	waitTime time.Duration
	// 自定义回调超时时间，默认三秒钟
	cbTimeout time.Duration

	cbs []ShutdownCallback
}

func WithServers(servers ...*Server) Option {
	return func(app *App) {
		app.servers = servers
	}
}

func WithShutDownTimeOut(shutdownTimeout time.Duration) Option {
	return func(app *App) {
		app.shutdownTimeout = shutdownTimeout
	}
}

func WithWaitTime(waitTime time.Duration) Option {
	return func(app *App) {
		app.waitTime = waitTime
	}
}

func WithCbTimeout(cbTimeout time.Duration) Option {
	return func(app *App) {
		app.cbTimeout = cbTimeout
	}
}

func WithShutdownCallbacks(cbs ...ShutdownCallback) Option {
	return func(app *App) {
		app.cbs = cbs
	}
}
