package util 

func PtrVal[T any](t *T) T { 
	if t == nil { 
		var newT T 
		return newT
	}
	return *t 
}
 
func Ptr[T any](t T) *T { 
	return &t
}

func ArraysEqual[T comparable](ar1 []T, ar2 []T) bool { 
	if len(ar1) != len(ar2) { 
		return false
	}

	for i, v1 := range ar1 { 
		v2 := ar2[i]
		if v1 != v2 { 
			return false 
		}
	}

	return true 
}