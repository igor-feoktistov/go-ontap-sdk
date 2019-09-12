package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type AggregateGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		AggregateGetOptions
	}
}

type AggregateGetOptions struct {
	DesiredAttributes *AggregateInfo `xml:"desired-attributes>aggr-attributes,omitempty"`
	MaxRecords        int            `xml:"max-records,omitempty"`
	Query             *AggregateInfo `xml:"query>aggr-attributes,omitempty"`
	Tag               string         `xml:"tag,omitempty"`
}

type AggregateInfo struct {
	AggregateName                string                        `xml:"aggregate-name,omitempty"`
	AggregateInodeAttributes     *AggregateInodeAttributes     `xml:"aggr-inode-attributes,omitempty"`
	AggregateSpaceAttributes     *AggregateSpaceAttributes     `xml:"aggr-space-attributes,omitempty"`
	AggregateOwnershipAttributes *AggregateOwnershipAttributes `xml:"aggr-ownership-attributes,omitempty"`
	AggregateRaidAttributes      *AggregateRaidAttributes      `xml:"aggr-raid-attributes,omitempty"`
}

type AggregateRaidAttributes struct {
	AggregateType      string `xml:"aggregate-type,omitempty"`
	CacheRaidGroupSize int    `xml:"cache-raid-group-size,omitempty"`
	ChecksumStatus     string `xml:"checksum-status,omitempty"`
	ChecksumStyle      string `xml:"checksum-style,omitempty"`
	DiskCount          int    `xml:"disk-count,omitempty"`
	EncryptionKeyID    string `xml:"encryption-key-id,omitempty"`
	HaPolicy           string `xml:"ha-policy,omitempty"`
	HasLocalRoot       *bool  `xml:"has-local-root"`
	HasPartnerRoot     *bool  `xml:"has-partner-root"`
	IsChecksumEnabled  *bool  `xml:"is-checksum-enabled"`
	IsEncrypted        *bool  `xml:"is-encrypted"`
	IsHybrid           *bool  `xml:"is-hybrid"`
	IsHybridEnabled    *bool  `xml:"is-hybrid-enabled"`
	IsInconsistent     *bool  `xml:"is-inconsistent"`
	IsMirrored         *bool  `xml:"is-mirrored"`
	IsRootAggregate    *bool  `xml:"is-root-aggregate"`
	MirrorStatus       string `xml:"mirror-status,omitempty"`
	MountState         string `xml:"mount-state,omitempty"`
	PlexCount          int    `xml:"plex-count,omitempty"`
	RaidLostWriteState string `xml:"raid-lost-write-state,omitempty"`
	RaidSize           int    `xml:"raid-size,omitempty"`
	RaidStatus         string `xml:"raid-status,omitempty"`
	RaidType           string `xml:"raid-type,omitempty"`
	State              string `xml:"state,omitempty"`
	UsesSharedDisks    *bool  `xml:"uses-shared-disks"`
}

type AggregateOwnershipAttributes struct {
	Cluster   string `xml:"cluster"`
	HomeID    int    `xml:"home-id"`
	HomeName  string `xml:"home-name"`
	OwnerID   int    `xml:"owner-id"`
	OwnerName string `xml:"owner-name"`
}

type AggregateInodeAttributes struct {
	FilesPrivateUsed         int `xml:"files-private-used"`
	FilesTotal               int `xml:"files-total"`
	FilesUsed                int `xml:"files-used"`
	InodefilePrivateCapacity int `xml:"inodefile-private-capacity"`
	InodefilePublicCapacity  int `xml:"inodefile-public-capacity"`
	MaxfilesAvailable        int `xml:"maxfiles-available"`
	MaxfilesPossible         int `xml:"maxfiles-possible"`
	MaxfilesUsed             int `xml:"maxfiles-used"`
	PercentInodeUsedCapacity int `xml:"percent-inode-used-capacity"`
}

type AggregateSpaceAttributes struct {
	AggregateMetadata            string `xml:"aggregate-metadata"`
	HybridCacheSizeTotal         string `xml:"hybrid-cache-size-total"`
	PercentUsedCapacity          string `xml:"percent-used-capacity"`
	PhysicalUsed                 int    `xml:"physical-used"`
	PhysicalUsedPercent          int    `xml:"physical-used-percent"`
	SizeAvailable                int    `xml:"size-available"`
	SizeTotal                    int    `xml:"size-total"`
	SizeUsed                     int    `xml:"size-used"`
	TotalReservedSpace           int    `xml:"total-reserved-space"`
	UsedIncludingSnapshotReserve string `xml:"used-including-snapshot-reserve"`
	VolumeFootprints             string `xml:"volume-footprints"`
}

type AggregateGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AggregateAttributes []AggregateInfo `xml:"attributes-list>aggr-attributes"`
		NextTag             string     `xml:"next-tag"`
	} `xml:"results"`
}

func (c *Client) AggregateGetAPI(options *AggregateGetOptions) (*AggregateGetResponse, *http.Response, error) {
	if c.AggregateGetIter == nil {
		c.AggregateGetIter = &AggregateGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.AggregateGetIter.Params.XMLName = xml.Name{Local: "aggr-get-iter"}
	}
	c.AggregateGetIter.Params.AggregateGetOptions = *options
	r := AggregateGetResponse{}
	res, err := c.AggregateGetIter.get(c.AggregateGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(AggregateGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) AggregateGetIterAPI(options *AggregateGetOptions) (responseAggrs []*AggregateGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.AggregateGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseAggrs = append(responseAggrs, r)
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
