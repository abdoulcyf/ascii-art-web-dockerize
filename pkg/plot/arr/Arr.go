package arr

import (
	//"fmt"
	//"fmt"
	"errors"
	//"fmt"
	"strings"
	//"bufio"
	//"fmt"
)

// ---------------------------
//
//	type bannerFile struct {
//		firstCh byte =
//		lastCh byte
//		chLength int
//	}
//
// ------------------------
func ContentToMap(content string, chLength int, firstCh, lastCh byte) (map[byte][]string, error) {

	contentArr := PatternsToArr(content)
	contentMap, err := ArrToMap(contentArr, chLength, firstCh, lastCh)
	return contentMap, err
}

// ---------------------------------------------
func ArrToMap(strArr []string, chLength int, firstCh, lastCh byte) (map[byte][]string, error) {
	resultMap := make(map[byte][]string)
	// ascii code start = 32 space
	// ascii code end = 126
	//var firstCh byte = 'a' //' ' //space = 32
	//var lastCh byte = 'b'  //'~' //last ~ = 126
	howManyCh := int(lastCh-firstCh) + 1
	howManyLine := howManyCh * chLength

	//fmt.Println(len(strArr), howManyLine)

	if !(howManyLine <= len(strArr)) {
		return nil, errors.New("no enough line of")
	}

	n := 0
	for i := firstCh; i <= lastCh; i++ {
		//------------------------------
		resultMap[i] = strArr[n : n+chLength]
		n += chLength
		//-----------------------------
	}
	return resultMap, nil
}

// ------------------------
func PatternsToArr(str string) []string {

	str = strings.Trim(str, "\n")
	str = strings.ReplaceAll(str, "\n\n", "\n")

	//fmt.Print("s>>>>\n",str)
	chArr := strings.Split(str, "\n")
	//fmt.Print("chArr=\n",chArr)
	//fmt.Println(">>len",len(chArr))
	// fmt.Println("chArr[0]",chArr[0])
	return chArr
}

//------------------------
