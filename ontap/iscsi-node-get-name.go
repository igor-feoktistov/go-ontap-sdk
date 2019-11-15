package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type IscsiNodeGetName struct {
	Base
	Params struct {
		XMLName xml.Name
	}
}

type IscsiNodeGetNameResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		NodeName string `xml:"node-name"`
	} `xml:"results"`
}

func (c *Client) IscsiNodeGetNameAPI() (*IscsiNodeGetNameResponse, *http.Response, error) {
	if c.IscsiNodeGetName == nil {
		c.IscsiNodeGetName = &IscsiNodeGetName {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.IscsiNodeGetName.Params.XMLName = xml.Name{Local: "iscsi-node-get-name"}
	}
	c.IscsiNodeGetName.Base.Name = c.vserver
	r := IscsiNodeGetNameResponse{}
	res, err := c.IscsiNodeGetName.get(c.IscsiNodeGetName, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(IscsiNodeGetNameAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
