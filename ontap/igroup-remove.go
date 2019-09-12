package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type IgroupRemove struct {
	Base
	Params struct {
		XMLName xml.Name
		Force      bool   `xml:"force"`
		Initiator  string `xml:"initiator"`
		IgroupName string `xml:"initiator-group-name"`
	}
}

func (c *Client) IgroupRemoveAPI(igroupName string, initiatorName string, force bool) (*SingleResultResponse, *http.Response, error) {
	if c.IgroupRemove == nil {
		c.IgroupRemove = &IgroupRemove {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.IgroupRemove.Params.XMLName = xml.Name{Local: "igroup-remove"}
	}
	c.IgroupRemove.Base.Name = c.vserver
	c.IgroupRemove.Params.IgroupName = igroupName
	c.IgroupRemove.Params.Initiator = initiatorName
	c.IgroupRemove.Params.Force = force
	r := SingleResultResponse{}
	res, err := c.IgroupRemove.get(c.IgroupRemove, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(IgroupRemoveAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
