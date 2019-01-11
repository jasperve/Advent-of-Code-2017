package main

import(
	"fmt"
)

func main() {

	fmt.Println(calculatePosition(-5, 4, -1, 2))
	fmt.Println(calculatePosition(-4, 2, 0, 2))
	fmt.Println(calculatePosition(-2, 1, 0, 2))
}


func calculatePosition(beginPosition int, velocity int, acceleration int, time int) int {

	return beginPosition + time * velocity + time * (time + 1) / 2 * acceleration

}

/*

x2 + 1x / 2 * - + x * 4 + -5 ==  -2 + x * 1 + x * (x + 1) / 2 * 0

=<-6,0,0>, v=< 3,0,0>, a=< 0,0,0>    
p=<-4,0,0>, v=< 2,0,0>, a=< 0,0,0>    -6 -5 -4 -3 -2 -1  0  1  2  3
p=<-2,0,0>, v=< 1,0,0>, a=< 0,0,0>
*/