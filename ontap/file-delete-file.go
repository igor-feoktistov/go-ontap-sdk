package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type FileDeleteFile struct {
	Base
	Params struct {
		XMLName xml.Name
		Path string `xml:"path"`
	}
}

func (c *Client) FileDeleteFileAPI(filePath string) (*SingleResultResponse, *http.Response, error) {
	if c.FileDeleteFile == nil {
		c.FileDeleteFile = &FileDeleteFile {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.FileDeleteFile.Params.XMLName = xml.Name{Local: "file-delete-file"}
	}
	c.FileDeleteFile.Base.Name = c.vserver
	c.FileDeleteFile.Params.Path = filePath
	r := SingleResultResponse{}
	res, err := c.FileDeleteFile.get(c.FileDeleteFile, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(FileDeleteFileAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
