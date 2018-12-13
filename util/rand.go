package util

import (
	"math/rand"
	"time"
)

func GenerateRangeNum(min,max int) int{
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min + 1) + min
	return randNum
}
