package fastconv

import "testing"

func TestEncode_Slice(t *testing.T) {
	encodeSuccess(t, []*bool{}, GetBuffer().Slice("", 0, 0))
	encodeSuccess(t, []interface{}{}, GetBuffer().Slice("", 0, 0))
	encodeSuccess(t, []*bool{nil, Ptr(false), nil, Ptr(true)}, GetBuffer().Slice("", 4, 1).Nil("").Bool("", false).Nil("").Bool("", true))
	encodeSuccess(t, []interface{}{nil, Ptr(false), nil, Ptr(3)}, GetBuffer().Slice("", 4, 1).Nil("").Bool("", false).Nil("").Int("", 3))
}
