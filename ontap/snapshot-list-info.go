package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type SnapshotListInfo struct {
	Base
	Params struct {
		XMLName xml.Name
		*SnapshotListInfoOptions
	}
}

type SnapshotListInfoOptions struct {
	Is7ModeSnapshot  bool   `xml:"is-7-mode-snapshot,omitempty"`
	LunCloneSnapshot bool   `xml:"lun-clone-snapshot,omitempty"`
	Snapowners       bool   `xml:"snapowners,omitempty"`
	TargetName       string `xml:"target-name,omitempty"`
	TargetType       string `xml:"target-type,omitempty"`
	Terse            bool   `xml:"terse,omitempty"`
	Volume           string `xml:"volume,omitempty"`
}

type SnapshotListInfoResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		Snapshots []SnapshotInfo `xml:"snapshots>snapshot-info"`
	} `xml:"results"`
}

func (c *Client) SnapshotListInfoAPI(options *SnapshotListInfoOptions) (*SnapshotListInfoResponse, *http.Response, error) {
	if c.SnapshotListInfo == nil {
		c.SnapshotListInfo = &SnapshotListInfo {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.SnapshotListInfo.Params.XMLName = xml.Name{Local: "snapshot-list-info"}
	}
	c.SnapshotListInfo.Base.Name = c.vserver
	c.SnapshotListInfo.Params.SnapshotListInfoOptions = options
	r := SnapshotListInfoResponse{}
	res, err := c.SnapshotListInfo.get(c.SnapshotListInfo, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(SnapshotListInfoAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
