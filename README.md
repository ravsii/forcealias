# forcealias

**`forcealias`** is a [Go static analysis](https://pkg.go.dev/golang.org/x/tools/go/analysis) tool that enforces specific import aliases in your codebase. It's especially useful for large projects or teams that want to maintain consistent import aliasing practices.

## Features

- Enforce specific aliases for imports
- Optionally ignore dot (`.`) and underscore (`_`) imports
- TODO: golangci-lint integration

## Usage

While it can be used as a standalone linter, do we recommend use it with
[golangci-lint](https://golangci-lint.run/). (when we add it there, haha)

### Standalone

#### Installation

You can install it via `go install`:

```bash
go install github.com/ravsii/forcealias/cmd/forcealias@latest
```

#### Running

```bash
forcealias --force-alias fmt=myfmt ./...
```

### golangci-lint

TODO

## Flags

| Flag                  | Description                                         |
| --------------------- | --------------------------------------------------- |
| `--force-alias`       | Comma-separated list of `importPath=alias` mappings |
| `--ignore-dot`        | Skip checking for `.` imports                       |
| `--ignore-underscore` | Skip checking for `_` imports                       |

### Example

```bash
forcealias --force-alias net/url=defaultUrl,fmt=myfmt --ignore-dot --ignore-underscore ./...
```

## License

MIT License. See [LICENSE](./LICENSE) for details.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
