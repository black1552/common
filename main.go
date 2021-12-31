package common

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gsession"
	"log"
	"os"
	"time"
)

// Start 创建无定时的http
func Start(upload, root string, time time.Duration) {
	s := g.Server()
	s.SetServerRoot(root)
	if !isDir(upload) {
		_ = os.MkdirAll(upload, os.ModePerm)
		s.AddStaticPath("/upload", upload)
	}
	s.SetFileServerEnabled(true)
	_ = s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time,
		"SessionStorage": gsession.NewStorageMemory(),
	})
	s.Use(MiddlewareError)
	s.Run()
}

// StartCorn 创建有定时任务的http
func StartCorn(upload, root string, time time.Duration, cronTime, name string, cron func()) {
	s := g.Server()
	s.SetServerRoot(root)
	if !isDir(upload) {
		_ = os.MkdirAll(upload, os.ModePerm)
		s.AddStaticPath("/upload", upload)
	}
	s.SetFileServerEnabled(true)
	_ = s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time,
		"SessionStorage": gsession.NewStorageMemory(),
	})
	CreateCron(cronTime, name, cron)
	s.Use(MiddlewareError)
	s.Run()
}

// StartTcp 创建一个TCP的http
func StartTcp(upload, root string, time time.Duration, tcpFun func()) {
	go func() {
		s := g.Server()
		s.SetServerRoot(root)
		if !isDir(upload) {
			_ = os.MkdirAll(upload, os.ModePerm)
			s.AddStaticPath("/upload", upload)
		}
		s.SetFileServerEnabled(true)
		_ = s.SetConfigWithMap(g.Map{
			"SessionMaxAge":  time,
			"SessionStorage": gsession.NewStorageMemory(),
		})
		s.Use(MiddlewareError)
		s.Run()
	}()

	go tcpFun()
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
