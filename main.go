package common

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gsession"
	"log"
	"os"
	"time"
)

func Start(upload, root string, time time.Time) {
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

func StartCorn(upload, root string, time time.Time, cronTime, name string, cron func()) {
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

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}
