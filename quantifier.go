package quango

func ForAll(list []int, pred func(int) bool) bool {
	val := true
	for i := 0; i < len(list); i++ {
		val = val && pred(list[i])
	}
	return val
}
