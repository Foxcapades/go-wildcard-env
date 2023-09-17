package wenv

// MatchResult represents an individual matched environment variable and value.
//
// This type contains methods to retrieve the raw environment name and the value
// of the environment variable.
type MatchResult interface {

	// Raw returns the whole matched environment variable name.
	Raw() string

	// Value returns the value of the matched environment variable.
	Value() string
}
