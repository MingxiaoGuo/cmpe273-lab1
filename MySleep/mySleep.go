package main
import "time"
import "fmt"

func mySleep(x int) {
    <-time.After(time.Second * time.Duration(x))
}

func main() {
	fmt.Println("hello");
	mySleep(3);
	fmt.Println("world");
}