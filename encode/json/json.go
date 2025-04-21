// Package json encodes a parser.Interface into json.
// This can be used for transcoding or marshaling.
package json

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/katydid/parser-go/parser"
)

// Encode encodes a parser.Interface into a byte slice containing valid json.
func Encode(p parser.Interface) ([]byte, error) {
	m, err := encode(p)
	if err != nil {
		return nil, err
	}
	buf, err := json.MarshalIndent(m, "", "\t")
	return buf, err
}

func encode(p parser.Interface) (interface{}, error) {
	var fields = make(map[string]interface{})
	var list []interface{}
	for {
		if err := p.Next(); err != nil {
			if err == io.EOF {
				if len(fields) > 0 {
					return fields, nil
				}
				if list != nil {
					return list, nil
				}
				return nil, nil
			}
			return nil, err
		}
		if p.IsLeaf() {
			if v, err := p.Int(); err == nil {
				return v, nil
			}
			if v, err := p.Uint(); err == nil {
				return v, nil
			}
			if v, err := p.Double(); err == nil {
				return v, nil
			}
			if v, err := p.Bool(); err == nil {
				return v, nil
			}
			if v, err := p.String(); err == nil {
				return v, nil
			}
			if v, err := p.Bytes(); err == nil {
				if bytes.Equal(v, []byte("null")) {
					return nil, nil
				}
			}
			return nil, nil
		}
		if _, err := p.Int(); err == nil {
			p.Down()
			item, err := encode(p)
			if err != nil {
				return nil, err
			}
			p.Up()
			list = append(list, item)
			continue
		}
		name, err := p.String()
		if err != nil {
			return nil, err
		}
		p.Down()
		value, err := encode(p)
		if err != nil {
			return nil, err
		}
		p.Up()
		fields[name] = value
	}
}
