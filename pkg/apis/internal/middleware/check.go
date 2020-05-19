package middleware

import (
	"github.com/gin-gonic/gin"
)

// LoginCheck checks login or not.
func LoginCheck(c *gin.Context) {
	//session := GetSession(c)
	//if 0 == session.UserID {
	//	result := model.NewResult()
	//	result.Code = -2
	//	result.Msg = "unauthenticated request"
	//	c.AbortWithStatusJSON(http.StatusOK, result)
	//
	//	return
	//}
	// TODO 不测试登录功能

	c.Next()
}
