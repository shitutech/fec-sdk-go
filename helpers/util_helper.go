package helpers

import (
	"math/rand"
	"time"
)

// GetRandomString 随机字符串
func GetRandomString(size int) string {
	kinds, result := [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		iKind := rand.Intn(3)
		scope, base := kinds[iKind][0], kinds[iKind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}

	return string(result)
}
