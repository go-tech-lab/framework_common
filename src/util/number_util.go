package util

import (
	"math/rand"
	"time"
)

const INT_MIN int = 0
const INT_MAX = 1<<63 - 1

const UINT_MIN uint = 0
const UINT_MAX uint = ^uint(0)

const UINT32_MIN uint32 = uint32(0)
const UINT32_MAX uint32 = 1<<32 - 1

const INT32_MIN int32 = int32(0)
const INT32_MAX int32 = 1<<31 - 1

const UINT64_MIN uint64 = uint64(0)
const UINT64_MAX uint64 = 1<<64 - 1

const INT64_MIN int64 = int64(0)
const INT64_MAX int64 = 1<<63 - 1

func RandomUint64() uint64 {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int63n(INT64_MAX)
	return uint64(randNum)
}

func RandomInt64() int64 {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int63n(INT64_MAX)
	return randNum
}

func RandomInt() int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int()
	return randNum
}

func RandomUint() uint {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int()
	return uint(randNum)
}

func RandomInt32() int32 {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int31n(INT32_MAX)
	return randNum
}

func RandomUint32() uint32 {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int31n(INT32_MAX)
	return uint32(randNum)
}
