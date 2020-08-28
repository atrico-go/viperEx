package viperEx

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var flags = make(map[string]*pflag.Flag)

// Used in testing to clear settings
func Reset() {
	viper.Reset()
	flags = make(map[string]*pflag.Flag)
}

// --------------------------------------------------------------------------------------------------------------
// Bool
// --------------------------------------------------------------------------------------------------------------
func AddBoolSetting(flagSet *pflag.FlagSet, name string, description string) {
	AddBoolSettingP(flagSet, name, "", description)
}

func AddBoolSettingP(flagSet *pflag.FlagSet, name string, shorthand string, description string) {
	addSetting(flagSet, name, func() { flagSet.BoolP(name, shorthand, false, description) })
}

// --------------------------------------------------------------------------------------------------------------
// Int
// --------------------------------------------------------------------------------------------------------------

func AddIntSetting(flagSet *pflag.FlagSet, name string, description string) {
	AddIntSettingP(flagSet, name, "", description)
}

func AddIntSettingP(flagSet *pflag.FlagSet, name string, shorthand string, description string) {
	AddIntSettingPD(flagSet, name, shorthand, 0, description)
}

func AddIntSettingD(flagSet *pflag.FlagSet, name string, defaultVal int, description string) {
	AddIntSettingPD(flagSet, name, "", defaultVal, description)
}

func AddIntSettingPD(flagSet *pflag.FlagSet, name string, shorthand string, defaultVal int, description string) {
	addSetting(flagSet, name, func() { flagSet.IntP(name, shorthand, defaultVal, description) })
}

// --------------------------------------------------------------------------------------------------------------
// String
// --------------------------------------------------------------------------------------------------------------
func AddStringSetting(flagSet *pflag.FlagSet, name string, description string) {
	AddStringSettingP(flagSet, name, "", description)
}

func AddStringSettingP(flagSet *pflag.FlagSet, name string, shorthand string, description string) {
	AddStringSettingPD(flagSet, name, shorthand, "", description)
}

func AddStringSettingD(flagSet *pflag.FlagSet, name string, defaultVal string, description string) {
	AddStringSettingPD(flagSet, name, "", defaultVal, description)
}

func AddStringSettingPD(flagSet *pflag.FlagSet, name string, shorthand string, defaultVal string, description string) {
	addSetting(flagSet, name, func() { flagSet.StringP(name, shorthand, defaultVal, description) })
}

// --------------------------------------------------------------------------------------------------------------
// StringArray
// --------------------------------------------------------------------------------------------------------------
func AddStringArraySetting(flagSet *pflag.FlagSet, name string, description string) {
	AddStringArraySettingP(flagSet, name, "", description)
}

func AddStringArraySettingP(flagSet *pflag.FlagSet, name string, shorthand string, description string) {
	AddStringArraySettingPD(flagSet, name, shorthand, []string{}, description)
}

func AddStringArraySettingD(flagSet *pflag.FlagSet, name string, defaultVal []string, description string) {
	AddStringArraySettingPD(flagSet, name, "", defaultVal, description)
}

func AddStringArraySettingPD(flagSet *pflag.FlagSet, name string, shorthand string, defaultVal []string, description string) {
	addSetting(flagSet, name, func() { flagSet.StringArrayP(name, shorthand, defaultVal, description) })
}

// --------------------------------------------------------------------------------------------------------------
// Common
// --------------------------------------------------------------------------------------------------------------
func addSetting(flagSet *pflag.FlagSet, name string, createFlag func()) {
	if flag, ok := flags[name]; ok {
		// TODO [Improvement] - check type is the same
		// Add existing flag
		flagSet.AddFlag(flag)
	} else {
		// Create new flag
		createFlag()
		flag = flagSet.Lookup(name)
		// Bind to viper
		if err := viper.BindPFlag(name, flag); err != nil {
			log.Fatal("Unable to bind flag:", err)
		}
		// Store for next time
		flags[name] = flag
	}
}
