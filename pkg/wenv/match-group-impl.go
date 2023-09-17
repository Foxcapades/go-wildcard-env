package wenv

import "fmt"

var merger = newKeyMerger()

func NewMatchGroup(name string) MatchGroup {
	return &matchGroup{
		name:     name,
		matchers: make([]KeyMatcher, 0, 8),
		required: make([]bool, 0, 8),
		results:  newMatchGroupMap(),
	}
}

// // // // // // // // // // // // // // // // // // // // // // // // // // //
//
//    Match Group Map
//
// // // // // // // // // // // // // // // // // // // // // // // // // // //

func newMatchGroupMap() (out matchGroupMap) {
	out.mp = make(map[string]map[string]MatchResult, 8)
	out.keys = make(map[string][]string, 8)
	return
}

type matchGroupMap struct {
	mp   map[string]map[string]MatchResult
	keys map[string][]string
}

func (m *matchGroupMap) put(keys []string, matcherName string, result MatchResult) {
	mergedKey := merger.merge(keys)

	m.keys[mergedKey] = keys

	if mp, ok := m.mp[mergedKey]; ok {
		mp[matcherName] = result
	} else {
		mp := make(map[string]MatchResult, 8)
		mp[matcherName] = result
		m.mp[mergedKey] = mp
	}
}

func (m *matchGroupMap) release() {
	m.mp = nil
	m.keys = nil
}

// // // // // // // // // // // // // // // // // // // // // // // // // // //
//
//    Match Group
//
// // // // // // // // // // // // // // // // // // // // // // // // // // //

type matchGroup struct {
	name     string
	matchers []KeyMatcher
	required []bool

	// results is a map of merged keys to maps of KeyMatcher names to match
	// results.
	results matchGroupMap
}

func (m *matchGroup) Name() string {
	return m.name
}

func (m *matchGroup) AddMatcher(matcher KeyMatcher, required bool) MatchGroup {
	m.matchers = append(m.matchers, matcher)
	m.required = append(m.required, required)
	return m
}

func (m *matchGroup) process(key, val string) (matched bool) {
	for _, km := range m.matchers {
		if km.Matches(key) {
			m.results.put(km.Process(key), km.Name(), &matchResult{key, val})
			matched = true
		}
	}

	return
}

func (m *matchGroup) result() (MatchGroupResults, []error) {
	results := make([]MatchGroupResult, 0, len(m.results.mp))

	for mergedKey, keyMatchers := range m.results.mp {
		keys := m.results.keys[mergedKey]
		results = append(results, newMatchGroupResult(m.name, keys, keyMatchers))
	}

	errors := make([]error, 0, 8)

	// Iterate through all the keys
	for i, req := range m.required {
		// filter down to only those that are required
		if req {
			// iterate through the result groups
			for _, res := range results {
				// If the result doesn't have a match for the required key
				if !res.Has(m.matchers[i].Name()) {
					errors = append(errors, fmt.Errorf("match group %s (keys: %s) does not have a match for required key %s", m.name, merger.merge(res.Keys()), m.matchers[i].Name()))
				}
			}
		}
	}

	return matchGroupResults(results), errors
}

func (m *matchGroup) release() {
	m.matchers = nil
	m.required = nil
	m.results.release()
}
