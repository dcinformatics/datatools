package datatools

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

func DownloadFiles(url string, path string, attr string, list string, tag string, prefix string, ext string) []string {
	var fileList []string

	matchPage := func(n *html.Node) bool {
		return Match(n, attr, list)
	}

	matchFile := func(n *html.Node) bool {
		return Match(n, attr, tag)
	}

	pages := scrape.FindAll(getPageAndParse(url+"/"+path), matchPage)
	for i, page := range pages {

		pagePath := scrape.Attr(page, attr)
		downloadPage := url + "/" + pagePath
		DebugVerbose(fmt.Sprintf("#%2d Year: %s %s)\n", i, scrape.Text(page), pagePath))

		files := scrape.FindAll(getPageAndParse(downloadPage), matchFile)

		for j, file := range files {
			fileName := GetInputFolder() + "/" + prefix + scrape.Text(file) + ext

			if _, err := os.Stat(fileName); os.IsNotExist(err) {
				Debug(fmt.Sprintf("Downloading to %s\n", fileName))

				filePath := scrape.Attr(file, attr)

				Debug(fmt.Sprintf("\t%2d: %s\n", i, filePath))

				zipFile := GetPage(filePath)

				file, err := os.Create(fileName)
				Check(err)

				_, err = io.Copy(file, zipFile.Body)
				Check(err)

				defer zipFile.Body.Close()
				file.Close()
			}

			j++

			fileList = append(fileList, fileName)
		}

	}

	return fileList
}

func CheckExtractFiles(moveFiles bool) error {
	var err error

	start := 1995
	now := time.Now()
	end := now.Year()

	directory := GetOutputFolder()

	for i := start; i <= end; i++ {
		if i <= 2011 {
			if _, err := os.Stat(directory + "/hosp_" + strconv.Itoa(i) + "_ALPHA.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp_%d_ALPHA.CSV", directory, i))
				if moveFiles {
					MoveHcrisCsvFile("hosp_" + strconv.Itoa(i) + "_ALPHA.CSV")
				}
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp_%d_ALPHA.CSV %s", directory, i))
			}

			if _, err := os.Stat(directory + "/hosp_" + strconv.Itoa(i) + "_NMRC.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp_%d_NMRC.CSV", directory, i))
				if moveFiles {
					MoveHcrisCsvFile("hosp_" + strconv.Itoa(i) + "_NMRC.CSV")
				}
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp_%d_NMRC.CSV", directory, i))
			}

			if _, err := os.Stat(directory + "/hosp_" + strconv.Itoa(i) + "_RPT.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp_%d_RPT.CSV", directory, i))
				if moveFiles {
					MoveHcrisCsvFile("hosp_" + strconv.Itoa(i) + "_RPT.CSV")
				}
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp_%d_RPT.CSV", directory, i))
			}

			if _, err := os.Stat(directory + "/hosp_" + strconv.Itoa(i) + "_ROLLUP.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp_%d_ROLLUP.CSV", directory, i))
				if moveFiles {
					MoveHcrisCsvFile("hosp_" + strconv.Itoa(i) + "_ROLLUP.CSV")
				}
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp_%d_ROLLUP.CSV", directory, i))
			}
		}

		if i >= 2010 {
			if _, err := os.Stat(directory + "/hosp10_" + strconv.Itoa(i) + "_ALPHA.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp10_%d_ALPHA.CSV", directory, i))
				if moveFiles {
					MoveHcrisCsvFile("hosp10_" + strconv.Itoa(i) + "_ALPHA.CSV")
				}
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp10_%d_ALPHA.CSV", directory, i))
			}

			if _, err := os.Stat(directory + "/hosp10_" + strconv.Itoa(i) + "_NMRC.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp10_%d_NMRC.CSV", directory, i))
				if moveFiles {
					MoveHcrisCsvFile("hosp10_" + strconv.Itoa(i) + "_NMRC.CSV")
				}
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp10_%d_NMRC.CSV", directory, i))
			}

			if _, err := os.Stat(directory + "/hosp10_" + strconv.Itoa(i) + "_RPT.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp10_%d_RPT.CSV", directory, i))
				if moveFiles {
					MoveHcrisCsvFile("hosp10_" + strconv.Itoa(i) + "_RPT.CSV")
				}
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp10_%d_RPT.CSV", directory, i))
			}
		}
	}

	return err
}

func CheckAndMoveExtractFiles() error {
	CheckExtractFiles(true)
	return nil
}

func MoveHcrisCsvFile(fileName string) error {
	currentTime := time.Now()
	checkTypes := "ALPHA,NMRC,RPT,ROLLUP"

	fileComponents := strings.Split(fileName, "_")

	fileType := fileComponents[2]
	fileYear := fileComponents[1]

	if strings.Index(checkTypes, fileType) == -1 {
		Fail(fmt.Sprintf("Invalid file type: %s", fileType))
	}

	if year, err := strconv.Atoi(fileYear); err == nil {
		if year < 1995 || year > currentTime.Year() {
			Fail(fmt.Sprintf("Invalid year: %s", fileYear))
		}
	}

	fileSource := GetOutputFolder() + "/" + fileName
	fileDest := GetOutputFolder() + "/" + fileType + "/" + fileYear + "/" + fileName

	if _, err := os.Stat(GetOutputFolder() + "/" + fileType); os.IsNotExist(err) {
		os.Mkdir(GetOutputFolder()+"/"+fileType, 0755)
	}

	if _, err := os.Stat(GetOutputFolder() + "/" + fileType + "/" + fileYear); os.IsNotExist(err) {
		os.Mkdir(GetOutputFolder()+"/"+fileType+"/"+fileYear, 0755)
	}

	Check(MoveFile(fileSource, fileDest))
	return nil
}
