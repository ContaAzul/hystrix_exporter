package hystrix

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	for _, f := range []string{"command", "threadpool"} {
		t.Run(f, func(t *testing.T) {
			var assert = assert.New(t)
			bts, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s.json", f))
			assert.NoError(err)
			data, err := Unmarshal(string(bts))
			assert.NoError(err)
			assert.NotNil(data)
			assert.NotEmpty(data.Type)
		})
	}
}

// SomeData is used to properly benchmark Unmarshal func.
var SomeData Data

func BenchmarkUnmarshal(b *testing.B) {
	bytes, err := ioutil.ReadFile("testdata/command.json")
	assert.NoError(b, err)

	for i := 0; i < b.N; i++ {
		var data Data
		data, err := Unmarshal(string(bytes))
		assert.NoError(b, err)
		SomeData = data
	}
}
