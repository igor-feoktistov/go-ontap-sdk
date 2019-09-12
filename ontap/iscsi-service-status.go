package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type IscsiServiceStatus struct {
	Base
	Params struct {
		XMLName xml.Name
	}
}

type IscsiServiceStatusResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		IsAvailable bool `xml:"is-available"`
	} `xml:"results"`
}

func (c *Client) IscsiServiceStatusAPI() (*IscsiServiceStatusResponse, *http.Response, error) {
	if c.IscsiServiceStatus == nil {
		c.IscsiServiceStatus = &IscsiServiceStatus {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.IscsiServiceStatus.Params.XMLName = xml.Name{Local: "iscsi-service-status"}
	}
	c.IscsiServiceStatus.Base.Name = c.vserver
	r := IscsiServiceStatusResponse{}
	res, err := c.IscsiServiceStatus.get(c.IscsiServiceStatus, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(IscsiServiceStatusAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
