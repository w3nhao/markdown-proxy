# markdown-proxy

A local HTTP proxy server for viewing Markdown files in a browser.

## Features

- Render local and remote Markdown files as HTML
- Support for GFM (GitHub Flavored Markdown)
- Embedded diagram rendering (Mermaid, PlantUML, SVG)
- GitHub/GitLab integration with authentication via git credential helper
- Multiple CSS themes (GitHub-like, Simple, Dark)
- Directory listing for local files
- Link rewriting for seamless navigation
- Top page with smart input (auto-detects file path or URL)
- Recently opened file history
- Localhost-only binding for security

## URL Scheme

| Type | URL Format |
|------|-----------|
| Local file | `http://localhost:9080/local/path/to/file.md` |
| Remote (HTTP) | `http://localhost:9080/http/server/path/to/file.md` |
| Remote (HTTPS) | `http://localhost:9080/https/server/path/to/file.md` |

## Usage

```bash
markdown-proxy [options]
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `--port`, `-p` | Listen port | `9080` |
| `--theme` | Default CSS theme | `github` |
| `--plantuml-server` | PlantUML server URL | `https://www.plantuml.com/plantuml` |

## Build

```bash
go build -o markdown-proxy ./cmd/markdown-proxy
```

### Cross-compile

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o markdown-proxy ./cmd/markdown-proxy

# Windows
GOOS=windows GOARCH=amd64 go build -o markdown-proxy.exe ./cmd/markdown-proxy
```
