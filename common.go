package common

import (
	"context"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/gsession"
	"github.com/gogf/gf/text/gstr"
	"net/http"
	"time"

	"github.com/gogf/gf/net/ghttp"
)

type PageSize struct {
	CurrentPage int         `json:"current_page"`
	Data        interface{} `json:"data"`
	LastPage    int         `json:"last_page"`
	PerPage     int         `json:"per_page"`
	Total       int         `json:"total"`
}

// SetPage 设置分页
func SetPage(page, limit, total int, data interface{}) *PageSize {
	var size = new(PageSize)
	if page == 1 {
		size.LastPage = 1
	} else {
		size.LastPage = page - 1
	}
	size.PerPage = limit
	size.Total = total
	size.CurrentPage = page
	size.Data = data
	return size
}

// MiddlewareError 异常处理中间件
func MiddlewareError(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.ClearBuffer()
		Error(r).SetMsg(err.Error()).End()
	}
}

// AuthBase 鉴权中间件，只有前端或者后端登录成功之后才能通过
func AuthBase(r *ghttp.Request) {
	info := r.Session.Get("admin", nil)
	if info != nil {
		r.Middleware.Next()
	} else {
		ErrorsLogin(r)
	}
}

// GetCapitalPass MD5化并转换为大写
func GetCapitalPass(val string) string {
	md5, err := gmd5.Encrypt(val)
	if err != nil {
		panic(err.Error())
	}
	return gstr.CaseCamel(md5)
}

// Transaction 简单封装事务操作
func Transaction(function func() error) {
	db := gdb.TX{}
	err := db.Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		return function()
	})
	if err != nil {
		panic(err.Error())
	}
}

// CreateCron 创建定时任务
func CreateCron(time string, name string, operate func()) {
	_, err := gcron.Add(time, func() {
		operate()
	}, name)
	if err != nil {
		panic(err.Error())
	}
}

// StopCron 紧停止指定定时任务
func StopCron(name string) {
	gcron.Stop(name)
}

// RemoveCron 停止并删除定时任务
func RemoveCron(name string) {
	gcron.Stop(name)
	gcron.Remove(name)
}

// PostResult 建立POST请求并返回结果
func PostResult(url string, data g.Map, header string, class string) *http.Response {
	if url == "" {
		panic("请求地址不可为空")
	}
	client := g.Client()
	if header != "" {
		client = client.HeaderRaw(header)
	}
	switch class {
	case "json":
		client = client.ContentJson()
	case "xml":
		client = client.ContentXml()
	default:
	}
	result, err := client.Post(url, data)
	if err != nil {
		panic(err.Error())
	}
	return result.Response
}

func GetResult(url string, data g.Map) *http.Response {
	client := g.Client()
	if url == "" {
		panic("请求地址不可为空")
	}
	result, err := client.Get(url, data)
	if err != nil {
		panic(err.Error())
	}
	return result.Response
}

func SetSession(time time.Time) g.Map {
	return g.Map{
		"SessionMaxAge":  time,
		"SessionStorage": gsession.NewStorageMemory(),
	}
}
