package model

// LoginReq 登录请求体
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginCheck 登录检测（为了方便这里就不适用gorm从数据库查询了）
func LoginCheck(req LoginReq) (bool,User,error)  {
	if req.Email=="123456@163.com" && req.Password=="123456" {
		return true, User{
			Name: "张三",
			Email: "123456@163.com",
		}, nil
	}else {
		return false, User{}, nil
	}
}
