package draw

import "testing"

// ======================================================
func Test_StrMaker(t *testing.T) {
	t.Skip()
	// ------------
	t.Run("1-return ", func(t *testing.T) {
		// -------setup-----
		cliStr := "ab"
		chLength := 2
		inpMap := map[byte][]string{
			'a': {"aaa", "AAA"},
			'b': {"bbb", "BBB"},
		}
		// -------ACT--------
		outStr := StrMaker(cliStr, inpMap, chLength)
		// ------expectation------
		expStr := "aaabbb\nAAABBB\n"
		assert_StrMaker(t, outStr, expStr)
	})
}

// =============assertion function ============
func assert_StrMaker(t testing.TB, outStr, expStr string) {
	t.Helper()
	if outStr != expStr {
		t.Errorf("\n>>>>>receive:\n%v-expect:\n%v\n", outStr, expStr)
	}
}

// ===============================================
