package logerror

import "fmt"


func ERROR(s string,e error) {
	if e !=nil {
		fmt.Println("❌🔥",s,e)
	}
	
}