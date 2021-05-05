// Interface xác định hành vi của đối tưởng. Trong Go interface là tập hợp các khai phương thức
package main

import "fmt"

// Định nghĩa 1 interface
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

//MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune  {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main()  {
	name := MyString ("Tung")
	var v VowelsFinder
	v = name // Điều này được chấp nhận MyString implement interface VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())
}