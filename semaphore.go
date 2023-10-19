package semaphore

type Semaphore struct {
	c chan struct{}
}

func New(size int) Semaphore {
	c := make(chan struct{}, size)
	return Semaphore{c}
}

func (s *Semaphore) Acquire() {
	s.c <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.c
}

func (s *Semaphore) Len() int {
	return len(s.c)
}

func (s *Semaphore) Close() {
	if s.c != nil {
		close(s.c)
	}
}
