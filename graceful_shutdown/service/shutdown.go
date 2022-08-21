package service

import (
	"context"
	"log"
	"sync"
)

// shutdown 你要设计这里面的执行步骤。
func (app *App) shutdown() {
	log.Println("开始关闭应用，停止接收新请求")
	// 你需要在这里让所有的 server 拒绝新请求

	// 在这里等待一段时间
	// 并发关闭服务器，同时要注意协调所有的 server 都关闭之后才能步入下一个阶段
	app.stopAllServer()

	log.Println("开始执行自定义回调")
	// 并发执行回调，要注意协调所有的回调都执行完才会步入下一个阶段
	app.handleAllCallBack()
	// 释放资源
	log.Println("开始释放资源")
	app.close()
}

func (app *App) stopAllServer() {
	wg := &sync.WaitGroup{}

	for _, s := range app.servers {
		wg.Add(1)

		go func(s *Server) {
			defer wg.Done()
			s.rejectReq()
			ctx, cancle := context.WithTimeout(context.Background(), app.waitTime)
			s.stop(ctx)
			defer cancle()
		}(s)
	}
	wg.Wait()
}

func (app *App) handleAllCallBack() {
	wg := &sync.WaitGroup{}

	for _, cb := range app.cbs {
		wg.Add(1)
		go func(cb ShutdownCallback) {
			defer wg.Done()
			ctx, cancle := context.WithTimeout(context.Background(), app.cbTimeout)
			cb(ctx)
			defer cancle()
		}(cb)
	}
	wg.Wait()
}
