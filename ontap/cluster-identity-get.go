package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type ClusterIdentityGet struct {
	Base
	Params struct {
		XMLName xml.Name
		*ClusterIdentityOptions
	}
}

type ClusterIdentityInfo struct {
	ClusterContact      string `xml:"cluster-contact,omitempty"`
	ClusterLocation     string `xml:"cluster-location"`
	ClusterName         string `xml:"cluster-name"`
	ClusterSerialNumber string `xml:"cluster-serial-number"`
	RdbUuid             string `xml:"rdb-uuid"`
	UUID                string `xml:"uuid"`
}

type ClusterIdentityOptions struct {
	DesiredAttributes *ClusterIdentityInfo `xml:"desired-attributes,omitempty"`
}

type ClusterIdentityResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		ClusterIdentityInfo []ClusterIdentityInfo `xml:"attributes>cluster-identity-info"`
	} `xml:"results"`
}

func (c *Client) ClusterIdentityGetAPI(options *ClusterIdentityOptions) (*ClusterIdentityResponse, *http.Response, error) {
	if c.ClusterIdentityGet == nil {
		c.ClusterIdentityGet = &ClusterIdentityGet {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.ClusterIdentityGet.Params.XMLName = xml.Name{Local: "cluster-identity-get"}
	}
	c.ClusterIdentityGet.Params.ClusterIdentityOptions = options
	r := ClusterIdentityResponse{}
	res, err := c.ClusterIdentityGet.get(c.ClusterIdentityGet, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(ClusterIdentityGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
