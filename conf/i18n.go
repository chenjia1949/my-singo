package conf

import (
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// 按语言初始化翻译文件
var Dictinary = make(map[string]*map[interface{}]interface{}, 2)

func InitLocales(lan, path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return err
	}
	Dictinary[lan] = &m
	return nil
}

// 自定义国际化
func Message(lan, key string) string {
	dic := *Dictinary[lan]
	keys := strings.Split(key, ".")
	for index, path := range keys {
		// 如果到达了最后一层，寻找目标翻译
		if len(keys) == (index + 1) {
			for k, v := range dic {
				if k, ok := k.(string); ok {
					if k == path {
						if value, ok := v.(string); ok {
							return value
						}
					}
				}
			}
			return path
		}
		// 如果还有下一层，继续寻找
		for k, v := range dic {
			if ks, ok := k.(string); ok {
				if ks == path {
					if dic, ok = v.(map[interface{}]interface{}); !ok {
						return path
					}
				}
			} else {
				return key
			}
		}
	}
	return key
}
