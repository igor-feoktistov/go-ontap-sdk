package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type FileWriteFile struct {
	Base
	Params struct {
		XMLName xml.Name
		FileWriteFileOptions
	}
}

type FileWriteFileOptions struct {
	Data       string `xml:"data"`
	Offset     int    `xml:"offset"`
	Overwrite  bool   `xml:"overwrite,omitempty"`
	Path       string `xml:"path"`
	StreamName string `xml:"stream-name,omitempty"`
}

type FileWriteFileResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		Length int `xml:"length"`
	} `xml:"results"`
}

func (c *Client) FileWriteFileAPI(options *FileWriteFileOptions) (*FileWriteFileResponse, *http.Response, error) {
	if c.FileWriteFile == nil {
		c.FileWriteFile = &FileWriteFile {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.FileWriteFile.Params.XMLName = xml.Name{Local: "file-write-file"}
	}
	c.FileWriteFile.Base.Name = c.vserver
	c.FileWriteFile.Params.FileWriteFileOptions = *options
	r := FileWriteFileResponse{}
	res, err := c.FileWriteFile.get(c.FileWriteFile, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(FileWriteFileAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
