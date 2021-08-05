package common

import "github.com/gogf/gf/net/ghttp"

type Json struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ApiResp struct {
	r    *ghttp.Request
	json *Json
}

func Errors(r *ghttp.Request, msg string) {
	r.Response.ClearBuffer()
	_ = r.Response.WriteJson(Json{402, nil, msg})
	r.Exit()
}

func ErrorsLogin(r *ghttp.Request) {
	_ = r.Response.WriteJson(Json{401, nil, "请登录后操作"})
	r.Exit()
}

func ResultVersion(r *ghttp.Request, data interface{}) {
	_ = r.Response.WriteJson(data)
	r.Exit()
}

func (a *ApiResp) SetDate(data interface{}) *ApiResp {
	a.json.Data = data
	return a
}

func (a *ApiResp) SetCode(code int) *ApiResp {
	a.json.Code = code
	return a
}

func (a *ApiResp) SetMsg(msg string) *ApiResp {
	a.json.Msg = msg
	return a
}

func Success(r *ghttp.Request) *ApiResp {
	json := Json{
		Code: 1,
	}
	var a = ApiResp{
		r:    r,
		json: &json,
	}
	return &a
}

func Error(r *ghttp.Request) *ApiResp {
	json := Json{
		Code: 402,
	}
	var a = ApiResp{
		r:    r,
		json: &json,
	}
	return &a
}

func (a *ApiResp) End() {
	_ = a.r.Response.WriteJsonExit(a.json)
}

