package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VolumeOffline struct {
	Base
	Params struct {
		XMLName xml.Name
		Name string `xml:"name"`
	}
}

func (c *Client) VolumeOfflineAPI(name string) (*SingleResultResponse, *http.Response, error) {
	if c.VolumeOffline == nil {
		c.VolumeOffline = &VolumeOffline {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VolumeOffline.Params.XMLName = xml.Name{Local: "volume-offline"}
	}
	c.VolumeOffline.Base.Name = c.vserver
	c.VolumeOffline.Params.Name = name
	r := SingleResultResponse{}
	res, err := c.VolumeOffline.get(c.VolumeOffline, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VolumeOfflineAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
