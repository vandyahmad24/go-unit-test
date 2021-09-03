package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTableHelloWorld(t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "HelloWorld(Vandy)",
			request: "vandy",
			expected: "Hello vandy",

		},
		{
			name: "HelloWorld(Ahmad)",
			request: "ahmad",
			expected: "Hello ahmad",
		},
	}

	for _, v := range tests{
		t.Run(v.name, func(t *testing.T) {
			result:=HelloWorld(v.request)
			assert.Equal(t, v.expected,result)
		})
	}
}

