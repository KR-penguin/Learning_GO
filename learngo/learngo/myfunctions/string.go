package myfunctions

import (
  "fmt"
  "strings"
)

func LenAddUpper_1(name string) (int, string) {
	return len(name), strings.ToUpper(name)
	// len 함수 = 인수의 길이를 반환
	// strings.ToUpper 함수 = string type의 인수의 소문자글자들은 대문자로 바꿔서 반환
}

func LenAddUpper_2(name string) (length int, uppercase string) {
	// 반환 type에 변수를 넣어 반환값을 미리 지정

	defer fmt.Println("Done") // defer 안에 있는 코드는 함수가 실행된 후 (return 한 후) 실행된다

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
