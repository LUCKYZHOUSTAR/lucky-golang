package config

import (
	"encoding/xml"
	"fmt"
	"go/frame/errors"
	x "go/frame/xml"
	"io/ioutil"
	"os"
	"strings"
)

type profileConfig struct {
	active   []string
	settings SettingMap
}

//判断结尾的文件是否生产配置
func (c *profileConfig) isOnlineEnv() bool {

	for _, p := range c.active {
		if strings.HasSuffix(p, "-std") || strings.HasSuffix(p, "-prd") {
			return true
		}
	}

	return false
}

func (c *profileConfig) filter(s string) string {
	if c.settings == nil {
		return s
	}

	return os.Expand(s, func(n string) string {
		if v, ok := c.settings[n]; ok {
			return v
		}

		return "${" + n + "}"
	})

}

func (c *profileConfig) loadXML(filePath string) (*x.Document, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	text := c.filter(string(bytes))

	doc := x.New()
	if err := doc.LoadString(text, nil); err != nil {
		return nil, err
	}

	return doc, nil
}

func (c *profileConfig) load(content []byte) error {
	// load profile settings
	type Profiles struct {
		Profiles []struct {
			Name     string `xml:"name,attr"`
			Settings []struct {
				Key   string `xml:"key,attr"`
				Value string `xml:"value,attr"`
			} `xml:"setting"`
		} `xml:"profile"`
	}

	var profiles Profiles
	err := xml.Unmarshal(content, &profiles)
	if err != nil {
		return errors.NewE("parse profile.conf failed", err)
	}

	fn := func(pn string) {
		for _, pf := range profiles.Profiles {
			if pf.Name == pn {
				for _, s := range pf.Settings {
					c.settings[s.Key] = s.Value
				}
				break
			}
		}
	}

	fmt.Println("config > load profile settings...")
	c.settings = SettingMap{}
	// load default profile first
	fn("")
	for _, n := range c.active {
		fn(n)
	}
	for k, v := range c.settings {
		fmt.Printf("\t%v=%v\n", k, v)
	}
	return nil
}
