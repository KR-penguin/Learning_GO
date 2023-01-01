package myfunctions

import "fmt"

func RepeatMe(words ...string) { // argument type안에 ...을 붙이면 해당 type의 인자를 원하는 만큼 넣을 수 있다.
	fmt.Println(words)
}

func SuperAdd(numbers ...int) int {
	// 반환 type에 변수를 넣어 반환값을 미리 지정

	defer fmt.Println("superAdd Done")
	// defer 안에 있는 코드는 함수가 실행된 후 (return 한 후) 실행된다

	/*
	  for i := 0; i < len(numbers); i++ {  <-- 와 밑에 코드(range를 사용한 부분)는 같은 역할을 함
	    ...
	  }
	*/

	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}