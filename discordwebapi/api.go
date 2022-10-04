package discordwebapi

const DiscordAPIURL = "https://discord.com/api"

type API struct {
	ApiKey  string
	BaseURL string
	Limiter RateLimiter
}
