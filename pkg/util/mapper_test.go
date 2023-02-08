package util

import (
	"testing"
)

const DefaultJsonPath = "digest-map.json"

func TestMapperWrite(t *testing.T) {
	mapper := NewMapper("")
	mapper.Data["a"] = "b"
	mapper.ToJsonFile()
}
