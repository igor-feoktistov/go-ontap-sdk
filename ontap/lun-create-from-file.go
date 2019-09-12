package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunCreateFromFile struct {
	Base
	Params struct {
		XMLName    xml.Name
		LunCreateFromFileOptions
	}
}

type LunCreateFromFileOptions struct {
	Class                   string `xml:"class,omitempty"`
	Comment                 string `xml:"comment,omitempty"`
	FileName                string `xml:"file-name"`
	OsType                  string `xml:"ostype,omitempty"`
	Path                    string `xml:"path"`
	PrefixSize              int    `xml:"prefix-size,omitempty"`
	QosAdaptivePolicyGroup  string `xml:"qos-adaptive-policy-group,omitempty"`
	QosPolicyGroup          string `xml:"qos-policy-group,omitempty"`
	SpaceAllocationEnabled  bool   `xml:"space-allocation-enabled,omitempty"`
	SpaceReservationEnabled bool   `xml:"space-reservation-enabled,omitempty"`
}

type LunCreateFromFileResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		ActualSize int `xml:"actual-size"`
	} `xml:"results"`
}

func (c *Client) LunCreateFromFileAPI(options *LunCreateFromFileOptions) (*LunCreateFromFileResponse, *http.Response, error) {
	if c.LunCreateFromFile == nil {
		c.LunCreateFromFile = &LunCreateFromFile {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunCreateFromFile.Params.XMLName = xml.Name{Local: "lun-create-from-file"}
	}
	c.LunCreateFromFile.Base.Name = c.vserver
	c.LunCreateFromFile.Params.LunCreateFromFileOptions = *options
	r := LunCreateFromFileResponse{}
	res, err := c.LunCreateFromFile.get(c.LunCreateFromFile, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunCreateFromFileAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
