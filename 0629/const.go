package main

import (
	"fmt"
	"reflect"
)

const x int64 = 10

const (
	idKey = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {

	const y = "hello"

	// const定义常量  Go中只有常量具有不可变性，变量是不能不可变的。
	//x = x +1
	//y = "bye"


	fmt.Println(x)
	fmt.Println(y)

//	 获取变量类型
	fmt.Println(reflect.TypeOf(y).Name())
	fmt.Println(reflect.TypeOf(y))
	fmt.Printf("%T\n",y)





	zz := 10
	fmt.Println(zz)
	zz = 90
}
