package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1) // Tambahkan counter waitgroup

	cond.L.Lock() // Mengunci sebelum menunggu kondisi
	cond.Wait()   // Menunggu sinyal kondisi
	fmt.Println("Done", value)
	cond.L.Unlock() // Membuka kunci setelah kondisi terpenuhi
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second) // Simulasi proses yang memakan waktu
			cond.Signal()               // Memberi sinyal pada satu goroutine yang menunggu
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast() // Memberi sinyal pada semua goroutine yang menunggu
	}()

	group.Wait()
}
