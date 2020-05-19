package apis

import (
	"github.com/Aoang/haiju/pkg/database"
	"github.com/Aoang/haiju/pkg/model"
	"github.com/Aoang/haiju/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUserAction(c *gin.Context) {
	result := model.NewResult()
	defer c.JSON(http.StatusOK, result)

	if c.Param("id") == "" {
		result.Data = database.GetAllUser()
		return
	}

	user := &model.User{}
	user.ID = util.Str2Uint64(c.Param("id"))
	if err := database.User.Get(user); err != nil {
		return
	}
	result.Data = user
}

func addUserAction(c *gin.Context) {
	result := model.NewResult()
	defer c.JSON(http.StatusOK, result)

}
