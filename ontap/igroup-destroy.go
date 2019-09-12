package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type IgroupDestroy struct {
	Base
	Params struct {
		XMLName xml.Name
		Force      bool   `xml:"force"`
		IgroupName string `xml:"initiator-group-name"`
	}
}

func (c *Client) IgroupDestroyAPI(igroupName string, force bool) (*SingleResultResponse, *http.Response, error) {
	if c.IgroupDestroy == nil {
		c.IgroupDestroy = &IgroupDestroy {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.IgroupDestroy.Params.XMLName = xml.Name{Local: "igroup-destroy"}
	}
	c.IgroupDestroy.Base.Name = c.vserver
	c.IgroupDestroy.Params.IgroupName = igroupName
	c.IgroupDestroy.Params.Force = force
	r := SingleResultResponse{}
	res, err := c.IgroupDestroy.get(c.IgroupDestroy, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(IgroupDestroyAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
