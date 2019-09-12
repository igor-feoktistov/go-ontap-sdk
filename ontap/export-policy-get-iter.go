package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type ExportPolicyGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*ExportPolicyGetOptions
	}
}

type ExportPolicyGetOptions struct {
	DesiredAttributes *ExportPolicyQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int                `xml:"max-records,omitempty"`
	Query             *ExportPolicyQuery `xml:"query,omitempty"`
	Tag               string             `xml:"tag,omitempty"`
}

type ExportPolicyQuery struct {
	ExportPolicyInfo *ExportPolicyInfo `xml:"export-policy-info,omitempty"`
}

type ExportPolicyInfo struct {
	PolicyId   int    `xml:"policy-id,omitempty"`
	PolicyName string `xml:"policy-name,omitempty"`
	Vserver    string `xml:"vserver,omitempty"`
}

type ExportPolicyGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			ExportPolicyAttributes []ExportPolicyInfo `xml:"export-policy-info,omitempty"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) ExportPolicyGetAPI(options *ExportPolicyGetOptions) (*ExportPolicyGetResponse, *http.Response, error) {
	if c.ExportPolicyGetIter == nil {
		c.ExportPolicyGetIter = &ExportPolicyGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.ExportPolicyGetIter.Params.XMLName = xml.Name{Local: "export-policy-get-iter"}
	}
	c.ExportPolicyGetIter.Base.Name = c.vserver
	c.ExportPolicyGetIter.Params.ExportPolicyGetOptions = options
	r := ExportPolicyGetResponse{}
	res, err := c.ExportPolicyGetIter.get(c.ExportPolicyGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(ExportPolicyGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) ExportPolicyGetIterAPI(options *ExportPolicyGetOptions) (responses []*ExportPolicyGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.ExportPolicyGetAPI(options)
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
