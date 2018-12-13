package internal

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Basepath   = filepath.Dir(filepath.Dir(b))
)

func LoadTestData(path string) ([]byte, error) {
	p := filepath.Join(Basepath, path)
	return ioutil.ReadFile(p)
}
