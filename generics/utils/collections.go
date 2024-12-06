package utils

func True(val bool) bool {
	return val == true
}

func False(val bool) bool {
	return val == false
}

func All[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

func One[T any](slice []T, predicate func(T) bool) bool {
	count := 0
	for _, v := range slice {
		if predicate(v) {
			count++
		}
	}
	return count == 1
}

func AtLeast[T any](slice []T, n int, predicate func(T) bool) bool {
	count := 0
	for _, v := range slice {
		if predicate(v) {
			count++
			if count >= n {
				return true
			}
		}
	}
	return false
}

func N[T any](slice []T, n int, predicate func(T) bool) bool {
	count := 0
	for _, v := range slice {
		if predicate(v) {
			count++
		}
	}
	return count == n
}

func Map[T any](slice []T, predicate func(T) (T, error)) ([]T, error) {
	for i, v := range slice {
		if t, err := predicate(v); err != nil {
			return nil, err
		} else {
			slice[i] = t
		}
	}
	return slice, nil
}

func Transform[T any, O any](slice []T, transform func(T) (O, error)) ([]O, error) {
	ret := make([]O, len(slice))
	for i, v := range slice {
		if o, err := transform(v); err != nil {
			return nil, err
		} else {
			ret[i] = o
		}
	}
	return ret, nil
}

func Compare[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
