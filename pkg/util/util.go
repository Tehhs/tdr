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