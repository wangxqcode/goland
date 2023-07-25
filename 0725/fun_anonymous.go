package main

import "fmt"

/*
// 匿名函数

func(参数)(返回值){
	函数体
}



func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) (c int) {
		fmt.Println(x + y)
		return x+y
	}
	bdd :=add(10, 20) // 通过变量调用匿名函数

	fmt.Println("bdd:",bdd)

	//自执行函数：匿名函数定义完加()直接执行
	cdd :=func(x, y int) (c int) {
		fmt.Println(x + y)
		return x+y
	}(10, 20)

	fmt.Println("cdd:",cdd)

}
*/

// 闭包
/*
func worker() func(int) int {
	var x int
	return func(y int) int {

		x += y
		return x
	}

}

func main()  {
	f := worker()
	fmt.Println(f(10))
	fmt.Println(f(20))
}


 */


// 闭包进阶1
func adder2(x int) func(int) int {
	fmt.Println("x==========",x)
	return func(y int) int {
		x += y

		fmt.Println("x------",x)
		fmt.Println("y------",y)
		return x

		}
}

func main() {
	// 变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效
	// f 为返回的闭包，f(10)为闭包中的y传参传入10

	var f = adder2(10)
	fmt.Println(f(10)) //20


	//fmt.Println(f(20)) //40
	//fmt.Println(f(30)) //70
	//
	//f1 := adder2(20)
	//fmt.Println(f1(40)) //60
	//fmt.Println(f1(50)) //110
}





// 闭包进阶2

/*
func makeSuffixFunc(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {

			return name + suffix

		}

		return name
	}
}

func main() {
	// jpgFunc 中传入的参数 .jpg 使得 suffix = .jpg,然后fmt.Println(jpgFunc("test"))操作 向闭包中传入test，即name=text,
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}

 */
