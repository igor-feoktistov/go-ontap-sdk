package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type FileCreateDirectory struct {
	Base
	Params struct {
		XMLName xml.Name
		Path string `xml:"path"`
		Perm string `xml:"perm"`
	}
}

func (c *Client) FileCreateDirectoryAPI(path string, perm string) (*SingleResultResponse, *http.Response, error) {
	if c.FileCreateDirectory == nil {
		c.FileCreateDirectory = &FileCreateDirectory {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.FileCreateDirectory.Params.XMLName = xml.Name{Local: "file-create-directory"}
	}
	c.FileCreateDirectory.Base.Name = c.vserver
	c.FileCreateDirectory.Params.Path = path
	c.FileCreateDirectory.Params.Perm = perm
	r := SingleResultResponse{}
	res, err := c.FileCreateDirectory.get(c.FileCreateDirectory, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(FileCreateDirectoryAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
