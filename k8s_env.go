package heroyaml

import (
	"fmt"

	"gopkg.in/yaml.v2"
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
	t := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(str), &t)
	if err != nil {
		return "", err
	}

	fmt.Println(t)
	return "", err
}

func (k K8sEnvConverter) convert(m map[string]interface{}) ([]string, error) {
	for key, val := range m {
		switch val.(type) {
		case []int:

		}
	}
}
