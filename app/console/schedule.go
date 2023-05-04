package console

import (
	"fmt"
	"thh/arms"
	"thh/arms/logger"
	"time"

	"github.com/robfig/cron/v3"
)

var c = cron.New()

func RunJob() {
	if entryID, err := c.AddFunc("* * * * *", upCmd(func() {
		logger.Info("HEART_IN_RUN_JOB", time.Now().Format("2006-01-02 15:04:05"))
	})); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("schedule entryId:", entryID, " set success")
	}

	c.Run()
}

func upCmd(cmd func()) func() {
	return func() {
		arms.MyTraceInit()
		defer arms.MyTraceClean()
		cmd()
	}
}

// 之后如果需要平滑关闭可以参考如下代码
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
