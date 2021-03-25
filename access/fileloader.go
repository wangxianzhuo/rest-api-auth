package access

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type FilePermissionLoader struct {
	fileName string
}

func (l FilePermissionLoader) PermissioLoad(id string) ([]string, error) {
	b, err := ioutil.ReadFile(l.fileName)
	if err != nil {
		return nil, err
	}
	log.Println("success loading " + l.fileName)

	var r map[string][]string
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	log.Println("success parsing " + l.fileName)

	return r[id], nil
}

func (l FilePermissionLoader) Close() error {
	return nil
}
