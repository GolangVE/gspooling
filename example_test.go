package gspooling_test

import (
	"fmt"
	"time"

	"github.com/leoxnidas/gspooling"
)

func Example() {
	q, _ := gspooling.NewQueue(10)
	p := make(chan bool, 1)
	g := make(chan bool, 1)

	go func() {
		for i := 0; i <= 100; i++ {
			time.Sleep(10 * time.Millisecond)

			if q.IsClosed() {
				break
			}

			err := q.Put(i)
			if err != nil {
				fmt.Println("Error ocurred when putting data to queue: ", err)
				i--
			} else {
				fmt.Println("Data sended: ", i)
			}
		}

		p <- true
	}()

	go func() {
		for {
			time.Sleep(30 * time.Millisecond)

			if q.IsClosed() {
				break
			}

			data, err := q.Get()
			if err != nil {
				fmt.Println("Error ocurred when getting data from queue: ", err)
			} else {
				fmt.Println("Data received from queue: ", data)
			}

			if data == 100 {
				break
			}
		}

		g <- true
	}()

	<-p
	<-g
	fmt.Println("Done!!")

	q.Close()
}
