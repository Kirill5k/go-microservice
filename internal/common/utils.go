package common

func If[T any](cond bool, ifTrue, ifFalse T) T {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func Map[T, U any](ts []T, f func(*T) U) []U {
	us := make([]U, len(ts))
	for i, item := range ts {
		us[i] = f(&item)
	}
	return us
}
