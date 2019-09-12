package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunResize struct {
	Base
	Params struct {
		XMLName    xml.Name
		LunResizeOptions
	}
}

type LunResizeOptions struct {
	Force bool   `xml:"force,omitempty"`
	Path  string `xml:"path"`
	Size  int    `xml:"size"`
}

type LunResizeResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		ActualSize int `xml:"actual-size"`
	} `xml:"results"`
}

func (c *Client) LunResizeAPI(options *LunResizeOptions) (*LunResizeResponse, *http.Response, error) {
	if c.LunResize == nil {
		c.LunResize = &LunResize {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunResize.Params.XMLName = xml.Name{Local: "lun-resize"}
	}
	c.LunResize.Base.Name = c.vserver
	c.LunResize.Params.LunResizeOptions = *options
	r := LunResizeResponse{}
	res, err := c.LunResize.get(c.LunResize, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunResizeAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
