package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VolumeSize struct {
	Base
	Params struct {
		XMLName xml.Name
		NewSize string `xml:"new-size,omitempty"`
		Volume  string `xml:"volume"`
	}
}

type VolumeSizeResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		IsFixedSizeFlexVolume bool   `xml:"is-fixed-size-flex-volume,omitempty"`
		IsReadonlyFlexVolume  bool   `xml:"is-readonly-flex-volume,omitempty"`
		IsReplicaFlexVolume   bool   `xml:"is-replica-flex-volume,omitempty"`
		VolumeSize            string `xml:"volume-size"`
	} `xml:"results"`
}

func (c *Client) VolumeSizeAPI(volumeName string, volumeSize string) (*VolumeSizeResponse, *http.Response, error) {
	if c.VolumeSize == nil {
		c.VolumeSize = &VolumeSize {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VolumeSize.Params.XMLName = xml.Name{Local: "volume-size"}
	}
	c.VolumeSize.Base.Name = c.vserver
	c.VolumeSize.Params.Volume = volumeName
	c.VolumeSize.Params.NewSize = volumeSize
	r := VolumeSizeResponse{}
	res, err := c.VolumeSize.get(c.VolumeSize, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VolumeSizeAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
