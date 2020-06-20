# Gandigo

Go API Client to interact with Gandi.net services. This implements the [RESTful HTTP API](https://doc.livedns.gandi.net/). Unfortunately this API will not have more future developments, so I stopped working on the Go client. However, this is still a good example on how to implement a Go API client.

The current implementation requires Go 1.14 and uses modules.

## Install 

Import `github.com/sgmac/gandigo` 


```go
client, err := gandigo.NewClient(nil)
if err != nil {
	log.Fatal("error")
}

resp, err := client.GetZones()
if err != nil {
	log.Fatal(err)
}

```

If required, an `OptsClient` struct can be passed to override the default base url.

## Examples

Adding a record:

```go
data := gandigo.Record{
 	RrsetType:   "A",
 	RrsetTTL:    300,
 	RrsetName:   "amazing-cli",
 	RrsetValues: []string{"18.185.88.103"},
 }
 err := client.CreateRecord(data, zoneID)
 if err != nil {
 	fmt.Println(err)
 }
```

Delete a record.

```go
 err = client.DeleteRecordType(zoneID, "amazing-cli", "A")
if err != nil {
	fmt.Println(err)
}
```

Operations with snapshots.

```go
snapshots, err := client.GetSnapshots(zoneID)
if err != nil {
	log.Fatal(err)
}

details, err := client.GetSnapshotDetails(zoneID, snapshotID)
if err != nil {
	log.Fatal(err)
}
fmt.Println(details)
```

## License

Copyright [2020] [Sergio Galvan]

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

