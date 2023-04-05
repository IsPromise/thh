package filequeue

import (
	"fmt"
	"io"
	"testing"
	"thh/arms"
	"thh/arms/app"

	"github.com/spf13/cast"
)

type Queue interface {
	Push(data string) error
	Pop() (string, error)
	Len() int64
	Clean() error
}

func TestData(t *testing.T) {
	data, _ := arms.FileGetContents("/Users/thh/workspace/about/tmp/b.b")
	fmt.Println(len(data))
	fmt.Println(data[0:])
}

func TestData2(t *testing.T) {
	_, err := OpenOrCreateFile("./storage/queue2/1_000_000_0020.q")
	if err != nil {
		fmt.Println(err)
	}
}

func TestCheckQueueData(t *testing.T) {
	// 64 + 128+ 128 = 256 + 64 = 320
	// 280 - 128 -128 =  24
	data, _ := arms.FileGetContents("./storage/queue/1_000_000_000.q")
	fmt.Println(len(data))
	fmt.Println(data[0:63])
	for i := 0; i < (len(data)-64)/128+1; i++ {
		start := 64 + i*128
		end := start + 128
		fmt.Println(data[start:end])
	}
}

type TestUnitData struct {
	Valid bool   `json:"valid"`
	Data  string `json:"data"`
}

func TestFqm(t *testing.T) {
	app.InitStart()
	var q Queue
	q, err := FqmStd("./storage/queue")

	if err != nil {
		t.Error(err)
	}

	maxTest := 1_000_000
	stopNum := 100_000

	for i := 1; i <= maxTest; i++ {
		err = q.Push(arms.JsonEncode(TestUnitData{true, cast.ToString(i) + "加个汉字"}))
		if err != nil {
			t.Error(err)
		}
		if i%stopNum == 0 {
			t.Log(app.GetRunTime())
		}
	}

	n := 0
	for {
		data, popErr := q.Pop()
		if popErr != nil {
			t.Log(err)
			break
		}
		n += 1
		if n%10 == 0 {
			t.Log(data)
			t.Log(app.GetRunTime())
			break
		}
	}
	t.Log("清理数据")

	err = q.Clean()
	if err != nil {
		t.Error(err)
	}
	arms.Together(func(goId int) {
		for {
			data, popErr := q.Pop()
			if popErr != nil {
				if popErr == io.EOF {
					break
				} else {
					t.Error(popErr)
					break
				}
			}
			n += 1
			if n%stopNum == 0 {
				t.Log(`n%`+cast.ToString(stopNum), data)
				t.Log("goId", app.GetRunTime())
			}
		}
	}, 3)

	t.Log("end:", app.GetRunTime())
	err = q.Clean()
	if err != nil {
		t.Error(err)
	}
}

func TestArrSet(t *testing.T) {
	block := make([]byte, 64)
	data := []byte(`汉字`)
	ReplaceData(block, data, 3)
	fmt.Println(block)
}
