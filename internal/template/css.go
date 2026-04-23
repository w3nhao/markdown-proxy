package template

const commonCSS = `
* { box-sizing: border-box; }
body { margin: 0; padding: 0; }
.toolbar {
  position: sticky; top: 0; z-index: 100;
  display: flex; justify-content: space-between; align-items: center;
  padding: 8px 20px;
  border-bottom: 1px solid #e1e4e8;
  background: #f6f8fa;
}
.theme-dark .toolbar { background: #1e1e1e; border-color: #444; }
.theme-academia .toolbar { background: #f5f5f0; border-color: #e0e0dc; }
.theme-academia-dark .toolbar { background: #1e1e1c; border-color: #2e2e2b; color: #b8b8b4; }
.home-link { text-decoration: none; font-weight: bold; font-size: 14px; }
.theme-github .home-link, .theme-simple .home-link { color: #0366d6; }
.theme-dark .home-link { color: #58a6ff; }
.theme-academia .home-link { color: #5f9b65; font-family: Tinos, Palatino, serif; letter-spacing: 0.02em; }
.theme-academia-dark .home-link { color: #7fd285; font-family: Tinos, Palatino, serif; letter-spacing: 0.02em; }
.toolbar-actions { display: flex; align-items: center; gap: 12px; }
.toolbar-link { font-size: 13px; text-decoration: none; }
.theme-github .toolbar-link, .theme-simple .toolbar-link { color: #0366d6; }
.theme-dark .toolbar-link { color: #58a6ff; }
.theme-academia .toolbar-link { color: #5f9b65; }
.theme-academia-dark .toolbar-link { color: #7fd285; }
.toolbar-link:hover { text-decoration: underline; }
.theme-switcher { display: flex; align-items: center; gap: 6px; font-size: 13px; }
.theme-switcher label { opacity: 0.7; }
.theme-switcher select { padding: 2px 6px; font-size: 13px; }
.theme-dark .theme-switcher select,
.theme-academia-dark .theme-switcher select {
  background: #2a2a2a;
  color: #d4d4d0;
  border: 1px solid #444;
  border-radius: 3px;
}
@media print {
  .toolbar { display: none; }
  table, pre, .math.display, img, blockquote, li { break-inside: avoid; }
  h1, h2, h3, h4, h5, h6 { break-after: avoid; }
}
.markdown-body pre.text-file {
  white-space: pre-wrap;
}
.markdown-body .frontmatter {
  margin: 0 0 1.5em 0;
  padding: 0;
  font-size: 0.8em;
  opacity: 0.55;
  border-left: 2px solid currentColor;
  padding-left: 12px;
}
.markdown-body .frontmatter:hover { opacity: 0.9; }
.markdown-body .frontmatter-body {
  margin: 0;
  padding: 0;
  background: transparent !important;
  border: none !important;
  font-family: inherit;
  white-space: pre-wrap;
  word-break: break-word;
  line-height: 1.5;
}
.theme-dark .markdown-body .frontmatter,
.theme-academia-dark .markdown-body .frontmatter { opacity: 0.5; }
.plantuml-notice {
  padding: 12px 16px;
  margin: 16px 0;
  border-radius: 6px;
  font-size: 14px;
  line-height: 1.5;
}
.plantuml-notice code {
  padding: .2em .4em;
  border-radius: 3px;
  font-size: 85%;
}
.theme-github .plantuml-notice,
.theme-simple .plantuml-notice {
  background: #fff8c5;
  border: 1px solid #d4a72c;
  color: #4d3800;
}
.theme-github .plantuml-notice code,
.theme-simple .plantuml-notice code {
  background: rgba(0,0,0,.08);
}
.theme-dark .plantuml-notice {
  background: #2d2a1e;
  border: 1px solid #966c00;
  color: #e3b341;
}
.theme-dark .plantuml-notice code {
  background: rgba(255,255,255,.1);
}
.markdown-body {
  max-width: 980px;
  margin: 0 auto;
  padding: 40px 20px;
}
.markdown-body table {
  border-collapse: collapse;
  width: 100%;
}
.markdown-body table th,
.markdown-body table td {
  border: 1px solid #dfe2e5;
  padding: 6px 13px;
}
.markdown-body table tr:nth-child(2n) {
  background-color: #f6f8fa;
}
.theme-dark .markdown-body table tr:nth-child(2n) {
  background-color: #2d2d2d;
}
.theme-dark .markdown-body table th,
.theme-dark .markdown-body table td {
  border-color: #444;
}
`

const githubCSS = `
.theme-github {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  color: #24292e;
  background: #fff;
}
.theme-github .markdown-body h1 { padding-bottom: .3em; border-bottom: 1px solid #eaecef; }
.theme-github .markdown-body h2 { padding-bottom: .3em; border-bottom: 1px solid #eaecef; }
.theme-github .markdown-body a { color: #0366d6; text-decoration: none; }
.theme-github .markdown-body a:hover { text-decoration: underline; }
.theme-github .markdown-body code {
  background: rgba(27,31,35,.05);
  padding: .2em .4em;
  border-radius: 3px;
  font-size: 85%;
}
.theme-github .markdown-body pre {
  background: #f6f8fa;
  padding: 16px;
  border-radius: 6px;
  border: 1px solid #e1e4e8;
  overflow: auto;
}
.theme-github .markdown-body pre code { background: none; padding: 0; font-size: 100%; }
.theme-github .markdown-body blockquote {
  color: #6a737d;
  border-left: .25em solid #dfe2e5;
  padding: 0 1em;
  margin: 0;
}
.theme-github .markdown-body img { max-width: 100%; }
`

const simpleCSS = `
.theme-simple {
  font-family: Georgia, "Times New Roman", serif;
  color: #333;
  background: #fefefe;
  line-height: 1.8;
}
.theme-simple .markdown-body a { color: #07c; }
.theme-simple .markdown-body code {
  background: #f0f0f0;
  padding: .15em .3em;
  border-radius: 2px;
}
.theme-simple .markdown-body pre {
  background: #f0f0f0;
  padding: 14px;
  border-radius: 4px;
  border: 1px solid #ddd;
  overflow: auto;
}
.theme-simple .markdown-body pre code { background: none; padding: 0; font-size: 100%; }
.theme-simple .markdown-body blockquote {
  color: #666;
  border-left: 3px solid #ccc;
  padding: 0 1em;
  margin: 0;
}
.theme-simple .markdown-body img { max-width: 100%; }
`

const darkCSS = `
.theme-dark {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  color: #c9d1d9;
  background: #0d1117;
}
.theme-dark .markdown-body h1 { padding-bottom: .3em; border-bottom: 1px solid #21262d; }
.theme-dark .markdown-body h2 { padding-bottom: .3em; border-bottom: 1px solid #21262d; }
.theme-dark .markdown-body a { color: #58a6ff; text-decoration: none; }
.theme-dark .markdown-body a:hover { text-decoration: underline; }
.theme-dark .markdown-body code {
  background: rgba(110,118,129,.4);
  padding: .2em .4em;
  border-radius: 3px;
  font-size: 85%;
}
.theme-dark .markdown-body pre {
  background: #161b22;
  padding: 16px;
  border-radius: 6px;
  border: 1px solid #30363d;
  overflow: auto;
}
.theme-dark .markdown-body pre code { background: none; padding: 0; font-size: 100%; }
.theme-dark .markdown-body blockquote {
  color: #8b949e;
  border-left: .25em solid #30363d;
  padding: 0 1em;
  margin: 0;
}
.theme-dark .markdown-body img { max-width: 100%; }
`

// Academia — Typora 主题移植，学术衬线风格
const academiaCSS = `
.theme-academia {
  --body-font: "Tinos Nerd Font Propo", "Tinos Nerd Font", Tinos, "Source Han Serif SC VF", "Songti SC", "Palatino", serif;
  --heading-font: "Tinos Nerd Font Propo", "Tinos Nerd Font", Tinos, "Source Han Serif SC VF", "Songti SC", "Palatino", serif;
  --code-font: "Sarasa Mono SC", "Sarasa Mono Slab SC", "JetBrainsMono Nerd Font", "JetBrains Mono", Menlo, monospace;
  --text-color: #2c2c2c;
  --heading-color: #1a1a1a;
  --link-color: #5f9b65;
  --link-hover: #3d7a42;
  --code-bg: #f7f7f4;
  --inline-code-bg: #f0efe9;
  --quote-border: #d0d0d0;
  --quote-text: #555;
  --table-rule: #333;
  --table-border-light: #e0e0e0;
  --bg-color: #ffffff;
  font-family: var(--body-font);
  color: var(--text-color);
  background: var(--bg-color);
}
.theme-academia .markdown-body {
  max-width: 900px;
  margin: 0 auto;
  padding: 40px 30px 120px;
  font-size: 17px;
  line-height: 1.8;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
.theme-academia .markdown-body h1,
.theme-academia .markdown-body h2,
.theme-academia .markdown-body h3,
.theme-academia .markdown-body h4,
.theme-academia .markdown-body h5,
.theme-academia .markdown-body h6 {
  font-family: var(--heading-font);
  color: var(--heading-color);
  font-weight: 600;
  line-height: 1.3;
}
.theme-academia .markdown-body h1 {
  font-size: 2.2em;
  text-align: center;
  margin-top: 1em;
  margin-bottom: 0.3em;
  font-weight: 700;
  letter-spacing: 0.01em;
  border-bottom: none;
  padding-bottom: 0;
}
.theme-academia .markdown-body h1 + p {
  text-align: center;
  color: #777;
  font-size: 0.95em;
  margin-top: 0;
}
.theme-academia .markdown-body h2 {
  font-size: 1.35em;
  margin-top: 2em;
  margin-bottom: 0.5em;
  padding-bottom: 0.3em;
  border-bottom: 1px solid #eaeaea;
}
.theme-academia .markdown-body h3 { font-size: 1.15em; margin-top: 1.7em; margin-bottom: 0.4em; }
.theme-academia .markdown-body h4 { font-size: 1.05em; margin-top: 1.4em; margin-bottom: 0.3em; font-style: italic; }
.theme-academia .markdown-body h5 { font-size: 1em; margin-top: 1.2em; margin-bottom: 0.25em; font-weight: 600; color: #555; }
.theme-academia .markdown-body h6 {
  font-size: 0.95em;
  margin-top: 1em;
  margin-bottom: 0.2em;
  color: #666;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}
.theme-academia .markdown-body p { margin: 0.9em 0; }
.theme-academia .markdown-body a {
  color: var(--link-color);
  text-decoration: none;
  border-bottom: 1px solid transparent;
  transition: border-color 0.2s;
}
.theme-academia .markdown-body a:hover {
  color: var(--link-hover);
  border-bottom-color: var(--link-hover);
}
.theme-academia .markdown-body strong { font-weight: 700; color: #1a1a1a; }
.theme-academia .markdown-body em { font-style: italic; }
.theme-academia .markdown-body mark {
  background: #fff3c4;
  padding: 1px 4px;
  border-radius: 2px;
}
.theme-academia .markdown-body ul,
.theme-academia .markdown-body ol { padding-left: 28px; margin: 0.8em 0; }
.theme-academia .markdown-body li { margin: 0.25em 0; }
.theme-academia .markdown-body li > ul,
.theme-academia .markdown-body li > ol { margin: 0.15em 0; }
.theme-academia .markdown-body blockquote {
  border-left: 3px solid var(--quote-border);
  padding: 0.4em 0 0.4em 20px;
  margin: 1.2em 0;
  color: var(--quote-text);
  font-size: 0.96em;
}
.theme-academia .markdown-body blockquote p { margin: 0.5em 0; }
.theme-academia .markdown-body code {
  font-family: var(--code-font);
  font-size: 0.85em;
  background: var(--inline-code-bg);
  border-radius: 3px;
  padding: 2px 6px;
  border: none;
  color: #3c3836;
}
.theme-academia .markdown-body pre {
  font-family: var(--code-font);
  font-size: 14.5px;
  line-height: 1.55;
  background: var(--code-bg);
  border: none;
  border-radius: 6px;
  padding: 16px 20px;
  margin: 1.2em 0;
  overflow-x: auto;
}
.theme-academia .markdown-body pre code { background: none; padding: 0; font-size: 100%; color: inherit; }
.theme-academia .markdown-body table {
  width: 100%;
  border-collapse: collapse;
  margin: 1.5em 0;
  font-size: 0.95em;
}
.theme-academia .markdown-body table th,
.theme-academia .markdown-body table td {
  padding: 8px 12px;
  text-align: left;
  border: none;
}
.theme-academia .markdown-body table thead {
  border-top: 2px solid var(--table-rule);
  border-bottom: 2px solid var(--table-rule);
  background: none;
}
.theme-academia .markdown-body table th {
  font-family: var(--heading-font);
  font-weight: 600;
  color: var(--heading-color);
}
.theme-academia .markdown-body table tbody tr { border-bottom: 0.5px solid var(--table-border-light); }
.theme-academia .markdown-body table tbody tr:last-child { border-bottom: 1.5px solid var(--table-rule); }
.theme-academia .markdown-body table tr:nth-child(2n) { background: none; }
.theme-academia .markdown-body hr {
  border: none;
  height: 0;
  border-top: 1px solid #d0d0d0;
  max-width: 30%;
  margin: 2.5em auto;
}
.theme-academia .markdown-body img {
  max-width: 100%;
  display: block;
  margin: 1.5em auto;
  border-radius: 2px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06);
}
.theme-academia .markdown-body .katex { font-size: 1.05em; }
`

// Academia Dark — academia 的深色变体
const academiaDarkCSS = `
.theme-academia-dark {
  --body-font: "Tinos Nerd Font Propo", "Tinos Nerd Font", Tinos, "Source Han Serif SC VF", "Songti SC", "Palatino", serif;
  --heading-font: "Tinos Nerd Font Propo", "Tinos Nerd Font", Tinos, "Source Han Serif SC VF", "Songti SC", "Palatino", serif;
  --code-font: "Sarasa Mono SC", "Sarasa Mono Slab SC", "JetBrainsMono Nerd Font", "JetBrains Mono", Menlo, monospace;
  --text-color: #d4d4d0;
  --heading-color: #f0f0ea;
  --link-color: #7fd285;
  --link-hover: #a3e0a8;
  --code-bg: #1e1e1c;
  --inline-code-bg: #2a2a28;
  --quote-border: #4a4a47;
  --quote-text: #a8a8a5;
  --table-rule: #c0c0b8;
  --table-border-light: #3a3a37;
  --bg-color: #17171a;
  font-family: var(--body-font);
  color: var(--text-color);
  background: var(--bg-color);
}
.theme-academia-dark .markdown-body {
  max-width: 900px;
  margin: 0 auto;
  padding: 40px 30px 120px;
  font-size: 17px;
  line-height: 1.8;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
.theme-academia-dark .markdown-body h1,
.theme-academia-dark .markdown-body h2,
.theme-academia-dark .markdown-body h3,
.theme-academia-dark .markdown-body h4,
.theme-academia-dark .markdown-body h5,
.theme-academia-dark .markdown-body h6 {
  font-family: var(--heading-font);
  color: var(--heading-color);
  font-weight: 600;
  line-height: 1.3;
}
.theme-academia-dark .markdown-body h1 {
  font-size: 2.2em;
  text-align: center;
  margin-top: 1em;
  margin-bottom: 0.3em;
  font-weight: 700;
  letter-spacing: 0.01em;
  border-bottom: none;
  padding-bottom: 0;
}
.theme-academia-dark .markdown-body h1 + p {
  text-align: center;
  color: #999;
  font-size: 0.95em;
  margin-top: 0;
}
.theme-academia-dark .markdown-body h2 {
  font-size: 1.35em;
  margin-top: 2em;
  margin-bottom: 0.5em;
  padding-bottom: 0.3em;
  border-bottom: 1px solid #2e2e2b;
}
.theme-academia-dark .markdown-body h3 { font-size: 1.15em; margin-top: 1.7em; margin-bottom: 0.4em; }
.theme-academia-dark .markdown-body h4 { font-size: 1.05em; margin-top: 1.4em; margin-bottom: 0.3em; font-style: italic; }
.theme-academia-dark .markdown-body h5 { font-size: 1em; margin-top: 1.2em; margin-bottom: 0.25em; font-weight: 600; color: #b0b0ad; }
.theme-academia-dark .markdown-body h6 {
  font-size: 0.95em;
  margin-top: 1em;
  margin-bottom: 0.2em;
  color: #a0a09c;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}
.theme-academia-dark .markdown-body p { margin: 0.9em 0; }
.theme-academia-dark .markdown-body a {
  color: var(--link-color);
  text-decoration: none;
  border-bottom: 1px solid transparent;
  transition: border-color 0.2s;
}
.theme-academia-dark .markdown-body a:hover {
  color: var(--link-hover);
  border-bottom-color: var(--link-hover);
}
.theme-academia-dark .markdown-body strong { font-weight: 700; color: #f5f5ef; }
.theme-academia-dark .markdown-body em { font-style: italic; }
.theme-academia-dark .markdown-body mark {
  background: #5c4a1e;
  color: #fff3c4;
  padding: 1px 4px;
  border-radius: 2px;
}
.theme-academia-dark .markdown-body ul,
.theme-academia-dark .markdown-body ol { padding-left: 28px; margin: 0.8em 0; }
.theme-academia-dark .markdown-body li { margin: 0.25em 0; }
.theme-academia-dark .markdown-body li > ul,
.theme-academia-dark .markdown-body li > ol { margin: 0.15em 0; }
.theme-academia-dark .markdown-body blockquote {
  border-left: 3px solid var(--quote-border);
  padding: 0.4em 0 0.4em 20px;
  margin: 1.2em 0;
  color: var(--quote-text);
  font-size: 0.96em;
}
.theme-academia-dark .markdown-body blockquote p { margin: 0.5em 0; }
.theme-academia-dark .markdown-body code {
  font-family: var(--code-font);
  font-size: 0.85em;
  background: var(--inline-code-bg);
  border-radius: 3px;
  padding: 2px 6px;
  border: none;
  color: #e8c89a;
}
.theme-academia-dark .markdown-body pre {
  font-family: var(--code-font);
  font-size: 14.5px;
  line-height: 1.55;
  background: var(--code-bg);
  border: none;
  border-radius: 6px;
  padding: 16px 20px;
  margin: 1.2em 0;
  overflow-x: auto;
}
.theme-academia-dark .markdown-body pre code { background: none; padding: 0; font-size: 100%; color: inherit; }
.theme-academia-dark .markdown-body table {
  width: 100%;
  border-collapse: collapse;
  margin: 1.5em 0;
  font-size: 0.95em;
}
.theme-academia-dark .markdown-body table th,
.theme-academia-dark .markdown-body table td {
  padding: 8px 12px;
  text-align: left;
  border: none;
}
.theme-academia-dark .markdown-body table thead {
  border-top: 2px solid var(--table-rule);
  border-bottom: 2px solid var(--table-rule);
  background: none;
}
.theme-academia-dark .markdown-body table th {
  font-family: var(--heading-font);
  font-weight: 600;
  color: var(--heading-color);
}
.theme-academia-dark .markdown-body table tbody tr { border-bottom: 0.5px solid var(--table-border-light); }
.theme-academia-dark .markdown-body table tbody tr:last-child { border-bottom: 1.5px solid var(--table-rule); }
.theme-academia-dark .markdown-body table tr:nth-child(2n) { background: none; }
.theme-academia-dark .markdown-body hr {
  border: none;
  height: 0;
  border-top: 1px solid #3a3a37;
  max-width: 30%;
  margin: 2.5em auto;
}
.theme-academia-dark .markdown-body img {
  max-width: 100%;
  display: block;
  margin: 1.5em auto;
  border-radius: 2px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.4);
}
.theme-academia-dark .markdown-body .katex { font-size: 1.05em; color: #e8e8e3; }
`
