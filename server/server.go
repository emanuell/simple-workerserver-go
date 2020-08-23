package server

type Work struct {
	Operation func(int, int) int
	A, B      int
	Reply     chan int
}

func New() (req chan<- *Work) {
	wc := make(chan *Work)
	go start(wc)
	return wc
}

func start(req <-chan *Work) {
	for w := range req {
		go run(w)
	}
}

func run(w *Work) {
	w.Reply <- w.Operation(w.A, w.B)
}
