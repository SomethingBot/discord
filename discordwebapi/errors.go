// Package discordwebapi is a ratelimited implementation of the discord REST API
package discordwebapi

// WebAPIError is an error returned by discord's REST api
type WebAPIError struct { //todo: convert to real error type
	JSONData []byte
}

// Error string from Discord
func (wae WebAPIError) Error() string {
	return string(wae.JSONData)
}
