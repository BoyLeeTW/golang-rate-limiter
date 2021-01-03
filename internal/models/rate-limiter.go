package models

type RateLimiterResponse struct {
	CurrentCount uint64
	IsExceed     bool
}
