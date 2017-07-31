package main

import (
	"path/filepath"
	"os"
)

func createNodeFunction(path string){

	indexJs := getStringFromBindata("data/nodejs/index.js")
	functionJson := getStringFromBindata("data/nodejs/function.json")

	absPath, _ := filepath.Abs("./functions")
	if _, err := os.Stat(absPath); err == nil {
		path = absPath+"/"+path
	}


	createDir(path)
	writeToFile(path + "/function.json", functionJson)
	writeToFile(path + "/index.js", indexJs)

}