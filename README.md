## Awesome Go Slog Formatter

This is a mash up of the following to project with a bit of my own sauce:

* https://github.com/lmittmann/tint
* https://github.com/dpotapov/slogpfx

### Usage

See ``example/``

```golang
options := &gocliawesomeness.Options{
    Format:     gocliawesomeness.Format(config.logFormat), // text or json
    Level:      gocliawesomeness.Level(config.logLevel), // debug, info, warn, err
    PrefixKeys: []string{"app", "service"},
}
slog.SetDefault(gocliawesomeness.NewAwesomeLogger(options))

logger := slog.With("app", "my-app")
```