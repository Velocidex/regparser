package appcompatcache

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/sebdah/goldie"
	"github.com/stretchr/testify/assert"
)

var (
	fixtures = []string{
		//"Win10Creators.bin",
		//		"Win10.bin",
		//		"Win81.bin",
		"Win80.bin",
	}
)

func TestCacheParsing(t *testing.T) {
	for _, fixture := range fixtures {
		fd, err := os.Open(filepath.Join("test_data", fixture))
		assert.NoError(t, err)

		buffer, err := ioutil.ReadAll(fd)
		assert.NoError(t, err)

		entries := ParseValueData(buffer)
		serialized, err := json.MarshalIndent(entries[:10], " ", " ")
		assert.NoError(t, err)

		goldie.Assert(t, "appcompatcache_"+fixture, serialized)
	}
}
