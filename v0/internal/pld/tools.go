package pld

import (
	"strings"
)

func WrapText(str string, maxLenght int) string {
	strList := strings.Split(str, "\n")
	for i, s := range strList {
		for j := 1; len(s) > (j * maxLenght); j++ {
			for k := (j * maxLenght) - 1; k >= 0; k-- {
				if k == ((j-1)*maxLenght - 1) {
					break
				}
				if strList[i][k] == ' ' {
					tmp := []byte(strList[i])
					tmp[k] = '\n'
					strList[i] = string(tmp)
					break
				}
			}
			//				strings.
			//				strList[i] = s
		}
	}
	return strings.Join(strList, "\n")
}
