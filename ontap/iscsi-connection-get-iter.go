package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type IscsiConnectionGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*IscsiConnectionGetOptions
	}
}

type IscsiConnectionQuery struct {
	IscsiConnectionListEntryInfo *IscsiConnectionListEntryInfo `xml:"iscsi-connection-list-entry-info,omitempty"`
}

type IscsiConnectionGetOptions struct {
	DesiredAttributes *IscsiConnectionQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int                   `xml:"max-records,omitempty"`
	Query             *IscsiConnectionQuery `xml:"query,omitempty"`
	Tag               string                `xml:"tag,omitempty"`
}

type IscsiConnectionListEntryInfo struct {
	ConnectionId    int    `xml:"connection-id,omitempty"`
	ConnectionState string `xml:"connection-state,omitempty"`
	HasSession      bool   `xml:"has-session,omitempty"`
	InterfaceName   string `xml:"interface-name,omitempty"`
	LocalIpAddress  string `xml:"local-ip-address,omitempty"`
	LocalIpPort     int    `xml:"local-ip-port,omitempty"`
	RemoteIpAddress string `xml:"remote-ip-address,omitempty"`
	RemoteIpPort    int    `xml:"remote-ip-port,omitempty"`
	SessionId       int    `xml:"session-id,omitempty"`
	TpgroupName     string `xml:"tpgroup-name,omitempty"`
	TpgroupTag      string `xml:"tpgroup-tag,omitempty"`
	Vserver         string `xml:"vserver,omitempty"`
}

type IscsiConnectionGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			IscsiConnectionAttributes []IscsiConnectionListEntryInfo `xml:"iscsi-connection-list-entry-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) IscsiConnectionGetAPI(options *IscsiConnectionGetOptions) (*IscsiConnectionGetResponse, *http.Response, error) {
	if c.IscsiConnectionGetIter == nil {
		c.IscsiConnectionGetIter = &IscsiConnectionGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.IscsiConnectionGetIter.Params.XMLName = xml.Name{Local: "iscsi-connection-get-iter"}
	}
	c.IscsiConnectionGetIter.Base.Name = c.vserver
	c.IscsiConnectionGetIter.Params.IscsiConnectionGetOptions = options
	r := IscsiConnectionGetResponse{}
	res, err := c.IscsiConnectionGetIter.get(c.IscsiConnectionGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(IscsiConnectionGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) IscsiConnectionGetIterAPI(options *IscsiConnectionGetOptions) (responseIscsiConnections []*IscsiConnectionGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.IscsiConnectionGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseIscsiConnections = append(responseIscsiConnections, r)
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
