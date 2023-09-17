package wenv

// MatchGroupResults is a list of MatchGroupResult elements for all distinct
// MatchGroup key matches.
//
// For example, given the following environment variables:
//   DB_FOO_USERNAME=someUser
//   DB_FOO_PASSWORD=somePassword
//   DB_BAR_USERNAME=someUser
//   DB_BAR_PASSWORD=somePassword
//
// the MatchGroupResults list would contain 2 MatchGroupResult elements, one for
// the distinct key "FOO" and one for the distinct key "BAR".
type MatchGroupResults interface {
	// Size returns the count of result groups in this MatchGroupResults list.
	Size() int

	// Get returns the MatchGroupResult at the given index.
	Get(index int) MatchGroupResult
}

// MatchGroupResult represents the match results for a single instance of a
// group match.
type MatchGroupResult interface {
	// Size returns the number of matched keys in this MatchGroupResult.
	Size() int

	// Name returns the name given to the MatchGroup parent of this
	// MatchGroupResult.
	Name() string

	// Keys returns the distinct keys found for this MatchGroupResult.
	Keys() []string

	// FirstKey returns the first key from the Keys for this MatchGroupResult.
	FirstKey() string

	// Has tests whether this MatchGroupResult contains a result for the target
	// KeyMatcher name.
	Has(matcherName string) bool

	// Get returns the MatchResult for the target KeyMatcher name.
	Get(matcherName string) MatchResult

	// Value returns the environment value from the key matched by the named
	// KeyMatcher.
	Value(matcherName string) string

	// ValueOr returns the environment value from the key matched by the named
	// KeyMatcher, or returns the fallback value if the target KeyMatcher did not
	// match any keys.
	ValueOr(matcherName, fallback string) string
}
