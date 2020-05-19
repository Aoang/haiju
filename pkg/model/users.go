package model

type User struct {
	Model

	Username       string // 用户名
	Password       string // 密码
	RoleID         uint64 // 角色
	Name           string // 真实姓名
	Sex            uint8  // 性别
	IdentityNumber string // 身份证号
	Phone          uint32 // 手机号
	WeiXinID       string // 微信 id
	ZhiFuBaoID     string // 支付宝 id
}

type Role struct {
	Model

	Name   string `json:"name" `
	Path   string `json:"path"`
	Method string `json:"method"`
}

//func (i *Role) Add() (bool, error) {
//	return casbin.Use().AddPolicy(i.Name, i.Path, i.Method)
//}
