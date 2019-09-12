package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VolumeContainer struct {
	Base
	Params struct {
		XMLName xml.Name
		Volume string `xml:"volume"`
	}
}

type VolumeContainerResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		ContainingAggregate string `xml:"containing-aggregate"`
	} `xml:"results"`
}

func (c *Client) VolumeContainerAPI(volumeName string) (*VolumeContainerResponse, *http.Response, error) {
	if c.VolumeContainer == nil {
		c.VolumeContainer = &VolumeContainer {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VolumeContainer.Params.XMLName = xml.Name{Local: "volume-container"}
	}
	c.VolumeContainer.Base.Name = c.vserver
	c.VolumeContainer.Params.Volume = volumeName
	r := VolumeContainerResponse{}
	res, err := c.VolumeContainer.get(c.VolumeContainer, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VolumeContainerAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
