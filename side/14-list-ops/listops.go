package listops

// IntList is just an alias for slice of int, created to support methods on the type
type IntList []int

type binFunc func(elem int, prev int) int

// Foldl folds the list from the left
func (l IntList) Foldl(f binFunc, init int) int {
	r := init
	for _, e := range l {
		r = f(r, e)
	}
	return r
}

// Foldr folds the list from the right
func (l IntList) Foldr(f binFunc, init int) int {
	r := init
	len := len(l)
	for i := len - 1; i >= 0; i-- {
		r = f(l[i], r)
	}
	return r
}

// Length returns the length of the list
func (l IntList) Length() int {
	return len(l)
}

type predFunc func(elem int) bool

// Filter returns the elements of l for which f is true
func (l IntList) Filter(f predFunc) IntList {
	out := IntList{} // Test expects non-nil slice for empty output
	for _, e := range l {
		if f(e) == true {
			out = append(out, e)
		}
	}
	return out
}

type unaryFunc func(int) int

// Map returns the elements of l transformed by f
func (l IntList) Map(f unaryFunc) IntList {
	out := make(IntList, len(l))
	for i, e := range l {
		out[i] = f(e)
	}
	return out
}

// Reverse returns reversed list
func (l IntList) Reverse() IntList {
	out := IntList{} // Test expects non-nil slice for empty output
	len := len(l)
	for i := len - 1; i >= 0; i-- {
		out = append(out, l[i])
	}
	return out
}

// Append returns the appended lists
func (l IntList) Append(a IntList) IntList {
	out := make(IntList, len(l)+len(a))
	copy(out, l)
	copy(out[len(l):], a)
	return out
}

// Concat returns the list that has all the lists in a concatenated together with l
func (l IntList) Concat(a []IntList) IntList {
	out := make(IntList, len(l))
	copy(out, l)

	for _, e := range a {
		out = append(out, e...)
	}
	return out
}
