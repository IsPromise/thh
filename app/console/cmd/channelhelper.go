package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "channelhelper",
		Short: "",
		Run:   runChannelhelper,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	}
	// cmd.Flags().String("param", "value", "input params")
	appendCommand(cmd)
}

func runChannelhelper(cmd *cobra.Command, args []string) {
	// param, _ := cmd.Flags().GetString("param")
	//fmt.Println("success")
	//
	//mych := newCH[int](1)
	//mych.Set(1)
	//data := mych.Get()
	//fmt.Println(data)
	demo2()
}

func demo2() {
	or, ow := newORW(make(chan string, 10))
	ow.Write("nihaonihao")
	fmt.Println(or.Read())
}

type Ch[T any] struct {
	myCh chan T
}

func newCH[T any](lenght int) Ch[T] {
	return Ch[T]{make(chan T, lenght)}
}

func (itself *Ch[T]) Set(data T) {
	itself.myCh <- data
}

func (itself *Ch[T]) Get() T {
	return <-itself.myCh
}

type OnlyRead[chanT any, T chan chanT] struct {
	myCh T
}

func (itself *OnlyRead[T, chanT]) Read() T {
	return <-itself.myCh
}

type OnlyWrite[chanT any, T chan chanT] struct {
	myCh T
}

func (itself *OnlyWrite[T, chanT]) Write(data T) {
	itself.myCh <- data
}

func newORW[chanT any, T chan chanT](ch T) (OnlyRead[chanT, T], OnlyWrite[chanT, T]) {
	return OnlyRead[chanT, T]{ch}, OnlyWrite[chanT, T]{ch}
}
