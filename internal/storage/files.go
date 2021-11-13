package storage

import (
	"KValDB/internal/model"
	"encoding/json"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}
	f, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var b = make([]byte, 2048)

	n, err := f.Read(b)
	if err != nil {
		return nil, err
	}

	return b[:n], nil
}

// Write to the file
// Receivees the db name - "PATH" and the value you want to write
// This will overwrite the whole file
func WriteFile(name string, val []byte) error {
	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write([]byte(string(val) + "\n"))
	return err
}

// Receives the db name - "PATH" and the model you want to populate with data
// This can be used to replace the ReadFile
func UnMarshal(db string, v *model.KValDBType) error {
	b, err := ReadFile(db)
	if err != nil {
		return nil
	}

	return json.Unmarshal(b, v)
}
