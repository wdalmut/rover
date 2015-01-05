package wrap

import (
	"bytes"
	"fmt"
)

type HtmlWrapper struct {
}

func (h *HtmlWrapper) Wrap(data []byte) []byte {
	pos := bytes.Index(data, []byte{0}) - 2
	js := fmt.Sprintf("<script>parent.postMessage('%s', location.origin)</script>", data[:pos])

	return []byte(js)
}
