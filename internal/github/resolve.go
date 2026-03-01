package github

import (
	"regexp"
	"strings"
)

// GitHub blob URL pattern: github.com/user/repo/blob/<ref>/path
var githubBlobRe = regexp.MustCompile(`^github\.com/([^/]+)/([^/]+)/blob/([^/]+)/(.+)$`)

// GitHub repo root pattern: github.com/user/repo (with optional trailing slash)
var githubRepoRe = regexp.MustCompile(`^github\.com/([^/]+)/([^/]+?)/?$`)

// GitLab blob URL pattern: <host>/<project>/-/blob/<ref>/path
// Matches any host with /-/blob/ (GitLab-specific URL structure), supporting subgroups.
var gitlabBlobRe = regexp.MustCompile(`^([^/]+)/(.+?)/-/blob/([^/]+)/(.+)$`)

// GitLab repo root pattern: gitlab.com/user/repo (with optional trailing slash)
var gitlabRepoRe = regexp.MustCompile(`^gitlab\.com/([^/]+)/([^/]+?)/?$`)

// ResolveRawURL converts a GitHub/GitLab blob URL path to a raw content URL.
// path is the host + path portion (e.g., "github.com/user/repo/blob/main/README.md").
// Returns the raw URL and true if conversion was performed, or the original path and false.
func ResolveRawURL(path string) (string, bool) {
	// GitHub: github.com/user/repo/blob/<ref>/path → raw.githubusercontent.com/user/repo/<ref>/path
	if m := githubBlobRe.FindStringSubmatch(path); m != nil {
		user, repo, ref, filePath := m[1], m[2], m[3], m[4]
		rawURL := "raw.githubusercontent.com/" + user + "/" + repo + "/" + ref + "/" + filePath
		return rawURL, true
	}

	// GitLab: <host>/<project>/-/blob/<ref>/path → <host>/api/v4/projects/<encoded_project>/repository/files/<encoded_file>/raw?ref=<ref>
	// Uses the GitLab API endpoint instead of /-/raw/ because the web endpoint does not
	// accept Bearer token authentication on some self-hosted instances.
	if m := gitlabBlobRe.FindStringSubmatch(path); m != nil {
		host, project, ref, filePath := m[1], m[2], m[3], m[4]
		encodedProject := strings.ReplaceAll(project, "/", "%2F")
		encodedFilePath := strings.ReplaceAll(filePath, "/", "%2F")
		return host + "/api/v4/projects/" + encodedProject + "/repository/files/" + encodedFilePath + "/raw?ref=" + ref, true
	}

	return path, false
}

// ResolveRepoRootURLs returns candidate raw URLs for a repository root.
// For repo root URLs (github.com/user/repo), returns README.md URLs for main and master branches.
// Returns nil if the path is not a repo root.
func ResolveRepoRootURLs(path string) []string {
	// GitHub: github.com/user/repo → try main, then master
	if m := githubRepoRe.FindStringSubmatch(path); m != nil {
		user, repo := m[1], m[2]
		return []string{
			"raw.githubusercontent.com/" + user + "/" + repo + "/main/README.md",
			"raw.githubusercontent.com/" + user + "/" + repo + "/master/README.md",
		}
	}

	// GitLab: gitlab.com/user/repo → try main, then master (via API endpoint)
	if m := gitlabRepoRe.FindStringSubmatch(path); m != nil {
		user, repo := m[1], m[2]
		encodedProject := user + "%2F" + repo
		return []string{
			"gitlab.com/api/v4/projects/" + encodedProject + "/repository/files/README.md/raw?ref=main",
			"gitlab.com/api/v4/projects/" + encodedProject + "/repository/files/README.md/raw?ref=master",
		}
	}

	return nil
}

// HostFromPath extracts the host portion from a URL path.
// e.g., "github.com/user/repo/blob/main/file.md" → "github.com"
func HostFromPath(path string) string {
	idx := strings.Index(path, "/")
	if idx < 0 {
		return path
	}
	return path[:idx]
}

// PathFromPath extracts the path portion (after the host) from a URL path.
// e.g., "github.com/user/repo/blob/main/file.md" → "user/repo/blob/main/file.md"
func PathFromPath(path string) string {
	idx := strings.Index(path, "/")
	if idx < 0 {
		return ""
	}
	return path[idx+1:]
}
