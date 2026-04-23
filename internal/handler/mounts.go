package handler

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type mountInfo struct {
	Source string // e.g. "user@host:/remote/path" or "10.0.0.1:/srv/x"
	Target string // e.g. "/Users/foo/mnt/bar"
	FSType string // e.g. "fuse.sshfs", "nfs"
}

// remoteFS lists filesystem types whose source field carries a useful
// "origin" for the mount point. Local filesystems like apfs/ext4/tmpfs
// are intentionally excluded so we fall back to hostname:path.
var remoteFS = map[string]bool{
	"fuse.sshfs":    true,
	"sshfs":         true,
	"nfs":           true,
	"nfs4":          true,
	"cifs":          true,
	"smbfs":         true,
	"smb3":          true,
	"9p":            true,
	"fuse.gcsfuse":  true,
	"fuse.s3fs":     true,
	"fuse.rclone":   true,
	"fuse.smbnetfs": true,
	"macfuse":       true,
	"osxfuse":       true,
}

func listMounts() []mountInfo {
	switch runtime.GOOS {
	case "linux":
		return parseProcMounts()
	case "darwin":
		return parseMountCmd()
	default:
		return nil
	}
}

func parseProcMounts() []mountInfo {
	data, err := os.ReadFile("/proc/mounts")
	if err != nil {
		return nil
	}
	var ms []mountInfo
	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		ms = append(ms, mountInfo{
			Source: unescapeProcField(fields[0]),
			Target: unescapeProcField(fields[1]),
			FSType: fields[2],
		})
	}
	return ms
}

// unescapeProcField decodes octal escapes like "\040" ('space') that
// /proc/mounts uses for whitespace in paths.
func unescapeProcField(s string) string {
	if !strings.Contains(s, `\`) {
		return s
	}
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+3 < len(s) {
			n := 0
			ok := true
			for j := 1; j <= 3; j++ {
				c := s[i+j]
				if c < '0' || c > '7' {
					ok = false
					break
				}
				n = n*8 + int(c-'0')
			}
			if ok {
				b.WriteByte(byte(n))
				i += 3
				continue
			}
		}
		b.WriteByte(s[i])
	}
	return b.String()
}

func parseMountCmd() []mountInfo {
	out, err := exec.Command("mount").Output()
	if err != nil {
		return nil
	}
	var ms []mountInfo
	for _, line := range strings.Split(string(out), "\n") {
		// Format: "source on target (fstype, options)"
		idx := strings.Index(line, " on ")
		if idx < 0 {
			continue
		}
		src := line[:idx]
		rest := line[idx+4:]
		paren := strings.LastIndex(rest, " (")
		if paren < 0 {
			continue
		}
		target := rest[:paren]
		inside := strings.TrimSuffix(rest[paren+2:], ")")
		fsType := inside
		if i := strings.IndexByte(inside, ','); i > 0 {
			fsType = inside[:i]
		}
		ms = append(ms, mountInfo{Source: src, Target: target, FSType: strings.TrimSpace(fsType)})
	}
	return ms
}

// findMountOrigin looks up the longest mount point containing path. If that
// mount is a remote filesystem (see remoteFS), return a best-effort origin
// string like "user@host:/remote/path"; otherwise return "" so the caller
// can fall back to its own default (typically hostname:path).
func findMountOrigin(path string, mounts []mountInfo) string {
	var best *mountInfo
	bestLen := -1
	for i := range mounts {
		m := &mounts[i]
		if m.Target == path || strings.HasPrefix(path, m.Target+"/") {
			if len(m.Target) > bestLen {
				best = m
				bestLen = len(m.Target)
			}
		}
	}
	if best == nil || !remoteFS[best.FSType] {
		return ""
	}
	src := best.Source
	if path != best.Target && strings.HasPrefix(path, best.Target+"/") {
		src = src + path[len(best.Target):]
	}
	return src
}
