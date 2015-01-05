package wrap

import (
	//"bytes"
	"fmt"
)

type HtmlWrapper struct {
}

func (h *HtmlWrapper) Wrap(data []byte) []byte {
	js := fmt.Sprintf("<script>parent.postMessage('%s', location.origin)</script>", data[:])

	return []byte(js)
}
