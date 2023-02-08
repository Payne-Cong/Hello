package api

func fb() bool {
	return false
}

func TestSwitch() {
	switch fb() {
	case true:
		println("1")
	case false:
		println("0")
	default:
		println("-1")
	}
}
