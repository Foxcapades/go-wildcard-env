package wenv

// MatchGroup defines a group of KeyMatchers that belong together, optionally
// requiring some or all of the KeyMatchers to be hit for a specific key.
//
// Example Env Group:
//   DB_EXAMPLE_1_USERNAME=someUser
//   DB_EXAMPLE_1_PASSWORD=somePassword
//   DB_EXAMPLE_2_USERNAME=someUser
//   DB_EXAMPLE_2_PASSWORD=somePassword
//
// For the above example environment variables, the keys would be matched using
// a wrapped KeyMatcher (NewWrappedKeyMatcher) with the prefix "DB_" and the
// suffixes "_USERNAME" and "_PASSWORD".  This would be parsed into a
// MatchGroupResult that contains matches for the key components "EXAMPLE_1" and
// "EXAMPLE_2".
type MatchGroup interface {
	// Name returns the name of this MatchGroup.  MatchGroup names are used to
	// reference/look up the matched results in the EnvMatchResult.
	Name() string

	// AddMatcher adds a new KeyMatcher to this MatchGroup.
	AddMatcher(matcher KeyMatcher, required bool) MatchGroup

	// process processes the given environment key and value.
	process(key, val string) bool

	// result returns the processing results of the given environment entries.
	result() (MatchGroupResults, []error)

	// release releases resources held by this MatchGroup.
	release()
}
