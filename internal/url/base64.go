package url

const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func EncodeBase62(num int64) string {
	if num == 0 {
		return "a"
	}
	result := ""
	for num > 0 {
		result = string(base62[num%62]) + result
		num /= 62
	}
	return result
}
