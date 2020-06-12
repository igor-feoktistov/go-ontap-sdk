package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type FileListDirectoryIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*FileListDirectoryOptions
	}
}

type FileListDirectoryOptions struct {
	DesiredAttributes []*FileInfo `xml:"desired-attributes,omitempty>file-info,omitempty"`
	MaxRecords        int         `xml:"max-records,omitempty"`
	Path              string      `xml:"path,omitempty"`
	Query             []*FileInfo `xml:"query,omitempty>file-info,omitempty"`
	Tag               string      `xml:"tag,omitempty"`
}

type FileListDirectoryResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			FileAttributes []FileInfo `xml:"file-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) FileListDirectoryAPI(options *FileListDirectoryOptions) (*FileListDirectoryResponse, *http.Response, error) {
	if c.FileListDirectoryIter == nil {
		c.FileListDirectoryIter = &FileListDirectoryIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.FileListDirectoryIter.Params.XMLName = xml.Name{Local: "file-list-directory-iter"}
	}
	c.FileListDirectoryIter.Base.Name = c.vserver
	c.FileListDirectoryIter.Params.FileListDirectoryOptions = options
	r := FileListDirectoryResponse{}
	res, err := c.FileListDirectoryIter.get(c.FileListDirectoryIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(FileListDirectoryAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) FileListDirectoryIterAPI(options *FileListDirectoryOptions) (responseListDirectory []*FileListDirectoryResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.FileListDirectoryAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseListDirectory = append(responseListDirectory, r)
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
