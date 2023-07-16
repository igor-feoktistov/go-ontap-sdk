package ontap

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type StorageDiskInfo struct {
	StorageDiskAttributes *StorageDiskAttributes `xml:"storage-disk-info,omitempty"`
}

type StorageDiskGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		StorageDiskGetOptions
	}
}

type StorageDiskGetOptions struct {
	DesiredAttributes *StorageDiskInfo `xml:"desired-attributes"`
	MaxRecords        int              `xml:"max-records,omitempty"`
	Query             *StorageDiskInfo `xml:"query"`
	Tag               string           `xml:"tag,omitempty"`
}

type StorageDiskGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		StorageDiskAttributes *StorageDiskGetIterResponseResultAttributesList `xml:"attributes-list"`
		NextTag               string                                          `xml:"next-tag"`
		NumRecords            int                                             `xml:"num-records"`
	} `xml:"results"`
}

type StorageDiskGetIterResponseResultAttributesList struct {
	XMLName         xml.Name                `xml:"attributes-list"`
	DiskDetailsInfo []StorageDiskAttributes `xml:"storage-disk-info"`
}

type StorageDiskAttributes struct {
	XMLName  xml.Name      `xml:"storage-disk-info"`
	DiskName string        `xml:"disk-name"`
	DiskRaid *DiskRaidInfo `xml:"disk-raid-info`
	DiskUid  string        `xml:"disk-uid"`
}

type DiskRaidInfo struct {
	XMLName           xml.Name `xml:"disk-raid-info"`
	ActiveNodeName    string   `xml:"active-node-name"`
	ContainerType     string   `xml:"container-type"`
	OutageIsInFDR     string   `xml:"outage-is-in-fdr"`
	OutageReason      string   `xml:"outage-reason"`
	EffectiveDiskType string   `xml:"effective-disk-type"`
	EffectiveRPM      string   `xml:"effective-rpm"`
	PhysicalBlocks    string   `xml:"physical-blocks"`
	Position          string   `xml:"position"`
	SparePool         string   `xml:"spare-pool"`
	StandardDiskType  string   `xml:"standard-disk-type"`
	UsedBlocks        string   `xml:"used-blocks"`
	ErrorTextList     string   `xml:"error-text-list"`
	ErrorType         string   `xml:"error-type"`
}

func (c *Client) StorageDiskGetAPI(options *StorageDiskGetOptions) (*StorageDiskGetResponse, *http.Response, error) {
	if c.StorageDiskGetIter == nil {
		c.StorageDiskGetIter = &StorageDiskGetIter{
			Base: Base{
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.StorageDiskGetIter.Params.XMLName = xml.Name{Local: "storage-disk-get-iter"}
	}
	c.StorageDiskGetIter.Params.StorageDiskGetOptions = *options
	r := StorageDiskGetResponse{}
	res, err := c.StorageDiskGetIter.get(c.StorageDiskGetIter, &r)
	if err == nil && r.Results.Passed() == false {
		err = fmt.Errorf("error(StorageDiskGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) StorageDiskGetIterAPI(options *StorageDiskGetOptions) (responseDisks []*StorageDiskGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.StorageDiskGetAPI(options)
		if err != nil {
			break
		} else {
			nextTag = r.Results.NextTag
			fmt.Printf("nextTag: %s", nextTag)
			fmt.Printf("%s", nextTag)
			responseDisks = append(responseDisks, r)
			if nextTag == "" {
				fmt.Print("nextTag is empty")
				break
			}
			options.Tag = nextTag
		}
	}

	return
}
