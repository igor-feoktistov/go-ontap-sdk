package util

import (
    "io"
    "fmt"
    "encoding/hex"
    "path/filepath"

    "github.com/igor-feoktistov/go-ontap-sdk/ontap"
    "github.com/vmware/go-nfs-client/nfs"
    "github.com/vmware/go-nfs-client/nfs/rpc"
)


func UploadFileAPI(c *ontap.Client, volumeName string, filePath string, r io.Reader) (bytesUploaded int64, err error) {
	err = createDirPath(c, volumeName, filePath)
	if err != nil {
		return
	}
	options := &ontap.FileWriteFileOptions {
			    Path: fmt.Sprintf("/vol/%s%s", volumeName, filePath),
                	    Offset: -1,
	}
	inb := make([]byte, 512*1024)
	for {
		n, read_err := r.Read(inb)
    		if n > 0 {
			options.Data = hex.EncodeToString(inb[0:n])
			_, _, write_err := c.FileWriteFileAPI(options)
			if write_err != nil {
				err = write_err
				break
			} else {
				bytesUploaded += int64(n)
			}
    		}
    		if read_err == io.EOF {
        		break
    		}
    		if read_err != nil {
    			err = read_err
        		break
    		}
	}
	return
}

func UploadFileNFS(c *ontap.Client, volumeName string, filePath string, r io.Reader) (bytesUploaded int64, err error) {
	clientIP, err := GetOutboundIP()
	if err != nil {
		return
	}
	options := &ontap.VolumeGetOptions {
		    MaxRecords: 1,
		    Query: &ontap.VolumeQuery {
                	    VolumeInfo: &ontap.VolumeInfo {
				    VolumeIDAttributes: &ontap.VolumeIDAttributes {
                        		    Name: volumeName,
                        	    },
                    	    },
		    },
	}
	response, _, err := c.VolumeGetAPI(options)
	if err != nil {
	    return
	}
	if response.Results.NumRecords != 1 {
		err = fmt.Errorf("UploadFileNFS: volume %s not found", volumeName)
		return
	}
	exportPolicy := response.Results.AttributesList[0].VolumeExportAttributes.Policy
	junctionPath := response.Results.AttributesList[0].VolumeIDAttributes.JunctionPath
	err = createDirPath(c, volumeName, filePath)
	if err != nil {
		return
	}
	lifs, err := DiscoverNfsLIFs(c, volumeName)
	if err != nil {
		return
	}
	serverIP := lifs[0].Address
	err = createExportPolicyRule(c, exportPolicy, clientIP.String())
	if err != nil {
		return
	}
	defer deleteExportPolicyRule(c, exportPolicy, clientIP.String())
	mount, err := nfs.DialMount(serverIP)
	if err != nil {
		return
	}
	defer mount.Close()
	auth := rpc.NewAuthUnix("root", 0, 0)
	v, err := mount.Mount(junctionPath, auth.Auth())
	if err != nil {
		return
	}
	defer v.Close()
	w, err := v.OpenFile(filePath, 0644)
	if err != nil {
		return
	}
	defer w.Close()
	bytesUploaded, err = io.Copy(w, r)
	return
}

func createDirPath(c *ontap.Client, volumeName string, filePath string) (err error) {
	var dirList []string
	for dir := filepath.Dir(filePath); dir != "/"; dir = filepath.Dir(dir) {
		dirList = append(dirList, dir)
	}
	for i := len(dirList) - 1; i >= 0; i-- {
		dirPath := fmt.Sprintf("/vol/%s%s", volumeName, dirList[i])
		response, _, err := c.FileGetFileInfoAPI(dirPath)
		if err != nil && response.Results.ErrorNo == 2 {
			_, _, err = c.FileCreateDirectoryAPI(dirPath, "0755")
		}
		if err != nil {
			return err
		}
	}
	return
}

func createExportPolicyRule(c *ontap.Client, policyName string, clientIP string) (err error) {
	options := &ontap.ExportRuleCreateOptions {
			PolicyName:           policyName,
			AnonymousUserId:      "0",
            		SuperUserSecurity:    &[]string{"any"},
            		Protocol:             &[]string{"nfs"},
            		IsAllowDevIsEnabled:  true,
            		IsAllowSetUidEnabled: true,
            		ClientMatch:          clientIP,
            		RwRule:               &[]string{"any"},
            		RoRule:               &[]string{"any"},
	}
	_, _, err = c.ExportRuleCreateAPI(options)
	return
}

func deleteExportPolicyRule(c *ontap.Client, policyName string, clientIP string) (err error) {
	options := &ontap.ExportRuleGetOptions {
			MaxRecords: 1024,
            		Query: &ontap.ExportRuleQuery {
                		ExportRuleInfo: &ontap.ExportRuleInfo {
                			PolicyName: policyName,
					ClientMatch: clientIP,
                            },
                    },
	}
	response, _, err := c.ExportRuleGetAPI(options)
	if err == nil {
		if response.Results.NumRecords > 0 {
			for _, rule := range response.Results.AttributesList.ExportRuleAttributes {
				_, _, err = c.ExportRuleDestroyAPI(policyName, rule.RuleIndex)
				if err != nil {
					break
				}
			}
		}
	}
	return
}
