## encode-go

Experimental encode-go package for Go.

This package takes a [Parser interface](https://github.com/katydid/parser-go) for any parser and encodes it into JSON, XML or a Reflect Go structure.
For example, encoding a JSON parser into a Reflect Go structure is the same as unmarshaling JSON.

How to transcode a parser to JSON:

```go
import "github.com/katydid/parser-go/parser"
import "github.com/katydid/encode-go/encode/json"
import "github.com/katydid/encode-go/encode/xml"

func encodeToJson(p parser.Interface) ([]byte, error) {
    return json.Encode(p)
}

func encodeToXml(p parser.Interface) ([]byte, error) {
    return xml.Encode(p)
}
```

Encoding a reflect Go structure requires also passing the structure into which will be encoded, or rather unmarshaled.

```go
import "github.com/katydid/parser-go/parser"
import reflectencoder "github.com/katydid/encode-go/encode/reflect"
import "reflect"

func unmarshal(p parser.Interface, a *A) error {
    v := reflect.ValueOf(a)
    return reflectencoder.Encode(p, v)
}
```
