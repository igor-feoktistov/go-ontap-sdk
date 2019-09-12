package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type IgroupLookupLun struct {
	Base
	Params struct {
		XMLName xml.Name
		IgroupName string `xml:"initiator-group-name"`
		LunId      int    `xml:"lun-id"`
	}
}

type IgroupLookupLunResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		LunPath string `xml:"path"`
	} `xml:"results"`
}

func (c *Client) IgroupLookupLunAPI(igroupName string, lunId int) (*IgroupLookupLunResponse, *http.Response, error) {
	if c.IgroupLookupLun == nil {
		c.IgroupLookupLun = &IgroupLookupLun {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.IgroupLookupLun.Params.XMLName = xml.Name{Local: "igroup-lookup-lun"}
	}
	c.IgroupLookupLun.Base.Name = c.vserver
	c.IgroupLookupLun.Params.IgroupName = igroupName
	c.IgroupLookupLun.Params.LunId = lunId
	r := IgroupLookupLunResponse{}
	res, err := c.IgroupLookupLun.get(c.IgroupLookupLun, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(IgroupLookupLunAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
