package iteration

func Repeat(character string, times int) string {
	repeated := ""
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
