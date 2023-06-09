package sugar_plus

func Try(userFn func(), catchFn func(err interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catchFn(err)
		}
	}()
}
