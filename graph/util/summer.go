package util

// This is called type set union
type Number interface {
	int32 | float64  
}

type UtilServices struct {}

// Generic to
func sumNumbers[T Number](t []T) (T,error){
	var result T
	var i int

	for i=range t {
		result += t[i]
	}

	return result,nil
}

