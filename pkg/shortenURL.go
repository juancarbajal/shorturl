package pkg

import "math/rand"

const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func ShortenURL(url string) string {
	var result []byte
	num := rand.Int63()
	for num > 0 {
		result = append(result, base62[num%62])
		num /= 62
	}
	return string(result)
}
