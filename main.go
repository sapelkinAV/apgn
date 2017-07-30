package main

import (
	"os"
	"strings"
	"path/filepath"
)

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
	 f.Close()
}
func getStringFromBindata(path string) string  {
	data, err := Asset(path)
	check(err)
	return string(data[:])
}

func createFunction(path string, inRootDir bool) {
	const packageTamplateName string = "nameOfThePackage"
	const functionTemplateName string = "Function"

	kotlinFunction := getStringFromBindata("data/Function.kt")
	var functionName string = strings.Title(path) +"Handler"
	var packageName string = "main"

	kotlinFunction = strings.Replace(kotlinFunction,functionTemplateName,functionName, 1)
	kotlinFunction = strings.Replace(kotlinFunction,packageTamplateName,packageName, 1)

	functionJson := getStringFromBindata("data/function.json")
	functionJson = strings.Replace(functionJson,packageTamplateName,packageName, 2)
	functionJson = strings.Replace(functionJson,functionTemplateName,functionName, 2)

	if inRootDir {
		absPath, _ := filepath.Abs("./functions")
		path = absPath+"/"+path
	}

	createDir(path)
	createDir(path + "/src")
	createDir(path+"/src/main")
	createDir(path+"/src/main/kotlin")







	writeToFile(path + "/build.gradle", getStringFromBindata("data/build.gradle"))
	writeToFile(path + "/src/main/kotlin/" + functionName + ".kt", kotlinFunction)
	writeToFile(path + "/function.json",functionJson)
}
func main() {

	nameOfFunc := os.Args[1]
	absPath, _ := filepath.Abs("./functions")
	if _, err := os.Stat(absPath); err == nil {
		// path/to/whatever exists
		createFunction(nameOfFunc,true)
	}else {
		createFunction(nameOfFunc,false)
	}



}
