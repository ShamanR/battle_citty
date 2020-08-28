package resource_manager

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shamanr/battle_citty/interfaces"
	"io/ioutil"
)

type JsonMap struct {
	Cells *interfaces.LevelMap `json:"cells"`
}

func (s *resourceManager) LoadMap(path string) *interfaces.LevelMap {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(errors.Wrap(err, "Unable to load map"))
	}

	test := string(content)
	fmt.Println(test)

	jsonMap := JsonMap{}
	err = json.Unmarshal(content, &jsonMap)
	if err != nil {
		panic(errors.Wrap(err, "Unable to load map"))
	}

	return jsonMap.Cells
}
