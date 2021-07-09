package mkTable

import (
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
	for _, row := range sheet.Rows {
		if row.Cells[0].String() != "COLUMN_NAME" {
			var arr []string
			arr = append(arr, row.Cells[0].String())
			arr = append(arr, row.Cells[1].String())
			arr = append(arr, row.Cells[2].String())
			arr = append(arr, row.Cells[3].String())
			arr = append(arr, row.Cells[4].String())
			resp = append(resp, arr)
		}
	}

	return resp
}
