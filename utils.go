package main

import (
	"os"
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

func addFuncToSettingsGradle(funcname string) {
	absPath, _ := filepath.Abs("./functions")
	if _, err := os.Stat(absPath); err == nil {
		f, err := os.OpenFile("settings.gradle", os.O_APPEND|os.O_WRONLY, 0600)
		check(err)
		if _, err = f.WriteString("include " + "'functions/"+funcname+"'"); err != nil {
			panic(err)
		}
	}


}