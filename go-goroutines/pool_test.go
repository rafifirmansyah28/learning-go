package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} { // fungsi yang akan dijalankan ketika pool kosong
			return "New"
		},
	}

	pool.Put("Test 1")
	pool.Put("Test 2")
	pool.Put("Test 3")

	for i := 1; i <= 10; i++ {
		go func() {
			data := pool.Get() // fungsi mengambil data dari pool
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data) // fungsi mengembalikan data ke pool
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Selesai")
}
