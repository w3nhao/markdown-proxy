package template

import _ "embed"

//go:embed assets/html-to-image.min.js
var htmlToImageJS string

const fabCSS = `
/* Floating action button */
.mp-fab {
  position: fixed; bottom: 22px; right: 22px;
  width: 44px; height: 44px; border-radius: 50%;
  background: rgba(255,255,255,0.92); color: #24292f;
  border: 1px solid rgba(0,0,0,0.08);
  box-shadow: 0 4px 14px rgba(0,0,0,0.12);
  display: flex; align-items: center; justify-content: center;
  cursor: pointer; z-index: 1000;
  backdrop-filter: blur(10px); -webkit-backdrop-filter: blur(10px);
  opacity: 0.55;
  transition: opacity 0.18s, transform 0.15s, background 0.2s;
}
.mp-fab:hover, .mp-fab.active { opacity: 1; transform: scale(1.04); }
.mp-fab svg { width: 22px; height: 22px; display: block; }
.theme-dark .mp-fab, .theme-academia-dark .mp-fab {
  background: rgba(38,38,38,0.92); color: #d4d4d0;
  border-color: rgba(255,255,255,0.10);
}

/* Main panel */
.mp-panel {
  position: fixed; bottom: 78px; right: 22px;
  width: 290px; max-height: calc(100vh - 120px);
  background: #ffffff; color: #24292f;
  border: 1px solid rgba(0,0,0,0.08); border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0,0,0,0.18);
  z-index: 999; display: none; overflow: hidden;
  font-size: 13px; text-align: left;
  flex-direction: column;
}
.mp-panel.open { display: flex; }
.theme-dark .mp-panel, .theme-academia-dark .mp-panel {
  background: #1e1e1e; color: #d4d4d0;
  border-color: rgba(255,255,255,0.10);
}
.mp-panel::after {
  content: ''; position: absolute; bottom: -7px; right: 32px;
  width: 14px; height: 14px; background: inherit;
  border-right: 1px solid rgba(0,0,0,0.08);
  border-bottom: 1px solid rgba(0,0,0,0.08);
  transform: rotate(45deg);
}
.theme-dark .mp-panel::after, .theme-academia-dark .mp-panel::after {
  border-color: rgba(255,255,255,0.10);
}

.mp-view { overflow-y: auto; flex: 1; }
.mp-view.hidden { display: none; }

.mp-section { padding: 6px 0; border-top: 1px solid rgba(128,128,128,0.12); }
.mp-section:first-child { border-top: none; }
.mp-caption {
  padding: 8px 16px 4px; font-size: 11px;
  letter-spacing: 0.06em; text-align: center;
  color: rgba(128,128,128,0.85); text-transform: uppercase;
}

.mp-row {
  display: flex; align-items: center; justify-content: space-between;
  padding: 8px 16px; gap: 10px; min-height: 32px;
}
.mp-row.action { cursor: pointer; user-select: none; }
.mp-row.action:hover { background: rgba(128,128,128,0.08); }
.mp-row > .mp-label { opacity: 0.85; flex-shrink: 0; }
.mp-row .chev { opacity: 0.4; font-size: 14px; }

.mp-row input[type="number"] {
  width: 78px; padding: 4px 8px; text-align: right;
  background: transparent; border: 1px solid rgba(128,128,128,0.3);
  border-radius: 5px; color: inherit; font: inherit;
}
.mp-row input[type="text"] {
  flex: 1; padding: 4px 8px;
  background: transparent; border: 1px solid rgba(128,128,128,0.3);
  border-radius: 5px; color: inherit; font: inherit;
}
.mp-row select {
  padding: 3px 8px; background: transparent;
  border: 1px solid rgba(128,128,128,0.3); border-radius: 5px;
  color: inherit; font: inherit;
}
.theme-dark .mp-row select option,
.theme-academia-dark .mp-row select option { background: #2a2a2a; }

/* iOS-style toggle */
.mp-toggle { position: relative; display: inline-block; width: 34px; height: 20px; flex-shrink: 0; }
.mp-toggle input { opacity: 0; width: 0; height: 0; }
.mp-toggle .slider {
  position: absolute; inset: 0; cursor: pointer;
  background: rgba(128,128,128,0.35); border-radius: 20px;
  transition: background 0.2s;
}
.mp-toggle .slider::before {
  content: ''; position: absolute;
  width: 16px; height: 16px; left: 2px; top: 2px;
  background: white; border-radius: 50%;
  box-shadow: 0 1px 2px rgba(0,0,0,0.2);
  transition: transform 0.2s;
}
.mp-toggle input:checked + .slider { background: #34c759; }
.mp-toggle input:checked + .slider::before { transform: translateX(14px); }

/* Segmented control */
.mp-seg {
  display: inline-flex; border: 1px solid rgba(128,128,128,0.3);
  border-radius: 5px; overflow: hidden;
}
.mp-seg button {
  background: transparent; border: none;
  padding: 3px 11px; font-size: 12px; cursor: pointer; color: inherit;
}
.mp-seg button.active { background: #0969da; color: white; }
.theme-dark .mp-seg button.active, .theme-academia-dark .mp-seg button.active {
  background: #58a6ff; color: #0d1117;
}

/* Preset buttons */
.mp-presets { display: flex; gap: 6px; padding: 6px 16px 4px; }
.mp-presets button {
  flex: 1; padding: 4px 8px; font-size: 12px;
  border: 1px solid rgba(128,128,128,0.3);
  background: transparent; color: inherit;
  border-radius: 5px; cursor: pointer;
}
.mp-presets button:hover { background: rgba(128,128,128,0.08); }

.mp-hint {
  padding: 2px 16px 8px; font-size: 11px;
  opacity: 0.55; text-align: right;
}

/* Back row */
.mp-back {
  display: flex; align-items: center; gap: 6px;
  padding: 10px 14px; cursor: pointer; font-size: 13px;
  border-bottom: 1px solid rgba(128,128,128,0.12);
  user-select: none;
}
.mp-back:hover { background: rgba(128,128,128,0.08); }

.mp-primary {
  display: block; width: calc(100% - 32px);
  margin: 10px 16px 14px; padding: 8px;
  background: #1f883d; color: white;
  border: none; border-radius: 6px;
  cursor: pointer; font-size: 13px; font-weight: 500;
  transition: background 0.15s;
}
.mp-primary:hover { background: #1a7f37; }
.mp-primary:disabled { background: #94d3a2; cursor: not-allowed; }

@media print { .mp-fab, .mp-panel { display: none !important; } }
`

// fabHTML uses Go template directives for theme-selected state and optional source link.
const fabHTML = `
<button class="mp-fab" id="mp-fab" aria-label="Settings" type="button">
  <svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true">
    <path d="M19.43 12.98c.04-.32.07-.64.07-.98 0-.34-.03-.66-.07-.98l2.11-1.65a.5.5 0 0 0 .12-.64l-2-3.46a.5.5 0 0 0-.61-.22l-2.49 1a7.3 7.3 0 0 0-1.69-.98l-.38-2.65A.49.49 0 0 0 14 2h-4a.49.49 0 0 0-.49.42l-.38 2.65c-.61.25-1.17.58-1.69.98l-2.49-1a.5.5 0 0 0-.61.22l-2 3.46a.5.5 0 0 0 .12.64l2.11 1.65c-.04.32-.07.64-.07.98 0 .34.03.66.07.98l-2.11 1.65a.5.5 0 0 0-.12.64l2 3.46c.14.24.43.34.68.22l2.49-1c.52.4 1.08.73 1.69.98l.38 2.65c.04.24.25.42.49.42h4c.24 0 .45-.18.49-.42l.38-2.65c.61-.25 1.17-.58 1.69-.98l2.49 1c.25.12.54.02.68-.22l2-3.46a.5.5 0 0 0-.12-.64l-2.11-1.65ZM12 15.5a3.5 3.5 0 1 1 0-7 3.5 3.5 0 0 1 0 7Z"/>
  </svg>
</button>

<div class="mp-panel" id="mp-panel" role="dialog" aria-label="Settings">
  <div class="mp-view" id="mp-view-main">
    <div class="mp-section">
      <div class="mp-caption">外观</div>
      <div class="mp-row">
        <span class="mp-label">主题</span>
        <select id="mp-theme">
          <option value="github"{{if eq .Theme "github"}} selected{{end}}>GitHub</option>
          <option value="simple"{{if eq .Theme "simple"}} selected{{end}}>Simple</option>
          <option value="dark"{{if eq .Theme "dark"}} selected{{end}}>Dark</option>
          <option value="academia"{{if eq .Theme "academia"}} selected{{end}}>Academia</option>
          <option value="academia-dark"{{if eq .Theme "academia-dark"}} selected{{end}}>Academia Dark</option>
        </select>
      </div>
    </div>
    <div class="mp-section">
      <div class="mp-caption">导航</div>
      <div class="mp-row action toc-toggle">
        <span class="mp-label">目录</span><span class="chev">›</span>
      </div>
      <a href="/" class="mp-row action" style="text-decoration:none;color:inherit">
        <span class="mp-label">首页</span><span class="chev">↩</span>
      </a>
      {{if .SourceURL}}
      <a href="{{.SourceURL}}" target="_blank" rel="noopener" class="mp-row action" style="text-decoration:none;color:inherit">
        <span class="mp-label">源文件</span><span class="chev">↗</span>
      </a>
      {{end}}
    </div>
    <div class="mp-section">
      <div class="mp-caption">操作</div>
      <div class="mp-row action" id="mp-print">
        <span class="mp-label">打印</span><span class="chev">⎙</span>
      </div>
      <div class="mp-row action" id="mp-export-open">
        <span class="mp-label">导出 PNG…</span><span class="chev">›</span>
      </div>
    </div>
  </div>

  <div class="mp-view hidden" id="mp-view-export">
    <div class="mp-back" id="mp-back">← 返回</div>
    <div class="mp-section">
      <div class="mp-caption">尺寸</div>
      <div class="mp-presets">
        <button type="button" data-preset="default">默认</button>
        <button type="button" data-preset="print">打印</button>
        <button type="button" data-preset="wide">宽</button>
      </div>
      <div class="mp-row">
        <span class="mp-label">宽度</span>
        <input type="number" id="export-width" value="760" min="300" max="2400" step="20" />
      </div>
      <div class="mp-row">
        <span class="mp-label">缩放</span>
        <div class="mp-seg" id="export-scale">
          <button type="button" data-scale="1">1&times;</button>
          <button type="button" data-scale="2" class="active">2&times;</button>
          <button type="button" data-scale="3">3&times;</button>
        </div>
      </div>
      <div class="mp-hint" id="export-pixels">&rarr; 1520 &times; auto px</div>
    </div>
    <div class="mp-section">
      <div class="mp-caption">内容</div>
      <div class="mp-row">
        <span class="mp-label">展开 &lt;details&gt;</span>
        <label class="mp-toggle"><input type="checkbox" id="export-expand-details" checked /><span class="slider"></span></label>
      </div>
      <div class="mp-row">
        <span class="mp-label">含 frontmatter</span>
        <label class="mp-toggle"><input type="checkbox" id="export-include-frontmatter" checked /><span class="slider"></span></label>
      </div>
      <div class="mp-row">
        <span class="mp-label">含 TOC 面板</span>
        <label class="mp-toggle"><input type="checkbox" id="export-include-toc" /><span class="slider"></span></label>
      </div>
      <div class="mp-row">
        <span class="mp-label">背景</span>
        <select id="export-bg">
          <option value="match" selected>跟随主题</option>
          <option value="white">白色</option>
          <option value="transparent">透明</option>
        </select>
      </div>
    </div>
    <div class="mp-section">
      <div class="mp-row">
        <span class="mp-label">文件名</span>
        <input type="text" id="export-filename" placeholder="auto" />
      </div>
    </div>
    <button class="mp-primary" type="button" id="export-download">下载 PNG</button>
  </div>
</div>
`

const fabJS = `
(function() {
  var fab = document.getElementById('mp-fab');
  var panel = document.getElementById('mp-panel');
  if (!fab || !panel) return;

  var viewMain = document.getElementById('mp-view-main');
  var viewExport = document.getElementById('mp-view-export');
  var themeSel = document.getElementById('mp-theme');
  var printBtn = document.getElementById('mp-print');
  var exportOpen = document.getElementById('mp-export-open');
  var backBtn = document.getElementById('mp-back');

  var widthInput = document.getElementById('export-width');
  var scaleGroup = document.getElementById('export-scale');
  var pixelsLabel = document.getElementById('export-pixels');
  var bgSel = document.getElementById('export-bg');
  var optExpand = document.getElementById('export-expand-details');
  var optFrontmatter = document.getElementById('export-include-frontmatter');
  var optToc = document.getElementById('export-include-toc');
  var filenameInput = document.getElementById('export-filename');
  var downloadBtn = document.getElementById('export-download');

  function showPanel(show) {
    panel.classList.toggle('open', show);
    fab.classList.toggle('active', show);
    if (!show) showView('main');
  }
  function showView(name) {
    viewMain.classList.toggle('hidden', name !== 'main');
    viewExport.classList.toggle('hidden', name !== 'export');
  }

  fab.addEventListener('click', function(e) {
    e.stopPropagation();
    showPanel(!panel.classList.contains('open'));
  });
  document.addEventListener('click', function(e) {
    if (!panel.contains(e.target) && !fab.contains(e.target)) showPanel(false);
  });
  document.addEventListener('keydown', function(e) {
    if (e.key === 'Escape' && panel.classList.contains('open')) showPanel(false);
  });

  // Theme switch (merged with saved-theme restore)
  function applyTheme(theme) {
    document.body.className = 'theme-' + theme;
    localStorage.setItem('mdproxy_theme', theme);
    if (window.mermaid) {
      try { mermaid.initialize({startOnLoad: false, theme: theme.indexOf('dark') !== -1 ? 'dark' : 'default'}); } catch(_) {}
    }
  }
  themeSel.addEventListener('change', function() { applyTheme(themeSel.value); });
  (function() {
    var saved = localStorage.getItem('mdproxy_theme');
    if (saved) {
      document.body.className = 'theme-' + saved;
      themeSel.value = saved;
    }
  })();

  printBtn.addEventListener('click', function() {
    var orig = document.title;
    document.title = orig.split(' - ')[0].replace(/\.(md|markdown)$/i, '');
    window.print();
    document.title = orig;
  });

  exportOpen.addEventListener('click', function() {
    showView('export');
    if (!filenameInput.value) filenameInput.value = defaultFilename();
  });
  backBtn.addEventListener('click', function() { showView('main'); });

  function defaultFilename() {
    var title = (document.title || 'export').split(' - ')[0];
    return title.replace(/\.(md|markdown)$/i, '') + '.png';
  }
  function getScale() {
    var a = scaleGroup.querySelector('button.active');
    return a ? parseInt(a.dataset.scale) : 2;
  }
  function setScale(s) {
    scaleGroup.querySelectorAll('button').forEach(function(b) {
      b.classList.toggle('active', parseInt(b.dataset.scale) === s);
    });
    updatePixels();
  }
  function updatePixels() {
    var w = parseInt(widthInput.value) || 760;
    pixelsLabel.innerHTML = '&rarr; ' + (w * getScale()) + ' &times; auto px';
  }
  function currentTheme() {
    var m = document.body.className.match(/theme-([\w-]+)/);
    return m ? m[1] : 'github';
  }
  function themeBg(theme) {
    return ({
      'github': '#ffffff', 'simple': '#ffffff', 'academia': '#fafaf7',
      'dark': '#0d1117', 'academia-dark': '#1a1a18'
    })[theme] || '#ffffff';
  }

  widthInput.addEventListener('input', updatePixels);
  scaleGroup.querySelectorAll('button').forEach(function(b) {
    b.addEventListener('click', function() { setScale(parseInt(b.dataset.scale)); });
  });
  document.querySelectorAll('.mp-presets button').forEach(function(b) {
    b.addEventListener('click', function() {
      var p = b.dataset.preset;
      if (p === 'default') { widthInput.value = 760; setScale(2); }
      else if (p === 'print') { widthInput.value = 900; setScale(2); }
      else if (p === 'wide') { widthInput.value = 1200; setScale(2); }
      updatePixels();
    });
  });

  downloadBtn.addEventListener('click', function() {
    if (typeof htmlToImage === 'undefined') {
      alert('html-to-image library did not load.');
      return;
    }
    downloadBtn.disabled = true;
    var origLabel = downloadBtn.textContent;
    downloadBtn.textContent = 'Rendering…';

    var width = parseInt(widthInput.value) || 760;
    var scale = getScale();
    var theme = currentTheme();
    var bgChoice = bgSel.value;
    var bgColor = (bgChoice === 'white') ? '#ffffff'
                : (bgChoice === 'transparent') ? undefined
                : themeBg(theme);

    var wrapper = document.createElement('div');
    wrapper.className = 'theme-' + theme;
    wrapper.style.position = 'fixed';
    wrapper.style.top = '0'; wrapper.style.left = '0';
    wrapper.style.zIndex = '2147483000';
    wrapper.style.width = width + 'px';
    wrapper.style.display = 'block';
    wrapper.style.background = bgColor || '#ffffff';

    var overlay = document.createElement('div');
    overlay.style.cssText = 'position:fixed;inset:0;z-index:2147483647;background:rgba(0,0,0,0.55);color:#fff;display:flex;align-items:center;justify-content:center;font:500 15px system-ui,sans-serif;';
    overlay.textContent = 'Rendering export…';

    Array.from(document.body.children).forEach(function(n) {
      if (n.id === 'mp-panel' || n.id === 'mp-fab') return;
      wrapper.appendChild(n.cloneNode(true));
    });
    var q = function(sel) { return wrapper.querySelectorAll(sel); };
    q('#mp-panel').forEach(function(n) { n.remove(); });
    q('#mp-fab').forEach(function(n) { n.remove(); });
    if (!optToc.checked) q('#toc-panel').forEach(function(n) { n.remove(); });
    if (!optFrontmatter.checked) q('.frontmatter').forEach(function(n) { n.remove(); });
    if (optExpand.checked) q('details').forEach(function(d) { d.open = true; d.setAttribute('open',''); });

    var mdBody = wrapper.querySelector('.markdown-body');
    if (mdBody) {
      mdBody.style.maxWidth = width + 'px';
      mdBody.style.margin = '0 auto';
      mdBody.style.padding = '32px 28px 48px';
      mdBody.style.boxSizing = 'border-box';
    }
    wrapper.querySelectorAll('pre').forEach(function(p) {
      p.style.whiteSpace = 'pre-wrap';
      p.style.wordBreak = 'break-word';
      p.style.overflowX = 'visible';
    });

    document.body.appendChild(wrapper);
    document.body.appendChild(overlay);

    var opts = {
      pixelRatio: scale, cacheBust: true,
      filter: function(node) { return !(node.tagName === 'LINK' && node.rel === 'stylesheet'); }
    };
    if (bgColor) opts.backgroundColor = bgColor;

    void wrapper.offsetHeight;
    var fontsReady = (document.fonts && document.fonts.ready) || Promise.resolve();
    fontsReady.then(function() { return htmlToImage.toPng(wrapper, opts); }).then(function(dataUrl) {
      var fname = filenameInput.value.trim() || defaultFilename();
      if (!/\.png$/i.test(fname)) fname += '.png';
      var a = document.createElement('a');
      a.href = dataUrl; a.download = fname;
      document.body.appendChild(a); a.click(); a.remove();
    }).catch(function(err) {
      console.error('Export failed:', err);
      alert('Export failed: ' + (err && err.message ? err.message : err));
    }).then(function() {
      wrapper.remove(); overlay.remove();
      downloadBtn.disabled = false;
      downloadBtn.textContent = origLabel;
      showPanel(false);
    });
  });

  updatePixels();
})();
`
