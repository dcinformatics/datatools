package datatools

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func ReadCsv(content string, data string) {
	r := csv.NewReader(strings.NewReader(data))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		Check(err)
		DebugVerbose(fmt.Sprintf("%s", record))

		switch content {
		case "alpha":
			DebugVerbose("ALPHA CSV OPTION")
		default:
			DebugVerbose("Could Not Determine Format")
		}

	}

}
