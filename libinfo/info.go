package libinfo

const (
	//BotLibraryUrl is the URL of the library
	BotLibraryUrl = "github.com/SomethingBot/dizzy"
	//BotLibraryVersion is the version of the library; this is the git tag
	BotLibraryVersion = "0.0.1"
	//BotUserAgent is the UserAgent of all library http requests
	BotUserAgent = "DiscordBot (" + BotLibraryUrl + ", " + BotLibraryVersion + ")"
)
