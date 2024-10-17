package contest

func New() Mutex {
	mutex := MyMutex{make(chan struct{}, 1)}
	mutex.Unlock()
	return &mutex
}

type MyMutex struct {
	mutexChan chan struct{}
}

func (m *MyMutex) Lock() {
	<-m.mutexChan
}

func (m *MyMutex) LockChannel() <-chan struct{} {
	return m.mutexChan
}

func (m *MyMutex) Unlock() {
	select {
	case m.mutexChan <- struct{}{}:
	default:
		panic("multiple unlock")
	}
}
