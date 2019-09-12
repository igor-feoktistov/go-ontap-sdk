package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunMap struct {
	Base
	Params struct {
		XMLName    xml.Name
		LunMapOptions
	}
}

type LunMapOptions struct {
	AdditionalReportingNode bool   `xml:"additional-reporting-node,omitempty"`
	Force                   bool   `xml:"force,omitempty"`
	InitiatorGroup          string `xml:"initiator-group"`
	LunId                   int    `xml:"lun-id,omitempty"`
	Path                    string `xml:"path"`
}

type LunMapResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		LunIdAssigned int `xml:"lun-id-assigned"`
	} `xml:"results"`
}

func (c *Client) LunMapAPI(options *LunMapOptions) (*LunMapResponse, *http.Response, error) {
	if c.LunMap == nil {
		c.LunMap = &LunMap {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunMap.Params.XMLName = xml.Name{Local: "lun-map"}
	}
	c.LunMap.Base.Name = c.vserver
	c.LunMap.Params.LunMapOptions = *options
	r := LunMapResponse{}
	res, err := c.LunMap.get(c.LunMap, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunMapAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
