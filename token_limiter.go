package rate_limiter

import (
	"time"
)

type TokenLimiter struct {
	maxCount int32
	tokenBucket chan bool
	interval float32
}

func (l *TokenLimiter) Run() {
	go func() {
		ticker := time.NewTicker(time.Duration(l.interval) * time.Second)
		for {
			<-ticker.C
			l.tokenBucket <- true
		}
	}()
}

func (l *TokenLimiter) GetToken() (bool) {
	<- l.tokenBucket
	return true
}
