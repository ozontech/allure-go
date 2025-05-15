package allure

import (
	"github.com/google/uuid"
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

// Ext returns file extension for this mime-type
func (mt MimeType) Ext() string {
	switch mt {
	case Text:
		return ".txt"
	case Csv:
		return ".csv"
	case Tsv:
		return ".tsv"
	case URIList:
		return ".uri"
	case HTML:
		return ".html"
	case XML:
		return ".xml"
	case JSON:
		return ".json"
	case Yaml:
		return ".yaml"
	case Pcap:
		return ".pcap"
	case Png:
		return ".png"
	case Jpg:
		return ".jpg"
	case Svg:
		return ".svg"
	case Gif:
		return ".gif"
	case Bmp:
		return ".bmp"
	case Tiff:
		return ".tiff"
	case Mp4:
		return ".mp4"
	case Ogg:
		return ".ogg"
	case Webm:
		return ".webm"
	case Mpeg:
		return ".mpeg"
	case Pdf:
		return ".pdf"
	case Xlsx:
		return ".xlsx"
	default:
		return ""
	}
}

// NewAttachment - Constructor. Returns pointer to new attachment object.
func NewAttachment(name string, mimeType MimeType, content []byte) *Attachment {
	id := uuid.New().String()

	return &Attachment{
		uuid:    id,
		content: content,
		Name:    name,
		Type:    mimeType,
		Source:  id + "-attachment" + mimeType.Ext(),
	}
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
