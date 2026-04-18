package template

const tocCSS = `
.toc-panel {
  position: fixed;
  top: 41px;
  right: 0;
  width: 280px;
  height: calc(100vh - 41px);
  overflow-y: auto;
  border-left: 1px solid #e1e4e8;
  background: #f6f8fa;
  transform: translateX(100%);
  transition: transform 0.2s ease-out;
  z-index: 90;
  font-size: 13px;
}
body.toc-visible .toc-panel { transform: translateX(0); }
body.toc-visible .markdown-body { margin-right: 280px; }
.toc-header {
  padding: 10px 16px;
  border-bottom: 1px solid #e1e4e8;
  font-weight: bold;
}
.toc-body { padding: 8px 0; }
.toc-list, .toc-list ul {
  list-style: none;
  padding-left: 0;
  margin: 0;
}
.toc-list ul { padding-left: 14px; }
.toc-list li { margin: 0; }
.toc-list a {
  display: block;
  padding: 4px 16px 4px 4px;
  text-decoration: none;
  color: inherit;
  border-left: 2px solid transparent;
  word-break: break-word;
}
.toc-list a:hover { background: rgba(0,0,0,0.05); }
.toc-toggle.disabled {
  opacity: 0.4;
  pointer-events: none;
}
.theme-dark .toc-panel {
  background: #161b22;
  border-left-color: #30363d;
}
.theme-dark .toc-header { border-bottom-color: #30363d; }
.theme-dark .toc-list a:hover { background: rgba(255,255,255,0.08); }
.theme-simple .toc-panel {
  background: #f5f5f5;
  border-left-color: #ddd;
}
.theme-simple .toc-header { border-bottom-color: #ddd; }
@media (max-width: 900px) {
  body.toc-visible .markdown-body { margin-right: 0; }
  .toc-panel { box-shadow: -2px 0 8px rgba(0,0,0,0.15); }
}
@media print {
  .toc-panel, .toc-toggle { display: none !important; }
  body.toc-visible .markdown-body { margin-right: 0 !important; }
}
`

const tocJS = `<script>
(function() {
  var STORAGE_KEY = 'mdproxy_toc_visible';
  var toggleBtn = document.querySelector('.toc-toggle');
  var panel = document.getElementById('toc-panel');
  if (!toggleBtn || !panel) return;

  var headings = document.querySelectorAll('.markdown-body h1, .markdown-body h2, .markdown-body h3, .markdown-body h4, .markdown-body h5, .markdown-body h6');
  if (headings.length === 0) {
    toggleBtn.classList.add('disabled');
    toggleBtn.setAttribute('title', 'No headings in this document');
    return;
  }

  var listRoot = panel.querySelector('.toc-list');
  headings.forEach(function(h, i) {
    if (!h.id) h.id = 'toc-' + i;
    var level = parseInt(h.tagName.substring(1), 10);
    var li = document.createElement('li');
    li.setAttribute('data-level', level);
    var a = document.createElement('a');
    a.href = '#' + h.id;
    a.textContent = h.textContent.trim();
    a.addEventListener('click', function(e) {
      e.preventDefault();
      h.scrollIntoView({behavior: 'smooth', block: 'start'});
      history.replaceState(null, '', '#' + h.id);
    });
    li.appendChild(a);
    listRoot.appendChild(li);
  });

  function setVisible(v) {
    document.body.classList.toggle('toc-visible', v);
    try { localStorage.setItem(STORAGE_KEY, v ? '1' : '0'); } catch (e) {}
  }
  toggleBtn.addEventListener('click', function() {
    setVisible(!document.body.classList.contains('toc-visible'));
  });

  var saved = null;
  try { saved = localStorage.getItem(STORAGE_KEY); } catch (e) {}
  if (saved === '1') setVisible(true);
})();
</script>`
