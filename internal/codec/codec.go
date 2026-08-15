
package codec

import "encoding/base64"

func EncodeURL(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}
