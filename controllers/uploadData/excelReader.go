package uploadData

import (
	"strings"

	"github.com/tealeg/xlsx"
)

// 엑셀 파일에 존재하는 테이블 이름 및 옵선을 추출
func TableDataReader(fileName string) [][]string {
	var resp [][]string

	// 파일 열기
	file, err := xlsx.OpenFile(fileName)
	if err != nil {
		panic(err)
	}

	// 엑셀 데이타 검색 후 추출
	sheet := file.Sheets[0]
	for idx, row := range sheet.Rows {
		var arr []string

		for _, cell := range row.Cells {
			if idx == 0 {
				arr = append(arr, strings.ToLower(cell.String()))
			} else {
				arr = append(arr, strings.Replace(cell.String(), "'", "", -1))
			}
		}

		resp = append(resp, arr)
	}

	return resp
}
