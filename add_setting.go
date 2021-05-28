package viperEx

import (
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
func AddBoolSetting(flagSet *pflag.FlagSet, name string, cmdline string, description string) {
	AddBoolSettingP(flagSet, name, cmdline, "", description)
}

func AddBoolSettingP(flagSet *pflag.FlagSet, name string, cmdline string, shorthand string, description string) {
	addSetting(flagSet, name, func() string { flagSet.BoolP(cmdline, shorthand, false, description); return cmdline })
}

// --------------------------------------------------------------------------------------------------------------
// Int
// --------------------------------------------------------------------------------------------------------------

func AddIntSetting(flagSet *pflag.FlagSet, name string, cmdline string, description string) {
	AddIntSettingP(flagSet, name, cmdline, "", description)
}

func AddIntSettingP(flagSet *pflag.FlagSet, name string, cmdline string, shorthand string, description string) {
	AddIntSettingPD(flagSet, name, cmdline, shorthand, 0, description)
}

func AddIntSettingD(flagSet *pflag.FlagSet, name string, cmdline string, defaultVal int, description string) {
	AddIntSettingPD(flagSet, name, cmdline, "", defaultVal, description)
}

func AddIntSettingPD(flagSet *pflag.FlagSet, name string, cmdline string, shorthand string, defaultVal int, description string) {
	addSetting(flagSet, name, func() string { flagSet.IntP(cmdline, shorthand, defaultVal, description); return cmdline })
}

// --------------------------------------------------------------------------------------------------------------
// String
// --------------------------------------------------------------------------------------------------------------
func AddStringSetting(flagSet *pflag.FlagSet, name string, cmdline string, description string) {
	AddStringSettingP(flagSet, name, cmdline, "", description)
}

func AddStringSettingP(flagSet *pflag.FlagSet, name string, cmdline string, shorthand string, description string) {
	AddStringSettingPD(flagSet, name, cmdline, shorthand, "", description)
}

func AddStringSettingD(flagSet *pflag.FlagSet, name string, cmdline string, defaultVal string, description string) {
	AddStringSettingPD(flagSet, name, cmdline, "", defaultVal, description)
}

func AddStringSettingPD(flagSet *pflag.FlagSet, name string, cmdline string, shorthand string, defaultVal string, description string) {
	addSetting(flagSet, name, func() string { flagSet.StringP(cmdline, shorthand, defaultVal, description); return cmdline })
}

// --------------------------------------------------------------------------------------------------------------
// StringArray
// --------------------------------------------------------------------------------------------------------------
func AddStringArraySetting(flagSet *pflag.FlagSet, name string, cmdline string, description string) {
	AddStringArraySettingP(flagSet, name, cmdline, "", description)
}

func AddStringArraySettingP(flagSet *pflag.FlagSet, name string, cmdline string, shorthand string, description string) {
	AddStringArraySettingPD(flagSet, name, cmdline, shorthand, []string{}, description)
}

func AddStringArraySettingD(flagSet *pflag.FlagSet, name string, cmdline string, defaultVal []string, description string) {
	AddStringArraySettingPD(flagSet, name, cmdline, "", defaultVal, description)
}

func AddStringArraySettingPD(flagSet *pflag.FlagSet, name string, cmdline string, shorthand string, defaultVal []string, description string) {
	addSetting(flagSet, name, func() string { flagSet.StringArrayP(cmdline, shorthand, defaultVal, description); return cmdline })
}

// --------------------------------------------------------------------------------------------------------------
// Common
// --------------------------------------------------------------------------------------------------------------
func addSetting(flagSet *pflag.FlagSet, name string, createFlag func() string) {
	if flag, ok := flags[name]; ok {
		// TODO [Improvement] - check type is the same

		// Store for next time
		flags[name] = flag
	}
}
