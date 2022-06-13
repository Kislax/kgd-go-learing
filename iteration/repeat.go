package iteration

func Repeat(charter string, counter int) string {
	var repeated string
	for i := 0; i < counter; i++ {
		repeated = repeated + charter
	}
	return repeated
}
