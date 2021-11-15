package goconcurrent

type LimitGoroutine struct {
	n int
	c chan struct{}
}

// initialization Glimit struct
func New(n int) *LimitGoroutine {
	return &LimitGoroutine{
		n: n,
		c: make(chan struct{}, n),
	}
}

// Run f in a new goroutine but with limit.
func (g *LimitGoroutine) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}
