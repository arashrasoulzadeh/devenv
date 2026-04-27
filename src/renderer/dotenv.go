package renderer

import "fmt"

func ParseDotEnv(m map[string]any) string {
	result := ""
	for k, v := range m {
		result += k + "=" + toString(v) + "\n"
	}
	return result
}

// helper function to convert any value to string
func toString(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case int, int32, int64, float32, float64:
		return fmt.Sprintf("%v", t)
	case bool:
		if t {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", t)
	}
}
