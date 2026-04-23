package template

import _ "embed"

//go:embed assets/html-to-image.min.js
var htmlToImageJS string

const exportCSS = `
.export-wrapper { position: relative; display: inline-block; }
.export-popover {
  position: absolute;
  top: calc(100% + 6px);
  right: 0;
  width: 300px;
  background: #ffffff;
  color: #24292f;
  border: 1px solid #d0d7de;
  border-radius: 6px;
  padding: 14px;
  box-shadow: 0 6px 18px rgba(0,0,0,0.15);
  z-index: 200;
  font-size: 13px;
  display: none;
  text-align: left;
}
.export-popover.open { display: block; }
.theme-dark .export-popover,
.theme-academia-dark .export-popover {
  background: #252525;
  color: #d4d4d0;
  border-color: #444;
}
.export-popover h3 { margin: 0 0 10px 0; font-size: 13px; font-weight: 600; }
.export-row { display: flex; align-items: center; gap: 8px; margin-bottom: 9px; }
.export-row > label { width: 70px; flex-shrink: 0; opacity: 0.65; }
.export-row input[type="number"] { width: 70px; padding: 3px 6px; font-size: 13px; }
.export-row input[type="text"] { flex: 1; padding: 3px 6px; font-size: 13px; }
.export-row select { flex: 1; padding: 2px 6px; font-size: 13px; }
.theme-dark .export-popover input,
.theme-dark .export-popover select,
.theme-academia-dark .export-popover input,
.theme-academia-dark .export-popover select {
  background: #2a2a2a; color: #d4d4d0; border: 1px solid #444; border-radius: 3px;
}
.export-segmented { display: inline-flex; border: 1px solid #d0d7de; border-radius: 4px; overflow: hidden; }
.export-segmented button {
  background: transparent; border: none; padding: 3px 11px; cursor: pointer;
  font-size: 12px; color: inherit;
}
.export-segmented button.active { background: #0969da; color: white; }
.theme-dark .export-segmented,
.theme-academia-dark .export-segmented { border-color: #444; }
.theme-dark .export-segmented button.active,
.theme-academia-dark .export-segmented button.active { background: #58a6ff; color: #0d1117; }
.export-presets { display: flex; gap: 6px; margin-bottom: 12px; }
.export-presets button {
  flex: 1; padding: 4px 8px; font-size: 12px;
  border: 1px solid #d0d7de; background: transparent; border-radius: 4px;
  cursor: pointer; color: inherit;
}
.export-presets button:hover { background: rgba(128,128,128,0.1); }
.theme-dark .export-presets button,
.theme-academia-dark .export-presets button { border-color: #444; }
.export-checks { margin: 10px 0; }
.export-checks label {
  display: block; margin-bottom: 5px; width: auto; opacity: 1;
  font-size: 12.5px; cursor: pointer;
}
.export-checks input[type="checkbox"] { margin-right: 6px; }
.export-pixels { font-size: 11px; opacity: 0.55; margin-top: -4px; margin-bottom: 10px; padding-left: 78px; }
.export-actions { display: flex; justify-content: flex-end; gap: 8px; margin-top: 12px; }
.export-actions button {
  padding: 5px 14px; font-size: 13px; border-radius: 4px; cursor: pointer;
  border: 1px solid #d0d7de; background: transparent; color: inherit;
}
.export-actions button.primary { background: #1f883d; color: white; border-color: #1f883d; }
.export-actions button.primary:hover { background: #1a7f37; }
.export-actions button.primary:disabled { background: #94d3a2; border-color: #94d3a2; cursor: not-allowed; }
.theme-dark .export-actions button,
.theme-academia-dark .export-actions button { border-color: #444; }
@media print { .export-popover, .export-wrapper { display: none !important; } }
`

const exportToolbarHTML = `
<span class="export-wrapper">
  <a href="javascript:void(0)" class="toolbar-link" id="export-btn">Export ▾</a>
  <div class="export-popover" id="export-popover">
    <h3>Export as PNG</h3>
    <div class="export-presets">
      <button type="button" data-preset="default">Default</button>
      <button type="button" data-preset="print">Print</button>
      <button type="button" data-preset="wide">Wide</button>
    </div>
    <div class="export-row">
      <label>Width</label>
      <input type="number" id="export-width" value="760" min="300" max="2400" step="20" list="export-width-presets" />
      <datalist id="export-width-presets">
        <option value="600"></option><option value="760"></option>
        <option value="900"></option><option value="1200"></option>
      </datalist>
      <span style="opacity:0.5">px</span>
    </div>
    <div class="export-row">
      <label>Scale</label>
      <div class="export-segmented" id="export-scale">
        <button type="button" data-scale="1">1&times;</button>
        <button type="button" data-scale="2" class="active">2&times;</button>
        <button type="button" data-scale="3">3&times;</button>
      </div>
    </div>
    <div class="export-pixels" id="export-pixels">&rarr; 1520 &times; auto px</div>
    <div class="export-row">
      <label>Theme</label>
      <select id="export-theme">
        <option value="current" selected>Current</option>
        <option value="github">GitHub</option>
        <option value="simple">Simple</option>
        <option value="dark">Dark</option>
        <option value="academia">Academia</option>
        <option value="academia-dark">Academia Dark</option>
      </select>
    </div>
    <div class="export-row">
      <label>Background</label>
      <select id="export-bg">
        <option value="match" selected>Match theme</option>
        <option value="white">White</option>
        <option value="transparent">Transparent</option>
      </select>
    </div>
    <div class="export-checks">
      <label><input type="checkbox" id="export-expand-details" checked /> Expand &lt;details&gt; blocks</label>
      <label><input type="checkbox" id="export-include-toolbar" /> Include toolbar</label>
      <label><input type="checkbox" id="export-include-frontmatter" checked /> Include frontmatter</label>
      <label><input type="checkbox" id="export-include-toc" /> Include TOC panel</label>
    </div>
    <div class="export-row">
      <label>Filename</label>
      <input type="text" id="export-filename" placeholder="auto" />
    </div>
    <div class="export-actions">
      <button type="button" id="export-cancel">Cancel</button>
      <button type="button" id="export-download" class="primary">Download</button>
    </div>
  </div>
</span>
`

const exportJS = `
(function() {
  var btn = document.getElementById('export-btn');
  var popover = document.getElementById('export-popover');
  if (!btn || !popover) return;

  var widthInput = document.getElementById('export-width');
  var scaleGroup = document.getElementById('export-scale');
  var pixelsLabel = document.getElementById('export-pixels');
  var themeSel = document.getElementById('export-theme');
  var bgSel = document.getElementById('export-bg');
  var optExpand = document.getElementById('export-expand-details');
  var optToolbar = document.getElementById('export-include-toolbar');
  var optFrontmatter = document.getElementById('export-include-frontmatter');
  var optToc = document.getElementById('export-include-toc');
  var filenameInput = document.getElementById('export-filename');
  var cancelBtn = document.getElementById('export-cancel');
  var downloadBtn = document.getElementById('export-download');

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
  function toggle(show) {
    popover.classList.toggle('open', show);
    if (show && !filenameInput.value) filenameInput.placeholder = defaultFilename();
  }

  btn.addEventListener('click', function(e) {
    e.stopPropagation();
    toggle(!popover.classList.contains('open'));
  });
  document.addEventListener('click', function(e) {
    if (!popover.contains(e.target) && e.target !== btn) toggle(false);
  });
  cancelBtn.addEventListener('click', function() { toggle(false); });
  widthInput.addEventListener('input', updatePixels);
  scaleGroup.querySelectorAll('button').forEach(function(b) {
    b.addEventListener('click', function() { setScale(parseInt(b.dataset.scale)); });
  });
  document.querySelectorAll('.export-presets button').forEach(function(b) {
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
    downloadBtn.textContent = 'Rendering...';

    var width = parseInt(widthInput.value) || 760;
    var scale = getScale();
    var theme = themeSel.value === 'current' ? currentTheme() : themeSel.value;
    var bgChoice = bgSel.value;
    var bgColor = (bgChoice === 'white') ? '#ffffff'
                : (bgChoice === 'transparent') ? undefined
                : themeBg(theme);

    // Wrapper is rendered on-screen; a covering overlay hides it from the user while html-to-image serializes.
    var wrapper = document.createElement('div');
    wrapper.className = 'theme-' + theme;
    wrapper.style.position = 'fixed';
    wrapper.style.top = '0';
    wrapper.style.left = '0';
    wrapper.style.zIndex = '2147483000';
    wrapper.style.width = width + 'px';
    wrapper.style.display = 'block';
    wrapper.style.background = bgColor || '#ffffff';

    var overlay = document.createElement('div');
    overlay.style.cssText = 'position:fixed;inset:0;z-index:2147483647;background:rgba(0,0,0,0.55);color:#fff;display:flex;align-items:center;justify-content:center;font:500 15px system-ui,sans-serif;';
    overlay.textContent = 'Rendering export...';

    // Clone the relevant parts of the current body into the wrapper
    var clonedNodes = [];
    var srcNodes = Array.from(document.body.children);
    srcNodes.forEach(function(n) {
      if (n.id === 'export-popover') return;
      // Skip the export wrapper itself if the toolbar is excluded
      var clone = n.cloneNode(true);
      clonedNodes.push({ src: n, clone: clone });
      wrapper.appendChild(clone);
    });
    var q = function(sel) { return wrapper.querySelectorAll(sel); };
    q('#export-popover').forEach(function(n) { n.remove(); });
    q('.export-wrapper').forEach(function(n) { n.remove(); });
    if (!optToolbar.checked) q('.toolbar').forEach(function(n) { n.remove(); });
    if (!optToc.checked) q('#toc-panel').forEach(function(n) { n.remove(); });
    if (!optFrontmatter.checked) q('.frontmatter').forEach(function(n) { n.remove(); });
    if (optExpand.checked) q('details').forEach(function(d) { d.open = true; d.setAttribute('open',''); });

    // Width constraint on markdown body inside the clone
    var mdBody = wrapper.querySelector('.markdown-body');
    if (mdBody) {
      mdBody.style.maxWidth = width + 'px';
      mdBody.style.margin = '0 auto';
      mdBody.style.padding = '32px 28px 48px';
      mdBody.style.boxSizing = 'border-box';
    }

    // Force code blocks to wrap instead of getting clipped by hidden horizontal scroll
    wrapper.querySelectorAll('pre').forEach(function(p) {
      p.style.whiteSpace = 'pre-wrap';
      p.style.wordBreak = 'break-word';
      p.style.overflowX = 'visible';
    });

    document.body.appendChild(wrapper);
    document.body.appendChild(overlay);

    var opts = {
      pixelRatio: scale,
      cacheBust: true,
      // Skip stylesheets we cannot read (CORS-blocked CDN sheets would throw otherwise).
      filter: function(node) { return !(node.tagName === 'LINK' && node.rel === 'stylesheet'); }
    };
    if (bgColor) opts.backgroundColor = bgColor;

    // Force layout + wait for fonts, then render
    void wrapper.offsetHeight;
    console.log('[export] wrapper size:', wrapper.offsetWidth, 'x', wrapper.offsetHeight, 'theme:', theme);
    var fontsReady = (document.fonts && document.fonts.ready) || Promise.resolve();
    fontsReady.then(function() { return htmlToImage.toPng(wrapper, opts); }).then(function(dataUrl) {
      console.log('[export] dataUrl length:', dataUrl.length);
      var fname = filenameInput.value.trim() || defaultFilename();
      if (!/\.png$/i.test(fname)) fname += '.png';
      var a = document.createElement('a');
      a.href = dataUrl; a.download = fname;
      document.body.appendChild(a); a.click(); a.remove();
    }).catch(function(err) {
      console.error('Export failed:', err);
      alert('Export failed: ' + (err && err.message ? err.message : err));
    }).then(function() {
      wrapper.remove();
      overlay.remove();
      downloadBtn.disabled = false;
      downloadBtn.textContent = origLabel;
      toggle(false);
    });
  });

  updatePixels();
})();
`
