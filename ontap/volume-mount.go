package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VolumeMount struct {
	Base
	Params struct {
		XMLName xml.Name
		ActivateJunction     bool   `xml:"activate-junction,omitempty"`
		ExportPolicyOverride bool   `xml:"export-policy-override,omitempty"`
		JunctionPath         string `xml:"junction-path"`
		VolumeName           string `xml:"volume-name"`
	}
}

func (c *Client) VolumeMountAPI(volumeName string, junctionPath string, exportPolicyOverride bool) (*SingleResultResponse, *http.Response, error) {
	if c.VolumeMount == nil {
		c.VolumeMount = &VolumeMount {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VolumeMount.Params.XMLName = xml.Name{Local: "volume-mount"}
	}
	c.VolumeMount.Base.Name = c.vserver
	c.VolumeMount.Params.VolumeName = volumeName
	c.VolumeMount.Params.JunctionPath = junctionPath
	c.VolumeMount.Params.ExportPolicyOverride = exportPolicyOverride
	r := SingleResultResponse{}
	res, err := c.VolumeMount.get(c.VolumeMount, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VolumeMountAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
