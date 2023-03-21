package console

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"thh/arms"
	"thh/arms/app"
	"thh/arms/logger"
	"thh/bundles/config"
	"thh/routes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// CmdServe represents the available web sub-command.
var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(_ *cobra.Command, _ []string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	logger.Info("Thousand-hand:start")
	logger.Info(fmt.Sprintf("Thousand-hand:useMem %d KB", m.Alloc/1024/8))

	go RunJob()

	// 初始化应用程序
	if config.GetBool("app.debug", true) {
		go func() {
			// go tool pprof http://localhost:6060/debug/pprof/profile
			//http://127.0.0.1:7070/debug/pprof/
			err := http.ListenAndServe("0.0.0.0:7070", nil)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	//fiberServe()
	ginServe()
}

const (
	ENV      = "env"
	EnvProd  = "production"
	EnvLocal = "local"
)

func ginServe() {

	var (
		port   = config.GetString("APP_PORT", 8080)
		isProd = config.Get("APP_ENV") == EnvProd
	)
	var engine *gin.Engine
	switch isProd {
	case true:
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		engine.Use(gin.Recovery())
		break
	default:
		engine = gin.Default()
		break
	}

	routes.RegisterByGin(engine)

	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.Info("Thousand-hand:listen " + port)
	fmt.Printf("use http://localhost:%s\n", port)
	fmt.Printf("use http://%v:%v\n", arms.GetLocalIp(), port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Println("start use:" + cast.ToString(app.GetUnitTime()))

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	_ = <-quit

	logger.Std().Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Std().Println("Server Shutdown:", err)
	}
	logger.Std().Println("Server exiting")
}