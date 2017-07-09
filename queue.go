// gspooling is an asynchronous, thread-safe
// fixed-size, buffered and easy to use fifo queue.
//
// 	func main() {
// 		queue := gspooling.NewQueue(2)
// 		queue.Put(10)
// 		queue.Put(2)
//
// 		v1, _ := queue.Get()
// 		v2, _ := queue.Get()
//
// 		fmt.Println("data 1: ", v1, " data 2: ", v2)
// 	}
//
package gspooling

import "errors"

var (
	// error throwed if the queue is closed.
	QueueAlreadyClosedErr = errors.New("Queue already closed.")
	NilDataErr            = errors.New("Data cannot be nil.")
)

type Queue struct {
	closed bool

	buffer *stackBuffer
	sh     *signalHandler

	eput chan error
	eget chan error

	put chan interface{}
	get chan interface{}
}

// Returns a new fixed size Queue
// otherwise returns an error.
func NewQueue(size int) (*Queue, error) {
	bff, err := newStackBuffer(size)
	if err != nil {
		return nil, err
	}

	// error channels
	eput := make(chan error, 1)
	eget := make(chan error, 1)

	// input and output channel
	put := make(chan interface{}, 1)
	get := make(chan interface{}, 1)

	queue := &Queue{
		put:    put,
		get:    get,
		eput:   eput,
		eget:   eget,
		buffer: bff,
		closed: false,
		sh:     newSignalHandler(),
	}

	queue.parallelizeStackBuffer()
	return queue, nil
}

// Put the buffer in a go thread
func (sq *Queue) parallelizeStackBuffer() {
	go func() {
		for {
			sg := sq.sh.handleNotification()
			if sg != nil {
				if sg.isPut() {
					err := sq.buffer.put(<-sq.put)
					sq.eput <- err
				} else if sg.isGet() {
					d, err := sq.buffer.get()
					sq.eget <- err
					sq.get <- d
				} else if sg.isClose() {
					break
				}
			}
		}
	}()
}

// Put data into the buffered queue.
// Return an error if buffered queue is full.
//
// Posible errors:
// 	- "queue already closed." -> queue was closed already.
// 	- "Buffer is full." -> queue is full of data.
//
func (sq *Queue) Put(data interface{}) error {
	if sq.IsClosed() {
		return QueueAlreadyClosedErr
	}

	if data == nil {
		return NilDataErr
	}

	sq.sh.notifyPut()
	sq.put <- data
	return <-sq.eput
}

// Get data from the buffered queue.
// Return an error if buffered queue is empty.
//
// Posible errors:
// 	- "queue already closed." -> queue was closed already.
// 	- "Buffer is empty." -> queue is empty
func (sq *Queue) Get() (interface{}, error) {
	if sq.IsClosed() {
		return nil, QueueAlreadyClosedErr
	}

	sq.sh.notifyGet()
	return <-sq.get, <-sq.eget
}

// It will try to close the Queue.
func (sq *Queue) Close() error {
	if !sq.IsClosed() {

		// notify the buffer
		// that need to be closed.
		sq.sh.notifyClose()
		close(sq.eget)
		close(sq.eput)
		close(sq.get)
		close(sq.put)
		close(sq.sh.sg)

		sq.closed = true

		return nil
	}

	return QueueAlreadyClosedErr
}

//verify if queue is closed.
func (sq *Queue) IsClosed() bool {
	return sq.closed
}
