package fib
import "testing"

func TestFib(t *testing.T){
	var result int
	result = fib(3)
	if result != 2 {

		t.Error("Expected 2, got ", result)
	}
}