## go-ontap-sdk: Go library for NetApp cDOT

`go-ontap-sdk` is a Go library for interfacing with NetApp cDOT API.

## Documentation

This library implements a subset of NetApp Manageability SDK 9.6 APIs

See Reference Manual for ONTAP 9.6 APIs at https://library.netapp.com/ecm/ecm_download_file/ECMLP2855863

## Installation

In order to install `go-ontap-sdk` execute the following command:

```
go get -v github.com/ifeoktistov/go-ontap-sdk
```

## Tests

```
TBD
```

## Examples

Check the included examples from this repository.

Please note that most of the examples create connection to vserver management LIF.
if you connect to cluster management LIF, make sure to remove comment 
from `//c.SetVserver("<your vserver name>")` to run API call in vserver scope.
