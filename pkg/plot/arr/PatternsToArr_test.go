package arr

import (
	"reflect"
	"testing"
	//"fmt"
)

// ======================================================
func Test_PatternsToArr(t *testing.T) {
	t.Skip()
	// ------------
	t.Run("1-return ", func(t *testing.T) {
		// -------setup-----

		inpStr := "\n---\n===\n+++\n\n***\n&&&\n"
		//fmt.Print("inpStr=\n",inpStr)
		// -------ACT--------
		outArr := PatternsToArr(inpStr)
		// ------expectation------
		expArr := []string{
			"---",
			"===",
			"+++",
			"***",
			"&&&",
		}
		//---------------------------
		assert_PatternsToArr(t, outArr, expArr)
	})
}

// =============assertion function ============
func assert_PatternsToArr(t testing.TB, outArr, expArr []string) {
	t.Helper()
	if !reflect.DeepEqual(outArr, expArr) {
		t.Errorf("\n>>>receive:%v<< \n>>>>expect:%v<<", outArr, expArr)
	}
}

// ===============================================
