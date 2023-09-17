package wenv

import (
	"strings"
	"sync"
)

// SplitEnvironment splits the raw environment slice returned by os.Environ into
// a map of keys and values as expected by the EnvironmentMatcher.
func SplitEnvironment(env []string) map[string]string {
	out := make(map[string]string, len(env))

	for _, pair := range env {
		i := strings.IndexByte(pair, '=')

		if i == -1 {
			out[pair] = ""
		} else {
			out[pair[:i]] = pair[i+1:]
		}
	}

	return out
}

func newKeyMerger() (out keyMerger) {
	out.sb.Grow(256)
	return
}

type keyMerger struct {
	sb strings.Builder
	lk sync.Mutex
}

func (k *keyMerger) merge(keys []string) string {
	k.lk.Lock()
	k.sb.Reset()

	if len(keys) == 0 {
		return ""
	}

	k.sb.WriteString(keys[0])
	for i := 1; i < len(keys); i++ {
		k.sb.WriteString(",")
		k.sb.WriteString(keys[i])
	}

	out := k.sb.String()
	k.lk.Unlock()

	return out
}
