package hystrix

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	for _, f := range []string{"command", "threadpool", "ping"} {
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
