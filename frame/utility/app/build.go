package app

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

// BuildInfo 程序构建信息
type BuildInfo struct {
	ProductVersion string `json:"productVersion"`
	ScmRevision    string `json:"scmRevision"`
	ScmBranch      string `json:"scmBranch"`
	ScmUser        string `json:"scmUser"`
	ScmMessage     string `json:"scmMessage"`
	ScmTime        string `json:"scmTime"`
	BuildAt        string `json:"buildAt"`
	BuildBy        string `json:"buildBy"`
	BuildOn        string `json:"buildOn"`
}

func (b *BuildInfo) String() string {
	bytes, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// GetBuildInfo 获取程序构建信息
func GetBuildInfo() (*BuildInfo, error) {
	dir, err := GetAppFolder()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, "build.json")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	bi := &BuildInfo{}
	err = json.Unmarshal(bytes, bi)
	if err != nil {
		return nil, err
	}
	return bi, nil
}
