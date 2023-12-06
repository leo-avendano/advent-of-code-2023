package common

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func IsNumber(b byte) bool {
	return (b >= 48 && b <= 57)
}
