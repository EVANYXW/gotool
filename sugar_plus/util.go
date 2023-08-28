package sugar_plus

import (
	"math"
	"math/rand"
	"time"
)

// GetRandomString 生成随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandInt 生成随机字符串
func GetRandInt(top int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rs := r.Intn(top)
	t := time.Now().UnixNano()
	if t%2 == 0 {
		return rs
	} else {
		return -rs
	}
}

func BKDRHash(str string) uint32 {
	seed := uint32(131) // 31 131 1313 13131 131313 etc..
	hash := uint32(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint32(str[i])
	}
	return hash & 0x7FFFFFFF
}

// Round 四舍五入，ROUND_HALF_UP 模式实现
// 返回将 val 根据指定精度 precision（十进制小数点后数字的数目）进行四舍五入的结果。precision 也可以是负数或零。
func Round(val float64, precision int) float64 {
	if precision == 0 {
		return math.Round(val)
	}

	p := math.Pow10(precision)
	if precision < 0 {
		return math.Floor(val*p+0.5) * math.Pow10(-precision)
	}

	return math.Floor(val*p+0.5) / p
}



func Try(userFn func(), catchFn func(err interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catchFn(err)
		}
	}()
}

