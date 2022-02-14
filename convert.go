package heroyaml

type Converter interface {
	Convert(str string) (string, error)
}
