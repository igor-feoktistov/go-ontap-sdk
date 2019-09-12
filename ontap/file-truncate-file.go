package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type FileTruncateFile struct {
	Base
	Params struct {
		XMLName xml.Name
		Path string `xml:"path"`
		Size int    `xml:"size,omitempty"`
	}
}

func (c *Client) FileTruncateFileAPI(filePath string, fileSize int) (*SingleResultResponse, *http.Response, error) {
	if c.FileTruncateFile == nil {
		c.FileTruncateFile = &FileTruncateFile {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.FileTruncateFile.Params.XMLName = xml.Name{Local: "file-truncate-file"}
	}
	c.FileTruncateFile.Base.Name = c.vserver
	c.FileTruncateFile.Params.Path = filePath
	c.FileTruncateFile.Params.Size = fileSize
	r := SingleResultResponse{}
	res, err := c.FileTruncateFile.get(c.FileTruncateFile, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(FileTruncateFileAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
