package primitives

import (
	"errors"
	"time"
	"unicode/utf8"
)

//EmbedType is documented at https://discord.com/developers/docs/resources/channel#embed-object-embed-types
type EmbedType string

const (
	EmbedTypeNil     EmbedType = "null"
	EmbedTypeRich    EmbedType = "rich"
	EmbedTypeImage   EmbedType = "image"
	EmbedTypeVideo   EmbedType = "video"
	EmbedTypeGifv    EmbedType = "gifv"
	EmbedTypeArticle EmbedType = "article"
	EmbedTypeLink    EmbedType = "link"
)

//IsValid EmbedType
func (e EmbedType) IsValid() bool {
	switch e {
	case EmbedTypeRich,
		EmbedTypeImage,
		EmbedTypeVideo,
		EmbedTypeGifv,
		EmbedTypeArticle,
		EmbedTypeLink:
		return true
	default:
		return false
	}
}

//EmbedFooter documented at https://discord.com/developers/docs/resources/channel#embed-object-embed-footer-structure
type EmbedFooter struct {
	//Text of EmbedFooter
	Text string `json:"text"`
	//IconURL of EmbedFooter
	IconURL string `json:"icon_url"`
	//ProxyIconUrl for EmbedFooter
	ProxyIconURL string `json:"proxy_icon_url"`
}

//EmbedImage documented at https://discord.com/developers/docs/resources/channel#embed-object-embed-image-structure
type EmbedImage struct {
	//URL for EmbedImage
	URL string `json:"url"`
	//ProxyURL for EmbedImage
	ProxyURL string `json:"proxy_url"`
	//Height of EmbedImage
	Height int `json:"height"`
	//Width of EmbedImage
	Width int `json:"width"`
}

//EmbedThumbnail documented at https://discord.com/developers/docs/resources/channel#embed-object-embed-thumbnail-structure
type EmbedThumbnail struct {
	//URL for EmbedThumbnail
	URL string `json:"url"`
	//ProxyURL for EmbedThumbnail
	ProxyURL string `json:"proxy_url"`
	//Height of EmbedThumbnail
	Height int `json:"height"`
	//Width of EmbedThumbnail
	Width int `json:"width"`
}

//EmbedVideo documented at https://discord.com/developers/docs/resources/channel#embed-object-embed-video-structure
type EmbedVideo struct {
	//URL for EmbedVideo
	URL string `json:"url"`
	//ProxyURL for EmbedVideo
	ProxyURL string `json:"proxy_url"`
	//Height of EmbedVideo
	Height int `json:"height"`
	//Width of EmbedVideo
	Width int `json:"width"`
}

//EmbedProvider documented at https://discord.com/developers/docs/resources/channel#embed-object-embed-provider-structure
type EmbedProvider struct {
	//Name of EmbedProvider
	Name string `json:"name"`
	//URL for EmbedProvider
	URL string `json:"url"`
}

//EmbedAuthor documented at https://discord.com/developers/docs/resources/channel#embed-object-embed-author-structure
type EmbedAuthor struct {
	//Name of EmbedAuthor
	Name string `json:"name"`
	//URL for EmbedAuthor
	URL string `json:"url"`
	//IconURL of EmbedAuthor
	IconURL string `json:"icon_url"`
	//ProxyIconUrl for EmbedAuthor
	ProxyIconURL string `json:"proxy_icon_url"`
}

//EmbedField documented at https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
type EmbedField struct {
	//Name of EmbedField
	Name string `json:"name"`
	//Value of EmbedField
	Value string `json:"value"`
	//IsInline displayed
	IsInline bool `json:"inline"`
}

//Embed documented https://discord.com/developers/docs/resources/channel#embed-object
type Embed struct {
	//Title of Embed
	Title string `json:"title"`
	//Type of Embed, always EmbedTypeRich for a webhook Embed
	Type EmbedType `json:"type"`
	//Description of Embed
	Description string `json:"description"`
	//URL of Embed
	URL string `json:"URL"`
	//Timestamp of Embed
	Timestamp time.Time `json:"timestamp"`
	//ColorCode of Embed
	ColorCode int `json:"color"`
	//Footer of Embed
	Footer EmbedFooter `json:"footer"`
	//Image of Embed
	Image EmbedImage `json:"image"`
	//Thumbnail of Embed
	Thumbnail EmbedThumbnail `json:"thumbnail"`
	//Video of Embed
	Video EmbedVideo `json:"video"`
	//Provider of Embed
	Provider EmbedProvider `json:"provider"`
	//Author of Embed
	Author EmbedAuthor `json:"author"`
	//Fields of Embed
	Fields []EmbedField `json:"fields"`
}

var (
	//ErrorInvalidUTF8 is when a message does not contain valid utf8
	ErrorInvalidUTF8 = errors.New("primitives: message contains invalid utf8")
	//ErrorEmbedTooLarge is when the combined sum of characters in Embed.Title, Embed.Description, all EmbedFields (EmbedField.Name, EmbedField.Value), EmbedFooter.Text, and EmbedAuthor.Name fields across all embeds attached to a message exceeds 6000 characters
	ErrorEmbedTooLarge = errors.New("primitives: embed data too large")
	//ErrorEmbedTitleTooLarge is when Embed.Title is over 256 characters
	ErrorEmbedTitleTooLarge = errors.New("primitives: embed title too large")
	//ErrorEmbedDescriptionTooLarge is when Embed.Description is over 4096 characters
	ErrorEmbedDescriptionTooLarge = errors.New("primitives: embed description too large")
	//ErrorEmbedFieldsTooLarge is when len(Embed.Fields) > 25
	ErrorEmbedFieldsTooLarge = errors.New("primitives: too many embed fields")
	//ErrorEmbedFieldNameTooLarge is when EmbedField.Name is over 256 characters
	ErrorEmbedFieldNameTooLarge = errors.New("primitives: embedfield name too large")
	//ErrorEmbedFieldValueTooLarge is when EmbedField.Value is over 1024 characters
	ErrorEmbedFieldValueTooLarge = errors.New("primitives: embedfield value too large")
	//ErrorEmbedFooterTextTooLarge is when EmbedFooter.Text is over 2048 characters
	ErrorEmbedFooterTextTooLarge = errors.New("primitives: embed footer text too large")
	//ErrorEmbedAuthorNameTooLarge is when EmbedAuthor.Name is over 256 characters
	ErrorEmbedAuthorNameTooLarge = errors.New("primitives: embed author name too large")
)

//IsValid Embed following https://discord.com/developers/docs/resources/channel#embed-object-embed-limits, nil if valid, don't know if discord considers a rune a character or an utf-8 character, assuming utf-8
func (e Embed) IsValid() error {
	var embedSize int

	if !utf8.ValidString(e.Title) {
		return ErrorInvalidUTF8
	}
	titleSize := len([]byte(e.Title))
	if titleSize > 256 {
		return ErrorEmbedTitleTooLarge
	}
	embedSize += titleSize

	if !utf8.ValidString(e.Description) {
		return ErrorInvalidUTF8
	}
	descriptionSize := len([]byte(e.Description))
	if descriptionSize > 4096 {
		return ErrorEmbedDescriptionTooLarge
	}
	embedSize += descriptionSize

	if len(e.Fields) > 25 {
		return ErrorEmbedFieldsTooLarge
	}

	var fieldNameLength, fieldValueLength int
	for _, f := range e.Fields {
		if !utf8.ValidString(f.Name) {
			return ErrorInvalidUTF8
		}
		fieldNameLength = len(f.Name)
		if fieldNameLength > 256 {
			return ErrorEmbedFieldNameTooLarge
		}

		if !utf8.ValidString(f.Value) {
			return ErrorInvalidUTF8
		}
		fieldValueLength = len(f.Value)
		if fieldValueLength > 1024 {
			return ErrorEmbedFieldValueTooLarge
		}
		embedSize += fieldNameLength + fieldValueLength
	}

	if !utf8.ValidString(e.Footer.Text) {
		return ErrorInvalidUTF8
	}
	footerSize := len(e.Footer.Text)
	if footerSize > 2048 {
		return ErrorEmbedFooterTextTooLarge
	}
	embedSize += footerSize

	if !utf8.ValidString(e.Author.Name) {
		return ErrorInvalidUTF8
	}
	authorNameSize := len(e.Author.Name)
	if authorNameSize > 256 {
		return ErrorEmbedAuthorNameTooLarge
	}
	embedSize += footerSize

	if embedSize > 6000 {
		return ErrorEmbedTooLarge
	}

	return nil
}
