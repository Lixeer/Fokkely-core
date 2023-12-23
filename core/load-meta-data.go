package core

import (
	"encoding/json"
	"io/ioutil"
)

type Skill struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// List []string
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
