package main

import "regexp"

func isLowerCase(ascii int) bool {
	return ascii >= START_LOWER && ascii <= END_LOWER 	
}

func isLettersOnly(data string) bool {
	var lettersOnly = regexp.MustCompile("^[a-zA-z]*$")
	return lettersOnly.MatchString(data)
}