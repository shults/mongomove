package config

type FakeReader struct {
	cnf Config
}

func (f *FakeReader) Read(path string) (*Config, error) {
	return &f.cnf, nil
}

func (f *FakeReader) SetConfig(cnf Config) {
	f.cnf = cnf
}

func NewFakeReader(cnf Config) *FakeReader {
	return &FakeReader{
		cnf: cnf,
	}
}
