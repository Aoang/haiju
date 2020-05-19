package apis

import (
	"github.com/Aoang/haiju/pkg/apis/internal/middleware"
	"github.com/Aoang/haiju/pkg/database"
	"github.com/Aoang/haiju/pkg/model"
	"github.com/Aoang/haiju/pkg/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO 登录方式应当有多种，微信、支付宝登录都未进行对接
// loginAction login a user.
func loginAction(c *gin.Context) {
	result := model.NewResult()
	defer c.JSON(http.StatusOK, result)

	var arg map[string]string
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "登录信息错误"

		return
	}

	// TODO 这里应该根据规则校验信息，避免攻击
	user := &model.User{
		Username: arg["username"],
	}

	if err := database.User.Get(user); err != nil {
		// TODO log
		result.Code = -1
		result.Msg = "登录信息错误"

		return
	}

	if !util.Check(user.Password, arg["password"]) {
		result.Code = -1
		result.Msg = "密码不正确"
		return
	}

	sd := middleware.SessionData{
		UserID: user.ID,
		RoleID: user.RoleID,
	}

	if err := sd.Save(c); nil != err {
		log.Errorf("saves session failed: " + err.Error())
	}
}

// logoutAction logout a user.
func logoutAction(c *gin.Context) {
	result := model.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})

	if err := middleware.ClearSession(c); nil != err {
		log.Errorf("saves session failed: " + err.Error())
	}
}
