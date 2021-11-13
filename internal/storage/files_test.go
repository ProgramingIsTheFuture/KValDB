package storage

import (
	"KValDB/internal/model"
	"encoding/json"
	"fmt"
	"testing"
)

const TestFolder = "test_files"

func TestReadFile(t *testing.T) {
	v, err := ReadFile(fmt.Sprintf("../../%s/not_exist_file.json", TestFolder))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log(v)
}

func TestWriteFile(t *testing.T) {

	var kval = model.KValDBType{
		"key1": "val1",
		"key2": "val2",
	}

	b, _ := json.Marshal(kval)

	err := WriteFile(fmt.Sprintf("../../%s/not_exist_file.json", TestFolder), b)
	if err != nil {
		t.Fail()
	}
}

func TestUnMarshal(t *testing.T) {
	var kval = model.KValDBType{}
	err := UnMarshal(fmt.Sprintf("../../%s/not_exist_file.json", TestFolder), &kval)

	if err != nil {
		t.Fail()
	}

	t.Log(kval)
}

/*
func TestDeleteFiles(t *testing.T) {
	os.RemoveAll(fmt.Sprintf("../../%s/", TestFolder))
}
*/
