package api

import (
	"errors"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"my-hello/app/module"
	"my-hello/app/service"
	"my-hello/app/utils"
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
@router /user/register [POST]
@success 200 {object} response.JsonResponse "执行结果"
*/
func (receiver *userApi) Register(r *ghttp.Request) {
	var (
		apiReq     module.UserApiSignUpReq
		serviceReq module.User
	)
	//解析表单内容，映射表单内容字段类型，error不为空，表单提交失败
	if error := r.ParseForm(&apiReq); error != nil {
		response.JsonExit(r, 1, error.Error())
	}
	//映射表单字段 与service层的结构体映射
	if error := gconv.Struct(&apiReq, &serviceReq); error != nil {
		response.JsonExit(r, 1, error.Error())
	}
	//真正处理后端业务逻辑，处理注册 逻辑 ，操作数据库
	if err := service.User.Register(&serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		//成功
		response.JsonExit(r, 200, "success")
	}

}

/**
用户登录接口
@tags 用户服务
@produce json
@param entity body  module.UserApiSignUpReq true "登录请求"
@router /user/login [POST]
@success 200 {object} response.JsonResponse "执行结果"
*/
func (receiver *userApi) Login(r *ghttp.Request) {
	var (
		apiReq     module.UserApiSignInReq
		serviceReq module.User
	)
	//解析表单内容，映射表单内容字段类型，error不为空，表单提交失败
	if error := r.ParseForm(&apiReq); error != nil {
		response.JsonExit(r, 1, error.Error())
	}
	//映射表单字段 与service层的结构体映射
	if error := gconv.Struct(&apiReq, &serviceReq); error != nil {
		response.JsonExit(r, 1, error.Error())
	}
	//真正处理后端业务逻辑，处理注册 逻辑 ，操作数据库
	if user := service.User.Login(r, &serviceReq); user != nil {
		response.JsonExit(r, 200, "success", user)
	} else {
		//失败
		response.JsonExit(r, 1, errors.New("账号或密码错误").Error())
	}
}

//@summary 获取用户详情信息fromDb
//@tags 用户服务
//@produce json
//@router /user/profile [GET]
// @success 200 {object} model.User "用户信息"
func (receiver *userApi) GetUserFromDb(r *ghttp.Request) {
	passport := r.GetString("passport")
	if userDb := service.User.GetUserFromDb(passport); userDb != nil {
		response.JsonExit(r, 200, "success", userDb)
	} else {
		response.JsonExit(r, 1, errors.New("账号不存在").Error())
	}
}

//@summary 获取用户详情信息fromSession
//@tags 用户服务
//@produce json
//@router /user/profile [GET]
// @success 200 {object} model.User "用户信息"
func (receiver *userApi) GetUserFromSession(r *ghttp.Request) {
	userSession := service.Session.GetUser(r.Context())
	response.JsonExit(r, 200, "success", userSession)
}

//@summary 获取所有用户
//@tags 用户服务
//@produce json
//@router /user/list [GET]
// @success 200 {object} model.User "用户信息"
func (receiver *userApi) List(r *ghttp.Request) {
	userLists := service.User.List(r)
	response.JsonExit(r, 200, "success", userLists)
}

//@summary 分页获取用户
//@tags 用户服务
//@produce json
//@router /user/list [GET]
// @success 200 {object} model.User "用户信息"
func (receiver *userApi) PageList(r *ghttp.Request) {
	page := r.GetInt("page")
	limit := r.GetInt("limit")
	userLists := service.User.PageList(page, limit)
	response.JsonExit(r, 200, "success", userLists)
}

func (receiver *userApi) RedisSet(r *ghttp.Request) {
	if err := service.RedisService.SetVal(r.GetString("key"), r.GetString("value")); err != nil {
		response.JsonExit(r, -1, "error", err.Error())
	} else {
		response.JsonExit(r, 200, "success")
	}
}

func (receiver *userApi) RedisGet(r *ghttp.Request) {
	value := service.RedisService.GetVal(r.GetString("key"))
	response.JsonExit(r, 200, "success", value)
}

func (receiver *userApi) RedisHSet(r *ghttp.Request) {
	if err := service.RedisService.Hset(r.GetString("key"), r.GetString("value")); err != nil {
		response.JsonExit(r, -1, "error", err.Error())
	} else {
		response.JsonExit(r, 200, "success")
	}
}

func (receiver *userApi) RedisHGet(r *ghttp.Request) {
	value := service.RedisService.HgetAll()
	response.JsonExit(r, 200, "success", value)
}

func (receiver *userApi) RedisHMSet(r *ghttp.Request) {
	//var reqMap g.Map
	//if err := r.Parse(&reqMap); err != nil {
	//	response.JsonExit(r,-1,errors.New("map类型转换失败！").Error())
	//	return
	//}
	//reqMap:=r.Get("map")
	reqMap:=r.GetVar("map").Map()
	if err := service.RedisService.HMset(reqMap); err != nil {
		response.JsonExit(r, 500, "error", err.Error())
	} else {
		response.JsonExit(r, 200, "success")
	}
}
func (receiver *userApi) RedisHMGet(r *ghttp.Request) {
	value := service.RedisService.HMget(r.GetString("key"))
	response.JsonExit(r, 200, "success", value)
}

func (receiver *userApi) HighRedisSet(r *ghttp.Request) {
	if err := utils.RediApi.Set(r.GetString("key"), r.GetString("value")); err != nil {
		response.JsonExit(r, 500, "error", err.Error())
	} else {
		response.JsonExit(r, 200, "success")
	}
}
func (receiver *userApi) HighRedisGet(r *ghttp.Request) {
	value := utils.RediApi.Get(r.GetString("key"))
	response.JsonExit(r, 200, "success", value)
}
