package main

import (
	"github.com/Aoang/haiju/pkg/model"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Aoang/haiju/pkg/apis"
	"github.com/Aoang/haiju/pkg/database"
	"github.com/Aoang/haiju/pkg/logger"
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
)

var log *logger.Logger

func init() {
	logger.SetLevel("debug")
	log = logger.NewLogger(os.Stdout)

	model.LoadConf()

	if "dev" == model.Conf.RuntimeMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
}

func main() {
	database.ConnectDB()

	router := apis.MapRoutes()
	server := &http.Server{
		Addr:    model.Conf.Server,
		Handler: router,
	}
	handleSignal(server)

	log.Info("Server is running ")
	go server.ListenAndServe()
	http.ListenAndServe("127.0.0.1:6060", nil)
}

// handleSignal handles system signal for graceful shutdown.
func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		log.Infof("got signal [%s], exiting now", s)
		if err := server.Close(); nil != err {
			log.Errorf("server close failed: " + err.Error())
		}

		database.DisconnectDB()

		log.Infof("Server exited")
		os.Exit(0)
	}()
}
