package wenv

type matchGroupResults []MatchGroupResult

func (m matchGroupResults) Size() int {
	return len(m)
}

func (m matchGroupResults) Get(index int) MatchGroupResult {
	return m[index]
}

func newMatchGroupResult(name string, keys []string, results map[string]MatchResult) MatchGroupResult {
	return &matchGroupResult{
		results: results,
		name:    name,
		keys:    keys,
	}
}

type matchGroupResult struct {
	results map[string]MatchResult
	name    string
	keys    []string
}

func (m *matchGroupResult) Size() int {
	return len(m.results)
}

func (m *matchGroupResult) Name() string {
	return m.name
}

func (m *matchGroupResult) Keys() []string {
	return m.keys
}

func (m *matchGroupResult) FirstKey() string {
	return m.keys[0]
}

func (m *matchGroupResult) Has(matcherName string) bool {
	_, ok := m.results[matcherName]
	return ok
}

func (m *matchGroupResult) Get(matcherName string) MatchResult {
	return m.results[matcherName]
}

func (m *matchGroupResult) Value(matcherName string) string {
	return m.results[matcherName].Value()
}

func (m *matchGroupResult) ValueOr(matcherName, fallback string) string {
	if r, ok := m.results[matcherName]; ok {
		return r.Value()
	} else {
		return fallback
	}
}
