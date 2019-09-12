package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VolumeOnline struct {
	Base
	Params struct {
		XMLName xml.Name
		Name string `xml:"name"`
	}
}

func (c *Client) VolumeOnlineAPI(name string) (*SingleResultResponse, *http.Response, error) {
	if c.VolumeOnline == nil {
		c.VolumeOnline = &VolumeOnline {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VolumeOnline.Params.XMLName = xml.Name{Local: "volume-online"}
	}
	c.VolumeOnline.Base.Name = c.vserver
	c.VolumeOnline.Params.Name = name
	r := SingleResultResponse{}
	res, err := c.VolumeOnline.get(c.VolumeOnline, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VolumeOnlineAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
