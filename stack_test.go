package gspooling

import "testing"

func TestWrongBufferSize(t *testing.T) {
	bff, err := newStackBuffer(-1)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(bff.put(1))
	t.Log("Buffer size: ", bff.csize)
}

func TestFullBuffer(t *testing.T) {
	bff, _ := newStackBuffer(1)
	err := bff.put(1)
	err = bff.put(2)
	if err != nil {
		t.Log(err)
	} else {
		t.Error("Buffer should be full, something wrong happend.")
	}
}

func TestPutDataIntoBuffer(t *testing.T) {
	size := 10
	bff, _ := newStackBuffer(size)
	for i := 0; i < size; i++ {
		err := bff.put(i + 1)
		if err != nil {
			t.Error(err)
			return
		} else {
			t.Log(bff.stBff)
		}
	}

	if bff.msize == bff.csize {
		t.Log("Stack is full of data")
	}

	t.Log("Stack size: ", bff.csize)
}

func TestBufferEmpty(t *testing.T) {
	bff, _ := newStackBuffer(10)
	_, err := bff.get()
	if err != nil {
		t.Log(err)
	} else {
		t.Error("Stack is empty, but error does not came out.")
	}
}

func TestPutThenGetDataBuffer(t *testing.T) {
	size := 10
	bff, _ := newStackBuffer(size)
	for i := 0; i < size; i++ {
		err := bff.put(i + 1)
		if err != nil {
			t.Error(err)
			return
		} else {
			t.Log(bff.stBff)
		}
	}

	for j := 0; j < size; j++ {
		d, err := bff.get()
		if err != nil {
			t.Error(err)
			return
		} else {
			t.Log("Data from the stack: ", d)
			t.Log("stack: ", bff.stBff)
			t.Log("stack size: ", bff.csize)
		}
	}
}

func TestFixBufferFunc(t *testing.T) {
	rawbff := make([]interface{}, 10)
	rawbff[9] = 1
	rawbff[4] = 2
	rawbff[1] = 3
	fixStackBuffer(rawbff)

	if rawbff[0] == 3 && rawbff[1] == 2 && rawbff[2] == 1 {
		t.Log("fixed raw stack", rawbff)
	} else {
		t.Error("fixStack func error")
	}
}
