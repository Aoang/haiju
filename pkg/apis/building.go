package apis

import (
	"github.com/Aoang/haiju/pkg/database"
	"github.com/Aoang/haiju/pkg/model"
	"github.com/Aoang/haiju/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getBuildingAction(c *gin.Context) {
	result := model.NewResult()
	defer c.JSON(http.StatusOK, result)

	if c.Param("id") == "" {
		return
	}

	// TODO 这里只获取了基本数据，没有将小区信息等等加入
	//		还需要将图片资源加入
	b, s := database.GetLimitBuilding(util.Str2int(c.Param("id")))
	result.Data = map[string]interface{}{
		"zs":     s,
		"buding": b,
	}
}

func addBuildingAction(c *gin.Context) {
	result := model.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := &model.Building{}
	if err := c.BindJSON(arg); nil != err {
		result.Code = -1
		result.Msg = "parses add user request failed"

		return
	}

	if database.Building.Add(arg) != nil {
		result.Code = -2
	}

}

func updateBuildingAction(c *gin.Context) {
	result := model.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := &model.Building{}
	if err := c.BindJSON(arg); nil != err {
		result.Code = -1
		result.Msg = "parses add user request failed"

		return
	}

	if database.Building.Update(arg) != nil {
		result.Code = -2
	}

}

func deleteBuildingAction(c *gin.Context) {
	result := model.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := model.Building{
		Model: model.Model{
			ID: util.Str2Uint64(c.Query("id")),
		},
	}

	if database.Building.Delete(&arg) != nil {
		result.Code = -2
	}
}
