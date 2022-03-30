package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"runtime"
)

var (
	a = 1
	b = "kkk"
	c = true
) 

func initialzation() {
	var (
		a int = 5
		b string = "bbb"
		c string = "kkk"
	)
	fmt.Printf("%d %q\n %q", a, b, c)
}

func initialzation1() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a,b,c,s)
}

func comple() {
	a := 3 + 5i
	fmt.Println(a) // (3+5i)
}

// func euler() {
// 	a := cmplx.Exp(1i*math.Pi + 1)
// 	fmt.Printf("%.3f", a)
// }

func tera() {
	a, b := 3,4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("%d",c)
}

func enums() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)
	fmt.Println(b, kb, mb, gb) // 1 1024 1048576 1073741824
}

// func checkFile() {
// 	const filename = "abc.txt"
// 	if content, err := ioutil.ReadFile(filename);
// 	// nil 是一个pointer, 用于判断是否指向正确
// 	err != nil {
// 		fmt.Println(string(content))
// 	} else {
// 		fmt.Println("could not print file contents:", err)
// 	}
// }

func checkFile() {
	const filename = "abc.txt"
	content, error := ioutil.ReadFile(filename) 
	if error == nil {
		fmt.Println("could not print file contents:", error)
	} else {
		fmt.Println(content)
	}
}

func iftest() {
	const a = 6
	if a < 5 {
		println("less")
	} else if a > 5{
		println("reject")
	} else {
		println("ok")
	}
}

func fortest(i int) {
	for ; i < 10; i++ {
		println(i)
	}
}

func fortest2() {
	for {
		fmt.Println("abc")
	}
}

func eval(a, b int, op string) (r int, e error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		r2, _ := div(a, b)
		return r2, nil
	default:
		return 0, fmt.Errorf("no")
	}
}

func div(a,b int) (r int, e error) {
	return a / b, e
}

func sum (numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func apply(op func(int, int) int, a,b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println(opName)
	return op(a,b)
}

// func main() {
// 	fmt.Println(
// 		apply(func(a int, b int) int {
// 			return int(math.Pow(
// 				float64(a), float64(b)))
// 		}, 3,5))
// 		a,b := 3,5
// 		swap(&a,&b)
// }

func swap(a,b *int) {
	*b, *a = *a, *b
}