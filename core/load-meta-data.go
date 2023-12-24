package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Skill struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// List []string
}

type SkillMap struct {
	DataMap
	Item map[string]Skill
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

func (skm *SkillMap) Get(k string) (Skill, error) {

	return skm.Item[k], nil

}

func GetSkillData() (map[string]Skill, error) {
	// 读取文件内容
	fileContent, err := ioutil.ReadFile("core/meta-data/skill.json")

	//fmt.Print(ct)
	if err != nil {
		return nil, err
	}

	// 解析 JSON 数据
	var skills map[string]Skill
	//fmt.Println("skills ", skills)
	err = json.Unmarshal(fileContent, &skills)
	if err != nil {
		return nil, err
	}

	return skills, nil
}
