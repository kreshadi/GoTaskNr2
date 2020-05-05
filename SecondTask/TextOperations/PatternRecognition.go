package main

//checks if w contains patterns xxx at any position

func check(w string) string {
	str := "The word W does not contain cube or patterns like xxx"
	str2 := "The word W contains cube or patterns like xxx"
	for d := 1; d < 16; d++ {
		for i := 0; i < len(w)-3*d; i++ {
			u := w[i:(i + d)]
			fu := w[i+d : (i + 2*d)]
			gu := w[i+2*d : (i + 3*d)]

			if (u == fu) && (fu == gu) {
				return str2
			}
		}
	}
	return str
}

