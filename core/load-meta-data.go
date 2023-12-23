package core

import (
	"encoding/json"
	"io/ioutil"
)

type Skill struct {
	ID   string `json:"name"`
	Name string `json:"name"`
}

func GetSkillData() (map[string]Skill, error) {
	// 读取文件内容
	fileContent, err := ioutil.ReadFile("meta-data/skill.json")
	if err != nil {
		return nil, err
	}

	// 解析 JSON 数据
	var skills map[string]Skill
	err = json.Unmarshal(fileContent, &skills)
	if err != nil {
		return nil, err
	}

	return skills, nil
}
