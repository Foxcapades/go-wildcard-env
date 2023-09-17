package wenv

// EnvMatchResult contains the results of the environment matching.  The result
// is a map of group names to the results for the named MatchGroup.
type EnvMatchResult interface {

	// Size returns the count of MatchGroupResults instances that are in this
	// EnvMatchResult instance.
	Size() int

	// Has tests whether this EnvMatchResults contains any hits for the named
	// MatchGroup.
	Has(groupName string) bool

	// Get returns the MatchGroupResults instance for the named MatchGroup.
	//
	// If no such group was found in the match results, this method will return
	// nil.
	//
	// If this method returns a MatchGroupResults instance, then that instance is
	// guaranteed to have at least one hit.
	Get(groupName string) MatchGroupResults

	// Errors returns the errors that were encountered while attempting to parse
	// and match the environment variables.  These errors will be for variables
	// or groups that were required but were not present in the environment.
	//
	// If the environment parsing had no errors, this method will return nil.
	Errors() MatcherErrors
}
