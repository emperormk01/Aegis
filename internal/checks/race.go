package checks

import (
	"bytes"
	"io"
	"net/http"
	"sync"
	"time"
)

type RaceResult struct {
	Status int               `json:"status"`
	Length int               `json:"length"`
	Header map[string]string `json:"headers,omitempty"`
	Error  string            `json:"error,omitempty"`
}

func Race(method, url string, count int, body []byte, headers map[string]string, timeout time.Duration) []RaceResult {
	results := make([]RaceResult, count)
	var wg sync.WaitGroup
	client := &http.Client{Timeout: timeout}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(idx int) {
			defer wg.Done()
			req, err := http.NewRequest(method, url, bytes.NewReader(body))
			if err != nil {
				results[idx] = RaceResult{Error: err.Error()}
				return
			}
			for k, v := range headers {
				req.Header.Set(k, v)
			}
			resp, err := client.Do(req)
			if err != nil {
				results[idx] = RaceResult{Error: err.Error()}
				return
			}
			defer resp.Body.Close()
			b, _ := io.ReadAll(resp.Body)
			h := map[string]string{}
			for k, v := range resp.Header {
				if len(v) > 0 {
					h[k] = v[0]
				}
			}
			results[idx] = RaceResult{Status: resp.StatusCode, Length: len(b), Header: h}
		}(i)
	}
	wg.Wait()
	return results
}

type RaceSummary struct {
	Total         int         `json:"total"`
	ByStatus      map[int]int `json:"by_status"`
	UniqueLengths []int       `json:"unique_lengths"`
	Diverged      bool        `json:"diverged"`
}

func AnalyzeRace(results []RaceResult) RaceSummary {
	by := map[int]int{}
	lens := map[int]bool{}
	for _, r := range results {
		if r.Error != "" {
			by[0]++
			continue
		}
		by[r.Status]++
		lens[r.Length] = true
	}
	var uniq []int
	for k := range lens {
		uniq = append(uniq, k)
	}
	return RaceSummary{Total: len(results), ByStatus: by, UniqueLengths: uniq, Diverged: len(uniq) > 1}
}
