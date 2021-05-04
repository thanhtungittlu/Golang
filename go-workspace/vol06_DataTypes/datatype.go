package main

import "fmt"

func main() {
	// Boolean
	fmt.Println()
	var test bool = true
	test = 1 > 2
	fmt.Printf("%v, %T\n", test, test)
	fmt.Println()
	// Numberics
	var a int8 = 10 // 1010
	var b int8 = 3  // 0011
	// Các phép toán trên bit
	fmt.Printf("%v, %T\n", a&b, a&b)   //0010 -> 2
	fmt.Printf("%v, %T\n", a|b, a|b)   //1011 -> 11
	fmt.Printf("%v, %T\n", a^b, a^b)   //1001 -> 9
	fmt.Printf("%v, %T\n", a&^b, a&^b) //0100 -> 8
	fmt.Println()
	fmt.Println()
	var c complex64 = 1 + 2i // var c complex64 = complex(5,12)
	fmt.Printf("%v, %T\n", imag(c), imag(c))
	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))
	fmt.Println()
	fmt.Println()
	var string1 string = "Hello"
	var string2 string = ", Cam on nhieu"
	var mang []byte = []byte(string1)
	fmt.Printf("%v, %T\n", string1+string2, string1+string2)
	fmt.Printf("%v, %T\n", mang, mang)
	fmt.Printf("%v, %T\n", string(mang), mang) // chuyển kiểu dữ liệu

}

/* 1. Boolean // Giá trị mặc định của Bool là false
   2. Numberics: Integer - Float - Complex
    - Integer
			signed      int8	int16	int32	int64
			unsigned	nint8	uint16 	uint32  uint6464
	- Float
			float32		float64
	- Complex -- Số phức
			complex64 	complex128
   3. Text: String - Byte - Rune
	- String
	- Byte -> UTF-8
	- Rune -> UTF-32

*/
