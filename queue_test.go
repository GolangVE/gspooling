package gspooling

import (
	"testing"
	"time"
)

func TestWrongSizeQueue(t *testing.T) {
	_, err := NewQueue(-1)
	if err != nil {
		t.Log(err)
	} else {
		t.Error("wrong size passed!!")
	}
}

func TestNilDataQueue(t *testing.T) {
	sq, _ := NewQueue(1)
	err := sq.Put(nil)
	if err != nil {
		t.Log(err)
	} else {
		t.Error("internal error, nil data cannot be in the queue.")
	}
}

func TestFullQueue(t *testing.T) {
	sq, _ := NewQueue(10)
	i := 1
	for {
		if !sq.IsClosed() {
			t.Log("queue is not closed.!")
			err := sq.Put(i)
			if err != nil {
				t.Log(err)
				i--
				sq.Close()
				break
			} else {
				t.Log("data sended: ", i)
				i++
			}
		}

	}

	if i > 10 {
		t.Error("Internal Queue error. ", i)
	}
}

func TestEmptyQueue(t *testing.T) {
	sq, _ := NewQueue(10)
	_, err := sq.Get()
	if err != nil {
		t.Log(err)
	} else {
		t.Error("queue most be empty..")
	}
}

func TestClosedQueue(t *testing.T) {
	sq, _ := NewQueue(10)
	sq.Close()

	err := sq.Put(1)
	if err != nil {
		if err == QueueAlreadyClosedErr {
			t.Log(err)
		} else {
			t.Error("QueueAlreadyClosedErr should be retorned, internal error.")
		}
	} else {
		t.Error("internal error, queue most be empty")
	}
}

func TestPutGetQueue(t *testing.T) {
	sq, _ := NewQueue(10)
	c1 := make(chan bool, 1)
	c2 := make(chan bool, 2)

	go func() {
		i := 1
		for {
			if !sq.IsClosed() {
				err := sq.Put(i)
				if err != nil {
					t.Log(err)
				} else {
					t.Log("data sended: ", i)
					i++
				}
			} else {
				t.Log("queue closed")
				break
			}
		}

		c1 <- true
	}()

	go func() {
		i := 1
		for {
			time.Sleep(4 * time.Millisecond)
			if !sq.IsClosed() {
				d, err := sq.Get()
				if err != nil {
					t.Log(err)
				} else {
					if d != nil {
						t.Log("Data received: ", d)
						i++
					} else {
						t.Error("Data received cannot be null")
						sq.Close()
						break
					}
				}
			} else {
				t.Log("queue closed.")
				break
			}

			if i > 10 {
				sq.Close()
				break
			}
		}

		c2 <- true
	}()

	<-c1
	<-c2
}
