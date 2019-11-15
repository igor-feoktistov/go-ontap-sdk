package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VserverGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		VserverGetOptions
	}
}

type VserverGetOptions struct {
	DesiredAttributes *VserverInfo `xml:"desired-attributes>vserver-info,omitempty"`
	MaxRecords        int            `xml:"max-records,omitempty"`
	Query             *VserverInfo `xml:"query>vserver-info,omitempty"`
	Tag               string         `xml:"tag,omitempty"`
}

type VserverInfo struct {
	AggrList                      *[]string `xml:"aggr-list>aggr-name,omitempty"`
	AllowedProtocols              *[]string `xml:"allowed-protocols>protocol,omitempty"`
	AntivirusOnAccessPolicy       string    `xml:"antivirus-on-access-policy,omitempty"`
	CachingPolicy                 string    `xml:"caching-policy,omitempty"`
	Comment                       string    `xml:",commentomitempty"`
	DisallowedProtocols           *[]string `xml:"disallowed-protocols>protocol,omitempty"`
	Ipspace                       string    `xml:"ipspace,omitempty"`
	IsConfigLockedForChanges      bool      `xml:"is-config-locked-for-changes,omitempty"`
	IsRepositoryVserver           bool      `xml:"is-repository-vserver,omitempty"`
	IsSpaceEnforcementLogical     bool      `xml:"is-space-enforcement-logical,omitempty"`
	IsSpaceReportingLogical       bool      `xml:"is-space-reporting-logical,omitempty"`
	IsVserverProtected            bool      `xml:"is-vserver-protected,omitempty"`
	Language                      string    `xml:"language,omitempty"`
	LdapDomain                    string    `xml:"ldap-domain,omitempty"`
	MaxVolumes                    string    `xml:"max-volumes,omitempty"`
	NameMappingSwitch             *[]string `xml:"name-mapping-switch>nsswitch,omitempty"`
	NameServerSwitch              *[]string `xml:"name-server-switch>nsswitch,omitempty"`
	NisDomain                     string    `xml:"nis-domain,omitempty"`
	OperationalState              string    `xml:"operational-state,omitempty"`
	OperationalStateStoppedReason string    `xml:"operational-state-stopped-reason,omitempty"`
	QosPolicyGroup                string    `xml:"qos-policy-group,omitempty"`
	QuotaPolicy                   string    `xml:"quota-policy,omitempty"`
	RootVolume                    string    `xml:"root-volume,omitempty"`
	RootVolumeAggregate           string    `xml:"root-volume-aggregate,omitempty"`
	RootVolumeSecurityStyle       string    `xml:"root-volume-security-style,omitempty"`
	SnapshotPolicy                string    `xml:"snapshot-policy,omitempty"`
	State                         string    `xml:"state,omitempty"`
	Uuid                          string    `xml:"uuid,omitempty"`
	VolumeDeleteRetentionHours    string    `xml:"volume-delete-retention-hours,omitempty"`
	VserverAggrInfoList           string    `xml:"vserver-aggr-info-list,omitempty"`
	VserverName                   string    `xml:"vserver-name,omitempty"`
	VserverSubtype                string    `xml:"vserver-subtype,omitempty"`
	VserverType                   string    `xml:"vserver-type,omitempty"`
}

type VserverGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		VserverAttributes []VserverInfo `xml:"attributes-list>vserver-info"`
		NextTag           string        `xml:"next-tag"`
	} `xml:"results"`
}

func (c *Client) VserverGetAPI(options *VserverGetOptions) (*VserverGetResponse, *http.Response, error) {
	if c.VserverGetIter == nil {
		c.VserverGetIter = &VserverGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VserverGetIter.Params.XMLName = xml.Name{Local: "vserver-get-iter"}
	}
	c.VserverGetIter.Params.VserverGetOptions = *options
	r := VserverGetResponse{}
	res, err := c.VserverGetIter.get(c.VserverGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VserverGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) VserverGetIterAPI(options *VserverGetOptions) (responseVservers []*VserverGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.VserverGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseVservers = append(responseVservers, r)
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
