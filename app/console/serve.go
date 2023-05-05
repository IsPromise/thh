package console

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"thh/bundles/kernel"
	"thh/bundles/logging"
	"thh/routes"
	"time"

	"github.com/leancodebox/goose/serverinfo"

	"github.com/leancodebox/goose/preferences"

	"github.com/spf13/cast"

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
	info("Thousand-hand:start")
	info(fmt.Sprintf("Thousand-hand:useMem %d KB", m.Alloc/1024/8))

	go RunJob()

	// 初始化应用程序
	if preferences.GetBool("app.debug", true) {
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
		port   = preferences.GetString("app.port", 8080)
		isProd = preferences.Get("app.env") == EnvProd
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
	info("Thousand-hand:listen " + port)
	fmt.Printf("use http://localhost:%s\n", port)
	ip, _ := serverinfo.GetLocalIp()
	fmt.Printf("use %v://%v:%v\n", "http", ip, port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Println("start use:" + cast.ToString(kernel.GetUnitTime()))

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	_ = <-quit

	logging.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.Println("Server Shutdown:", err)
	}
	logging.Println("Server exiting")
}
