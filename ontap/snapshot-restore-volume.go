package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type SnapshotRestoreVolume struct {
	Base
	Params struct {
		XMLName    xml.Name
		SnapshotRestoreVolumeOptions
	}
}

type SnapshotRestoreVolumeOptions struct {
	Force                bool   `xml:"force,omitempty"`
	PreserveLunIds       bool   `xml:"preserve-lun-ids,omitempty"`
	Snapshot             string `xml:"snapshot"`
	SnapshotInstanceUuid string `xml:"snapshot-instance-uuid,omitempty"`
	Volume               string `xml:"volume"`
}

func (c *Client) SnapshotRestoreVolumeAPI(options *SnapshotRestoreVolumeOptions) (*SingleResultResponse, *http.Response, error) {
	if c.SnapshotRestoreVolume == nil {
		c.SnapshotRestoreVolume = &SnapshotRestoreVolume {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.SnapshotRestoreVolume.Params.XMLName = xml.Name{Local: "snapshot-restore-volume"}
	}
	c.SnapshotRestoreVolume.Base.Name = c.vserver
	c.SnapshotRestoreVolume.Params.SnapshotRestoreVolumeOptions = *options
	r := SingleResultResponse{}
	res, err := c.SnapshotRestoreVolume.get(c.SnapshotRestoreVolume, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(SnapshotRestoreVolumeAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
