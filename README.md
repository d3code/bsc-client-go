# BSC Scan API Client

A Go client for the [BSC Scan API](https://coinmarketcap.com/api/documentation/v1/#section/Quick-Start-Guide)

### Install

```shell
go get github.com/d3code/bsc-client-go
```

### Import

```go
import (
    "github.com/d3code/bsc-client-go/bsc"
)
```

## Client

To use the client, create a BSC Client.

```go
package main

import (
    "github.com/d3code/bsc-client-go/bsc"
)

func main() {
    client := bsc.Client("<your-api-key>")
}
```

### Logging

There is no logging in this library, an `error` is returned if there are request / unmarshalling errors and can be handled accordingly.

To print the raw response, call `PrintResponse(true)` on the client.

```go
bsc.Client("").PrintResponse(true)
```