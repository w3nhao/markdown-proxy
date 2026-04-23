package handler

import (
	"html/template"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/patakuti/markdown-proxy/internal/config"
)

// firstNonLoopbackIP returns the first non-loopback IPv4 address of any
// interface, or an empty string when none is available. Used as a last-resort
// hostname label when os.Hostname() returns something useless.
func firstNonLoopbackIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, a := range addrs {
		n, ok := a.(*net.IPNet)
		if !ok || n.IP.IsLoopback() {
			continue
		}
		ip4 := n.IP.To4()
		if ip4 != nil {
			return ip4.String()
		}
	}
	return ""
}

type TopHandler struct {
	cfg  *config.Config
	tmpl *template.Template
}

func NewTopHandler(cfg *config.Config) *TopHandler {
	tmpl := template.Must(template.New("top").Parse(topPageTmpl))
	return &TopHandler{cfg: cfg, tmpl: tmpl}
}

type rootView struct {
	Path   string
	Label  string
	Origin string
}

func (h *TopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	host, _ := os.Hostname()
	if host == "" || host == "(none)" || host == "localhost" {
		host = firstNonLoopbackIP()
	}
	mounts := listMounts()
	roots := make([]rootView, 0, len(h.cfg.Roots))
	for _, rc := range h.cfg.Roots {
		label := rc.Label
		if label == "" {
			label = filepath.Base(rc.Path)
		}
		origin := rc.Origin
		if origin == "" {
			if mo := findMountOrigin(rc.Path, mounts); mo != "" {
				origin = mo
			} else if host != "" {
				origin = host + ":" + rc.Path
			}
		}
		roots = append(roots, rootView{Path: rc.Path, Label: label, Origin: origin})
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	h.tmpl.Execute(w, map[string]interface{}{
		"RemoteMode": h.cfg.IsRemoteMode(),
		"Roots":      roots,
		"Host":       host,
	})
}

const topPageTmpl = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>markdown-proxy</title>
<style>
  :root {
    --bg: #ffffff;
    --fg: #24292e;
    --muted: #586069;
    --border: #d1d5da;
    --border-soft: #eaecef;
    --accent: #0366d6;
    --accent-hover: #0256b9;
    --card-bg: transparent;
    --card-hover: rgba(0,0,0,0.03);
  }
  body.theme-dark {
    --bg: #0d1117; --fg: #c9d1d9; --muted: #8b949e;
    --border: #30363d; --border-soft: #21262d;
    --accent: #58a6ff; --accent-hover: #79b8ff;
    --card-hover: rgba(255,255,255,0.04);
  }
  body.theme-academia-dark {
    --bg: #1a1a18; --fg: #d4d4d0; --muted: #999992;
    --border: #2e2e2b; --border-soft: #242421;
    --accent: #7fd285; --accent-hover: #9fe2a3;
    --card-hover: rgba(255,255,255,0.04);
  }
  body.theme-academia {
    --bg: #fafaf7; --fg: #2a2a25; --muted: #6b6b5e;
    --border: #e0e0dc; --border-soft: #ecece8;
    --accent: #5f9b65; --accent-hover: #4c8252;
    --card-hover: rgba(0,0,0,0.025);
  }

  html, body { height: 100%; }
  body {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
    max-width: 820px;
    margin: 64px auto;
    padding: 0 24px;
    color: var(--fg);
    background: var(--bg);
    transition: background 0.15s, color 0.15s;
  }
  h1 { text-align: center; font-size: 1.9em; margin-bottom: 0.3em; }
  .subtitle { text-align: center; color: var(--muted); margin-bottom: 2em; }
  .input-group { display: flex; gap: 8px; margin-bottom: 1.5em; }
  input[type="text"] {
    flex: 1; padding: 10px 14px; font-size: 15px;
    border: 1px solid var(--border); border-radius: 6px;
    outline: none; color: inherit; background: transparent;
  }
  input[type="text"]:focus {
    border-color: var(--accent);
    box-shadow: 0 0 0 3px color-mix(in srgb, var(--accent) 25%, transparent);
  }
  button.primary {
    padding: 10px 20px; font-size: 15px;
    background: var(--accent); color: #fff;
    border: none; border-radius: 6px; cursor: pointer;
  }
  button.primary:hover { background: var(--accent-hover); }

  .section { margin-top: 2.2em; }
  .section h2 {
    font-size: 0.95em; font-weight: 600; color: var(--muted);
    text-transform: uppercase; letter-spacing: 0.06em;
    margin-bottom: 0.8em;
  }
  .section ul { list-style: none; padding: 0; margin: 0; }

  .roots li {
    border: 1px solid var(--border-soft); border-radius: 8px;
    margin-bottom: 8px; background: var(--card-bg);
    transition: background 0.12s;
  }
  .roots li:hover { background: var(--card-hover); }
  .roots a {
    display: block; padding: 10px 14px;
    color: inherit; text-decoration: none;
  }
  .roots .label {
    font-weight: 600; font-size: 15px; color: var(--accent);
  }
  .roots .path {
    display: block; font-size: 12px; color: var(--muted);
    font-family: "SF Mono", "JetBrains Mono", Consolas, monospace;
    margin-top: 2px; word-break: break-all;
  }
  .roots .origin {
    display: block; font-size: 11px; color: var(--muted);
    margin-top: 3px; opacity: 0.75;
  }

  .history li {
    padding: 7px 2px; border-bottom: 1px solid var(--border-soft);
  }
  .history a { color: var(--accent); text-decoration: none; }
  .history a:hover { text-decoration: underline; }
  .history .clear-btn {
    background: none; border: none; color: var(--muted);
    cursor: pointer; font-size: 0.9em; padding: 4px 8px;
    margin-left: 6px;
  }
  .history .clear-btn:hover { color: #d73a49; }

  .host-pill {
    display: inline-block; padding: 1px 8px; margin-left: 6px;
    font-size: 11px; border-radius: 10px;
    background: color-mix(in srgb, var(--accent) 18%, transparent);
    color: var(--accent); font-weight: 500;
    vertical-align: 2px;
  }

  /* Theme toggle */
  .theme-picker {
    position: fixed; top: 14px; right: 14px;
    display: flex; gap: 4px;
    background: var(--card-hover); padding: 4px;
    border: 1px solid var(--border-soft); border-radius: 8px;
    font-size: 12px;
  }
  .theme-picker button {
    background: transparent; color: inherit;
    border: none; padding: 3px 8px; border-radius: 5px;
    cursor: pointer; font-size: 12px;
  }
  .theme-picker button.active {
    background: var(--accent); color: #fff;
  }
</style>
</head>
<body>
<div class="theme-picker">
  <button type="button" data-theme="light">Light</button>
  <button type="button" data-theme="dark">Dark</button>
</div>
<h1>markdown-proxy</h1>
{{if .RemoteMode}}
<p class="subtitle">Enter a URL to view Markdown</p>
<div class="input-group">
  <input type="text" id="path-input" placeholder="https://example.com/doc.md" autofocus>
  <button class="primary" onclick="navigate()">Open</button>
</div>
{{else}}
<p class="subtitle">Enter a file path or URL to view Markdown</p>
<div class="input-group">
  <input type="text" id="path-input" placeholder="/path/to/file.md or https://example.com/doc.md" autofocus>
  <button class="primary" onclick="navigate()">Open</button>
</div>
{{end}}
{{if and (not .RemoteMode) .Roots}}
<div class="section roots">
  <h2>Mounted directories{{if .Host}}<span class="host-pill">{{.Host}}</span>{{end}}</h2>
  <ul>
  {{range .Roots}}
    <li><a href="/local{{.Path}}/">
      <span class="label">{{.Label}}</span>
      <span class="path">{{.Path}}</span>
      {{if .Origin}}<span class="origin">&larr; {{.Origin}}</span>{{end}}
    </a></li>
  {{end}}
  </ul>
</div>
{{end}}
<div class="section history" id="history-section" style="display:none;">
  <h2>Recent files<button class="clear-btn" onclick="clearHistory()">clear</button></h2>
  <ul id="history-list"></ul>
</div>
<script>
var remoteMode = {{.RemoteMode}};

// Theme: honor saved preference, else prefers-color-scheme.
(function() {
  var saved = localStorage.getItem('mdproxy_theme');
  if (saved) {
    document.body.className = 'theme-' + saved;
  } else if (window.matchMedia && matchMedia('(prefers-color-scheme: dark)').matches) {
    document.body.className = 'theme-dark';
  }
  syncThemePicker();
})();
function syncThemePicker() {
  var cur = (document.body.className.match(/theme-([\w-]+)/) || [,''])[1];
  var isDark = cur.indexOf('dark') !== -1;
  document.querySelectorAll('.theme-picker button').forEach(function(b) {
    b.classList.toggle('active', (b.dataset.theme === 'dark') === isDark);
  });
}
document.querySelectorAll('.theme-picker button').forEach(function(b) {
  b.addEventListener('click', function() {
    var t = b.dataset.theme;
    // Preserve an academia variant across light/dark if already chosen.
    var cur = (document.body.className.match(/theme-([\w-]+)/) || [,''])[1];
    var next;
    if (cur.indexOf('academia') !== -1) {
      next = t === 'dark' ? 'academia-dark' : 'academia';
    } else {
      next = t === 'dark' ? 'dark' : 'github';
    }
    document.body.className = 'theme-' + next;
    localStorage.setItem('mdproxy_theme', next);
    syncThemePicker();
  });
});

function navigate() {
  var input = document.getElementById('path-input').value.trim();
  if (!input) return;
  var url = '';
  if (input.startsWith('http://')) {
    url = '/http/' + input.substring(7);
  } else if (input.startsWith('https://')) {
    url = '/https/' + input.substring(8);
  } else if (!remoteMode) {
    if (input.startsWith('/') || input.startsWith('~') || /^[A-Za-z]:\\/.test(input)) {
      url = '/local' + (input.startsWith('/') ? '' : '/') + input;
    } else {
      url = '/local/' + input;
    }
  } else {
    url = '/https/' + input;
  }
  saveHistory(input, url);
  window.location.href = url;
}

document.getElementById('path-input').addEventListener('keydown', function(e) {
  if (e.key === 'Enter') navigate();
});

function getHistory() {
  try { return JSON.parse(localStorage.getItem('mdproxy_history') || '[]'); }
  catch(e) { return []; }
}
function saveHistory(input, url) {
  var history = getHistory().filter(function(h) { return h.input !== input; });
  history.unshift({input: input, url: url, time: Date.now()});
  if (history.length > 20) history = history.slice(0, 20);
  localStorage.setItem('mdproxy_history', JSON.stringify(history));
}
function clearHistory() {
  localStorage.removeItem('mdproxy_history');
  renderHistory();
}
function renderHistory() {
  var history = getHistory();
  if (remoteMode) history = history.filter(function(h) { return !h.url.startsWith('/local'); });
  var section = document.getElementById('history-section');
  var list = document.getElementById('history-list');
  if (history.length === 0) { section.style.display = 'none'; return; }
  section.style.display = 'block';
  list.innerHTML = '';
  history.forEach(function(h) {
    var li = document.createElement('li');
    var a = document.createElement('a');
    a.href = h.url;
    a.textContent = h.input;
    li.appendChild(a);
    list.appendChild(li);
  });
}
renderHistory();
</script>
</body>
</html>`
