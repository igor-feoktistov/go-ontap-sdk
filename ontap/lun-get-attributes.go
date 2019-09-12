package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunGetAttributes struct {
	Base
	Params struct {
		XMLName    xml.Name
		Name  string `xml:"name,omitempty"`
		Path  string `xml:"path"`
	}
}

type LunGetAttributesResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		Attributes struct {
			LunAttributes []LunAttributeInfo `xml:"lun-attribute-info"`
		} `xml:"attributes"`
	} `xml:"results"`
}

type LunAttributeInfo struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

func (c *Client) LunGetAttributesAPI(lunPath string, attrName string) (*LunGetAttributesResponse, *http.Response, error) {
	if c.LunGetAttributes == nil {
		c.LunGetAttributes = &LunGetAttributes {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunGetAttributes.Params.XMLName = xml.Name{Local: "lun-get-attributes"}
	}
	c.LunGetAttributes.Base.Name = c.vserver
	c.LunGetAttributes.Params.Path = lunPath
	c.LunGetAttributes.Params.Name = attrName
	r := LunGetAttributesResponse{}
	res, err := c.LunGetAttributes.get(c.LunGetAttributes, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunGetAttributesAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
