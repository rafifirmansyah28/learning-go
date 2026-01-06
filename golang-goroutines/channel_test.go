package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// channel := make(chan string) // create channel

	// channel <- "Hello Channel" // send data to channel

	// data := <-channel // receive data from channel

	// fmt.Println("data", data)

	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello Channel"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println("data", data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello Channel"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println("data", data)

	time.Sleep(5 * time.Second)
}

// tipe of params: chan <- string // only send data to channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello Channel"
}

// tipe of params: <- chan string // only receive data from channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // buffered channel with capacity 3. Buffered channel not block when sending data until capacity is full
	defer close(channel)

	fmt.Println("Capacity channel", cap(channel))

	go func() {
		channel <- "Message 1"
		channel <- "Message 2"
		channel <- "Message 3"
		fmt.Println("All messages sent to buffered channel")
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("Message %d", i)
		}
		close(channel)
	}()

	// read data from channel until channel is closed
	for data := range channel {
		fmt.Println("Received:", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Received from channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Received from channel 2:", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Received from channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Received from channel 2:", data)
			counter++
		default:
			fmt.Println("Waiting for data...")
		}

		if counter == 2 {
			break
		}
	}
}
