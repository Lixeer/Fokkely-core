package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Skill struct {
	ID   string `json:"name"`
	Name string `json:"name"`
}

func GetSkillData(skillID string) (*Skill, error) {
	file, err := os.Open("meta-skill/skill.json")
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 读取文件内容
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("无法读取文件内容: %v", err)
	}

	// 创建一个Person结构体实例
	var skill Skill

	// 解析JSON数据到结构体
	err = json.Unmarshal(data, &skill)
	if err != nil {
		return nil, fmt.Errorf("解析JSON时发生错误: %v", err)
	}

	return &skill, nil

}
