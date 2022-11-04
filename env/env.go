package env

import (
	"fmt"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/tidwall/gjson"
)

type env struct {
	Name, Value string
}

func (e env) String() string {
	name := strings.ToUpper(e.Name)
	return fmt.Sprintf("%s=%s", name, e.Value)
}

func PrintEnv(data []byte, ext, prefix string) error {
	var json gjson.Result
	switch ext {
	case "yaml", "yml":
		jsonData, err := yaml.YAMLToJSON(data)
		if err != nil {
			return err
		}

		json = gjson.Parse(string(jsonData))
	case "json":
		json = gjson.Parse(string(data))
	default:
		return fmt.Errorf(".%s type does not support", ext)
	}

	printEnv(json, prefix)
	return nil
}

func printEnv(json gjson.Result, preKey string) {
	json.ForEach(func(key, value gjson.Result) bool {
		k := key.String()
		if len(preKey) > 0 {
			k = fmt.Sprintf("%s_%s", preKey, k)
		}

		if value.IsObject() {
			printEnv(value, k)
			return true
		}

		v := value.String()
		if value.IsArray() {
			arr := value.Array()
			strs := make([]string, len(arr))
			for i, result := range arr {
				strs[i] = result.String()
			}

			v = strings.Join(strs, ",")
		}

		kv := env{
			Name:  k,
			Value: v,
		}
		fmt.Println(kv)
		return true
	})
}
