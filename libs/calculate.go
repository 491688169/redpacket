package libs

import (
	"math/rand"
	"time"
)

func DoubleAverage2(mount, count int64) int64 {
	if count == 1 {
		return mount
	}
	const MIN = 1
	max := mount - count*MIN
	avarage := max / count
	avarage2 := avarage*2 + MIN
	rand.Seed(time.Now().UnixNano())
	result := rand.Int63n(avarage2)
	return result
}
