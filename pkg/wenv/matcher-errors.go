package wenv

import (
	"bufio"
	"fmt"
	"io"
)

type MatcherErrors []error

func (m MatcherErrors) Size() int {
	return len(m)
}

func (m MatcherErrors) Get(index int) error {
	return m[index]
}

func (m MatcherErrors) IsEmpty() bool {
	return len(m) == 0
}

func (m MatcherErrors) HasErrors() bool {
	return len(m) > 0
}

func (m MatcherErrors) Error() string {
	return fmt.Sprintf("encountered %d environment parsing errors", len(m))
}

func (m MatcherErrors) WriteLines(w io.Writer) (written int, err error) {
	var buf *bufio.Writer

	if m.IsEmpty() {
		return 0, nil
	}

	if b, ok := w.(*bufio.Writer); ok {
		buf = b
	} else {
		buf = bufio.NewWriter(w)
	}

	if w, e := buf.WriteString(m[0].Error()); e != nil {
		written += w
		err = e
		return
	} else {
		written += w
	}

	for i := 1; i < len(m); i++ {
		if e := buf.WriteByte('\n'); e != nil {
			written++
			err = e
			return
		} else {
			written++
		}

		if w, e := buf.WriteString(m[i].Error()); e != nil {
			written += w
			err = e
			return
		} else {
			written += w
		}
	}

	buf.Flush()

	return
}
