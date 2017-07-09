<p align="center">
  <h3 align="center">gspooling</h3>
  <p align="center">Asynchronous, thread-safe, fixed-size, buffered and easy to use fifo queue.</p>
</p>

---

The purpose of this project is to offer yet another fully asynchronous thread-safe fixed size queue. This experimental library borned after i was reading an Operating System book, I'm trying to implement [spooling algorithm](https://www.wikiwand.com/en/Spooling).

If you wan to check changelogs, please go to [CHANGELOGS file](./CHANGELOGS).

If you want to contribute or give me any advice, be welcome =).

## Quick Start

**Download and install**

```go get github.com/leoxnidas/gspooling```

**Implementation**

This is a simple example, where we'll put 2 values into the queue and then
we'll get those values from the queue. Easy!

```go
package main

import (
	"fmt"
    "github.com/leoxnidas/gspooling"
)

func main() {
	queue := gspooling.NewQueue(2)
    queue.Put(10)
    queue.Put(2)
    
    v1, _ := queue.Get()
    v2, _ := queue.Get()
    
    fmt.Println("data 1: ", v1, " data 2: ", v2)
    
    queue.Close()
}
```

## License

gspooling library is license under [MIT license](./LICENSE).