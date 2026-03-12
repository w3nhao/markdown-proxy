package template

const lineAnchorJS = `<script>
(function() {
  function highlightLines() {
    // Clear previous highlights
    document.querySelectorAll('.line-highlight').forEach(function(el) {
      el.classList.remove('line-highlight');
    });

    var hash = location.hash;
    if (!hash) return;

    // Parse #L12 or #L12-L34
    var m = hash.match(/^#L(\d+)(?:-L(\d+))?$/);
    if (!m) return;

    var startLine = parseInt(m[1], 10);
    var endLine = m[2] ? parseInt(m[2], 10) : startLine;

    // Find all line anchor elements in range
    var anchors = [];
    for (var i = startLine; i <= endLine; i++) {
      var el = document.getElementById('L' + i);
      if (el) anchors.push(el);
    }
    if (anchors.length === 0) {
      // Try to find the nearest anchor before startLine
      for (var i = startLine - 1; i >= 1; i--) {
        var el = document.getElementById('L' + i);
        if (el) {
          el.scrollIntoView({behavior: 'smooth', block: 'center'});
          return;
        }
      }
      return;
    }

    // Scroll to first anchor
    anchors[0].scrollIntoView({behavior: 'smooth', block: 'center'});

    // Highlight: find the parent block element of each anchor and highlight it
    var highlighted = new Set();
    anchors.forEach(function(anchor) {
      // Walk up to find the nearest block-level parent that contains content
      var parent = anchor.parentElement;
      while (parent && parent.classList.contains('markdown-body')) {
        parent = null;
        break;
      }
      if (parent && !highlighted.has(parent)) {
        highlighted.add(parent);
        parent.classList.add('line-highlight');
      }
    });
  }

  // Run on page load and hash change
  window.addEventListener('hashchange', highlightLines);
  window.addEventListener('DOMContentLoaded', highlightLines);
})();
</script>`

const lineAnchorCSS = `
.line-highlight {
  background-color: rgba(255, 255, 0, 0.2) !important;
  border-left: 3px solid #f0c040 !important;
  transition: background-color 0.3s ease;
}
.theme-dark .line-highlight {
  background-color: rgba(255, 255, 0, 0.1) !important;
  border-left: 3px solid #b08820 !important;
}
@media print {
  .line-highlight {
    background-color: transparent !important;
    border-left: none !important;
  }
}
`
