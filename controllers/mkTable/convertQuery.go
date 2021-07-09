package mkTable

import (
	"fmt"
)

func CreateQuery(tn string, data_list [][]string) string {

	var q string

	q = fmt.Sprintf("CREATE TABLE %s (", tn)

	for idx, data := range data_list {
		sq := fmt.Sprintf("%s %s", data[0], data[1])
		if data[2] == "1" {
			sq = fmt.Sprintf("%s PRIMARY KEY", sq)
		}

		if data[3] == "1" {
			sq = fmt.Sprintf("%s UNIQUE", sq)
		}

		if data[4] == "1" {
			sq = fmt.Sprintf("%s NOT NULL", sq)
		}

		if (idx + 1) < len(data_list) {
			sq = fmt.Sprintf("%s,", sq)
		}

		q = fmt.Sprintf("%s %s", q, sq)
	}

	q = q + " )"

	return q

}
