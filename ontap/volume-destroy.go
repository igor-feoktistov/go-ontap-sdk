package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VolumeDestroy struct {
	Base
	Params struct {
		XMLName xml.Name
		Name string `xml:"name"`
	}
}

func (c *Client) VolumeDestroyAPI(name string) (*SingleResultResponse, *http.Response, error) {
	if c.VolumeDestroy == nil {
		c.VolumeDestroy = &VolumeDestroy {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VolumeDestroy.Params.XMLName = xml.Name{Local: "volume-destroy"}
	}
	c.VolumeDestroy.Base.Name = c.vserver
	c.VolumeDestroy.Params.Name = name
	r := SingleResultResponse{}
	res, err := c.VolumeDestroy.get(c.VolumeDestroy, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VolumeDestroyAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
