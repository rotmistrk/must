package must

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Must2[P any, Q any](p P, q Q, err error) (P, Q) {
	if err != nil {
		panic(err)
	}
	return p, q
}
