package wenv

type envMatchResult struct {
	results map[string]MatchGroupResults
	errors  []error
}

func (e *envMatchResult) Size() int {
	return len(e.results)
}

func (e *envMatchResult) Has(groupName string) bool {
	_, ok := e.results[groupName]
	return ok
}

func (e *envMatchResult) Get(groupName string) MatchGroupResults {
	if res, ok := e.results[groupName]; ok {
		return res
	} else {
		return nil
	}
}

func (e *envMatchResult) Errors() MatcherErrors {
	return e.errors
}
