package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"my-hello/app/module"
	"my-hello/library/response"
)

//user服务api
var User = userApi{}

type userApi struct {
}

/**
用户注册接口
@tags 用户服务
@produce json
@param entity body  module.UserApiSignUpReq true "注册请求"
@router /user/signup [POST]
@success 200 {object} response.JsonResponse "执行结果"
*/
func (receiver *userApi) SignUp(r ghttp.Request) {

	var (
		apiReq     module.UserApiSignUpReq
		serviceReq module.UserServiceSignUpReq
	)

	//解析表单内容，映射表单内容字段类型，error不为空，表单提交失败
	if error := r.ParseForm(&apiReq); error != nil {
		response.JsonExit(&r, 1, error.Error())
	}
	//映射表单字段 与service层的结构体映射
	if error := gconv.Struct(&apiReq, &serviceReq); error != nil {
		response.JsonExit(&r, 1, error.Error())
	}
	//真正处理后端业务逻辑，处理注册 逻辑 ，操作数据库

}

//@summary 获取用户详情信息
//@tags 用户服务
//@produce json
//@router /user/profile [GET]
// @success 200 {object} model.User "用户信息"
func (receiver *userApi) Profile(r *ghttp.Request) {




}



//@summary 获取所有用户
//@tags 用户服务
//@produce json
//@router /user/list [GET]
// @success 200 {object} model.User "用户信息"
func (receiver *userApi) List(r *ghttp.Request) {





}


