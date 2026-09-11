package checks

import (
	"net/http"
	"strings"
	"time"
)

var permutations = []string{
	"{base}",
	"{base}-static",
	"{base}-assets",
	"{base}-uploads",
	"{base}-logs",
	"{base}-backup-2026",
	"{base}-prod",
	"{base}-dev",
	"{base}-staging",
}

type S3Result struct {
	Bucket   string `json:"bucket"`
	Status   string `json:"status"`
	HTTP     int    `json:"http"`
	Listable bool   `json:"listable"`
	URL      string `json:"url"`
}

func CheckBucket(bucket string, timeout time.Duration) S3Result {
	url := "https://" + bucket + ".s3.amazonaws.com/"
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		return S3Result{Bucket: bucket, Status: "error", URL: url}
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	n, _ := resp.Body.Read(buf)
	body := strings.ToLower(string(buf[:n]))
	listable := strings.Contains(body, "<listbucketresult") || strings.Contains(body, "<contents>")
	switch resp.StatusCode {
	case 200:
		return S3Result{Bucket: bucket, Status: "exists", HTTP: 200, Listable: listable, URL: url}
	case 403:
		return S3Result{Bucket: bucket, Status: "exists", HTTP: 403, Listable: false, URL: url}
	case 404:
		return S3Result{Bucket: bucket, Status: "not_found", HTTP: 404, URL: url}
	default:
		return S3Result{Bucket: bucket, Status: "unknown", HTTP: resp.StatusCode, URL: url}
	}
}

func EnumerateS3(base string, timeout time.Duration) []S3Result {
	var out []S3Result
	for _, pat := range permutations {
		bucket := strings.ReplaceAll(pat, "{base}", base)
		out = append(out, CheckBucket(bucket, timeout))
	}
	return out
}
