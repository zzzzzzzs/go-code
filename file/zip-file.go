package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// goland 压缩文件
func main() {
	srcDirPath := "C:\\Users\\zs\\AppData\\Roaming\\JetBrains\\IntelliJIdea2022.3"
	destZipPath := "C:\\Users\\zs\\AppData\\Roaming\\JetBrains\\\\IntelliJIdea2022.3.zip"
	zipDirName := "IntelliJIdea2022.3"

	// 创建目标文件
	destFile, err := os.Create(destZipPath)
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	// 创建 Zip Writer
	zipWriter := zip.NewWriter(destFile)
	defer zipWriter.Close()

	// 遍历文件夹中的文件和子文件夹
	err = filepath.Walk(srcDirPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 创建 Zip FileHeader
		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}

		// 将文件路径转换为压缩文件中的相对路径
		relPath, err := filepath.Rel(srcDirPath, filePath)
		if err != nil {
			return err
		}

		// 如果是文件夹，将路径末尾添加分隔符
		if fileInfo.IsDir() {
			relPath = relPath + string(os.PathSeparator)
		}

		// 设置压缩文件中的文件路径
		header.Name = filepath.Join(zipDirName, relPath)

		// 如果是文件，将其压缩并写入 Zip Writer
		if !fileInfo.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}

			if _, err = io.Copy(writer, file); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
