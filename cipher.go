package main

import (
	"fmt"
	"strings"
)

const START_UPPER = 65
const START_LOWER = 97
const END_UPPER = 90
const END_LOWER = 122

func main() {
	fmt.Println("Ceasar Cipher Decoder and Encoder")
	fmt.Println("Choose operation")
	fmt.Println("1) Encode string")
	fmt.Println("2) Decode string")

	var choice int32
	for {
		fmt.Print("Choice: ")
		fmt.Scanln(&choice)
		if choice == 1 || choice == 2 {
			break
		}
	}
	
	var data string
	var shift int = 4
	if choice == 1 {
		fmt.Print("Enter shift: ")
		fmt.Scanln(&shift)
		fmt.Print("Enter string to encode: ")
		fmt.Scanln(&data)

		encoded := encode(data, shift)
		fmt.Println(encoded)
	} else if choice == 2 {
		fmt.Print("Enter shift: ")
		fmt.Scanln(&shift)
		fmt.Print("Enter string to decode: ")
		fmt.Scanln(&data)
		
		decoded := decode(data, shift)
		fmt.Println(decoded)
	}
}

func decode(data string, shift int) string {
	var decoded strings.Builder
	for i := 0; i < len(data); i++ {
		index := int(data[i])
		num := index - shift

		if isLowerCase(index) {
			if num < START_LOWER {
				diff := num - START_LOWER + 1
				num = END_LOWER - diff
			}
		} else {
			if num > END_UPPER {
				num = (START_UPPER - 1) - shift
			}
		}

		var char = string(rune(num))
		decoded.WriteString(char)
	}

	return decoded.String()
}

func encode(data string, shift int) string {
	var encoded strings.Builder
	for i := 0; i < len(data); i++ {
		index := int(data[i])
		num := index + shift
		
		if isLowerCase(index) {
			if num > END_LOWER {
				num = (START_LOWER - 1) + shift	
			}
		} else {
			if num > END_UPPER {
				num = (START_UPPER - 1) + shift
			}
		}
		
		var char = string(rune(num))
		encoded.WriteString(char)
	}

	return encoded.String()
}

func isLowerCase(ascii int) bool {
	return ascii >= START_LOWER && ascii <= END_LOWER 	
}