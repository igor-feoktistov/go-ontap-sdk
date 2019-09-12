package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VolumeUnmount struct {
	Base
	Params struct {
		XMLName xml.Name
		Force      bool   `xml:"force,omitempty"`
		VolumeName string `xml:"volume-name"`
	}
}

func (c *Client) VolumeUnmountAPI(volumeName string, force bool) (*SingleResultResponse, *http.Response, error) {
	if c.VolumeUnmount == nil {
		c.VolumeUnmount = &VolumeUnmount {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VolumeUnmount.Params.XMLName = xml.Name{Local: "volume-unmount"}
	}
	c.VolumeUnmount.Base.Name = c.vserver
	c.VolumeUnmount.Params.VolumeName = volumeName
	c.VolumeUnmount.Params.Force = force
	r := SingleResultResponse{}
	res, err := c.VolumeUnmount.get(c.VolumeUnmount, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VolumeUnmountAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
