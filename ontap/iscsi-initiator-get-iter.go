package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type IscsiInitiatorGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*IscsiInitiatorGetOptions
	}
}

type IscsiInitiatorQuery struct {
	IscsiInitiatorListEntryInfo *IscsiInitiatorListEntryInfo `xml:"iscsi-initiator-list-entry-info,omitempty"`
}

type IscsiInitiatorGetOptions struct {
	DesiredAttributes *IscsiInitiatorQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int                  `xml:"max-records,omitempty"`
	Query             *IscsiInitiatorQuery `xml:"query,omitempty"`
	Tag               string               `xml:"tag,omitempty"`
}

type IscsiInitiatorListEntryInfo struct {
	InitiatorAliasname    string `xml:"initiator-aliasname,omitempty"`
	InitiatorGroupList []struct {
                InitiatorName string `xml:"initiator-group-name"`
        }                            `xml:"initiator-group-list>initiator-list-info,omitempty"`
	InitiatorNodename     string `xml:"initiator-nodename,omitempty"`
	Isid                  string `xml:"isid,omitempty"`
	TargetSessionId       int    `xml:"target-session-id,omitempty"`
	TpgroupName           string `xml:"tpgroup-name,omitempty"`
	TpgroupTag            int    `xml:"tpgroup-tag,omitempty"`
	Vserver               string `xml:"vserver,omitempty"`
}

type IscsiInitiatorGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			IscsiInitiatorAttributes []IscsiInitiatorListEntryInfo `xml:"iscsi-initiator-list-entry-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) IscsiInitiatorGetAPI(options *IscsiInitiatorGetOptions) (*IscsiInitiatorGetResponse, *http.Response, error) {
	if c.IscsiInitiatorGetIter == nil {
		c.IscsiInitiatorGetIter = &IscsiInitiatorGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.IscsiInitiatorGetIter.Params.XMLName = xml.Name{Local: "iscsi-initiator-get-iter"}
	}
	c.IscsiInitiatorGetIter.Base.Name = c.vserver
	c.IscsiInitiatorGetIter.Params.IscsiInitiatorGetOptions = options
	r := IscsiInitiatorGetResponse{}
	res, err := c.IscsiInitiatorGetIter.get(c.IscsiInitiatorGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(IscsiInitiatorGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) IscsiInitiatorGetIterAPI(options *IscsiInitiatorGetOptions) (responseIscsiInitiators []*IscsiInitiatorGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.IscsiInitiatorGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseIscsiInitiators = append(responseIscsiInitiators, r)
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
