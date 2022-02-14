package heroyaml

import (
	"fmt"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/tidwall/gjson"
)

type k8sEnv struct {
	Name, Value string
}

func (k k8sEnv) String() string {
	return fmt.Sprintf(`
- name: "%s"
  value: "%s"
`, k.Name, k.Value)
}

type K8sEnvConverter struct{}

var _ Converter = (*K8sEnvConverter)(nil)

func (k K8sEnvConverter) Convert(str string) (string, error) {
	data, err := yaml.YAMLToJSON([]byte(str))
	if err != nil {
		return "", err
	}

	json := gjson.Parse(string(data))
	printJSON(json, "")

	return "", err
}

func printJSON(json gjson.Result, preKey string) {
	json.ForEach(func(key, value gjson.Result) bool {
		k := key.String()
		if len(preKey) > 0 {
			k = fmt.Sprintf("%s.%s", preKey, k)
		}

		if value.IsObject() {
			printJSON(value, k)
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

		printKV(k, v)
		return true
	})
}

func printKV(k, v string) {
	k = strings.ToUpper(k)
	fmt.Printf(`
- name: "%s"
  value: "%s"
`, k, v)
}
