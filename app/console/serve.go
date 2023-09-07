package console

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"thh/app/bundles/kernel"
	"thh/app/bundles/logging"
	"thh/app/routes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leancodebox/goose/preferences"
	"github.com/leancodebox/goose/serverinfo"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

// CmdServe represents the available web sub-command.
var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

const (
	ENV      = "env"
	EnvProd  = "production"
	EnvLocal = "local"
)

var (
	withSchedule = preferences.GetBool("app.withSchedule", false)
)

func runWeb(_ *cobra.Command, _ []string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	slog.Info("Thousand-hand:start")
	slog.Info(fmt.Sprintf("Thousand-hand:useMem %d KB", m.Alloc/1024/8))

	if withSchedule {
		go RunJob()
	}
	// 初始化应用程序
	debug4pprof()
	ginServe()
}

func debug4pprof() {
	if preferences.GetBool("app.debug", false) {
		go func() {
			// go tool pprof http://localhost:6060/debug/pprof/profile
			//http://127.0.0.1:7070/debug/pprof/
			err := http.ListenAndServe("0.0.0.0:7070", nil)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
}

func ginServe() {
	var (
		port   = preferences.GetString("app.port", 8080)
		isProd = preferences.Get("app.env") == EnvProd
	)
	var engine *gin.Engine
	if isProd {
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		engine.Use(gin.Recovery())
	} else {
		engine = gin.Default()
	}

	routes.RegisterByGin(engine)

	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	ip, _ := serverinfo.GetLocalIp()
	slog.Info("Thousand-hand:listen " + port)
	slog.Info("http://localhost:" + port)
	slog.Info(fmt.Sprintf("%v://%v:%v", "http", ip, port))

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	slog.Info("start use:" + cast.ToString(kernel.GetUnitTime()))

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logging.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.Info("Server Shutdown:", err)
	}
	logging.Info("Server exiting")
	logging.Shutdown()
}
