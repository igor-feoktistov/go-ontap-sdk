package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type SnapshotCreate struct {
	Base
	Params struct {
		XMLName    xml.Name
		SnapshotCreateOptions
	}
}

type SnapshotCreateOptions struct {
	Async           bool   `xml:"async,omitempty"`
	Comment         string `xml:"comment,omitempty"`
	ExpiryTime      int    `xml:"expiry-time,omitempty"`
	SnapmirrorLabel string `xml:"snapmirror-label,omitempty"`
	Snapshot        string `xml:"snapshot"`
	Volume          string `xml:"volume"`
}

func (c *Client) SnapshotCreateAPI(options *SnapshotCreateOptions) (*SingleResultResponse, *http.Response, error) {
	if c.SnapshotCreate == nil {
		c.SnapshotCreate = &SnapshotCreate {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.SnapshotCreate.Params.XMLName = xml.Name{Local: "snapshot-create"}
	}
	c.SnapshotCreate.Base.Name = c.vserver
	c.SnapshotCreate.Params.SnapshotCreateOptions = *options
	r := SingleResultResponse{}
	res, err := c.SnapshotCreate.get(c.SnapshotCreate, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(SnapshotCreateAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
