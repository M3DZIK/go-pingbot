package backend

func StopAfterPing() int {
	cache(0)
	num := ping()

	for {
		if int64(num) == Status.Success+Status.Error {
			break
		}
	}

	return num
}
