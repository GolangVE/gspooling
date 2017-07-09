package gspooling

type signalHandler struct {
	sg chan *signal
}

func newSignalHandler() *signalHandler {
	return &signalHandler{sg: make(chan *signal, 1)}
}

func (sh *signalHandler) notifyGet() {
	sh.sg <- &signal{s: notiget}
}

func (sh *signalHandler) notifyPut() {
	sh.sg <- &signal{s: notiput}
}

func (sh *signalHandler) notifyClose() {
	sh.sg <- &signal{s: noticlose}
}

func (sh *signalHandler) handleNotification() *signal {
	select {
	case signal := <-sh.sg:
		return signal
	default:
		return nil
	}
}
