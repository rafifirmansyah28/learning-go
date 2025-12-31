package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second) // membuat ticker dengan interval 1 detik

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop() // menghentikan ticker setelah 5 detik
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second) // membuat ticker dengan interval 1 detik

	for time := range channel {
		fmt.Println(time)
	}
}
