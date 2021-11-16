package go_pool

type Pool struct {
	size    int
	channel chan func()
}

func (p *Pool) Exec(fun func()) {
	p.channel <- fun
}

func New(size int) *Pool {
	p := *&Pool{size: size, channel: make(chan func(), 999999)}
	for i := 0; i < size; i++ {
		go func() {
			for {
				(<-p.channel)()
			}
		}()
	}
	return &p
}
