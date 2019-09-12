package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunOffline struct {
	Base
	Params struct {
		XMLName    xml.Name
		Path string `xml:"path"`
	}
}

func (c *Client) LunOfflineAPI(lunPath string) (*SingleResultResponse, *http.Response, error) {
	if c.LunOffline == nil {
		c.LunOffline = &LunOffline {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunOffline.Params.XMLName = xml.Name{Local: "lun-offline"}
	}
	c.LunOffline.Base.Name = c.vserver
	c.LunOffline.Params.Path = lunPath
	r := SingleResultResponse{}
	res, err := c.LunOffline.get(c.LunOffline, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunOfflineAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
