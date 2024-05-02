package threadMeta

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"sync"
	"testing"
)

func TestSetMeta(t *testing.T) {
	printData := func() {
		_, data := GetMeta()
		fmt.Println(getg().goid, data)
	}

	SetMeta(nil, "test_data")
	printData()
	exitChan := make(chan struct{})

	go func() {
		printData()
		SetMeta(nil, "test_data2")
		printData()
		exitChan <- struct{}{}
	}()
	<-exitChan
	printData()
}

func Test_GetMeta(t *testing.T) {
	const tCnt = 100000
	memStats := new(runtime.MemStats)

	var wg sync.WaitGroup
	wg.Add(tCnt)
	for i := 0; i < tCnt; i++ {
		go func(d int) {
			defer wg.Done()
			want := d

			SetMeta(nil, want)
			_, got := GetMeta()
			assert.Equal(t, want, got, "want", want, "got", got)

			//AddLabel("uid", want)
			//uid := GetLabel("uid")
			//assert.Equal(t, want, uid, "want", want, "got", uid)
			//
			//_, got = GetMeta()
			//assert.Equal(t, want, got, "want", want, "got", got)
		}(i)
	}
	wg.Wait()

	runtime.ReadMemStats(memStats)
	println("Mallocs:", memStats.Mallocs, "HeapObjects:", memStats.HeapObjects)

	runtime.GC()
	runtime.ReadMemStats(memStats)
	println("Frees:", memStats.Frees, "HeapObjects:", memStats.HeapObjects)

	//Mallocs: 627573 HeapObjects: 152111
	//Frees: 579544 HeapObjects: 48029

}
