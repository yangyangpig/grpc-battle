package synchro

import "sync"

/**
目标：实现多个goroutine读和多个goroutine写的同步，读goroutine要等待写goroutine通知才运行
*/

type SynchroConfOpt func(s *SynchroCond)
type SynchroCond struct {
	mu     *sync.Mutex
	ga     func(p interface{}) // 先执行完ga
	gb     func(p interface{}) // 再执行gb
	mucond *sync.Cond
}

func NewSynchroCond(opts ...SynchroConfOpt) *SynchroCond {
	s := &SynchroCond{}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *SynchroCond) ExecuteGa(a interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	go func(a interface{}) {
		s.ga(a)
		s.mucond.Broadcast()
	}(a)
}

func (s *SynchroCond) ExecuteGb(b interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	go func(b interface{}) {
		s.mucond.Wait()
		s.gb(b)

	}(b)
}



