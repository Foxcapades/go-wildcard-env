package wenv

import (
	"fmt"
	"regexp"
	"strings"
)

// KeyMatcher defines a type that may be used to attempt to match keys in a
// given environment.
type KeyMatcher interface {

	// Name returns the name of the KeyMatcher.  KeyMatcher names are used to
	// reference the target keys in the match results.
	Name() string

	// Matches tests whether the given environment key matches this KeyMatcher's
	// configured filter(s).
	Matches(key string) bool

	// Process processes the given key, returning the keys matched in the given
	// environment key.
	Process(key string) []string
}

// // // // // // // // // // // // // // // // // // // // // // // // // // //
//
//    Prefix Key Matcher
//
// // // // // // // // // // // // // // // // // // // // // // // // // // //

// NewPrefixMatcher constructs a new KeyMatcher instance that uses the given
// prefix to match environment variable names and extract a single key from
// those names.
//
// Example:
//   matcher := NewPrefixKeyMatcher("myMatcher", "MY_KEY_PREFIX_")
//
// This type of matcher is useful if the variable or wildcard part of the target
// environment variables is at the very end of the environment variable names,
// after a common prefix.
//
// An example of such an environment expectation might be:
//   PLUGIN_NAME_ORANGE=My Orange Plugin
//   PLUGIN_PATH_ORANGE=/opt/app/plugins/orange
//   PLUGIN_NAME_PURPLE=My Purple Plugin
//   PLUGIN_PATH_PURPLE=/opt/app/plugins/purple
// In this example, the common prefixes are "PLUGIN_NAME_" and "PLUGIN_PATH_"
// with the wildcard keys being "ORANGE" and "PURPLE".
func NewPrefixMatcher(name, prefix string) KeyMatcher {
	return &prefixKeyMatcher{name, prefix}
}

type prefixKeyMatcher struct{ name, prefix string }

func (p *prefixKeyMatcher) Name() string {
	return p.name
}

func (p *prefixKeyMatcher) Matches(key string) bool {
	return len(key) > len(p.prefix) && strings.HasPrefix(key, p.prefix)
}

func (p *prefixKeyMatcher) Process(key string) []string {
	return []string{key[len(p.prefix):]}
}

// // // // // // // // // // // // // // // // // // // // // // // // // // //
//
//    Suffix Key Matcher
//
// // // // // // // // // // // // // // // // // // // // // // // // // // //

// NewSuffixMatcher constructs a new KeyMatcher instance that uses the given
// suffix to match environment variable names and extract a single key from
// those names.
//
// Example:
//   matcher := NewSuffixKeyMatcher("my matcher", "_MY_KEY_SUFFIX")
//
// This type of matcher is useful if the variable or wildcard part of the target
// environment variables is at the beginning of the environment variable names,
// followed by a common suffix.
//
// An example of such an environment expectation might be:
//   ORANGE_PLUGIN_NAME=My Orange Plugin
//   ORANGE_PLUGIN_PATH=/opt/app/plugins/orange
//   PURPLE_PLUGIN_NAME=My Purple Plugin
//   PURPLE_PLUGIN_PATH=/opt/app/plugins/purple
// In this example, the common suffixes are "_PLUGIN_NAME" and "_PLUGIN_PATH"
// with the wildcard keys being "ORANGE" and "PURPLE".
func NewSuffixMatcher(name, suffix string) KeyMatcher {
	return &suffixKeyMatcher{name, suffix}
}

type suffixKeyMatcher struct{ name, suffix string }

func (s *suffixKeyMatcher) Name() string {
	return s.name
}

func (s *suffixKeyMatcher) Matches(key string) bool {
	return len(key) > len(s.suffix) && strings.HasSuffix(key, s.suffix)
}

func (s *suffixKeyMatcher) Process(key string) []string {
	return []string{key[:len(key)-len(s.suffix)]}
}

// // // // // // // // // // // // // // // // // // // // // // // // // // //
//
//    Wrapped Key Matcher
//
// // // // // // // // // // // // // // // // // // // // // // // // // // //

// NewWrappedMatcher constructs a new KeyMatcher instance that uses the given
// prefix and suffix to match environment variable names and extract a single
// key from those names.
//
// Example:
//   matcher := NewWrappedKeyMatcher("myMatcher", "MY_PREFIX_", "_MY_SUFFIX")
//
// This type of matcher is useful if the variable or wildcard part of the target
// environment variables is in between common prefixes and suffixes.
//
// An example of such an environment expectation might be:
//   PLUGIN_ORANGE_NAME=My Orange Plugin
//   PLUGIN_ORANGE_PATH=/opt/app/plugins/orange
//   PLUGIN_PURPLE_NAME=My Purple Plugin
//   PLUGIN_PURPLE_PATH=/opt/app/plugins/purple
// In this example, the common prefix is "PLUGIN_" and the common suffixes are
// "_NAME" and "_PATH".
func NewWrappedMatcher(name, prefix, suffix string) KeyMatcher {
	return &wrappedKeyMatcher{name, prefix, suffix}
}

type wrappedKeyMatcher struct{ name, prefix, suffix string }

func (w *wrappedKeyMatcher) Name() string {
	return w.name
}

func (w *wrappedKeyMatcher) Matches(key string) bool {
	return len(key) > len(w.prefix)+len(w.suffix) &&
		strings.HasPrefix(key, w.prefix) &&
		strings.HasSuffix(key, w.suffix)
}

func (w *wrappedKeyMatcher) Process(key string) []string {
	return []string{key[len(w.prefix) : len(key)-len(w.suffix)]}
}

// // // // // // // // // // // // // // // // // // // // // // // // // // //
//
//    Regex Key Matcher
//
// // // // // // // // // // // // // // // // // // // // // // // // // // //

// NewRegexMatcher constructs a new KeyMatcher instance that uses the given
// regex to match environment variable names and extract keys from those names.
//
// The given regex must contain at least one matching group, otherwise this key
// matcher will error when used.
//
// Examples:
//   matcher := NewRegexMatcher(regexp.MustCompile(`^PLUGIN_(\w+)_NAME$`))
//   matcher := NewRegexMatcher(regexp.MustCompile(`^PLUGIN_(\w+)_(\w+)_NAME`))
//
// This type of matcher is useful if the environment variable matching is
// complex or multiple wildcard keys need to be parsed from the environment
// variable names.
//
// An example of such an environment expectation might be:
//   FRUIT_PAIR_ORANGE_BANANA=Orange,Banana
//   FRUIT_PAIR_GRAPE_MANGO=Grape,Mango
// In this example, the regex used to match and parse the above environment
// variables would be `^FRUIT_PAIR_(\w+)_(\w+)$`.
func NewRegexMatcher(name string, regex *regexp.Regexp) KeyMatcher {
	return &regexKeyMatcher{name, regex}
}

type regexKeyMatcher struct {
	name  string
	regex *regexp.Regexp
}

func (r *regexKeyMatcher) Name() string {
	return r.name
}

func (r *regexKeyMatcher) Matches(key string) bool {
	return r.regex.MatchString(key)
}

func (r *regexKeyMatcher) Process(key string) []string {
	matches := r.regex.FindStringSubmatch(key)

	if len(matches) == 0 {
		panic(fmt.Errorf("illegal state: no matches were found for key matcher %s", r.name))
	}

	if len(matches) == 1 {
		panic(fmt.Errorf("illegal state: no regex matching groups for key matcher %s", r.name))
	}

	return matches[1:]
}
