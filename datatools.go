package datatools

import (
	"archive/zip"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var OutputFile *os.File
var OutputQueue []string

func getPageAndParse(url string) *html.Node {
	resp := GetPage(url)
	root := ParseContent(resp)
	return root
}

func ParseContent(response *http.Response) *html.Node {
	root, err := html.Parse(response.Body)
	Check(err)

	return root
}

func Match(n *html.Node, tag string, text string) bool {
	if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
		link := string(scrape.Attr(n, tag))
		matched, err := regexp.MatchString(text, link)
		Check(err)

		DebugVerbose(fmt.Sprintf("Link: %s %s\n", link, matched))

		if matched {
			DebugVerbose(fmt.Sprintf("*** Matched: %s\n", link))
			return matched
		}
	}
	return false
}

// ExtractFile - Open a zip archive for reading.
func ExtractFile(file string) {
	OutputQueue = nil

	r, err := zip.OpenReader(file)
	Check(err)

	defer r.Close()

	for _, f := range r.File {
		FileDescriptors := regexp.MustCompile("_").Split(f.Name, -1)
		FileContent := strings.ToUpper(FileDescriptors[len(FileDescriptors)-1])
		fileType := strings.ToLower(regexp.MustCompile(regexp.QuoteMeta(".")).Split(FileContent, -1)[0])

		Debug(fmt.Sprintf("Extracting from %s", f.Name))
		rc, err := f.Open()
		Check(err)

		scanner := bufio.NewScanner(rc)
		for scanner.Scan() {
			ReadCsv(fileType, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			Check(err)
		}

		rc.Close()

		//TODO: Write Somewhere.

		//panic("Development Stop in Content.go")
	}

}
