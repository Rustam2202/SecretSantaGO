package postgres

import "strconv"

func makeSQLArray(items []int) string {
	arrStr := "{"
	if len(items) > 0 {
		for i := 0; i < len(items)-1; i++ {
			arrStr += strconv.Itoa(items[i]) + ","
		}
		arrStr += strconv.Itoa(items[len(items)-1])
	}
	arrStr += "}"
	return arrStr
}
