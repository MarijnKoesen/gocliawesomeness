package gocliawesomeness

import (
	"golang.org/x/term"
	"log/slog"
	"os"
)

type LogFormat string

const (
	FormatText LogFormat = "text"
	FormatJson LogFormat = "json"
	FormatAuto LogFormat = "auto"
)

func NewAwesomeLogger(options *Options) *slog.Logger {
	switch options.Format {
	case FormatText:
		return NewAwesomeTextLogger(options)
	case FormatJson:
		return NewAwesomeJsonLogger(options)
	default:
		if term.IsTerminal(int(os.Stdout.Fd())) {
			return NewAwesomeTextLogger(options)
		} else {
			return NewAwesomeJsonLogger(options)
		}
	}
}

func NewAwesomeJsonLogger(options *Options) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   options.AddSource,
		Level:       options.Level,
		ReplaceAttr: options.ReplaceAttr,
	})

	return slog.New(handler)
}

func NewAwesomeTextLogger(options *Options) *slog.Logger {
	tinted := NewTintedHandler(os.Stdout, options)
	prefixed := NewPrefixedHandler(tinted, &HandlerOptions{
		PrefixKeys: options.PrefixKeys,
	})

	return slog.New(prefixed)
}

func Level(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func Format(format string) LogFormat {
	switch format {
	case string(FormatText):
		return FormatText
	case string(FormatJson):
		return FormatJson
	default:
		return FormatAuto
	}
}
