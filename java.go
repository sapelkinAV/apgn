package main

import (
	"strings"
	"path/filepath"
	"os"
)

func createJavaFunction(path string) {
	const packageTamplateName string = "nameOfThePackage"
	const functionTemplateName string = "Function"

	javaFunction := getStringFromBindata("data/java/Function.java")
	var functionName string = strings.Title(path) +"Handler"
	var packageName string = "main"

	javaFunction = strings.Replace(javaFunction,functionTemplateName,functionName, 1)
	javaFunction = strings.Replace(javaFunction,packageTamplateName,packageName, 1)

	functionJson := getStringFromBindata("data/java/function.json")
	functionJson = strings.Replace(functionJson,packageTamplateName,packageName, 2)
	functionJson = strings.Replace(functionJson,functionTemplateName,functionName, 2)

	absPath, _ := filepath.Abs("./functions")
	if _, err := os.Stat(absPath); err == nil {
		path = absPath+"/"+path
	}


	createDir(path)
	createDir(path + "/src")
	createDir(path+"/src/main")
	createDir(path+"/src/main/java")


	writeToFile(path + "/build.gradle", getStringFromBindata("data/java/build.gradle"))
	writeToFile(path + "/src/main/java/" + functionName + ".java", javaFunction)
	writeToFile(path + "/function.json",functionJson)
}
