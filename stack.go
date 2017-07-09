package gspooling

import (
	"errors"
	"sync"
)

var (
	// wrong stack size error
	WrongStackBufferSizeErr = errors.New("Stack Buffer size most be positive number.")

	// stack erros
	StackBufferFullErr  = errors.New("Stack Buffer is full.")
	StackBufferEmptyErr = errors.New("Stack Buffer is empty.")
)

// fixed-size ordered stack that implements
// fifo(first in, first out)
type stackBuffer struct {
	csize int           // max size
	msize int           // current size
	mx    sync.Mutex    // mutex object
	stBff []interface{} // stack buffer ;)
}

func newStackBuffer(s int) (*stackBuffer, error) {
	if s <= 0 {
		return nil, WrongStackBufferSizeErr
	}

	st := make([]interface{}, s)
	return &stackBuffer{
		msize: s,
		csize: 0,
		stBff: st,
	}, nil
}

func (b *stackBuffer) put(d interface{}) error {
	if b.csize >= b.msize {
		return StackBufferFullErr
	}

	b.mx.Lock()
	b.stBff[b.csize] = d
	b.csize++
	b.mx.Unlock()
	return nil
}

func (b *stackBuffer) get() (interface{}, error) {
	if b.csize == 0 {
		return nil, StackBufferEmptyErr
	}

	var d interface{}
	b.mx.Lock()
	d = b.stBff[0]
	b.stBff[0] = nil
	fixStackBuffer(b.stBff)
	b.csize--
	b.mx.Unlock()
	return d, nil
}

// putting together all non-nil values
// at the beginning of the stack.
func fixStackBuffer(bff []interface{}) {
	for i := 0; i < len(bff); i++ {
		for j := 0; j < len(bff)-1; j++ {
			if bff[j] == nil {
				tmp := bff[j]
				bff[j] = bff[j+1]
				bff[j+1] = tmp
			}
		}
	}
}
