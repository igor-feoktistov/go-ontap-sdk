package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VolumeAutosizeSet struct {
	Base
	Params struct {
		XMLName xml.Name
		VolumeAutosizeOptions
	}
}

type VolumeAutosizeOptions struct {
	GrowThresholdPercent   int    `xml:"grow-threshold-percent,omitempty"`
	IncrementSize          string `xml:"increment-size,omitempty"`
	IsEnabled              bool   `xml:"is-enabled,omitempty"`
	MaximumSize            string `xml:"maximum-size,omitempty"`
	MinimumSize            string `xml:"minimum-size,omitempty"`
	Mode                   string `xml:"mode,omitempty"`
	Reset                  bool   `xml:"reset,omitempty"`
	ShrinkThresholdPercent int    `xml:"shrink-threshold-percent,omitempty"`
	Volume                 string `xml:"volume"`
}

func (c *Client) VolumeAutosizeSetAPI(options *VolumeAutosizeOptions) (*SingleResultResponse, *http.Response, error) {
	if c.VolumeAutosizeSet == nil {
		c.VolumeAutosizeSet = &VolumeAutosizeSet {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VolumeAutosizeSet.Params.XMLName = xml.Name{Local: "volume-autosize-set"}
	}
	c.VolumeAutosizeSet.Base.Name = c.vserver
	c.VolumeAutosizeSet.Params.VolumeAutosizeOptions = *options
	r := SingleResultResponse{}
	res, err := c.VolumeAutosizeSet.get(c.VolumeAutosizeSet, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VolumeAutosizeSetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
