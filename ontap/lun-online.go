package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunOnline struct {
	Base
	Params struct {
		XMLName    xml.Name
		Force bool   `xml:"force,omitempty"`
		Path  string `xml:"path"`
	}
}

func (c *Client) LunOnlineAPI(lunPath string, force bool) (*SingleResultResponse, *http.Response, error) {
	if c.LunOnline == nil {
		c.LunOnline = &LunOnline {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunOnline.Params.XMLName = xml.Name{Local: "lun-online"}
	}
	c.LunOnline.Base.Name = c.vserver
	c.LunOnline.Params.Path = lunPath
	c.LunOnline.Params.Force = force
	r := SingleResultResponse{}
	res, err := c.LunOnline.get(c.LunOnline, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunOnlineAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
