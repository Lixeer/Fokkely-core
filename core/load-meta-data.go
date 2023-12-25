package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Skill struct {
	Item map[string]interface{}
	// List []string
}

type SkillMap struct {
	DataMap
	Item map[string]map[string]interface{}
}

func (skm *SkillMap) Load() error {

	fileContent, err := ioutil.ReadFile("core/meta-data/skill.json")

	//fmt.Print(ct)
	if err != nil {
		return err
	}

	// 解析 JSON 数据

	//fmt.Println("skills ", skills)
	err = json.Unmarshal(fileContent, &skm.Item) //test pass
	fmt.Println("tip", skm.Item)
	if err != nil {
		return err
	}

	return nil

}

func (skm *SkillMap) Get(k string) (map[string]interface{}, error) {

	return skm.Item[k], nil

}
