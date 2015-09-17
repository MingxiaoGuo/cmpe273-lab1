package main
import "testing"
//import "math"

func Testmain(t *testing.T){
	var result float64

	r := Rectangle{width: 3, length: 2}
	result = r.perimeter()
	if result != (3 + 2) * 2 {
		t.Error("Expected ", (3 + 2) * 2, "got ", result)
	}
}