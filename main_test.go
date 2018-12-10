package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkReport(b *testing.B) {
	cluster := "foo"
	bytes, err := ioutil.ReadFile("hystrix/testdata/command.json")
	assert.NoError(b, err)
	line := string(bytes)

	for i := 0; i < b.N; i++ {
		report(cluster, line)
	}
}
