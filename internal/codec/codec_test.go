
package codec

import (
	"strings"
	"testing"
)

func TestURLSafe(t *testing.T) {
	s := EncodeURL([]byte{0xfb, 0xff, 0xfe})
	if strings.ContainsAny(s, "+/") {
		t.Fatalf("not url-safe: %s", s)
	}
}
