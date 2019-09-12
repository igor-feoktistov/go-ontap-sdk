package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type IgroupAdd struct {
	Base
	Params struct {
		XMLName xml.Name
		Force      bool   `xml:"force"`
		Initiator  string `xml:"initiator"`
		IgroupName string `xml:"initiator-group-name"`
	}
}

func (c *Client) IgroupAddAPI(igroupName string, initiatorName string, force bool) (*SingleResultResponse, *http.Response, error) {
	if c.IgroupAdd == nil {
		c.IgroupAdd = &IgroupAdd {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.IgroupAdd.Params.XMLName = xml.Name{Local: "igroup-add"}
	}
	c.IgroupAdd.Base.Name = c.vserver
	c.IgroupAdd.Params.IgroupName = igroupName
	c.IgroupAdd.Params.Initiator = initiatorName
	c.IgroupAdd.Params.Force = force
	r := SingleResultResponse{}
	res, err := c.IgroupAdd.get(c.IgroupAdd, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(IgroupAddAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
