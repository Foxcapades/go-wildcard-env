package wenv

type matchResult struct {
	raw   string
	value string
}

func (m *matchResult) Raw() string {
	return m.raw
}

func (m *matchResult) Value() string {
	return m.value
}
