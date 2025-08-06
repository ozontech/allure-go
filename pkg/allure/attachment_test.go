package allure

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var mimeTypes = []MimeType{
	Text,
	Csv,
	Tsv,
	URIList,
	HTML,
	XML,
	JSON,
	Yaml,
	Pcap,
	Png,
	Jpg,
	Svg,
	Gif,
	Bmp,
	Tiff,
	Mp4,
	Ogg,
	Webm,
	Mpeg,
	Pdf,
	Xlsx,
}

func TestNewAttachment(t *testing.T) {
	for _, mt := range mimeTypes {
		testAttachName := fmt.Sprintf("Test init fileType: %s", mt)
		t.Run(testAttachName, func(t *testing.T) {
			content := []byte("some content")
			attachment := NewAttachment(testAttachName, mt, content)
			require.NotNil(t, attachment.UUID)
			require.Equal(t, mt, attachment.Type, "mime type should be same")
			require.Equal(t, testAttachName, attachment.Name)
			require.Equal(t, content, attachment.content)
			require.Contains(t, attachment.Source, mt.Ext())
		})
	}
}

func TestAttachment_Print(t *testing.T) {
	const testFolder = "allure-results"

	err := os.MkdirAll(testFolder, os.ModePerm)
	require.NoError(t, err)

	defer os.RemoveAll(testFolder)

	for _, mt := range mimeTypes {
		testAttachName := fmt.Sprintf("Test init fileType: %s", mt)
		t.Run(testAttachName, func(t *testing.T) {
			content := []byte("some content")
			attachment := NewAttachment(testAttachName, mt, content)
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
