package rate_limiter

type Limiter interface {
	Run()		// 存入令牌
	GetToken() (bool)	// 获取令牌
}


