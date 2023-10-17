package main

import (
	"fmt"
	"os"

	"shtb.dev/go/files/fileutils"
)

func main() {
	rootPath,_ := os.Getwd()
	filePath := rootPath+"/data/text.txt"

	content, err := fileutils.ReadTextFile(rootPath+"/data/text.txt")
	if err == nil {
		fmt.Println(content)
		newContent := fmt.Sprintf("OriginalL %v\n Double orginal: %v %v",content,content,content)
		fileutils.WriteToFile(filePath,newContent)
	}else{
		fmt.Printf("Error Panic: %v",err)
	}
}