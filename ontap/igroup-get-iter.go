package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type IgroupGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*IgroupGetOptions
	}
}

type IgroupQuery struct {
	IgroupInfo *IgroupInfo `xml:"initiator-group-info,omitempty"`
}

type IgroupGetOptions struct {
	DesiredAttributes *IgroupQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int          `xml:"max-records,omitempty"`
	Query             *IgroupQuery `xml:"query,omitempty"`
	Tag               string       `xml:"tag,omitempty"`
}

type IgroupInfo struct {
	InitiatorGroupAluaEnabled     bool           `xml:"initiator-group-alua-enabled,omitempty"`
	InitiatorGroupDeleteOnUnmap   bool           `xml:"initiator-group-delete-on-unmap,omitempty"`
	InitiatorGroupName            string         `xml:"initiator-group-name,omitempty"`
	InitiatorGroupOsType          string         `xml:"initiator-group-os-type,omitempty"`
	InitiatorGroupPortsetName     string         `xml:"initiator-group-portset-name,omitempty"`
	InitiatorGroupThrottleBorrow  bool           `xml:"initiator-group-throttle-borrow,omitempty"`
	InitiatorGroupThrottleReserve int            `xml:"initiator-group-throttle-reserve,omitempty"`
	InitiatorGroupType            string         `xml:"initiator-group-type,omitempty"`
	InitiatorGroupUsePartner      bool           `xml:"initiator-group-use-partner,omitempty"`
	InitiatorGroupUuid            string         `xml:"initiator-group-uuid,omitempty"`
	InitiatorGroupVsaEnabled      bool           `xml:"initiator-group-vsa-enabled,omitempty"`
	Initiators []struct {
		InitiatorName string `xml:"initiator-name"`
        }                                            `xml:"initiators>initiator-info"`
	LunId                         int            `xml:"lun-id,omitempty"`
	Vserver                       string         `xml:"vserver,omitempty"`
}

type IgroupGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			IgroupAttributes []IgroupInfo `xml:"initiator-group-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) IgroupGetAPI(options *IgroupGetOptions) (*IgroupGetResponse, *http.Response, error) {
	if c.IgroupGetIter == nil {
		c.IgroupGetIter = &IgroupGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.IgroupGetIter.Params.XMLName = xml.Name{Local: "igroup-get-iter"}
	}
	c.IgroupGetIter.Base.Name = c.vserver
	c.IgroupGetIter.Params.IgroupGetOptions = options
	r := IgroupGetResponse{}
	res, err := c.IgroupGetIter.get(c.IgroupGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(IgroupGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) IgroupGetIterAPI(options *IgroupGetOptions) (responseIgroups []*IgroupGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.IgroupGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseIgroups = append(responseIgroups, r)
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
