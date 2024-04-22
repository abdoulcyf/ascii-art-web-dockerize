package arr

import (
	"reflect"
	"testing"
	//"fmt"
)

// ======================================================
func Test_ContentToMap(t *testing.T) {
	//t.Skip()
	// -------------------------------------
	t.Run("1-return ", func(t *testing.T) {
		// -------setup-----
		inpStr := "\n---\n===\n+++\n\n***\n"

		// -------ACT--------
		outMap, _ := ContentToMap(inpStr, 2, 'a', 'b')
		// ------expectation------
		expMap := map[byte][]string{
			'a': {"---", "==="},
			'b': {"+++", "***"},
		}
		assert_ContentToMap(t, outMap, expMap)
	})
	// -------------------------------------
	t.Run("2-return ", func(t *testing.T) {
		// -------setup-----
		inpStr := "\n---\n===\n+++\n\n***\n"

		// -------ACT--------
		outMap, _ := ContentToMap(inpStr, 1, 'a', 'd')
		// ------expectation------
		expMap := map[byte][]string{
			'a': {"---"},
			'b': {"==="},
			'c': {"+++"},
			'd': {"***"},
		}
		assert_ContentToMap(t, outMap, expMap)
	})

	////////////////////////////////////////////
}

// =============assertion function ============
func assert_ContentToMap(t testing.TB, outMap, expMap map[byte][]string) {
	t.Helper()
	if !reflect.DeepEqual(outMap, expMap) {
		t.Errorf("\n>>>receive:%v<< \n>>>>expect:%v<<", outMap, expMap)
	}
}

// ===============================================
