package main

import (
	"bufio"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
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

	outputFile, err := os.Create("compressed_" + ratioStr + "_" + filename)
	if err != nil {
		fmt.Println("Error creating output file: ", err)
		return
	}
	defer outputFile.Close()

	bufferedWriter := bufio.NewWriter(outputFile)
	defer bufferedWriter.Flush()

	switch format {
	case "jpeg":
		err = jpeg.Encode(bufferedWriter, img, &jpeg.Options{Quality: int(ratio)})
		if err != nil {
			fmt.Println("Error encoding JPEG: ", err)
			return
		}
	case "png":
		err = png.Encode(bufferedWriter, img)
		if err != nil {
			fmt.Println("Error encoding PNG: ", err)
			return
		}
	case "gif":
		err = gif.Encode(bufferedWriter, img, &gif.Options{NumColors: int(ratio)})
		if err != nil {
			fmt.Println("Error encoding GIF: ", err)
			return
		}
	default:
		fmt.Println("Unsupported image format")
		return
	}

	fmt.Println("Image compressed successfully")
}
