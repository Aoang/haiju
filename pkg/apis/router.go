package apis

import (
	"github.com/Aoang/haiju/pkg/model"
	"os"
	"strings"

	"github.com/Aoang/haiju/pkg/apis/internal/middleware"
	"github.com/Aoang/haiju/pkg/apis/internal/path"
	"github.com/Aoang/haiju/pkg/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// TODO 业务逻辑没打 log

var log = logger.NewLogger(os.Stdout)

// MapRoutes returns a gin engine and binds controllers with request URLs
func MapRoutes() *gin.Engine {
	ret := gin.New()
	ret.Use(gin.Recovery())

	store := cookie.NewStore([]byte("haiju"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   1200,
		Secure:   strings.HasPrefix(model.Conf.Server, "https"),
		HttpOnly: true,
	})
	ret.Use(sessions.Sessions("haiju", store))

	api := ret.Group(path.APIPath)

	api.POST(path.Login, loginAction)
	api.POST(path.Logout, logoutAction)

	// 用户管理
	userGroup := api.Group(path.User)
	userGroup.Use(middleware.LoginCheck)
	userGroup.GET("/:id", getUserAction)
	userGroup.POST("/", addUserAction)

	// 房产管理
	buildingGroup := api.Group(path.Building)
	buildingGroup.GET("/:id", getBuildingAction)
	buildingGroup.POST("/", addBuildingAction)
	buildingGroup.PUT("/:id", updateBuildingAction)
	buildingGroup.DELETE("/:id", deleteBuildingAction)

	// 文件管理
	fileGroup := api.Group(path.File)
	fileGroup.Static("/",model.Conf.StaticRoot)
	fileGroup.POST("/:id", addFileAction)
	fileGroup.DELETE("/:id", deleteFileAction)

	return ret
}
