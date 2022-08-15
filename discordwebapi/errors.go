package discordwebapi

type WebAPIError struct { //todo: convert to real error type
	JsonData []byte
}

func (wae WebAPIError) Error() string {
	return string(wae.JsonData)
}
