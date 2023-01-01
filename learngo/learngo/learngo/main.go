package main

import (
	"fmt"
	/*
		  "main/learngo/myfunctions"
			"strings"
	*/)

// struct(구조체) 선언방식
// type name struct { ... }

type developer struct {
	name     string // 구조체의 멤버변수 선언 (name type)
	age      int
	height   int
	weight   float32
	uselangs []string
}

func main() {

	/*
	  var a int = 0
	  var b int = 0

	  fmt.Scanln(&a)
	  fmt.Scanln(&b)

	  fmt.Println(multiply(a, b))
	*/

	/*

	     totalLength, upperName := lenAddUpper("kopeng")
	   //totalLength, _ := lenAddUpper("kopeng") '_'는 무시된 값으로 2번쨰 반환값을 무시한다는 말이다.
	     fmt.Println(totalLength, upperName)

	     repeatMe("kopeng", "mading", "reamy", "ggubal", "chniming", "ngshin")

	*/

	/*

		var number int = 0
		var number2 int = 0

		fmt.Scanln(&number)
		fmt.Scanln(&number2)

		plus, minus, multiplication, division := myfunctions.Calculate(number, number2)
		fmt.Println(plus, minus, multiplication, division)

		totalLength, uppercase := lenAddUpper("kopeng")
		fmt.Println(totalLength, uppercase)

	*/

	/*

		result := myfunctions.SuperAdd(1, 2, 3, 4, 5, 6)
		fmt.Println(result)

	*/

	/*

		fmt.Println(myfunctions.CanIdrinkAcohol(13))
		fmt.Println(myfunctions.CanIdrinkCoffee(13))

	*/

	/*

	  var a int = 10
	  var b int = 20

	  c := &a

	  // golang에서 포인터 변수를 선언하는 방법 (C와 흡사하다)

	  var d *int = &b

	  fmt.Println(a, b, c, *d)

	  myfunctions.Swap(&a, &b)

	  fmt.Println(a, b, c, *d)

	*/

	/*

	  // golang에서 array의 초기값을 초기화할때는 '='을 쓰지 않는다.

	  names := [5]string{"kopeng", "mading", "reamy"}

	  numbers := [5]int{1, 2, 3, 4, 5}

	  var words [5]string // 이렇게 선언하면 초기값을 초기화할 수 없다.
	  words[0] = "apple"
	  words[1] = "banana"
	  words[2] = "orange"

	*/

	/*

	  // golang에서 원하는 만큼 인덱스를 가지고 싶을때는 length를 빼고 선언하면 된다.

	  langs := []string{"C++", "C#", "Java", "JavaScript"}  // 하지만 이 array의 길이는 4이므로
	  값을 더 추가할 때는 'append()' function을 사용해야 한다.


	   // append() function의 첫번째 argument는 수정할 array(or slice)이고 두번째 argument는 추
	      가할 값이다
	   // * append() function은 새롭게 수정된 array(or slice)를 반환하기 때문에 반드시 기존 배열에     대입해줘야 한다

	  langs = append(langs, "TypeScript")


	  fmt.Println(names)
	  fmt.Println(numbers)
	  fmt.Println(words)
	  fmt.Println(langs)

	*/

	/*

			// map은 key와 value로 이루어진 자료구조이다.
			// 참고 : https://memostack.tistory.com/234

			kopeng := map[string]string{"name": "kopeng", "age": "12"}

			mading := map[string]int{"age": 15, "height": 156, "weight": 80}

		  // map을 range로 전달시켜 loop 시킬 수도 있다.

		  for key, value := range kopeng {
		    println(key, value)
		  }

			fmt.Println(kopeng)
			fmt.Println(mading)

	*/

	UseProgrammingLanguages := []string{"C/C++", "Python", "LuaScript", "JavaScript", "Go"}

	kopeng := developer{name: "kopeng", age: 15, height: 167, weight: 58.6, uselangs: UseProgrammingLanguages}

	fmt.Println(kopeng)
}
