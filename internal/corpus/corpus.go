package corpus

import (
	"embed"
	"io/fs"
	"path"
	"strings"
)

//go:embed bugbounty-findings/*.md kambegoye-scan/* security-research/*.md bugbounty-skill/*
var FS embed.FS

func List() []string {
	var out []string
	fs.WalkDir(FS, ".", func(p string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			out = append(out, p)
		}
		return nil
	})
	return out
}

func Read(p string) (string, error) {
	b, err := fs.ReadFile(FS, p)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func Search(keyword string) []string {
	kw := strings.ToLower(keyword)
	var hits []string
	for _, p := range List() {
		content, _ := Read(p)
		if strings.Contains(strings.ToLower(content), kw) || strings.Contains(strings.ToLower(path.Base(p)), kw) {
			hits = append(hits, p)
		}
	}
	return hits
}
