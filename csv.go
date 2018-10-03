package datatools

import (
	"encoding/csv"
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

		switch content {
		case "alpha":
			DebugVerbose("ALPHA CSV OPTION")
		default:
			DebugVerbose("Could Not Determine Format")
		}

	}

}
