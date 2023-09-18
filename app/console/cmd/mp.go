package cmd

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	cmd := &cobra.Command{
		Use:   "mp",
		Short: "",
		Run:   runMp,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	}
	// cmd.Flags().String("param", "value", "input params")
	appendCommand(cmd)
}

func runMp(cmd *cobra.Command, args []string) {
	// param, _ := cmd.Flags().GetString("param")

	currentTime := time.Now()
	twoAM := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 14, 0, 0, 0, currentTime.Location())
	twoThirtyAM := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 14, 30, 0, 0, currentTime.Location())

	if currentTime.After(twoAM) && currentTime.Before(twoThirtyAM) {
		for {
			if currentTime.After(twoThirtyAM) {
				break
			}

			if currentTime.Before(twoAM.Add(20 * time.Minute)) {
				// 当前时间在2:00到2:20之间，每5分钟输出一次 "hello"
				thhSay("2点了 半小时内要抢票了")
				time.Sleep(3 * time.Minute)
			} else {
				// 当前时间在2:20到2:30之间，输出剩余时间到2:30
				remainingTime := twoThirtyAM.Sub(currentTime)
				thhSay(fmt.Sprintf("剩余时间：%v\n", remainingTime))
				time.Sleep(1 * time.Minute)
			}

			currentTime = time.Now()
		}
	} else {
		fmt.Println("当前时间不在指定范围内")
	}

	thhSay("结束了")
}

func thhSay(msg string) {
	err := beeep.Notify("thh", msg, "")
	if err != nil {
		fmt.Println(err)
	}
}
