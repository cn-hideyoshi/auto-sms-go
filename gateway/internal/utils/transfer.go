package utils

type Transfer struct {
	Addr string

	remoteAddr string
}

func (t *Transfer) DialService() {
	t.getRemoteAddr()
}

func (t *Transfer) getRemoteAddr() {
	t.remoteAddr = t.Addr
}
