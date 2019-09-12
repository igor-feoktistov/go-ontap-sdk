package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunCopyGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*LunCopyGetOptions
	}
}

type LunCopyGetOptions struct {
	DesiredAttributes *LunCopyInfo `xml:"desired-attributes>lun-copy-info,omitempty"`
	MaxRecords        int          `xml:"max-records,omitempty"`
	Query             *LunCopyInfo `xml:"query>lun-copy-info,omitempty"`
	Tag               string       `xml:"tag,omitempty"`
}

type LunCopyInfo struct {
	CutoverTime                int    `xml:"cutover-time,omitempty"`
	DestinationPath            string `xml:"destination-path,omitempty"`
	DestinationVserver         string `xml:"destination-vserver,omitempty"`
	ElapsedTime                int    `xml:"elapsed-time,omitempty"`
	IsDestinationReady         bool   `xml:"is-destination-ready,omitempty"`
	IsPromotedEarly            bool   `xml:"is-promoted-early,omitempty"`
	IsSnapshotFenced           bool   `xml:"is-snapshot-fenced,omitempty"`
	JobStatus                  string `xml:"job-status,omitempty"`
	JobUuid                    string `xml:"job-uuid,omitempty"`
	LastFailureReason          string `xml:"last-failure-reason,omitempty"`
	LunIndex                   int    `xml:"lun-index,omitempty"`
	MaxThroughput              int    `xml:"max-throughput,omitempty"`
	ProgressPercent            int    `xml:"progress-percent,omitempty"`
	ScannerProgress            int    `xml:"scanner-progress,omitempty"`
	ScannerTotal               int    `xml:"scanner-total,omitempty"`
	SourcePath                 string `xml:"source-path,omitempty"`
	SourceSnapshotInstanceUuid string `xml:"source-snapshot-instance-uuid,omitempty"`
	SourceVserver              string `xml:"source-vserver,omitempty"`
}

type LunCopyGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			LunCopyAttributes []LunCopyInfo `xml:"lun-copy-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) LunCopyGetAPI(options *LunCopyGetOptions) (*LunCopyGetResponse, *http.Response, error) {
	if c.LunCopyGetIter == nil {
		c.LunCopyGetIter = &LunCopyGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunCopyGetIter.Params.XMLName = xml.Name{Local: "lun-copy-get-iter"}
	}
	c.LunCopyGetIter.Base.Name = c.vserver
	c.LunCopyGetIter.Params.LunCopyGetOptions = options
	r := LunCopyGetResponse{}
	res, err := c.LunCopyGetIter.get(c.LunCopyGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunCopyGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) LunCopyGetIterAPI(options *LunCopyGetOptions) (responses []*LunCopyGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.LunCopyGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responses = append(responses, r)
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
