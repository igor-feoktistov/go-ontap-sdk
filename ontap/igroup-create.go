package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type IgroupCreate struct {
	Base
	Params struct {
		XMLName xml.Name
		BindPortset string `xml:"bind-portset,omitempty"`
		IgroupName  string `xml:"initiator-group-name"`
		IgroupType  string `xml:"initiator-group-type"`
		OsType      string `xml:"os-type"`
	}
}

func (c *Client) IgroupCreateAPI(igroupName string, igroupType string, osType string, bindPortset string) (*SingleResultResponse, *http.Response, error) {
	if c.IgroupCreate == nil {
		c.IgroupCreate = &IgroupCreate {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.IgroupCreate.Params.XMLName = xml.Name{Local: "igroup-create"}
	}
	c.IgroupCreate.Base.Name = c.vserver
	c.IgroupCreate.Params.IgroupName = igroupName
	c.IgroupCreate.Params.IgroupType = igroupType
	c.IgroupCreate.Params.OsType = osType
	c.IgroupCreate.Params.BindPortset = bindPortset
	r := SingleResultResponse{}
	res, err := c.IgroupCreate.get(c.IgroupCreate, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(IgroupCreateAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
