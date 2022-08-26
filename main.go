package main

import (
	"fmt"
	"regexp"
	"strings"
)

type arguments struct {
	pascalCase bool
	upperCase  bool
}

func camelCase(word string, p arguments) string {

	var lowercase = ""
	for _, chr := range word {
		if chr >= 65 && chr <= 90 {
			chr += 32
		}
		lowercase += fmt.Sprintf("%c", chr)
	}

	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	whiteSpaceWord := reg.ReplaceAllString(lowercase, " ")
	trimmed := strings.Trim(whiteSpaceWord, " ")

	var bArr []string
	var aArr []string
	var cArr []string
	var dArr []string

	if p.upperCase {
		bArr = strings.Split(trimmed, " ")
		bArr[len(bArr)-1] = strings.ToUpper(bArr[len(bArr)-1])
		aArr = strings.Split(strings.Join(bArr, " "), "")
	} else {
		aArr = strings.Split(trimmed, "")
	}

	for i := 0; i < len(aArr); i++ {
		element := aArr[i]
		if element == " " {
			cArr = append(aArr[:i], strings.ToUpper(aArr[i+1]))
			dArr = append(cArr, aArr[i+2:]...)
		}
	}

	if len(dArr) < 1 {
		return word
	}

	camel := strings.Join(dArr, "")
	upChar := strings.ToUpper(string(camel[0]))
	lowChar := strings.ToLower(string(camel[0]))
	if p.pascalCase {
		if lowChar == string(camel[0]) {
			arr := strings.Split(camel, "")
			arr[0] = strings.ToUpper(arr[0])
			camelFix := strings.Join(arr, "")
			return camelFix
		}
		return camel
	} else {
		if upChar == string(camel[0]) {
			arr := strings.Split(camel, "")
			arr[0] = strings.ToLower(arr[0])
			camelFix := strings.Join(arr, "")
			return camelFix
		}
	}

	return camel
}

func main() {
	fmt.Println(camelCase("foo-bar", arguments{pascalCase: false, upperCase: false}))
	fmt.Println(camelCase("foo_bar", arguments{pascalCase: false, upperCase: false}))
	fmt.Println(camelCase("Foo-bar", arguments{pascalCase: false, upperCase: false}))
	fmt.Println(camelCase("Foo-bar", arguments{pascalCase: true, upperCase: false}))
	fmt.Println(camelCase("--foo.bar", arguments{pascalCase: false, upperCase: false}))
	fmt.Println(camelCase("Foo-BAR-ba", arguments{pascalCase: false, upperCase: true}))
	fmt.Println(camelCase("foobar", arguments{pascalCase: false, upperCase: false}))
	fmt.Println(camelCase("fooBAR", arguments{pascalCase: true, upperCase: true}))
	fmt.Println(camelCase("foo bar", arguments{pascalCase: false, upperCase: false}))
}
