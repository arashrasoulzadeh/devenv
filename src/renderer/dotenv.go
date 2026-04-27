package renderer

import (
	"fmt"
	"sort"
	"strings"
)

func ParseDotEnv(m map[string]any) string {
	if len(m) == 0 {
		return ""
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var b strings.Builder

	for i, k := range keys {
		v := m[k]
		if v == nil {
			continue
		}

		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(formatValue(v))

		if i < len(keys)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func formatValue(v any) string {
	switch t := v.(type) {

	case string:
		t = strings.ReplaceAll(t, `"`, `\"`)
		return `"` + t + `"`

	case bool:
		if t {
			return "true"
		}
		return "false"

	default:
		return fmt.Sprintf("%v", t)
	}
}
