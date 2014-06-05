package gocommons

import "strings"

// If str is empty or just contains whitespace, returns def.
func DefaultIfBlank(str string, def string) string {
    if len(strings.TrimSpace(str)) > 0 {
        return str
    } else {
        return def
    }
}
