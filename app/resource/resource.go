package resource

import "embed"

//go:embed html
var html embed.FS

func GetHtmlResource() embed.FS {
	return html
}
