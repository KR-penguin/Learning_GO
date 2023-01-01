package myfunctions

import "fmt"

func CanIdrinkAcohol(age int) bool {

	defer fmt.Println("CanIdrink done")
	/*

	  if age < 20 {
	    return false
	  } else {
	    return true
	  }

	*/

	/*
	  if age < 20 {
	    return false // 함수는 return에서 끝나기 때문에 굳이 else문을 쓸 필요가 없다
	  }
	  return true

	*/

	// go에서는 if문에서 변수를 생성할 수 있다.

	if Koreanage := age + 2; Koreanage < 20 {
		return false
	}
	return true
}

// go에서는 switch 문에서 변수를 생성할 수 있다.

func CanIdrinkCoffee(age int) bool {

	defer fmt.Println("CanIdrinkCoffee Done")

	/*
	   switch Hoage := age - 2; Hoage {
	     case 12:
	       return false
	     case 18:
	       return true
	   }
	*/
	switch {
	case age > 20:
		return false
	case age == 20:
		return true
	case age > 70:
		return false
	default:
		return true
	}
}
