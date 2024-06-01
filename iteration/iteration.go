package iteration

func Repeat(c string, num int) string {
	str := make([]byte, 0, num)
	for i := 0; i < num; i++ {
		str = append(str, []byte(c)[0])
	}
	return string(str)
}
