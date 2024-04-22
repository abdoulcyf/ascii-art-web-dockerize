package draw

import (
	"strings"
)

func StrMaker(cliStr string, patternMap map[byte][]string, chLength int) string {
	var finalStr string

	cliStrArr := strings.Split(cliStr, "\r\n")

	for _, strItem := range cliStrArr {
		if strItem == "" {
			finalStr += "\n"
			continue
		}

		for row := 1; row <= chLength; row++ {
			for _, v := range strItem {
				finalStr += patternMap[byte(v)][row-1]
			}
			finalStr += "\n"
		}
	}

	if ((len(cliStrArr)) == 2) && (cliStrArr[0] == "") && cliStrArr[1] == "" {
		finalStr = finalStr[:len(finalStr)-2]
	} else {

		finalStr = finalStr[:len(finalStr)-1]
	}

	return finalStr
}
