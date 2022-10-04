package discordwebapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/SomethingBot/multierror"
)

// Bucket is the ratelimit bucket information from discord
type Bucket struct {
	limit      int
	remaining  int
	reset      time.Time
	resetAfter time.Duration
	bucket     string
	isGlobal   bool
	scope      string
}

// MarshalString from Bucket
func (b Bucket) MarshalString() string {
	bd := &bytes.Buffer{}
	fmt.Fprintf(bd, "limit %v\n", b.limit)
	fmt.Fprintf(bd, "remaining %v\n", b.remaining)
	fmt.Fprintf(bd, "reset %v\n", b.reset.String())
	fmt.Fprintf(bd, "resetAfter %v\n", b.resetAfter.String())
	fmt.Fprintf(bd, "bucket %v\n", b.bucket)
	fmt.Fprintf(bd, "isGlobal %v\n", b.isGlobal)
	fmt.Fprintf(bd, "scope %v\n", b.scope)
	return bd.String()
}

// UnmarshalRateLimitHeaders from the http headers
func UnmarshalRateLimitHeaders(headers http.Header) (Bucket, error) {
	h := headers.Get("X-RateLimit-Limit")
	if h == "" {
		return Bucket{}, nil
	}
	limit, err := strconv.Atoi(h)
	if err != nil {
		return Bucket{}, err
	}

	var b Bucket
	b.limit = limit

	h = headers.Get("X-RateLimit-Remaining")
	if h != "" {
		remaining, err := strconv.Atoi(h)
		if err != nil {
			return Bucket{}, err
		}
		b.remaining = remaining
	}

	h = headers.Get("X-RateLimit-Reset")
	if h != "" {
		resetf, err := strconv.ParseFloat(h, 64)
		if err != nil {
			return Bucket{}, err
		}
		i, d := math.Modf(resetf)
		b.reset = time.Unix(int64(i), int64(d))
	}

	h = headers.Get("X-RateLimit-Reset-After")
	if h != "" {
		resetf, err := strconv.ParseFloat(h, 64)
		if err != nil {
			return Bucket{}, err
		}
		b.resetAfter = time.Duration(resetf * 1e10)
	}

	h = headers.Get("X-RateLimit-Bucket")
	if h != "" {
		b.bucket = h
	}

	h = headers.Get("X-RateLimit-Global")
	if h != "" {
		b.isGlobal = true
	}

	h = headers.Get("X-RateLimit-Scope")
	if h != "" {
		b.scope = h
	}

	return b, nil
}

// RateLimiter is a function called to limit a request
type RateLimiter func(bucket string, f func()) error

// MemoryBucket describes a specific rate limited route
type MemoryBucket struct {
	Bucket
	mu sync.Mutex
}

// MemoryRateLimiter for multi-bucket support
type MemoryRateLimiter struct {
	mu      sync.Mutex
	buckets map[string]Bucket
}

// Run f() according to Bucket. If no room to run, blocks until room is avaiable and then runs f(), always returns nil
func (mb *MemoryBucket) Run(bucket string, f func()) error {
	return nil
}

// DoRequest and automatically handle ratelimiting, if err != nil you must close req as normal (ex: body.Close())
func DoRequest(ctx context.Context, limiter RateLimiter, httpClient *http.Client, req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	var bucket Bucket
	for {
		resp, err = httpClient.Do(req)
		if err != nil {
			break
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			//todo: currently we just hit and wait for the error, we need to premetively check for a bucket, although discord provides no satisfactory way to do this, so we dont do anything other than retry for now
			bucket, err = UnmarshalRateLimitHeaders(resp.Header)
			if err != nil {
				break
			}
			var body []byte
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				break
			}
			err = resp.Body.Close()
			if err != nil {
				break
			}
			jsonBody := struct{ RetryAfter time.Duration }{}
			err = json.Unmarshal(body, &jsonBody)
			if err != nil {
				break
			}
			fmt.Printf("rate limited: retrying after (%v)\n", jsonBody.RetryAfter)
			time.Sleep(jsonBody.RetryAfter)
		}

		break
	}
	if err != nil {
		err = fmt.Errorf("ratelimiter: headers (%#+v) error (%w)", bucket, err)
		err2 := resp.Body.Close() //todo: check if body already closed
		if err2 != nil {
			return nil, multierror.Append(err, err2)
		}
	}

	return resp, nil
}
