package main

import (
	"os"
)


func main() {


	nameOfFunc := os.Args[2]
	switch os.Args[1] {
	case "kt":
		createKotlinFunction(nameOfFunc)
	case "jv":
		createJavaFunction(nameOfFunc)
	case "js":
		createNodeFunction(nameOfFunc)
	}

}
