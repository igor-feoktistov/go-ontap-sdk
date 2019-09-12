package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunGetAttribute struct {
	Base
	Params struct {
		XMLName    xml.Name
		Name  string `xml:"name"`
		Path  string `xml:"path"`
	}
}

type LunGetAttributeResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		Value string `xml:"value"`
	} `xml:"results"`
}

func (c *Client) LunGetAttributeAPI(lunPath string, attrName string) (*LunGetAttributeResponse, *http.Response, error) {
	if c.LunGetAttribute == nil {
		c.LunGetAttribute = &LunGetAttribute {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunGetAttribute.Params.XMLName = xml.Name{Local: "lun-get-attribute"}
	}
	c.LunGetAttribute.Base.Name = c.vserver
	c.LunGetAttribute.Params.Path = lunPath
	c.LunGetAttribute.Params.Name = attrName
	r := LunGetAttributeResponse{}
	res, err := c.LunGetAttribute.get(c.LunGetAttribute, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunGetAttributeAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
