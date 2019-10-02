package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*LunGetOptions
	}
}

type LunQuery struct {
	LunInfo *LunInfo `xml:"lun-info,omitempty"`
}

type LunGetOptions struct {
	DesiredAttributes *LunQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int       `xml:"max-records,omitempty"`
	Query             *LunQuery `xml:"query,omitempty"`
	Tag               string    `xml:"tag,omitempty"`
}

type LunInfo struct {
	Alignment                 string `xml:"alignment,omitempty"`
	BackingSnapshot           string `xml:"backing-snapshot,omitempty"`
	BlockSize                 int    `xml:"block-size,omitempty"`
	Class                     string `xml:"class,omitempty"`
	CloneBackingSnapshot      string `xml:"clone-backing-snapshot,omitempty"`
	Comment                   string `xml:"comment,omitempty"`
	CreationTimestamp         int    `xml:"creation-timestamp,omitempty"`
	DeviceBinaryId            string `xml:"device-binary-id,omitempty"`
	DeviceId                  int    `xml:"device-id,omitempty"`
	DeviceTextId              string `xml:"device-text-id,omitempty"`
	IsClone                   bool   `xml:"is-clone,omitempty"`
	IsCloneAutodeleteEnabled  bool   `xml:"is-clone-autodelete-enabled,omitempty"`
	IsInconsistentImport      bool   `xml:"is-inconsistent-import,omitempty"`
	IsRestoreInaccessible     bool   `xml:"is-restore-inaccessible,omitempty"`
	IsSpaceAllocEnabled       bool   `xml:"is-space-alloc-enabled,omitempty"`
	IsSpaceReservationEnabled bool   `xml:"is-space-reservation-enabled,omitempty"`
	Mapped                    bool   `xml:"mapped,omitempty"`
	MultiprotocolType         string `xml:"multiprotocol-type,omitempty"`
	Node                      string `xml:"node,omitempty"`
	Online                    bool   `xml:"online,omitempty"`
	Path                      string `xml:"path,omitempty"`
	PrefixSize                int    `xml:"prefix-size,omitempty"`
	QosPolicyGroup            string `xml:"qos-policy-group,omitempty"`
	Qtree                     string `xml:"qtree,omitempty"`
	ReadOnly                  bool   `xml:"read-only,omitempty"`
	Serial7Mode               string `xml:"serial-7-mode,omitempty"`
	SerialNumber              string `xml:"serial-number,omitempty"`
	ShareState                string `xml:"share-state,omitempty"`
	Size                      int    `xml:"size,omitempty"`
	SizeUsed                  int    `xml:"size-used,omitempty"`
	Staging                   bool   `xml:"staging,omitempty"`
	State                     string `xml:"state,omitempty"`
	SuffixSize                int    `xml:"suffix-size,omitempty"`
	Uuid                      string `xml:"uuid,omitempty"`
	Volume                    string `xml:"volume,omitempty"`
	Vserver                   string `xml:"vserver,omitempty"`
}

type LunGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			LunAttributes []LunInfo `xml:"lun-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) LunGetAPI(options *LunGetOptions) (*LunGetResponse, *http.Response, error) {
	if c.LunGetIter == nil {
		c.LunGetIter = &LunGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunGetIter.Params.XMLName = xml.Name{Local: "lun-get-iter"}
	}
	c.LunGetIter.Base.Name = c.vserver
	c.LunGetIter.Params.LunGetOptions = options
	r := LunGetResponse{}
	res, err := c.LunGetIter.get(c.LunGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) LunGetIterAPI(options *LunGetOptions) (responseLuns []*LunGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.LunGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseLuns = append(responseLuns, r)
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
