package datatools

import (
	"fmt"
	"io"
	"os"
	"strconv"
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

func CheckExtractFiles() error {
	var err error
	start := 1995
	now := time.Now()
	end := now.Year()

	directory := GetOutputFolder()

	for i := start; i <= end; i++ {
		if i <= 2011 {
			if _, err := os.Stat(directory + "/hosp_" + strconv.Itoa(i) + "_ALPHA.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp_%d_ALPHA.CSV", directory, i))
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp_%d_ALPHA.CSV %s", directory, i))
			}

			if _, err := os.Stat(directory + "/hosp_" + strconv.Itoa(i) + "_NMRC.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp_%d_NMRC.CSV", directory, i))
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp_%d_NMRC.CSV", directory, i))
			}

			if _, err := os.Stat(directory + "/hosp_" + strconv.Itoa(i) + "_RPT.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp_%d_RPT.CSV", directory, i))
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp_%d_RPT.CSV", directory, i))
			}

			if _, err := os.Stat(directory + "/hosp_" + strconv.Itoa(i) + "_ROLLUP.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp_%d_ROLLUP.CSV", directory, i))
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp_%d_ROLLUP.CSV", directory, i))
			}
		}

		if i >= 2010 {
			if _, err := os.Stat(directory + "/hosp10_" + strconv.Itoa(i) + "_ALPHA.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp10_%d_ALPHA.CSV", directory, i))
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp10_%d_ALPHA.CSV", directory, i))
			}

			if _, err := os.Stat(directory + "/hosp10_" + strconv.Itoa(i) + "_NMRC.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp10_%d_NMRC.CSV", directory, i))
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp10_%d_NMRC.CSV", directory, i))
			}

			if _, err := os.Stat(directory + "/hosp10_" + strconv.Itoa(i) + "_RPT.CSV"); err == nil {
				Pass(fmt.Sprintf("Exists %s/hosp10_%d_RPT.CSV", directory, i))
			} else {
				Fail(fmt.Sprintf("Does NOT Exist %s/hosp10_%d_RPT.CSV", directory, i))
			}
		}
	}

	return err
}
