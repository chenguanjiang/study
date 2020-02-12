package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func a1(num int) int {
	if num > 100 {
		return num
	} else if num < 0 {
		return 0
	} else {
		return num
	}
}

func a2() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", contents)
	}
}

func a3() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
		fmt.Println("=================")
		fmt.Println(reflect.TypeOf(err))
	} else {
		fmt.Println(contents)
		fmt.Println(reflect.TypeOf(contents))
		fmt.Println(string(contents[:]))
	}
}

func a4(a, b int, meth string) {
	var result int
	switch meth {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("not support this method:" + meth)
	}
	fmt.Println(result)
}

func a5(grade int) string {
	switch {
	case grade >= 90:
		return "A"
	case grade < 90 && grade >= 70:
		return "B"
	case grade < 70 && grade >= 60:
		return "C"
	case grade < 0:
		panic("grade cannot be fushu")
	default:
		return "D"
	}
}

func a6() int {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}

func a7() int {
	sum := 0
	var i int
	for i = 1; i <= 100; i++ {
		sum += i
	}
	return sum
}

func a8(num int) string {
	if num == 0 {
		return "0"
	} else {
		var result = ""
		for ; num > 0; num /= 2 {
			temp := num % 2
			result = strconv.Itoa(temp) + result
		}
		return result
	}
}

func a9(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		PrintFileContext(file)
	}
}

func PrintFileContext(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func PrintString(s string) {
	r := strings.NewReader(s)
	PrintFileContext(r)
}

func main() {
	//fmt.Println(a1(-10))
	//a4(2,0,"a")
	//fmt.Println(a5(-91))
	//fmt.Println(a7())
	//fmt.Println(a8(10))
	a9("abc.txt")
	s := `asdjg'
	'dasdasdas
	dasdsadasdsa`
	PrintString(s)
}
