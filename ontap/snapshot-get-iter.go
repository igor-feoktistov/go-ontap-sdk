package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type SnapshotGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*SnapshotGetOptions
	}
}

type SnapshotGetOptions struct {
	DesiredAttributes *SnapshotQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int            `xml:"max-records,omitempty"`
	Query             *SnapshotQuery `xml:"query,omitempty"`
	Tag               string         `xml:"tag,omitempty"`
}

type SnapshotQuery struct {
	SnapshotInfo *SnapshotInfo `xml:"snapshot-info,omitempty"`
}

type SnapshotOwner struct {
	Owner string `xml:"owner"`
}

type SnapshotInfo struct {
	AccessTime                        int        `xml:"access-time,omitempty"`
	AfsUsed                           int        `xml:"afs-used,omitempty"`
	Busy                              bool       `xml:"busy,omitempty"`
	Comment                           string     `xml:"comment,omitempty"`
	CompressSavings                   int        `xml:"compress-savings,omitempty"`
	CompressionType                   string     `xml:"compression-type,omitempty"`
	ContainsLunClones                 bool       `xml:"contains-lun-clones,omitempty"`
	CumulativePercentageOfTotalBlocks int        `xml:"cumulative-percentage-of-total-blocks,omitempty"`
	CumulativePercentageOfUsedBlocks  int        `xml:"cumulative-percentage-of-used-blocks,omitempty"`
	CumulativeTotal                   int        `xml:"cumulative-total,omitempty"`
	DedupSavings                      int        `xml:"dedup-savings,omitempty"`
	Dependency                        string     `xml:"dependency,omitempty"`
	ExpiryTime                        int        `xml:"expiry-time,omitempty"`
	InfiniteSnaplockExpiryTime        bool       `xml:"infinite-snaplock-expiry-time,omitempty"`
	InofileVersion                    int        `xml:"inofile-version,omitempty"`
	Is7ModeSnapshot                   bool       `xml:"is-7-mode-snapshot,omitempty"`
	IsConstituentSnapshot             bool       `xml:"is-constituent-snapshot,omitempty"`
	Name                              string     `xml:"name,omitempty"`
	PercentageOfTotalBlocks           int        `xml:"percentage-of-total-blocks,omitempty"`
	PercentageOfUsedBlocks            int        `xml:"percentage-of-used-blocks,omitempty"`
	PerformanceMetadata               int        `xml:"performance-metadata,omitempty"`
	SnaplockExpiryTime                int        `xml:"snaplock-expiry-time,omitempty"`
	SnapmirrorLabel                   string     `xml:"snapmirror-label,omitempty"`
	SnapshotInstanceUuid              string     `xml:"snapshot-instance-uuid,omitempty"`
	SnapshotOwnersList                []struct {
    		Owner                     string     `xml:"owner"`
        }                                            `xml:"snapshot-owners-list,omitempty>snapshot-owner,omitempty"`
	SnapshotVersionUuid               string     `xml:"snapshot-version-uuid,omitempty"`
	State                             string     `xml:"state,omitempty"`
	Total                             int        `xml:"total,omitempty"`
	Vbn0Savings                       int        `xml:"vbn0-savings,omitempty"`
	VolumeProvenanceUuid              string     `xml:"volume-provenance-uuid,omitempty"`
	Vserver                           string     `xml:"vserver,omitempty"`
}

type SnapshotGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList []SnapshotInfo `xml:"attributes-list>snapshot-info"`
		NextTag        string       `xml:"next-tag"`
		NumRecords     int          `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) SnapshotGetAPI(options *SnapshotGetOptions) (*SnapshotGetResponse, *http.Response, error) {
	if c.SnapshotGetIter == nil {
		c.SnapshotGetIter = &SnapshotGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.SnapshotGetIter.Params.XMLName = xml.Name{Local: "snapshot-get-iter"}
	}
	c.SnapshotGetIter.Base.Name = c.vserver
	c.SnapshotGetIter.Params.SnapshotGetOptions = options
	r := SnapshotGetResponse{}
	res, err := c.SnapshotGetIter.get(c.SnapshotGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(SnapshotGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) SnapshotGetIterAPI(options *SnapshotGetOptions) (responseSnapshots []*SnapshotGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.SnapshotGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseSnapshots = append(responseSnapshots, r)
			if nextTag == "" {
				break
			}
			options.Tag = nextTag
		} else {
			break
		}
	}
	return
}
