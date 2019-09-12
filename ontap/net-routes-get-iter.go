package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type NetRoutesGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*NetRoutesGetOptions
	}
}

type NetRoutesQuery struct {
	NetVsRoutesInfo *NetVsRoutesInfo `xml:"net-vs-routes-info,omitempty"`
}

type NetRoutesGetOptions struct {
	DesiredAttributes *NetRoutesQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int             `xml:"max-records,omitempty"`
	Query             *NetRoutesQuery `xml:"query,omitempty"`
	Tag               string          `xml:"tag,omitempty"`
}

type NetVsRoutesInfo struct {
	AddressFamily string `xml:"address-family,omitempty"`
	Destination   string `xml:"destination,omitempty"`
	Gateway       string `xml:"gateway,omitempty"`
	Metric        int    `xml:"metric,omitempty"`
	Vserver       string `xml:"vserver,omitempty"`
}

type NetRoutesGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			NetRoutesAttributes []NetVsRoutesInfo `xml:"net-vs-routes-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) NetRoutesGetAPI(options *NetRoutesGetOptions) (*NetRoutesGetResponse, *http.Response, error) {
	if c.NetRoutesGetIter == nil {
		c.NetRoutesGetIter = &NetRoutesGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.NetRoutesGetIter.Params.XMLName = xml.Name{Local: "net-routes-get-iter"}
	}
	c.NetRoutesGetIter.Base.Name = c.vserver
	c.NetRoutesGetIter.Params.NetRoutesGetOptions = options
	r := NetRoutesGetResponse{}
	res, err := c.NetRoutesGetIter.get(c.NetRoutesGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(NetRoutesGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) NetRoutesGetIterAPI(options *NetRoutesGetOptions) (responseNetRoutes []*NetRoutesGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.NetRoutesGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseNetRoutes = append(responseNetRoutes, r)
			if nextTag == "" {
				break
			}
			options.Tag = nextTag
		} else {
			break
		}
	}
	return
}
