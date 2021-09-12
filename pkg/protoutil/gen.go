package protoutil

type GeneratorFunc func() error

func (f GeneratorFunc) Generate() error {
	return f()
}
