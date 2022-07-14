package gojq_test

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/smartcontractkit/gojq"
)

func TestPreview(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected string
	}{
		{
			nil,
			"null",
		},
		{
			false,
			"false",
		},
		{
			true,
			"true",
		},
		{
			0,
			"0",
		},
		{
			3.14,
			"3.14",
		},
		{
			math.NaN(),
			"null",
		},
		{
			math.Inf(1),
			"1.7976931348623157e+308",
		},
		{
			math.Inf(-1),
			"-1.7976931348623157e+308",
		},
		{
			big.NewInt(9223372036854775807),
			"9223372036854775807",
		},
		{
			new(big.Int).SetBytes([]byte("\x0c\x9f\x2c\x9c\xd0\x46\x74\xed\xea\x3f\xff\xff\xff")),
			"999999999999999999999999999999",
		},
		{
			new(big.Int).SetBytes([]byte("\x0c\x9f\x2c\x9c\xd0\x46\x74\xed\xea\x40\x00\x00\x00")),
			"10000000000000000000000000 ...",
		},
		{
			"0 1 2 3 4 5 6 7 8 9 10 11 12",
			`"0 1 2 3 4 5 6 7 8 9 10 11 12"`,
		},
		{
			"0 1 2 3 4 5 6 7 8 9 10 11 12 13",
			`"0 1 2 3 4 5 6 7 8 9 10 1 ..."`,
		},
		{
			"０１２３４５６７８９",
			`"０１２３４５６７ ..."`,
		},
		{
			[]interface{}{},
			"[]",
		},
		{
			[]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			"[0,1,2,3,4,5,6,7,8,9,10,11,12]",
		},
		{
			[]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			"[0,1,2,3,4,5,6,7,8,9,10,1 ...]",
		},
		{
			[]interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{nil, nil, nil}}}}}}}},
			"[[[[[[[[null,null,null]]]]]]]]",
		},
		{
			[]interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{nil, nil, nil, nil}}}}}}}},
			"[[[[[[[[null,null,null,nu ...]",
		},
		{
			map[string]interface{}{},
			"{}",
		},
		{
			map[string]interface{}{"0": map[string]interface{}{"1": map[string]interface{}{"2": map[string]interface{}{"3": []interface{}{nil}}}}},
			`{"0":{"1":{"2":{"3":[null]}}}}`,
		},
		{
			map[string]interface{}{"0": map[string]interface{}{"1": map[string]interface{}{"2": map[string]interface{}{"3": map[string]interface{}{"4": map[string]interface{}{}}}}}},
			`{"0":{"1":{"2":{"3":{"4": ...}`,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.value), func(t *testing.T) {
			got := gojq.Preview(tc.value)
			if got != tc.expected {
				t.Errorf("Preview(%v): got %s, expected %s", tc.value, got, tc.expected)
			}
		})
	}
}
