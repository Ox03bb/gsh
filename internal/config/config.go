package config

import "os"

type Config struct {
	Prompt    string
	Theme     string
	ColorMode bool

	stdin  *os.File
	stdout *os.File
	stderr *os.File

	HistoryFile string
}

func SetDefault[typ any](conf *Config, field *typ, val typ, def typ) {
	if _, ok := any(val).(string); ok {
		if any(val).(string) == "" {
			*field = def
		} else {
			*field = val
		}
	} else {
		if any(val) == nil {
			*field = def
		} else {
			*field = val
		}
	}
}

func (cfg *Config) NewConfig(
	prompt string,
	theme string,
	colorMode bool,
	stdin *os.File,
	stdout *os.File,
	stderr *os.File,
	historyFile string,
) {
	// Theme
	SetDefault(cfg, &cfg.Prompt, prompt, DefaultPrompt)
	SetDefault(cfg, &cfg.Theme, theme, "")
	SetDefault(cfg, &cfg.ColorMode, colorMode, true)

	// File descriptors
	SetDefault(cfg, &cfg.stdin, stdin, os.Stdin)
	SetDefault(cfg, &cfg.stdout, stdout, os.Stdout)
	SetDefault(cfg, &cfg.stderr, stderr, os.Stderr)

	// History
	SetDefault(cfg, &cfg.HistoryFile, historyFile, "")
}
