package util

import (
	"errors"
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

func ContentToMap(content string, chLength int) (map[byte][]string, error) {
	contentArr := PatternsToArr(content)
	contentMap, err := ToMap(contentArr, chLength)
	return contentMap, err
}

func ToMap(strArr []string, chLength int) (map[byte][]string, error) {
	if len(strArr) < 94*chLength {
		return nil, errors.New("not enough lines")
	}
	resultMap := make(map[byte][]string)
	for i := ' '; i <= '~'; i++ {
		startIndex := int(i-' ') * (chLength)
		endIndex := startIndex + chLength
		resultMap[byte(i)] = strArr[startIndex:endIndex]
	}
	return resultMap, nil
}

func PatternsToArr(str string) []string {
	str = strings.Trim(str, "\n")
	str = strings.ReplaceAll(str, "\n\n", "\n")
	chArr := strings.Split(str, "\n")
	return chArr
}
