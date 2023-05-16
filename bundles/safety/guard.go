package safety

func Guard(action func(), errHandle func(any)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				errHandle(err)
			}
		}()
		action()
	}()
}
