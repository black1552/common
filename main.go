package common

import (
	"log"
	"os"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gsession"
)

// Start 创建无定时的http
func Start(time time.Duration) {
	s := g.Server()
	s.SetServerRoot(gfile.MainPkgPath())
	upload := gfile.MainPkgPath() + "/public/upload"
	if !isDir(upload) {
		_ = os.MkdirAll(upload, os.ModePerm)
		s.AddStaticPath("/upload", upload)
	}
	s.SetFileServerEnabled(true)
	if time != 0 {
		_ = s.SetConfigWithMap(g.Map{
			"SessionMaxAge":  time,
			"SessionStorage": gsession.NewStorageMemory(),
			"SessionPath":    gfile.MainPkgPath() + "/public/session",
		})
	}
	s.Use(MiddlewareError)
	s.Run()
}

// StartCorn 创建有定时任务的http
func StartCorn(time time.Duration, cronTime, name string, cron func()) {
	s := g.Server()
	s.SetServerRoot(gfile.MainPkgPath())
	upload := gfile.MainPkgPath() + "/public/upload"
	if !isDir(upload) {
		_ = os.MkdirAll(upload, os.ModePerm)
		s.AddStaticPath("/upload", upload)
	}
	s.SetFileServerEnabled(true)
	_ = s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time,
		"SessionStorage": gsession.NewStorageMemory(),
		"SessionPath":    gfile.MainPkgPath() + "/public/session",
	})
	CreateCron(cronTime, name, cron)
	s.Use(MiddlewareError)
	s.Run()
}

// StartTcp 创建一个TCP的http
func StartTcp(time time.Duration, tcpFun func()) {
	go Start(time)

	go tcpFun()
	g.Wait()
}

// StartWebSocket 创建一个带websocket的http
func StartWebSocket(time time.Duration, webFun func()) {
	go Start(time)
	go webFun()

	g.Wait()
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}
