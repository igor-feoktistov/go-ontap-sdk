package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunMapListInfo struct {
	Base
	Params struct {
		XMLName    xml.Name
		Path string `xml:"path"`
	}
}

type LunMapListInfoResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		InitiatorGroups struct {
			IgroupAttributes []IgroupInfo `xml:"initiator-group-info"`
		} `xml:"initiator-groups"`
	} `xml:"results"`
}

func (c *Client) LunMapListInfoAPI(lunPath string) (*LunMapListInfoResponse, *http.Response, error) {
	if c.LunMapListInfo == nil {
		c.LunMapListInfo = &LunMapListInfo {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunMapListInfo.Params.XMLName = xml.Name{Local: "lun-map-list-info"}
	}
	c.LunMapListInfo.Base.Name = c.vserver
	c.LunMapListInfo.Params.Path = lunPath
	r := LunMapListInfoResponse{}
	res, err := c.LunMapListInfo.get(c.LunMapListInfo, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunMapListInfoAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
