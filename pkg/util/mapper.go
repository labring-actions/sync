package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gogf/gf/os/glog"
)

type Mapper struct {
	FilePath string
	Data     map[string]string
}

func NewMapper(jsonPath string) Mapper {
	mapper := Mapper{FilePath: jsonPath, Data: map[string]string{}}
	err := mapper.FromJsonFile()
	if err != nil {
		glog.Error(err)
		return mapper
	}
	return mapper
}

func (m *Mapper) ToJsonFile() error {
	if m.FilePath == "" {
		return fmt.Errorf("no outpul file path")
	}
	jsonData, err := json.Marshal(m.Data)
	if err != nil {
		return err
	}
	jsonFile, err := os.Create(m.FilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	if _, err = jsonFile.Write(jsonData); err != nil {
		return err
	}
	return nil
}

func (m *Mapper) FromJsonFile() error {
	if m.FilePath == "" {
		return fmt.Errorf("no input file")
	}
	jsonFile, err := os.Open(m.FilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	jsonData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(jsonData, &m.Data)
	if err != nil {
		return err
	}
	return nil
}
