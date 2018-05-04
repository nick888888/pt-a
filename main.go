package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "NO FARE FOR CLASS USED qqq"
	fmt.Println(errorContainSuccessString(a))

}


//  测试strings.Contains
func errorContainSuccessString(a string) (isContain bool) {
	esUpper:=strings.ToUpper(a)
	successString := []string{
		"No complete journey can be built in IF2/ADVJR1",
		"NO AVAILABLE FLIGHT SCHEDULES",
		"This response version is deprecated and will be decommissioned once a newer version is released.",
		"NO COMBINABLE SCHEDULES RETURNED. ALL SCHEDULES FILTERED OUT",
		"NO FARE FOR CLASS USED",
		"NO FLIGHT SCHEDULES FOR QUALIFIERS USED",
		"No Availability",
		"NO FLIGHTS FOUND FOR",
	}
	for _, v := range successString {
		//  fmt.Println(strings.Contains("seafood", "foo")) //true 前面的参数是否包含后面的参数
		if strings.Contains(esUpper, strings.ToUpper(v)){
			fmt.Println("true----1111")
			return true
		}
	}
	fmt.Println("false----222")
	return
}
