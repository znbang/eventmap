package mvc

func PageSize(p, s int32) (page, size int) {
	page = 1
	size = 10

	if p > 0 {
		page = int(p)
	}

	if s > 0 {
		size = int(s)
	}

	return page, size
}
