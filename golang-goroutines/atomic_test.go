package golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			/**
			Jika terjadi error: panic: sync: WaitGroup is reused before previous Wait has returned
			Itu artinya, go routine belum selesai menjalankan kode group.Add(1), namun go goroutine sudah menjalankan group.Wait(),
			group tidak boleh di add ketika sudah di Wait(), hal ini biasanya terjadi jika resource hardware kurang cepat ketika menjalankan
			go routine di awal. Jika hal ini terjadi, silahkan pindahkan kode group.Add(1), ke baris 15 sebelum memanggil go func().
			*/
			// group.Add(1) // Jadi pindah ke baris 15
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1) // fungsi atomic menambahkan nilai x secara aman (dari race condition) di dalam go routine
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter =", x)
}
