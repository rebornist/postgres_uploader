package uploadData

import (
	"fmt"
	"strings"
)

func CreateQuery(tn string, data_list [][]string) []string {

	var resp []string

	for idx, data := range data_list {
		if idx == 0 {
			continue
		}

		var q string
		q = fmt.Sprintf("INSERT INTO %s (%s)", tn, strings.Join(data_list[0], ", "))

		var sq string
		var ar []string

		for _, d := range data {
			t := fmt.Sprintf("'%s'", d)
			ar = append(ar, t)
		}

		if len(data) < len(data_list[0]) {
			for i := 0; i < len(data_list[0])-len(data); i++ {
				ar = append(ar, "''")
			}
		}

		sq = fmt.Sprintf("VALUES (%s)", strings.Join(ar, ", "))

		q = fmt.Sprintf("%s %s", q, sq)

		resp = append(resp, q)
	}

	return resp

}

func CreateDocument(tn string, data_list [][]string) []map[string]string {

	var resp []map[string]string

	fmt.Println(resp)

	for idx, data := range data_list {
		if idx == 0 {
			continue
		}

		doc := make(map[string]string)
		keys := data_list[0]

		for i, d := range data {
			doc[keys[i]] = d
		}

		resp = append(resp, doc)
	}

	return resp
}
