package datatools

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func MoveFile(fromFile string, toFile string) error {
	Check(os.Rename(fromFile, toFile))
	return nil
}

func GetInputFolder() string {
	// 01/02 03:04:05PM '06 -0700

	root := AppConfig.Settings.Input
	folder := time.Now().Format("2006-01-02")

	if AppConfig.Settings.FixedDate != "" {
		folder = AppConfig.Settings.FixedDate
	}

	if _, err := os.Stat(root); os.IsNotExist(err) {
		os.Mkdir(root, 0755)
	}

	if _, err := os.Stat(root + "/" + folder); os.IsNotExist(err) {
		os.Mkdir(root+"/"+folder, 0755)
	}

	fmt.Sprintf("%d", "Data Folder: "+root+"/"+folder)

	return root + "/" + folder
}

func GetOutputFolder() string {
	// 01/02 03:04:05PM '06 -0700

	root := AppConfig.Settings.Output
	folder := time.Now().Format("2006-01-02")

	if AppConfig.Settings.FixedDate != "" {
		folder = AppConfig.Settings.FixedDate
	}

	if _, err := os.Stat(root); os.IsNotExist(err) {
		os.Mkdir(root, 0755)
	}

	if _, err := os.Stat(root + "/" + folder); os.IsNotExist(err) {
		os.Mkdir(root+"/"+folder, 0755)
	}

	fmt.Sprintf("%d", "Store Folder: "+root+"/"+folder)

	return root + "/" + folder
}

func ExtractToFolder(file string, outputDirectory string) {
	Check(unzip(file, outputDirectory))
}

func WriteStringToFile(output string, file string) error {
	OutputFile, err := os.Create(GetOutputFolder() + "/" + file)
	Check(err)

	defer OutputFile.Close()

	result, err := OutputFile.WriteString(output)

	Check(err)
	fmt.Sprintf("%d", result)

	return err
}

// Ref: http://blog.ralch.com/tutorial/golang-working-with-zip/
func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
