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
	Alignment                 string `xml:"alignment"`
	BackingSnapshot           string `xml:"backing-snapshot"`
	BlockSize                 int    `xml:"block-size"`
	Class                     string `xml:"class"`
	CloneBackingSnapshot      string `xml:"clone-backing-snapshot"`
	Comment                   string `xml:"comment"`
	CreationTimestamp         int    `xml:"creation-timestamp"`
	DeviceBinaryId            string `xml:"device-binary-id"`
	DeviceId                  int    `xml:"device-id"`
	DeviceTextId              string `xml:"device-text-id"`
	IsClone                   bool   `xml:"is-clone"`
	IsCloneAutodeleteEnabled  bool   `xml:"is-clone-autodelete-enabled"`
	IsInconsistentImport      bool   `xml:"is-inconsistent-import"`
	IsRestoreInaccessible     bool   `xml:"is-restore-inaccessible"`
	IsSpaceAllocEnabled       bool   `xml:"is-space-alloc-enabled"`
	IsSpaceReservationEnabled bool   `xml:"is-space-reservation-enabled"`
	Mapped                    bool   `xml:"mapped"`
	MultiprotocolType         string `xml:"multiprotocol-type"`
	Node                      string `xml:"node"`
	Online                    bool   `xml:"online"`
	Path                      string `xml:"path"`
	PrefixSize                int    `xml:"prefix-size"`
	QosPolicyGroup            string `xml:"qos-policy-group"`
	Qtree                     string `xml:"qtree"`
	ReadOnly                  bool   `xml:"read-only"`
	Serial7Mode               string `xml:"serial-7-mode"`
	SerialNumber              string `xml:"serial-number"`
	ShareState                string `xml:"share-state"`
	Size                      int    `xml:"size"`
	SizeUsed                  int    `xml:"size-used"`
	Staging                   bool   `xml:"staging"`
	State                     string `xml:"state"`
	SuffixSize                int    `xml:"suffix-size"`
	Uuid                      string `xml:"uuid"`
	Volume                    string `xml:"volume"`
	Vserver                   string `xml:"vserver"`
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
