package component

import (
	"net/http"
	"thh/app/models/Users"
)

const (
	SUCCESS = 1
	FAIL    = 0
)

type RequestContext struct {
	UserId   uint64
	userSet  bool
	userInfo Users.Users
}

func (r *RequestContext) GetUser() (Users.Users, error) {
	if r.userSet != false {
		return r.userInfo, nil
	}
	user, _ := Users.Get(r.UserId)

	r.userSet = true
	r.userInfo = user
	return r.userInfo, nil
}

type Response struct {
	Code int
	Data any
}

type DataMap map[string]interface{}

func BuildResponse(code int, data any) Response {
	return Response{code, data}
}

func SuccessResponse(data any) Response {
	return BuildResponse(http.StatusOK,
		map[string]any{
			"msg":  nil,
			"data": data,
			"code": SUCCESS,
		},
	)
}

func FailResponse(msg any) Response {
	return BuildResponse(http.StatusOK,
		map[string]any{
			"msg":  msg,
			"data": nil,
			"code": FAIL,
		},
	)
}