package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type SnapshotDelete struct {
	Base
	Params struct {
		XMLName    xml.Name
		SnapshotDeleteOptions
	}
}

type SnapshotDeleteOptions struct {
	IgnoreOwners         bool   `xml:"force,omitempty"`
	Snapshot             string `xml:"snapshot"`
	SnapshotInstanceUuid string `xml:"snapshot-instance-uuid,omitempty"`
	Volume               string `xml:"volume"`
}

func (c *Client) SnapshotDeleteAPI(options *SnapshotDeleteOptions) (*SingleResultResponse, *http.Response, error) {
	if c.SnapshotDelete == nil {
		c.SnapshotDelete = &SnapshotDelete {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.SnapshotDelete.Params.XMLName = xml.Name{Local: "snapshot-delete"}
	}
	c.SnapshotDelete.Base.Name = c.vserver
	c.SnapshotDelete.Params.SnapshotDeleteOptions = *options
	r := SingleResultResponse{}
	res, err := c.SnapshotDelete.get(c.SnapshotDelete, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(SnapshotDeleteAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
