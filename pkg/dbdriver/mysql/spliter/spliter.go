package spliter

import (
	"fmt"
	"grpc-battle/pkg/dbdriver/mysql/algo"
)

// TODO 分离器，主要用于分库分表，其中会结合各种算法

type SetSpliter func(s *Spliter)

func WithSetAlgorithm(creator func() algo.Algo) SetSpliter {
	return func(s *Spliter) {
		s.SpliterAlgorithm = creator()
	}
}

type Spliter struct {
	SpliterAlgorithm algo.Algo
}

func NewSpliter(setter ...SetSpliter) *Spliter {
	s := &Spliter{
		// TODO maybe need default spliter algorithm
	}
	if len(setter) > 0 {
		for _, set := range setter {
			set(s)
		}
	}
	return s

}

func (s *Spliter) DatabaseName(dpref string, raw uint32) string {
	suffix := s.SpliterAlgorithm.CreateSuffixLabel(raw)
	return fmt.Sprintf("%s_%s", dpref, suffix)
}

func (s *Spliter) TableName(tperf string, raw uint32) string {
	suffix := s.SpliterAlgorithm.CreateSuffixLabel(raw)
	return fmt.Sprintf("%s_%s", tperf, suffix)
}

func (s *Spliter) Router() {

}
