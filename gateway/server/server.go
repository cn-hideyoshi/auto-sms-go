package server

import (
	"blog.hideyoshi.top/gateway/rpc"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(r *gin.Engine, servName string, addr string) {
	serv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		log.Printf("%s启动,端口号%s", servName, serv.Addr)
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("%s启动异常:%v", servName, err)
		}
	}()

	quit := make(chan os.Signal)
	// SIGINT 用户发送INTR字符(Ctrl+C)触发 kill -2
	// SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("正在关闭程序%s", servName)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatalf("%s异常关闭,因为:%v", servName, err)
	}
	select {
	case <-ctx.Done():
		log.Printf("等待%s关闭...", servName)
	}
	rpc.Server.ResolverClose()
	log.Printf("%s关闭成功", servName)
}
