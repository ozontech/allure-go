package allure

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAttachment(t *testing.T) {
	for mimeType, fileType := range mimeTypeMap {
		testAttachName := fmt.Sprintf("Test init fileType: %s", mimeType)
		t.Run(testAttachName, func(t *testing.T) {
			content := []byte("some content")
			attachment := NewAttachment(testAttachName, mimeType, content)
			require.NotNil(t, attachment.uuid)
			require.Equal(t, mimeType, attachment.Type, "mime type should be same")
			require.Equal(t, testAttachName, attachment.Name)
			require.Equal(t, content, attachment.content)
			require.Contains(t, attachment.Source, fileType)
		})
	}
}

func TestAttachment_Print(t *testing.T) {
	testFolder := "./allure-results/"
	isExists, err := exists(testFolder)
	if err != nil {
		panic(err)
	}
	if !isExists {
		_ = os.MkdirAll(testFolder, os.ModePerm)
	}
	defer os.RemoveAll(testFolder)
	for mimeType := range mimeTypeMap {
		testAttachName := fmt.Sprintf("Test init fileType: %s", mimeType)
		t.Run(testAttachName, func(t *testing.T) {
			content := []byte("some content")
			attachment := NewAttachment(testAttachName, mimeType, content)
			err := attachment.Print()
			require.Nil(t, err, "No errors expected")
			require.FileExists(t, fmt.Sprintf("./allure-results/%s", attachment.Source))
		})
	}
}

func TestAttachment_GetUUID(t *testing.T) {
	content := []byte("some content")
	testAttachName := fmt.Sprintf("Test init fileType: %s", Text)

	attachment := NewAttachment(testAttachName, Text, content)
	require.NotNil(t, attachment.GetUUID())
}

func TestAttachment_GetContent(t *testing.T) {
	content := []byte("some content")
	testAttachName := fmt.Sprintf("Test init fileType: %s", Text)

	attachment := NewAttachment(testAttachName, Text, content)
	require.Equal(t, content, attachment.GetContent())
}

func TestMimeTypeMap(t *testing.T) {
	t.Run("File format not changed", func(t *testing.T) {
		var _mimeTypeMap = map[MimeType]string{
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
		assert.Equal(t, len(_mimeTypeMap), len(mimeTypeMap), "Miss Some!")
		for _type, _format := range _mimeTypeMap {
			assert.Equal(t, _format, mimeTypeMap[_type], "should not ever be changed")
		}
	})

	t.Run("MimeTypes still same", func(t *testing.T) {
		var _Text MimeType = "text/plain"
		assert.Equal(t, _Text, Text)

		var _Csv MimeType = "text/csv"
		assert.Equal(t, _Csv, Csv)

		var _Tsv MimeType = "text/tab-separated-values"
		assert.Equal(t, _Tsv, Tsv)

		var _URIList MimeType = "text/uri-list"
		assert.Equal(t, _URIList, URIList)

		var _HTML MimeType = "text/html"
		assert.Equal(t, _HTML, HTML)

		var _XML MimeType = "application/xml"
		assert.Equal(t, _XML, XML)

		var _JSON MimeType = "application/json"
		assert.Equal(t, _JSON, JSON)

		var _Yaml MimeType = "application/yaml"
		assert.Equal(t, _Yaml, Yaml)

		var _Pcap MimeType = "application/vnd.tcpdump.pcap"
		assert.Equal(t, _Pcap, Pcap)

		var _Png MimeType = "image/png"
		assert.Equal(t, _Png, Png)

		var _Jpg MimeType = "image/jpg"
		assert.Equal(t, _Jpg, Jpg)

		var _Svg MimeType = "image/svg-xml"
		assert.Equal(t, _Svg, Svg)

		var _Gif MimeType = "image/gif"
		assert.Equal(t, _Gif, Gif)

		var _Bmp MimeType = "image/bmp"
		assert.Equal(t, _Bmp, Bmp)

		var _Tiff MimeType = "image/tiff"
		assert.Equal(t, _Tiff, Tiff)

		var _Mp4 MimeType = "video/mp4"
		assert.Equal(t, _Mp4, Mp4)

		var _Ogg MimeType = "video/ogg"
		assert.Equal(t, _Ogg, Ogg)

		var _Webm MimeType = "video/webm"
		assert.Equal(t, _Webm, Webm)

		var _Mpeg MimeType = "video/mpeg"
		assert.Equal(t, _Mpeg, Mpeg)

		var _Pdf MimeType = "application/pdf"
		assert.Equal(t, _Pdf, Pdf)

		var _Xlsx MimeType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
		assert.Equal(t, _Xlsx, Xlsx)
	})
}
