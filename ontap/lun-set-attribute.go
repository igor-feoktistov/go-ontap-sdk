package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunSetAttribute struct {
	Base
	Params struct {
		XMLName    xml.Name
		Name  string `xml:"name"`
		Path  string `xml:"path"`
		Value string `xml:"value"`
	}
}

func (c *Client) LunSetAttributeAPI(lunPath string, attrName string, attrValue string) (*SingleResultResponse, *http.Response, error) {
	if c.LunSetAttribute == nil {
		c.LunSetAttribute = &LunSetAttribute {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunSetAttribute.Params.XMLName = xml.Name{Local: "lun-set-attribute"}
	}
	c.LunSetAttribute.Base.Name = c.vserver
	c.LunSetAttribute.Params.Path = lunPath
	c.LunSetAttribute.Params.Name = attrName
	c.LunSetAttribute.Params.Value = attrValue
	r := SingleResultResponse{}
	res, err := c.LunSetAttribute.get(c.LunSetAttribute, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunSetAttributeAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
