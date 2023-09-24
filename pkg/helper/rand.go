package helper

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	charset = "abcdefhjkmnpqrst23456789"

	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var l = len(charset)
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

//var sha1Hasher = sha1.New()

func nextRand(seed, base int) int {
	a := 9
	b := 7

	seed = (a*seed + b) % base
	if seed < 0 {
		seed += base
	}

	return seed
}

func RandStr(length int) string {
	b := make([]byte, length)
	r := seededRand.Int()

	for i := range b {
		r = nextRand(r, 100000000)
		b[i] = charset[r%l]
	}

	return string(b)
}

// GetRandomIntString 生成随机数字字符串
func GetRandomIntString(length int) string {

	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成随机验证码
func GetRandString(n int, str ...string) string {
	var meta string
	b := make([]byte, n)
	if len(str) != 0 {
		meta = str[0]
	} else {
		meta = letterBytes
	}
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(meta) {
			b[i] = meta[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// 产生n为数字的随机字符串
func RandNum(n int) string {
	max := int32(math.Pow10(n))
	num := seededRand.Int31n(max)
	format := fmt.Sprintf("%%0%dd", n)
	return fmt.Sprintf(format, num)
}

func GetRowId(length int, prefix string) string {
	date := prefix + time.Now().Format("20060102150405")
	if length < len(date) {
		return date
	}
	sLen := length - len(date)

	return date + GetRandString(sLen)
}

func OrderNo(prefix string) string {
	t := time.Now()
	//return prefix + t.Format("060102150405") + RandStr(3)
	r := prefix + t.Format("060102150405.000") + RandStr(4)
	return strings.Replace(r, ".", "", 1)
	//return r
}
