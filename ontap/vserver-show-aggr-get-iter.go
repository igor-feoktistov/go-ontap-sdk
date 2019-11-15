package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VserverShowAggrGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		VserverShowAggrGetOptions
	}
}

type VserverShowAggrGetOptions struct {
	DesiredAttributes *ShowAggregates `xml:"desired-attributes>show-aggregates,omitempty"`
	MaxRecords        int             `xml:"max-records,omitempty"`
	Query             *ShowAggregates  `xml:"query>show-aggregates,omitempty"`
	Tag               string          `xml:"tag,omitempty"`
	Vserver           string          `xml:"vserver,omitempty"`
}

type ShowAggregates struct {
	AggregateName string `xml:"aggregate-name,omitempty"`
	AggregateType string `xml:"aggregate-type,omitempty"`
	AvailableSize int    `xml:"available-size,omitempty"`
	IsNveCapable  bool   `xml:"is-nve-capable,omitempty"`
	SnaplockType  string `xml:"snaplock-type,omitempty"`
	VserverName   string `xml:"vserver-name,omitempty"`
}

type VserverShowAggrGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AggrAttributes []ShowAggregates `xml:"attributes-list>show-aggregates"`
		NextTag        string           `xml:"next-tag"`
	} `xml:"results"`
}

func (c *Client) VserverShowAggrGetAPI(options *VserverShowAggrGetOptions) (*VserverShowAggrGetResponse, *http.Response, error) {
	if c.VserverShowAggrGetIter == nil {
		c.VserverShowAggrGetIter = &VserverShowAggrGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VserverShowAggrGetIter.Params.XMLName = xml.Name{Local: "vserver-show-aggr-get-iter"}
	}
	c.VserverShowAggrGetIter.Params.VserverShowAggrGetOptions = *options
	r := VserverShowAggrGetResponse{}
	res, err := c.VserverShowAggrGetIter.get(c.VserverShowAggrGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VserverShowAggrGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) VserverShowAggrGetIterAPI(options *VserverShowAggrGetOptions) (responseAggregates []*VserverShowAggrGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.VserverShowAggrGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseAggregates = append(responseAggregates, r)
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
