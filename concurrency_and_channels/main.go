package main

import (
	"fmt"
	"time"
)

// func f1(s string) {
// 	fmt.Printf("In goroutine printing %s\n", s)
// }

// func compressit(fileName string) {
// 	fmt.Println("Compressing", fileName)
// }

// the main function in itself is a goroutine,
// that get started when the program starts up
func main() {
	//// 6. Select ////

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1000 * time.Millisecond)
		c1 <- "first"
	}()
	go func() {
		time.Sleep(2000 * time.Millisecond)
		c2 <- "second"
	}()

	// The syntax is interesting, the case is executed when the channel
	// has something to provide, and we assign it to msg1 or msg2 within the case.

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("got the", msg1)
		case msg2 := <-c2:
			fmt.Println("got the", msg2)
		}
	}

	// c1 := make(chan string)

	// go func() {
	// 	defer close(c1)
	// 	time.Sleep(1000 * time.Millisecond)
	// 	c1 <- "first"
	// }()

	// select {
	// case msg1 := <-c1:
	// 	fmt.Println("Received", msg1)
	// }

	//// 5. Channels ////

	// buffered channel
	// limit := 5
	// intStream := make(chan rune, limit)
	// go func() {
	// 	defer close(intStream)
	// 	defer fmt.Println("Rune sending Done.")
	// 	for i := 0; i < limit; i++ {
	// 		r := 'A' + rune(i)
	// 		fmt.Printf("Sending - %c\n", r)
	// 		intStream <- r
	// 	}
	// }()

	// for rune := range intStream {
	// 	fmt.Printf("Received - %c\n", rune)
	// }

	// unbuffered channel
	// strStream := make(chan string) // channel are typed,
	// // so we can only send strings through this channel
	// go func() {
	// 	defer close(strStream)
	// 	for i := 0; i < 26; i++ {
	// 		chr := string(int('A') + i)
	// 		strStream <- string(chr)
	// 	}
	// }()

	// for str := range strStream {
	// 	fmt.Printf("%s ", str)
	// }

	/// 4. Mutex ////

	// var str string
	// var mutex sync.Mutex

	// inc := func() {
	// 	mutex.Lock()
	// 	defer mutex.Unlock()
	// 	str += "A"
	// 	fmt.Println("Appending:", str)
	// }

	// dec := func() {
	// 	mutex.Lock()
	// 	defer mutex.Unlock()
	// 	if len(str) > 0 {
	// 		str = str[:len(str)-1]
	// 		fmt.Println("Removing Last:", str)
	// 	} else {
	// 		fmt.Println("Not modifying empty string")
	// 	}
	// }

	// upperLimit := 5
	// var wg sync.WaitGroup
	// wg.Add(upperLimit)
	// for i := 0; i < upperLimit; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		inc()
	// 	}()
	// }

	// wg.Add(upperLimit)
	// for i := 0; i < upperLimit; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		dec()
	// 	}()
	// }

	// wg.Wait()
	// fmt.Println("Function main complete.")

	/////* 3. compressit */////

	// var wg sync.WaitGroup

	// var i int = -1
	// var file string
	// for i, file = range os.Args[1:] {
	// 	wg.Add(1)
	// 	go func(fileName string) {
	// 		defer wg.Done()
	// 		compressit(fileName)
	// 	}(file)

	// }
	// wg.Wait()
	// fmt.Printf("Compressed %d files\n", i+1)

	/////* 2. WaitGroup */////
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("1st goroutine executing")
	// 	time.Sleep(1000 * time.Millisecond)
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("2nd goroutine executing")
	// 	time.Sleep(2000 * time.Millisecond)
	// }()

	// wg.Wait()
	// fmt.Println("All goroutines finished executing")

	/////* 1. Simple goroutine */////

	// cores := runtime.NumCPU()
	// println("Number of cores: ", cores)

	// var wg sync.WaitGroup
	// for _, s := range []string{"1", "2", "3"} {
	// 	wg.Add(1)
	// 	go func(s string) {
	// 		defer wg.Done()
	// 		fmt.Println(s)
	// 	}(s)
	// 	wg.Wait()
	// }

	// if we don't sleep, the main goroutine will exit
	// time.Sleep(time.Microsecond * 1000)
}
