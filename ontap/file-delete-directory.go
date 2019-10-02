package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type FileDeleteDirectory struct {
	Base
	Params struct {
		XMLName xml.Name
		Path string `xml:"path"`
	}
}

func (c *Client) FileDeleteDirectoryAPI(path string) (*SingleResultResponse, *http.Response, error) {
	if c.FileDeleteDirectory == nil {
		c.FileDeleteDirectory = &FileDeleteDirectory {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.FileDeleteDirectory.Params.XMLName = xml.Name{Local: "file-delete-directory"}
	}
	c.FileDeleteDirectory.Base.Name = c.vserver
	c.FileDeleteDirectory.Params.Path = path
	r := SingleResultResponse{}
	res, err := c.FileDeleteDirectory.get(c.FileDeleteDirectory, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(FileDeleteDirectoryAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
