package viperEx

import (
	"strings"

	"github.com/spf13/viper"
)

// Get a string slice from viper
// Fixes the pflag binding problem
func GetStringSlice(name string) []string {
	result := viper.GetStringSlice(name)
	if len(result) == 1 && strings.HasPrefix(result[0], "[") && strings.HasSuffix(result[0], "]") {
		result2 := strings.Trim(result[0], "[]")
		if result2 == "" {
			return []string{}
		}
		return strings.Split(result2, ",")
	}
	return result
}
