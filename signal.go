package gspooling

type signal_type uint8

const (
	notiput   signal_type = 1
	notiget   signal_type = 2
	noticlose signal_type = 99
)

type signal struct {
	s signal_type
}

func (sg *signal) isPut() bool {
	return sg.s == notiput
}

func (sg *signal) isGet() bool {
	return sg.s == notiget
}

func (sg *signal) isClose() bool {
	return sg.s == noticlose
}
