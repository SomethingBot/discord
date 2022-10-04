package discordwebapi

var noOpRateLimiter = func(bucket string, f func()) error { return nil }
