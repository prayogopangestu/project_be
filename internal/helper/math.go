package helper

func TofmInt(nominal float64) (nominalRes int) {
	res := nominal * 10000
	return int(res)
}

func TofmDesimal(nominal int) (nominalRes float64) {
	res := float64(nominal) / 10000
	return res
}
