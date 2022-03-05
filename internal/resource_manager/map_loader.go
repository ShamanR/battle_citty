package resource_manager

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/shamanr/battle_citty/internal/consts"
	"io/ioutil"
)

type JsonMap struct {
	Cells consts.LevelMap `json:"cells"`
}

func (s *resourceManager) LoadMap(path string) consts.LevelMap {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(errors.Wrap(err, "Unable to load map"))
	}

	jsonMap := JsonMap{}
	err = json.Unmarshal(content, &jsonMap)
	if err != nil {
		panic(errors.Wrap(err, "Unable to load map"))
	}

	return jsonMap.Cells
}
