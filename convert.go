package mo

type Converter interface {
	Convert(str string) (string, error)
}
