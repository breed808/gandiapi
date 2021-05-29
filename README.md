# Gandigo

Go API Client to interact with Gandi.net services. This implements part of the [LiveDNS API](https://api.gandi.net/docs/livedns/) and part of the [Domain API](https://api.gandi.net/docs/domains/).

This implementation was forked from the older [implementation by sgmac](https://github.com/sgmac/api-gandi)

The current implementation requires Go 1.14 and uses modules.

## Install 

`import "github.com/breed808/gandiapi"`

```go

client, err := gandiapi.NewClient(nil)
if err != nil {
	log.Fatal("error")
}

resp, err := client.GetDomains()
if err != nil {
	log.Fatal(err)
}

```

If required, an `OptsClient` struct can be passed to override the default base url.

## Examples

Adding a record:

```go
data := gandiapi.Record{
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
 err = client.DeleteRecord("example.com", "amazing-cli", "A")
if err != nil {
	fmt.Println(err)
}
```

## License

Copyright [2021] [Ben Reedy]
Copyright [2020] [Sergio Galvan]

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

