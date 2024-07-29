package iterations

func Repeat(ch string, iter int) string {
	var repeated string
	for i := 0; i < iter; i++ {
		repeated += ch
	}
	return repeated
}
