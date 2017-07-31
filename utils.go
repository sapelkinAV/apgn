package main

import "os"

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
	data, err := Asset(path)
	check(err)
	return string(data[:])
}