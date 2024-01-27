package fastconv

import (
	"testing"
)

func TestEncode_Interface(t *testing.T) {
	m := map[int]interface{}{0: nil}
	encodeSuccess(t, m, GetBuffer().Map("", 1, 1).Nil("0"))
}
