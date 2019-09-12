package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunDestroy struct {
	Base
	Params struct {
		XMLName    xml.Name
		LunDestroyOptions
	}
}

type LunDestroyOptions struct {
	DestroyApplicationLun bool   `xml:"destroy-application-lun,omitempty"`
	DestroyFencedLun      bool   `xml:"destroy-fenced-lun,omitempty"`
	Force                 bool   `xml:"force,omitempty"`
	Path                  string `xml:"path"`
}

func (c *Client) LunDestroyAPI(options *LunDestroyOptions) (*SingleResultResponse, *http.Response, error) {
	if c.LunDestroy == nil {
		c.LunDestroy = &LunDestroy {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunDestroy.Params.XMLName = xml.Name{Local: "lun-destroy"}
	}
	c.LunDestroy.Base.Name = c.vserver
	c.LunDestroy.Params.LunDestroyOptions = *options
	r := SingleResultResponse{}
	res, err := c.LunDestroy.get(c.LunDestroy, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunDestroyAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
