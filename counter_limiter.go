package rate_limiter

import (
	"sync"
	"time"
)

type CounterLimiter struct {
	interval	float32
	counter		int32
	mutex 		sync.Mutex
	maxCount	int32
}

// 获取token
func (l *CounterLimiter) GetToken() (bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.counter < l.maxCount {
		return true
	} else {
		return false
	}
}

// 运行limiter
func (l *CounterLimiter) Run() {
	go func() {
		ticker := time.NewTicker(time.Duration(l.interval)*time.Second)
		for {
			<- ticker.C
			l.mutex.Lock()
			l.counter = 0
			l.mutex.Unlock()
		}
	}()
}
