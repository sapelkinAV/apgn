package main

import (
	"strings"
	"path/filepath"
	"os"
)




func createKotlinFunction(path string) {
	const packageTamplateName string = "nameOfThePackage"
	const functionTemplateName string = "Function"

	kotlinFunction := getStringFromBindata("data/kotlin/Function.kt")
	var functionName string = strings.Title(path) +"Handler"
	var packageName string = "main"

	kotlinFunction = strings.Replace(kotlinFunction,functionTemplateName,functionName, 1)
	kotlinFunction = strings.Replace(kotlinFunction,packageTamplateName,packageName, 1)

	functionJson := getStringFromBindata("data/kotlin/function.json")
	functionJson = strings.Replace(functionJson,packageTamplateName,packageName, 2)
	functionJson = strings.Replace(functionJson,functionTemplateName,functionName, 2)

	addFuncToSettingsGradle(path)
	absPath, _ := filepath.Abs("./functions")
	if _, err := os.Stat(absPath); err == nil {
		path = absPath+"/"+path
	}


	createDir(path)
	createDir(path + "/src")
	createDir(path+"/src/main")
	createDir(path+"/src/main/kotlin")


	writeToFile(path + "/build.gradle", getStringFromBindata("data/kotlin/build.gradle"))
	writeToFile(path + "/src/main/kotlin/" + functionName + ".kt", kotlinFunction)
	writeToFile(path + "/function.json",functionJson)
}