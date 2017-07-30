package main

import (
	"os"
	"strings"
	"regexp"
)

// func main() {

// 	// `os.Args` provides access to raw command-line
// 	// arguments. Note that the first value in this slice
// 	// is the path to the program, and `os.Args[1:]`
// 	// holds the arguments to the program.
// 	argsWithProg := os.Args
// 	argsWithoutProg := os.Args[1:]

// 	// You can get individual args with normal indexing.
// 	arg := os.Args[1]

// 	fmt.Println(argsWithProg)
// 	fmt.Println(argsWithoutProg)
// 	fmt.Println(arg)
// }
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func createDir(path string){
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path,0777)

	}
}

func writeToFile(path string,content string) {
	f,err := os.Create(path)
	check(err)
	f.WriteString(content)
	defer f.Close()
}
func getStringFromBindata(path string) string  {
	data, err := Asset("data/build.gradle")
	check(err)
	return string(data[:])
}
func main() {
	const packageTamplateName string = "nameOfThePackage"
	path := os.Args[1]
	createDir(path)
	createDir(path + "/src")
	createDir(path + "/src/main/kotlin")
	createDir(path + "/src/main/java")


	kotlinFunction := getStringFromBindata("data/Function.kt")
	var functionName string = strings.Title(path) +"Handler"
	var packageName string = "main"
	kotlinFunction = strings.Replace(kotlinFunction,"Function",functionName, 1)
	kotlinFunction = strings.Replace(kotlinFunction,packageTamplateName,packageName, 1)

	functionJson := getStringFromBindata("data/function.json")



	writeToFile(path + "/build.gradle", getStringFromBindata("data/build.gradle"))








}
