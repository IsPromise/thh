package console

import (
	"fmt"
	"github.com/spf13/cobra"
	"thh/app/service/cqclient"
	"thh/app/service/forbiden"
	"thh/bundles/logging"
	"time"

	"github.com/robfig/cron/v3"
)

var scheduleAction = &cobra.Command{
	Use:   "schedule",
	Short: "Start web server",
	Run: func(_ *cobra.Command, _ []string) {
		RunJob()
	},
	Args: cobra.NoArgs,
}

var c = cron.New()

func RunJob() {
	_, _ = c.AddFunc("* * * * *", upCmd(func() {
		logging.Info("HEART_IN_RUN_JOB", time.Now().Format("2006-01-02 15:04:05"))
	}))
	c.AddFunc("* 7 * * *", func() {
		if forbiden.Forbidden(time.Now().Format("20060102") + "am") {
			return
		}
		currentTime := time.Now()

		// 目标日期
		targetDate := time.Date(2023, 12, 22, 0, 0, 0, 0, time.UTC)

		// 计算天数差
		days := targetDate.Sub(currentTime).Hours() / 24
		cqclient.Send4group(622611442, fmt.Sprintf("同志们~冲起来，还有%v天了", days))

	})

	c.AddFunc("* 0 * * *", func() {
		if forbiden.Forbidden(time.Now().Format("20060102") + "pm") {
			return
		}
		currentTime := time.Now()

		// 目标日期
		targetDate := time.Date(2023, 12, 22, 0, 0, 0, 0, time.UTC)

		// 计算天数差
		days := targetDate.Sub(currentTime).Hours() / 24
		cqclient.Send4group(622611442, fmt.Sprintf("同志们~早点休息，还有%v天了 冲冲冲", days))

	})

	c.Run()
}

func StopJob() {
	c.Stop()
}

func upCmd(cmd func()) func() {
	return func() {
		cmd()
	}
}

////之后如果需要平滑关闭可以参考如下代码
//var needStop = false
//var needStopL = &sync.Mutex{}
//
//func demo(){
//	ctx:=c.Stop()
//	Stop()
//	<-ctx.Done()
//}
//
//func Stop() {
//	c.Run()
//	needStopL.Lock()
//	defer needStopL.Unlock()
//	needStop = true
//}
//
//func getStop() bool {
//	needStopL.Lock()
//	defer needStopL.Unlock()
//	return needStop
//}
