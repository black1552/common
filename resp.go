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

func ResultVersion(r *ghttp.Request, data interface{}) {
	_ = r.Response.WriteJson(data)
	r.Exit()
}

// SetCode 设置状态码
func (a *ApiResp) SetCode(code int) *ApiResp {
	a.json.Code = code
	return a
}

// SetDate 设置数据
func (a *ApiResp) SetDate(data interface{}) *ApiResp {
	a.json.Data = data
	return a
}

// SetMsg 设置信息
func (a *ApiResp) SetMsg(msg string) *ApiResp {
	a.json.Msg = msg
	return a
}

// Success 设置成功JSON
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

// Error 设置错误JSON
func Error(r *ghttp.Request) *ApiResp {
	json := Json{
		Code: 0,
	}
	var a = ApiResp{
		r:    r,
		json: &json,
	}
	return &a
}

// End 返回JSON
func (a *ApiResp) End() {
	a.r.Response.Status = 200
	_ = a.r.Response.WriteJsonExit(a.json)
}
