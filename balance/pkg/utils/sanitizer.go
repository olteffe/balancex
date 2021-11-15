package utils

import "github.com/microcosm-cc/bluemonday"

// SanitizeString Sanitize string
func SanitizeString(str string) string {
	ugcPolicy := bluemonday.UGCPolicy()
	return ugcPolicy.Sanitize(str)
}
