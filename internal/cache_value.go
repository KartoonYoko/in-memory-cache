package internal

import (
	"fmt"
	"time"
)

type cacheValue struct {
	value      any
	expiryDate time.Time
}

// вышло ли время существования значения
func (cv *cacheValue) isTimeExpired() bool {
	return time.Now().After(cv.expiryDate)
}

func (cv *cacheValue) String() string {
	return fmt.Sprintf("cache value: %s, time to expire: %s", cv.value, cv.expiryDate)
}