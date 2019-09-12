package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type FileReadFile struct {
	Base
	Params struct {
		XMLName xml.Name
		FileReadFileOptions
	}
}

type FileReadFileOptions struct {
	Length     int    `xml:"length"`
	Offset     int    `xml:"offset"`
	Path       string `xml:"path"`
	StreamName string `xml:"stream-name,omitempty"`
}

type FileReadFileResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		Data   string `xml:"data"`
		Length int    `xml:"length"`
	} `xml:"results"`
}

func (c *Client) FileReadFileAPI(options *FileReadFileOptions) (*FileReadFileResponse, *http.Response, error) {
	if c.FileReadFile == nil {
		c.FileReadFile = &FileReadFile {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.FileReadFile.Params.XMLName = xml.Name{Local: "file-read-file"}
	}
	c.FileReadFile.Base.Name = c.vserver
	c.FileReadFile.Params.FileReadFileOptions = *options
	r := FileReadFileResponse{}
	res, err := c.FileReadFile.get(c.FileReadFile, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(FileReadFileAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
