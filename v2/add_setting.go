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
type BoolSettingBuilder interface {
	Cmdline(cmdline string) BoolSettingBuilder
	CmdlineShortcut(shortcut rune) BoolSettingBuilder
	EnvVar(envVar string) BoolSettingBuilder
	DefaultVal(defaultValue bool) BoolSettingBuilder
	AddTo(flagSet *pflag.FlagSet)
}

func BoolSetting(name, description string) BoolSettingBuilder {
	return &boolSetting{name: name, description: description}
}

type boolSetting struct {
	name, description string
	cmdline, shortcut string
	envVar            string
	defaultVal        bool
}

func (b *boolSetting) Cmdline(cmdline string) BoolSettingBuilder {
	b.cmdline = cmdline
	return b
}

func (b *boolSetting) CmdlineShortcut(shortcut rune) BoolSettingBuilder {
	b.shortcut = string(shortcut)
	return b
}

func (b *boolSetting) EnvVar(envVar string) BoolSettingBuilder {
	b.envVar = envVar
	return b
}

func (b *boolSetting) DefaultVal(defaultValue bool) BoolSettingBuilder {
	b.defaultVal = defaultValue
	return b
}

func (b *boolSetting) AddTo(flagSet *pflag.FlagSet) {
	addSetting(flagSet, b.name, b.envVar, func() string { flagSet.BoolP(b.cmdline, b.shortcut, b.defaultVal, b.description); return b.cmdline })
}

// --------------------------------------------------------------------------------------------------------------
// Int
// --------------------------------------------------------------------------------------------------------------
type IntSettingBuilder interface {
	Cmdline(cmdline string) IntSettingBuilder
	CmdlineShortcut(shortcut rune) IntSettingBuilder
	EnvVar(envVar string) IntSettingBuilder
	DefaultVal(defaultValue int) IntSettingBuilder
	AddTo(flagSet *pflag.FlagSet)
}

func IntSetting(name, description string) IntSettingBuilder {
	return &intSetting{name: name, description: description}
}

type intSetting struct {
	name, description string
	cmdline, shortcut string
	envVar            string
	defaultVal        int
}

func (b *intSetting) Cmdline(cmdline string) IntSettingBuilder {
	b.cmdline = cmdline
	return b
}

func (b *intSetting) CmdlineShortcut(shortcut rune) IntSettingBuilder {
	b.shortcut = string(shortcut)
	return b
}

func (b *intSetting) EnvVar(envVar string) IntSettingBuilder {
	b.envVar = envVar
	return b
}

func (b *intSetting) DefaultVal(defaultValue int) IntSettingBuilder {
	b.defaultVal = defaultValue
	return b
}

func (b *intSetting) AddTo(flagSet *pflag.FlagSet) {
	addSetting(flagSet, b.envVar, b.name, func() string { flagSet.IntP(b.cmdline, b.shortcut, b.defaultVal, b.description); return b.cmdline })
}

// --------------------------------------------------------------------------------------------------------------
// String
// --------------------------------------------------------------------------------------------------------------
type StringSettingBuilder interface {
	Cmdline(cmdline string) StringSettingBuilder
	CmdlineShortcut(shortcut rune) StringSettingBuilder
	EnvVar(envVar string) StringSettingBuilder
	DefaultVal(defaultValue string) StringSettingBuilder
	AddTo(flagSet *pflag.FlagSet)
}

func StringSetting(name, description string) StringSettingBuilder {
	return &stringSetting{name: name, description: description}
}

type stringSetting struct {
	name, description string
	cmdline, shortcut string
	envVar            string
	defaultVal        string
}

func (b *stringSetting) Cmdline(cmdline string) StringSettingBuilder {
	b.cmdline = cmdline
	return b
}

func (b *stringSetting) CmdlineShortcut(shortcut rune) StringSettingBuilder {
	b.shortcut = string(shortcut)
	return b
}

func (b *stringSetting) EnvVar(envVar string) StringSettingBuilder {
	b.envVar = envVar
	return b
}

func (b *stringSetting) DefaultVal(defaultValue string) StringSettingBuilder {
	b.defaultVal = defaultValue
	return b
}

func (b *stringSetting) AddTo(flagSet *pflag.FlagSet) {
	addSetting(flagSet, b.name, b.envVar, func() string { flagSet.StringP(b.cmdline, b.shortcut, b.defaultVal, b.description); return b.cmdline })
}

// --------------------------------------------------------------------------------------------------------------
// StringArray
// --------------------------------------------------------------------------------------------------------------
type StringArraySettingBuilder interface {
	Cmdline(cmdline string) StringArraySettingBuilder
	CmdlineShortcut(shortcut rune) StringArraySettingBuilder
	EnvVar(envVar string) StringArraySettingBuilder
	DefaultVal(defaultValue []string) StringArraySettingBuilder
	AddTo(flagSet *pflag.FlagSet)
}

func StringArraySetting(name, description string) StringArraySettingBuilder {
	return &stringArraySetting{name: name, description: description, defaultVal: make([]string, 0)}
}

type stringArraySetting struct {
	name, description string
	cmdline, shortcut string
	envVar            string
	defaultVal        []string
}

func (b *stringArraySetting) Cmdline(cmdline string) StringArraySettingBuilder {
	b.cmdline = cmdline
	return b
}

func (b *stringArraySetting) CmdlineShortcut(shortcut rune) StringArraySettingBuilder {
	b.shortcut = string(shortcut)
	return b
}

func (b *stringArraySetting) EnvVar(envVar string) StringArraySettingBuilder {
	b.envVar = envVar
	return b
}

func (b *stringArraySetting) DefaultVal(defaultValue []string) StringArraySettingBuilder {
	b.defaultVal = defaultValue
	return b
}

func (b *stringArraySetting) AddTo(flagSet *pflag.FlagSet) {
	addSetting(flagSet, b.name, b.envVar, func() string { flagSet.StringArrayP(b.cmdline, b.shortcut, b.defaultVal, b.description); return b.cmdline })
}

// --------------------------------------------------------------------------------------------------------------
// Common
// --------------------------------------------------------------------------------------------------------------
func addSetting(flagSet *pflag.FlagSet, name string, envVar string, createFlag func() string) {
	if flag, ok := flags[name]; ok {
		// TODO [Improvement] - check type is the same
		// Add existing flag
		flagSet.AddFlag(flag)
	} else {
		// Create new flag
		flag = flagSet.Lookup(createFlag())
		// Bind to viper
		if err := viper.BindPFlag(name, flag); err != nil {
			log.Fatal("Unable to bind flag:", err)
		}
		// Store for next time
		flags[name] = flag
	}
	if envVar != "" {
		viper.BindEnv(name, envVar)
	}
}
