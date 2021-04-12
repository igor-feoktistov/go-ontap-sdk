package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunCreateBySize struct {
	Base
	Params struct {
		XMLName    xml.Name
		LunCreateBySizeOptions
	}
}

type LunCreateBySizeOptions struct {
	Class                   string `xml:"class,omitempty"`
	Comment                 string `xml:"comment,omitempty"`
	ForeignDisk             string `xml:"foreign-disk,omitempty"`
	OsType                  string `xml:"ostype,omitempty"`
	Path                    string `xml:"path"`
	PrefixSize              int    `xml:"prefix-size,omitempty"`
	QosAdaptivePolicyGroup  string `xml:"qos-adaptive-policy-group,omitempty"`
	QosPolicyGroup          string `xml:"qos-policy-group,omitempty"`
	Size                    int    `xml:"size,omitempty"`
	SpaceAllocationEnabled  bool   `xml:"space-allocation-enabled"`
	SpaceReservationEnabled bool   `xml:"space-reservation-enabled"`
	UseExactSize            bool   `xml:"use-exact-size,omitempty"`
}

type LunCreateBySizeResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		ActualSize int `xml:"actual-size"`
	} `xml:"results"`
}

func (c *Client) LunCreateBySizeAPI(options *LunCreateBySizeOptions) (*LunCreateBySizeResponse, *http.Response, error) {
	if c.LunCreateBySize == nil {
		c.LunCreateBySize = &LunCreateBySize {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunCreateBySize.Params.XMLName = xml.Name{Local: "lun-create-by-size"}
	}
	c.LunCreateBySize.Base.Name = c.vserver
	c.LunCreateBySize.Params.LunCreateBySizeOptions = *options
	r := LunCreateBySizeResponse{}
	res, err := c.LunCreateBySize.get(c.LunCreateBySize, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunCreateBySizeAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
