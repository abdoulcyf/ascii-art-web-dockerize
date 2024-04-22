package arr

import (
	"reflect"
	"testing"
	//"fmt"
)

// ======================================================
func Test_ArrToMap(t *testing.T) {
	t.Skip()
	// -------------------------------------
	t.Run("1-return ", func(t *testing.T) {
		// -------setup-----
		inpArr := []string{
			"---",
			"===",
			"+++",
			"***",
		}
		// -------ACT--------
		outMap, _ := ArrToMap(inpArr, 2, 'a', 'b')
		// ------expectation------
		expMap := map[byte][]string{
			'a': {"---", "==="},
			'b': {"+++", "***"},
		}
		assert_ArrToMap(t, outMap, expMap)
	})
	// -------------------------------------
	t.Run("2-return ", func(t *testing.T) {
		// -------setup-----
		inpArr := []string{
			"---",
			"===",
			"+++",
			"***",
		}
		// -------ACT--------
		outMap, _ := ArrToMap(inpArr, 1, 'a', 'd')
		// ------expectation------
		expMap := map[byte][]string{
			'a': {"---"},
			'b': {"==="},
			'c': {"+++"},
			'd': {"***"},
		}
		assert_ArrToMap(t, outMap, expMap)
	})

	////////////////////////////////////////////
}

// =============assertion function ============
func assert_ArrToMap(t testing.TB, outArr, expArr map[byte][]string) {
	t.Helper()
	if !reflect.DeepEqual(outArr, expArr) {
		t.Errorf("\n>>>receive:%v<< \n>>>>expect:%v<<", outArr, expArr)
	}
}

// ===============================================
