package wenv

// An EnvironmentMatcher may be used to parse construct a series of MatchGroups
// and parse a given environment map.
type EnvironmentMatcher interface {
	// AddGroup adds a matching group to this EnvironmentMatcher instance.
	//
	// If the MatchGroup is required, then at least one hit of the given
	// MatchGroup will be expected in the parsed environment, and if the group
	// does not match any environment keys, the EnvMatchResult returned by
	// ParseEnv will contain an error for the MatchGroup.
	AddGroup(group MatchGroup, required bool) EnvironmentMatcher

	// ParseEnv parses the given environment map against the configured
	// MatchGroups.
	ParseEnv(env map[string]string) EnvMatchResult
}
