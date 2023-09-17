package wenv

import "fmt"

// NewEnvironmentMatcher returns a new EnvironmentMatcher instance.
//
// Example:
//   envResults := NewEnvironmentMatcher().
//     AddGroup(NewMatchGroup("databases").
//       AddMatcher(NewWrappedMatcher("name", "DB_", "_NAME"), true).
//       AddMatcher(NewWrappedMatcher("address", "DB_", "_ADDRESS"), true).
//       AddMatcher(NewWrappedMatcher("port", "DB_", "_PORT"), true).
//       AddMatcher(NewWrappedMatcher("user", "DB_", "_USER"), true).
//       AddMatcher(NewWrappedMatcher("pass", "DB_", "_PASSWORD"), true).
//       AddMatcher(NewWrappedMatcher("poolSize", "DB_, "_POOL_SIZE"), false),
//       true).
//     ParseEnv(SplitEnvironment(os.Environ()))
func NewEnvironmentMatcher() EnvironmentMatcher {
	return &environmentMatcher{
		groups:   make([]MatchGroup, 0, 8),
		required: make([]bool, 0, 8),
	}
}

type environmentMatcher struct {
	groups   []MatchGroup
	required []bool
}

func (e *environmentMatcher) AddGroup(group MatchGroup, required bool) EnvironmentMatcher {
	e.groups = append(e.groups, group)
	e.required = append(e.required, required)
	return e
}

func (e *environmentMatcher) ParseEnv(env map[string]string) EnvMatchResult {
	result := &envMatchResult{
		results: make(map[string]MatchGroupResults),
	}

	// Make an error slice to contain any errors we encounter.
	// We don't set this on the result right now so we don't waste mem on an empty
	// slice if there are no errors.  In the case that there _are_ errors, then we
	// will set this error list on the result.  In the case that there are not any
	// errors, then the result's errors ref will remain nil.
	errors := make([]error, 0, 8)

	for _, group := range e.groups {
		for k, v := range env {
			group.process(k, v)
		}

		res, err := group.result()
		if res.Size() > 0 {
			result.results[group.Name()] = res
		}

		errors = append(errors, err...)
		group.release()
	}

	// For each requirement flag
	for i, req := range e.required {
		// if the flag is true (the matching group is required)
		if req {
			// ensure that we have that group.  If we don't...
			if !result.Has(e.groups[i].Name()) {
				// record an error for it
				errors = append(errors, fmt.Errorf("no environment matches found for environment group %s", e.groups[i].Name()))
			}
		}
	}

	// If we had any errors, then set them on the result.
	if len(errors) > 0 {
		result.errors = errors
	}

	return result
}
