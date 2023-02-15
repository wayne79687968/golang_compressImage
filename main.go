package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter file name (support jpeg, png, gif): ")
	var filename string
	fmt.Scanln(&filename)

	fmt.Print("Enter compressionratio (1-100): ")
	var ratioStr string
	fmt.Scanln(&ratioStr)
	ratio, _ := strconv.ParseFloat(ratioStr, 64)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image: ", err)
		return
	}

	outputFile, err := os.Create("compressed_" + filename)
	if err != nil {
		fmt.Println("Error creating output file: ", err)
		return
	}
	defer outputFile.Close()

	bufferedWriter := bufio.NewWriter(outputFile)
	defer bufferedWriter.Flush()
	err = jpeg.Encode(bufferedWriter, img, &jpeg.Options{Quality: int(ratio)})
	if err != nil {
		fmt.Println("Error encoding JPEG: ", err)
		return
	}

	fmt.Println("Image compressed successfully")
}
