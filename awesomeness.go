package gocliawesomeness

import (
	"log/slog"
	"os"
)

type LogFormat string

const (
	FormatText LogFormat = "text"
	FormatJson LogFormat = "json"
)

func NewAwesomeLogger(options *Options) *slog.Logger {
	if options.Format == FormatText {
		return NewAwesomeTextLogger(options)
	} else {
		return NewAwesomeJsonLogger(options)
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
	default:
		return FormatJson
	}
}
