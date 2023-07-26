package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

const START_UPPER = 65
const START_LOWER = 97
const END_UPPER = 90
const END_LOWER = 122

func main() {
	operationPrompt := promptui.Select{
		Label: "Choose operation",
		Items: []string{"Encode", "Decode"},
	}

	_, result, _ := operationPrompt.Run()
	
	shift, _ := promptShift()
	marks, _ := strconv.Atoi(shift)
	
	if result == "Encode" {
		data := promptData()
		encoded := encode(data, marks)
		fmt.Println(encoded)
	} else if result == "Decode" {
		data := promptData()
		decoded := decode(data, marks)
		fmt.Println(decoded)
	}
}

func promptData() string {
	validate := func(input string) error {
		lettersOnly := isLettersOnly(input)
		empty := len(input) <= 0
		if !lettersOnly {
			return errors.New("only letters are allowed")
		}
		if empty {
			return errors.New("empty string is not allowed")
		}
		return nil
	}

	// var data string
	dataPrompt := promptui.Prompt{
		Label: "Enter data string: ",
		Validate: validate,
	}
	data, _ := dataPrompt.Run()

	return data
}

func promptShift() (string, error){
	validate := func(input string) error {
		_, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	var result string
	shiftPrompt := promptui.Prompt{
		Label: "Numeric Shift: ",
		Validate: validate,
	}

	result, err := shiftPrompt.Run()
	if err != nil {
		return "", err
	}

	shift := result
	return shift, nil
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

