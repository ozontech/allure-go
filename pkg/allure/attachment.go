package allure

import (
	"fmt"
)

// Attachment - is an implementation of the attachments to the report in allure. It is most often used to contain
// screenshots, responses, files and other data obtained during the test.
type Attachment struct {
	Name    string   `json:"name,omitempty"`   // Attachment name
	Source  string   `json:"source,omitempty"` // Path to the Attachment file (name)
	Type    MimeType `json:"type,omitempty"`   // Mime-type of the Attachment
	uuid    string   // Unique identifier of the Attachment
	content []byte   // Attachment's content as bytes array
}

// MimeType is Attachment's mime type.
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
type MimeType string

// Attachment's MimeType constants
const (
	Text    MimeType = "text/plain"
	Csv     MimeType = "text/csv"
	Tsv     MimeType = "text/tab-separated-values"
	URIList MimeType = "text/uri-list"

	HTML MimeType = "text/html"
	XML  MimeType = "application/xml"
	JSON MimeType = "application/json"
	Yaml MimeType = "application/yaml"
	Pcap MimeType = "application/vnd.tcpdump.pcap"

	Png  MimeType = "image/png"
	Jpg  MimeType = "image/jpg"
	Svg  MimeType = "image/svg-xml"
	Gif  MimeType = "image/gif"
	Bmp  MimeType = "image/bmp"
	Tiff MimeType = "image/tiff"

	Mp4  MimeType = "video/mp4"
	Ogg  MimeType = "video/ogg"
	Webm MimeType = "video/webm"
	Mpeg MimeType = "video/mpeg"

	Pdf  MimeType = "application/pdf"
	Xlsx MimeType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
)

var mimeTypeMap = map[MimeType]string{
	Text:    "txt",
	Csv:     "csv",
	Tsv:     "tsv",
	URIList: "uri",
	HTML:    "html",
	XML:     "xml",
	JSON:    "json",
	Yaml:    "yaml",
	Pcap:    "pcap",
	Png:     "png",
	Jpg:     "jpg",
	Svg:     "svg",
	Gif:     "gif",
	Bmp:     "bmp",
	Tiff:    "tiff",
	Mp4:     "mp4",
	Ogg:     "ogg",
	Webm:    "webm",
	Mpeg:    "mpeg",
	Pdf:     "pdf",
	Xlsx:    "xlsx",
}

// NewAttachment - Constructor. Returns pointer to new attachment object.
func NewAttachment(name string, mimeType MimeType, content []byte) *Attachment {
	attachment := &Attachment{
		uuid:    getUUID().String(),
		content: content,
		Name:    name,
		Type:    mimeType,
	}
	attachment.Source = fmt.Sprintf("%s-attachment.%s", attachment.uuid, mimeTypeMap[attachment.Type])

	return attachment
}

func (a *Attachment) GetUUID() string {
	return a.uuid
}

func (a *Attachment) GetContent() []byte {
	return a.content
}

// Print - Creates a file from `Attachment.content`. The file type is determined by its `Attachment.mimeType`.
func (a *Attachment) Print() error {
	return NewFileManager().CreateFile(a.Source, a.content)
}
