package main

import (
	"fmt"
)

//checks if w contains patterns xxx at any position

func check(w string) {

	for d := 1; d < 16; d++ {
		for i := 0; i < len(w)-3*d; i++ {
			u := w[i:(i + d)]
			fu := w[i+d : (i + 2*d)]
			gu := w[i+2*d : (i + 3*d)]

			if (u == fu) && (fu == gu) {
				fmt.Println("The word W does not contain cube or patterns like xxx")
			}
		}
	}
}

func main() {

	check("012201210001021201")

}
