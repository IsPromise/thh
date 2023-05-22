package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"runtime"
	"strconv"
	"sync"
)

type Trace struct {
	traceId string
	number  int
}

func (itself *Trace) GetNextTrace() string {
	itself.number += 1
	return fmt.Sprintf("%v:%v", itself.traceId, itself.number)
}

func GetTrace() *Trace {
	return &Trace{
		traceId: uuid.NewString(),
	}
}

func GoID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

var traceManager sync.Map

func MyTraceClean() {
	id := GoID()
	traceManager.Delete(id)
}

func MyTrace() *Trace {
	id := GoID()
	trace, ok := traceManager.Load(id)
	if !ok {
		trace = GetTrace()
		traceManager.Store(id, trace)
	}
	return trace.(*Trace)
}

func MyTraceInit() *Trace {
	trace := GetTrace()
	id := GoID()
	traceManager.Store(id, trace)
	return trace
}

var (
	goroutinePrefix = []byte("goroutine ")
	errBadStack     = errors.New("invalid runtime.Stack output")
)

// This is terrible, slow, and should never be used.
func goid() (int, error) {
	buf := make([]byte, 32)
	runtime.Gosched()
	n := runtime.Stack(buf, false)
	buf = buf[:n]
	// goroutine 1 [running]: ...

	buf, ok := bytes.CutPrefix(buf, goroutinePrefix)
	if !ok {
		return 0, errBadStack
	}

	i := bytes.IndexByte(buf, ' ')
	if i < 0 {
		return 0, errBadStack
	}

	return strconv.Atoi(string(buf[:i]))
}
