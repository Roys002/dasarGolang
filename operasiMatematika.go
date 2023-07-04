package main

import "fmt"

func main() {
	var(
		 a int16 = 5
		 b int16 = 9
		 c int16 = 7
		 d int16 = 7 
		 e int16 = 7
		 f int16 = 7
	
		 sum = a + b
		 x = a * b
		 y = a / b
		 z = a % b
		 
		)
		// augmented assignment
		c += 10 /* sama dengan c = c + 10 */
		d *= 2 /* sama dengan d = d * 10 */
		e /= 2	/* sama dengan e = e / 10 */
		f %= 2	/* sama dengan f = f % 10 */
	fmt.Println(sum)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	// unary operation
	
}