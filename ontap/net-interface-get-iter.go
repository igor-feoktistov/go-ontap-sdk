package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type NetInterfaceGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*NetInterfaceGetOptions
	}
}

type NetInterfaceQuery struct {
	NetInterfaceInfo *NetInterfaceInfo `xml:"net-interface-info,omitempty"`
}

type NetInterfaceGetOptions struct {
	DesiredAttributes *NetInterfaceQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int                `xml:"max-records,omitempty"`
	Query             *NetInterfaceQuery `xml:"query,omitempty"`
	Tag               string             `xml:"tag,omitempty"`
}

type NetInterfaceInfo struct {
	Address                string    `xml:"address,omitempty"`
	AddressFamily          string    `xml:"address-family,omitempty"`
	AdministrativeStatus   string    `xml:"administrative-status,omitempty"`
	Comment                string    `xml:"comment,omitempty"`
	CurrentNode            string    `xml:"current-node,omitempty"`
	CurrentPort            string    `xml:"current-port,omitempty"`
	DataProtocols          *[]string `xml:"data-protocols>data-protocol,omitempty"`
	DnsDomainName          string    `xml:"dns-domain-name,omitempty"`
	ExtendedStatus         string    `xml:"extended-status,omitempty"`
	FailoverGroup          string    `xml:"failover-group,omitempty"`
	FailoverPolicy         string    `xml:"failover-policy,omitempty"`
	FirewallPolicy         string    `xml:"firewall-policy,omitempty"`
	ForceSubnetAssociation string    `xml:"force-subnet-association,omitempty"`
	HomeNode               string    `xml:"home-node,omitempty"`
	HomePort               string    `xml:"home-port,omitempty"`
	InterfaceName          string    `xml:"interface-name,omitempty"`
	Ipspace                string    `xml:"ipspace,omitempty"`
	IsAutoRevert           bool      `xml:"is-auto-revert,omitempty"`
	IsDnsUpdateEnabled     bool      `xml:"is-dns-update-enabled,omitempty"`
	IsHome                 bool      `xml:"is-home,omitempty"`
	IsIpv4LinkLocal        bool      `xml:"is-ipv4-link-local,omitempty"`
	IsVip                  bool      `xml:"is-vip,omitempty"`
	LifUuid                string    `xml:"lif-uuid,omitempty"`
	ListenForDnsQuery      bool      `xml:"listen-for-dns-query,omitempty"`
	Netmask                string    `xml:"netmask,omitempty"`
	NetmaskLength          int       `xml:"netmask-length,omitempty"`
	OperationalStatus      string    `xml:"operational-status,omitempty"`
	ProbePort              int       `xml:"probe-port,omitempty"`
	Role                   string    `xml:"role,omitempty"`
	RoutingGroupName       string    `xml:"routing-group-name,omitempty"`
	ServiceNames           *[]string `xml:"service-names>lif-service-name,omitempty"`
	ServicePolicy          string    `xml:"service-policy,omitempty"`
	SubnetName             string    `xml:"subnet-name,omitempty"`
	UseFailoverGroup       string    `xml:"use-failover-group,omitempty"`
	Vserver                string    `xml:"vserver,omitempty"`
	Wwpn                   string    `xml:"wwpn,omitempty"`
}

type NetInterfaceGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			NetInterfaceAttributes []NetInterfaceInfo `xml:"net-interface-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) NetInterfaceGetAPI(options *NetInterfaceGetOptions) (*NetInterfaceGetResponse, *http.Response, error) {
	if c.NetInterfaceGetIter == nil {
		c.NetInterfaceGetIter = &NetInterfaceGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.NetInterfaceGetIter.Params.XMLName = xml.Name{Local: "net-interface-get-iter"}
	}
	c.NetInterfaceGetIter.Base.Name = c.vserver
	c.NetInterfaceGetIter.Params.NetInterfaceGetOptions = options
	r := NetInterfaceGetResponse{}
	res, err := c.NetInterfaceGetIter.get(c.NetInterfaceGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(NetInterfaceGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) NetInterfaceGetIterAPI(options *NetInterfaceGetOptions) (responseNetInterfaces []*NetInterfaceGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.NetInterfaceGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseNetInterfaces = append(responseNetInterfaces, r)
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
