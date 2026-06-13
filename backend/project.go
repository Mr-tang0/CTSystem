package backend

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// CTProject CT项目配置
type CTProject struct {
	TotalScanAngle float32 `json:"total_scan_angle"` //扫描角度，扫描累计角度
	AngularStep    float32 `json:"angular_step"`     //扫描角度步长
}

type Project struct {
	UserName  string    `json:"user_name"`
	ProjectID string    `json:"project_id"`
	CTProject CTProject `json:"ct_project"` // CT项目
	FileName  string    `json:"file_name"`
	FilePath  string    `json:"file_path"`
}

func NewProject() *Project {
	project := &Project{
		UserName:  "user",
		ProjectID: time.Now().Format("20060102150405"),
		CTProject: CTProject{
			TotalScanAngle: 360,
			AngularStep:    1,
		},
		FileName: "CT_" + time.Now().Format("20060102150405"),
		FilePath: "./projects/",
	}

	project.LoadHistoryProject()
	return project
}

// 从本地读取历史项目文件
func (p *Project) LoadHistoryProject() *Project {
	file := os.Getenv("USERPROFILE") + "/NIMTE/CT/project.json"
	fmt.Println("[Project] 加载历史项目文件:", file)

	// 确保目录存在
	err := os.MkdirAll(os.Getenv("USERPROFILE")+"/NIMTE/CT", 0755)
	if err != nil {
		fmt.Println("[Project] 创建目录失败:", err.Error())
		return NewProject()
	}

	// 检查文件是否存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("[Project] 文件不存在，使用默认配置并保存")
		p.SaveHistoryProject(p)
		return p
	}

	// 读取文件
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("[Project] 读取文件失败:", err.Error())
		return p
	}

	// 检查文件是否为空
	if len(data) == 0 {
		fmt.Println("[Project] 文件为空，使用默认配置并保存")
		p.SaveHistoryProject(p)
		return p
	}

	// 解析 JSON
	err = json.Unmarshal(data, p)
	if err != nil {
		fmt.Println("[Project] 解析 JSON 失败:", err.Error())
		fmt.Println("[Project] 使用默认配置并保存")
		p.SaveHistoryProject(p)
		return p
	}

	fmt.Println("[Project] 加载历史项目成功:", p.ProjectID)
	return p
}

// 将项目信息保存到本地文件
func (p *Project) SaveHistoryProject(project *Project) *Project {
	file := os.Getenv("USERPROFILE") + "/NIMTE/CT/project.json"
	fmt.Println("[Project] 保存项目文件:", file)

	// 确保目录存在
	err := os.MkdirAll(os.Getenv("USERPROFILE")+"/NIMTE/CT", 0755)
	if err != nil {
		fmt.Println("[Project] 创建目录失败:", err.Error())
		return NewProject()
	}

	// 序列化 JSON
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("[Project] 序列化 JSON 失败:", err.Error())
		return p
	}

	// 写入文件
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		fmt.Println("[Project] 写入文件失败:", err.Error())
		return p
	}

	fmt.Println("[Project] 保存项目成功")
	return p
}
