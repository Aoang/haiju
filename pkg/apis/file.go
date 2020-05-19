package apis

import (
	"github.com/Aoang/haiju/pkg/database"
	"github.com/Aoang/haiju/pkg/model"
	"github.com/Aoang/haiju/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

/*
	TODO
		上传应该限制大小、获取文件后缀格式
		对于部分其他资源，应该增加上传到其他目录
		文件的获取应该增加中间件、设定 token 校验、权限处理
*/
func addFileAction(c *gin.Context) {
	result := model.NewResult()
	defer c.JSON(http.StatusOK, result)

	file, err := c.FormFile("file")
	if err != nil {
		result.Code = -1
		return
	}

	f := &model.File{
		OID:  util.Str2Uint64(c.Query("id")),
		UUID: uuid.New(),
	}
	if err := database.File.Add(f); err != nil {
		result.Code = -1
		return
	}

	if err := c.SaveUploadedFile(file, model.Conf.StaticRoot+f.UUID.String()); err != nil {
		result.Code = -1
		return
	}
}

func deleteFileAction(c *gin.Context) {
	result := model.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := &model.File{
		Model: model.Model{
			ID: util.Str2Uint64(c.Query("id")),
		},
	}

	if database.File.Delete(arg) != nil {
		result.Code = -2
	}

}
