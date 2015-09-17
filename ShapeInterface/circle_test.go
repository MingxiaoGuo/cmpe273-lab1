package main
import "math"
import "testing"

func Test456(t *testing.T){
	var result float64
	c := Circle{diameter: 3}
	result = c.perimeter()
	if result != math.Pi * 3 * 2 {
		t.Error("Expected ", math.Pi * 3 * 2,  "got ", result)
	}
}