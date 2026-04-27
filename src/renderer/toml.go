package renderer

import (
	"fmt"
	"sort"
	"strings"
)

func ParseTOML(m map[string]any) string {
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
		b.WriteString(" = ")
		b.WriteString(tomlFormatValue(v))

		if i < len(keys)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func tomlFormatValue(v any) string {
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
		// In TOML, numbers just render as-is
		return fmt.Sprintf("%v", t)
	}
}
