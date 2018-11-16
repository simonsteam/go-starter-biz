package utl_test

import (
	"github.com/stretchr/testify/assert"
	"local/biz/utl"
	"testing"
)

func TestSliceDistinct(t *testing.T) {
	type ts struct {
		A string
		B int
	}

	cases := []struct {
		param    *[]interface{}
		shouldBe *[]interface{}
	}{
		{
			param:    &[]interface{}{1, 2, 3, 4, 2, 5},
			shouldBe: &[]interface{}{1, 2, 3, 4, 5},
		},
		{
			param:    &[]interface{}{"233", "666", "666", "777"},
			shouldBe: &[]interface{}{"233", "666", "777"},
		},
		{
			param:    &[]interface{}{ts{"A", 123}, 233, "555", "233", 233, "555", ts{"A", 123}},
			shouldBe: &[]interface{}{ts{"A", 123}, 233, "555", "233"},
		},
	}

	for _, c := range cases {
		rs := utl.SliceDistinct(c.param)
		assert.Equal(t, len(*c.shouldBe), len(*rs))
		assert.Equal(t, c.shouldBe, rs, "they should be equal")
	}
}
