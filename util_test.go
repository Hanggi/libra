package libra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormUtil(t *testing.T) {
	// var vali Context.Validate
	vali := new(Context).Validate

	var (
		test1 bool
		test2 bool
		// test3 bool
	)

	test1 = vali.IsNumber("123")
	assert.Equal(t, test1, true, "IsNumber should be true!")
	test2 = vali.IsNumber("abcd")
	assert.Equal(t, test2, false, "IsNumber should be false!")
	// test3 = vali.IsNumber("123abc")
	// assert.Equal(t, test3, false, "IsNumber should be false!")

	test1 = vali.IsHan("汉字")
	assert.Equal(t, test1, true, "IsHan should be true!")
	test2 = vali.IsHan("abc123")
	assert.Equal(t, test2, false, "IsHan should be false!")
	// test3 = vali.IsHan("汉字abc123")
	// assert.Equal(t, test3, false, "IsHan should be false!")

	test1 = vali.IsEng("abc")
	assert.Equal(t, test1, true, "IsEng should be true!")
	test2 = vali.IsEng("123")
	assert.Equal(t, test2, false, "IsEng should be false!")
	// test3 = vali.IsEng("汉字abc123")
	// assert.Equal(t, test3, false, "IsEng should be false!")

	test1 = vali.IsEmail("abc@def.ghl")
	assert.Equal(t, test1, true, "IsEmail should be true!")
	test2 = vali.IsEmail("www.google.com")
	assert.Equal(t, test2, false, "IsEmail should be false!")
	// test3 = vali.IsEmail("vv@1.2")
	// assert.Equal(t, test3, false, "IsEmail should be false!")

}
