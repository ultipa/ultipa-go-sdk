package test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"
	"testing"
	"time"
	ultipa "ultipa-go-sdk/rpc"
)

func TestSlice(t *testing.T) {
	a := []int{2}

	deleteIndex := 0
	a = append(a[:deleteIndex], a[deleteIndex+1:]...)

	log.Fatalln(a)
}

func TestBitIsInclude(t *testing.T) {

	const (
		U uint32 = 0
		A uint32 = 1
		B uint32 = 2
		C uint32 = 4
		D uint32 = 8
	)

	AB := A | B

	log.Println(AB&A, AB&B, AB&C, AB&D)
}

func TestTimeZoneTimestamp(t *testing.T) {
	t1 := "2022-05-19T12:00:00Z"
	t2 := "2022-05-19T12:00:00Z"

	tt1, err := time.Parse(time.RFC3339, t1)

	if err != nil {
		log.Fatalln("E1", err)
	}

	tt2, err := time.Parse("2006-01-02T15:04:05Z07:00", t2)

	if err != nil {
		log.Fatalln("E2", err)
	}

	log.Println(tt1.UTC(), tt2.UTC())
}

type strToTimeParam struct {
	layout   string
	input    string
	hasError bool
}

func TestStrToTime(t *testing.T) {
	testParam := []*strToTimeParam{
		&strToTimeParam{"2006-1-2 15:04", "2020-1-1 7:14", false},
		&strToTimeParam{"2006-1-2 15:04", "2020-01-01 7:14", false},
		&strToTimeParam{"2006-1-2 15:04", "2020-1-18 7:14", false},
		&strToTimeParam{"2006-1-2 15:04", "2020-01-18 7:14", false},
		&strToTimeParam{"2006-1-2 15:04", "2020-10-11 7:14", false},
		&strToTimeParam{"2006-01-02 15:04", "2020-10-11 7:14", false},
		&strToTimeParam{"2006-01-02 15:04", "2020-01-01 7:14", false},
		&strToTimeParam{"2006-01-02 15:04", "2020-1-1 7:14", true}, //error

		&strToTimeParam{"2006/1/2 15:04", "2020/1/1 7:14", false},
		&strToTimeParam{"2006/1/2 15:04", "2020/01/01 7:14", false},
		&strToTimeParam{"2006/1/2 15:04", "2020/10/11 7:14", false},

		&strToTimeParam{"2006/01/02 15:04", "2020/10/11 7:14", false},
		&strToTimeParam{"2006/01/02 15:04", "2020/01/01 7:14", false},
		&strToTimeParam{"2006/01/02 15:04", "2020/1/1 7:14", true}, //error

		&strToTimeParam{"2006-1-2T15:04:05Z07:00", "2022-05-19T12:00:00Z", false},
		&strToTimeParam{"2006-1-2T15:04:05Z07:00", "2022-10-1T12:00:00Z", false},
		&strToTimeParam{"2006-1-2T15:04:05Z07:00", "2022-10-10T12:00:00Z", false},
		&strToTimeParam{"2006/1/2T15:04:05Z07:00", "2022/01/1T12:00:00Z", false},
		&strToTimeParam{"2006/1/2T15:04:05Z07:00", "2022/01/10T12:00:00Z", false},
		&strToTimeParam{"2006/1/2T15:04:05Z07:00", "2022/10/10T12:00:00Z", false},
		&strToTimeParam{"2006-1-2T15:04:05Z07:00", "2020/1/1 7:14", true}, //error
	}

	for i, param := range testParam {
		time, err := time.Parse(param.layout, param.input)
		if err != nil {
			t.Logf(`Failed to parse to time for row %d,%v`, i, err)
			if !param.hasError {
				t.Fail()
			}
		} else {
			t.Log(time)
		}
	}
}

func TestFormatDate(t *testing.T) {
	testParam := [][]string{
		{"2006-1-2 15:04", "2020-1-1 7:14"},
		{"2006-1-2 15:04", "2020-01-01 7:14"},
		{"2006-1-2 15:04", "2020-1-18 7:14"},
		{"2006-1-2 15:04", "2020-01-18 7:14"},
		{"2006-1-2 15:04", "2020-10-11 7:14"},
		{"2006-01-02 15:04", "2020-10-11 7:14"},
		{"2006-01-02 15:04", "2020-01-01 7:14"},
	}

	for i, param := range testParam {
		time, err := time.Parse(param[0], param[1])
		if err != nil {
			t.Errorf("failed to format row %d,%v", i, err)
		} else {
			t.Logf("%s", time.Format(`2006-01-02 15:04`))
		}

	}
}

func exit04() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("退出协程")
				return
			default:
				fmt.Println("监控01")
				time.Sleep(1 * time.Second)

			}

		}

	}()

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("退出程序")

}

func TestWg(t *testing.T) {
	wg := sync.WaitGroup{}

	err := funcName(wg)

	wg.Wait()
	if err != nil {
		log.Println(err.Error())
	}

}

func funcName(wg sync.WaitGroup) (err error) {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i1 int) {
			defer wg.Done()
			if i1%2 == 0 {
				cancel()
				err = errors.New("context.WithCancel(context.Background())")
			}
			log.Println(i1)
		}(i)
		select {
		case <-ctx.Done():
			//fmt.Println("退出 ，停止了。。。")
			return err
		default:
			//time.Sleep(1 * time.Millisecond)
			fmt.Println("运行中。。。")
		}
	}
	return err
}

func TestEnumName(t *testing.T) {
	r := ultipa.ResultType_RESULT_TYPE_ATTR
	fmt.Println(fmt.Sprintf("resultType:%s", r))

	fmt.Println(uint64(math.MaxUint64))
	fmt.Println(uint32(math.MaxUint32))
	fmt.Println(strconv.FormatUint(uint64(math.MaxUint64), 2))
	fmt.Println(strconv.FormatUint(uint64(math.MaxUint32), 2))
}
