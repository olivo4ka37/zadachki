package main

import "fmt"

func main() {
	fmt.Println(compareStrings(fullResult, testFullResult))
	fmt.Println(compareStrings(dirResult, testDirResult))
}

func compareStrings(str1, str2 string) string {
	length := len(str1)
	if len(str2) < length {
		length = len(str2)
	}

	for i := 0; i < length; i++ {
		if str1[i] != str2[i] {
			return fmt.Sprintf("Difference at position %d: '%c' vs '%c'", i, str1[i], str2[i])
		}
	}

	if len(str1) != len(str2) {
		return fmt.Sprintf("Strings differ in length: %d vs %d", len(str1), len(str2))
	}

	return "Strings are identical"
}

const testFullResult = `├───project
│	├───file.txt (19b)
│	└───gopher.png (70372b)
├───static
│	├───a_lorem
│	│	├───dolor.txt (empty)
│	│	├───gopher.png (70372b)
│	│	└───ipsum
│	│		└───gopher.png (70372b)
│	├───css
│	│	└───body.css (28b)
│	├───empty.txt (empty)
│	├───html
│	│	└───index.html (57b)
│	├───js
│	│	└───site.js (10b)
│	└───z_lorem
│		├───dolor.txt (empty)
│		├───gopher.png (70372b)
│		└───ipsum
│			└───gopher.png (70372b)
├───zline
│	├───empty.txt (empty)
│	└───lorem
│		├───dolor.txt (empty)
│		├───gopher.png (70372b)
│		└───ipsum
│			└───gopher.png (70372b)
└───zzfile.txt (empty)
`

const testDirResult = `├───project
├───static
│	├───a_lorem
│	│	└───ipsum
│	├───css
│	├───html
│	├───js
│	└───z_lorem
│		└───ipsum
└───zline
	└───lorem
		└───ipsum
`

const fullResult = `├───project
        │       ├───file.txt (19b)
        │       └───gopher.png (70372b)
        ├───static
        │       ├───a_lorem
        │       │       ├───dolor.txt (empty)
        │       │       ├───gopher.png (70372b)
        │       │       └───ipsum
        │       │               └───gopher.png (70372b)
        │       ├───css
        │       │       └───body.css (28b)
        │       ├───empty.txt (empty)
        │       ├───html
        │       │       └───index.html (57b)
        │       ├───js
        │       │       └───site.js (10b)
        │       └───z_lorem
        │               ├───dolor.txt (empty)
        │               ├───gopher.png (70372b)
        │               └───ipsum
        │                       └───gopher.png (70372b)
        ├───zline
        │       ├───empty.txt (empty)
        │       └───lorem
        │               ├───dolor.txt (empty)
        │               ├───gopher.png (70372b)
        │               └───ipsum
        │                       └───gopher.png (70372b)
        └───zzfile.txt (empty)

`

const dirResult = `├───project
        ├───static
        │       ├───a_lorem
        │       │       └───ipsum
        │       ├───css
        │       ├───html
        │       ├───js
        │       └───z_lorem
        │               └───ipsum
        └───zline
                └───lorem
                        └───ipsum

`
