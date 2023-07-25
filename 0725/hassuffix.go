

package main

import (
"fmt"
"strings"

)

//golang字符串操作
func main(){
	s := "hello world hello world"
	//str := "wo"
	//var s = []string{"11","22","33"}

	//判断字符串s是否以prefix结尾
	ret := strings.HasSuffix(s,"rl")
	fmt.Println(ret) //false
}
