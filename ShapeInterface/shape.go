package main
import ( "math")
type Shape interface {
	perimeter() float64 
}

type Circle struct {
	diameter float64;
}

type Rectangle struct {
	length float64;
	width float64;
}

func (c Circle) perimeter() float64{
	if c.diameter <= 0 {
		return 0
	}
	return math.Pi * c.diameter * 2;
}

func (r Rectangle) perimeter() float64{
	if r.length <= 0 || r.width <= 0 {
		return 0
	}
	return 2 * (r.length + r.width);
}

func main(){
	
}