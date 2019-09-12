package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VolumeSetOption struct {
	Base
	Params struct {
		XMLName xml.Name
		OptionName  string `xml:"option-name"`
		OptionValue string `xml:"option-value"`
		Volume      string `xml:"volume"`
	}
}

func (c *Client) VolumeSetOptionAPI(volumeName string, optionName string, optionValue string) (*SingleResultResponse, *http.Response, error) {
	if c.VolumeSetOption == nil {
		c.VolumeSetOption = &VolumeSetOption {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VolumeSetOption.Params.XMLName = xml.Name{Local: "volume-set-option"}
	}
	c.VolumeSetOption.Base.Name = c.vserver
	c.VolumeSetOption.Params.Volume = volumeName
	c.VolumeSetOption.Params.OptionName = optionName
	c.VolumeSetOption.Params.OptionValue = optionValue
	r := SingleResultResponse{}
	res, err := c.VolumeSetOption.get(c.VolumeSetOption, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VolumeSetOptionAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
