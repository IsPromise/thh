package cmd

import (
	"fmt"
	"thh/bundles/app"

	"github.com/leancodebox/goose/filequeue"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{Use: "tool:fileQueue", Short: "fileQueue", Run: fileQueue})
}

func fileQueue(_ *cobra.Command, _ []string) {

	f, err := filequeue.NewDefaultFileQueue("./storage/queue")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= 1000000; i++ {
		err = f.Push(cast.ToString(i) + "加个汉字")
		if err != nil {
			fmt.Println(err)
		}
		if i%100000 == 0 {
			fmt.Println(app.GetRunTime())
		}
	}
	n := 0
	for {
		data, popErr := f.Pop()
		if popErr != nil {
			fmt.Println(err)
			break
		}
		n += 1
		if n%10 == 0 {
			fmt.Println(data)
			fmt.Println(app.GetRunTime())
			break
		}
	}
	fmt.Println("清理数据")

	err = f.Clean()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		data, popErr := f.Pop()
		if popErr != nil {
			fmt.Println(err)
			break
		}
		n += 1
		//fmt.Println(data)
		if n%100000 == 0 {
			fmt.Println(`n%100000`, data)
			fmt.Println(app.GetRunTime())
		}
	}
	_ = f.Clean()
}
