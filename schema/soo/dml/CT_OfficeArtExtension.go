// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/928799934/gooxml"
)

type CT_OfficeArtExtension struct {
	UriAttr string
	Any     []gooxml.Any
}

func NewCT_OfficeArtExtension() *CT_OfficeArtExtension {
	ret := &CT_OfficeArtExtension{}
	return ret
}

func (m *CT_OfficeArtExtension) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uri"},
		Value: fmt.Sprintf("%v", m.UriAttr)})
	e.EncodeToken(start)
	if m.Any != nil {
		for _, c := range m.Any {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OfficeArtExtension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uri" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UriAttr = parsed
			continue
		}
	}
lCT_OfficeArtExtension:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			default:
				if anyEl, err := gooxml.CreateElement(el); err != nil {
					return err
				} else {
					if err := d.DecodeElement(anyEl, &el); err != nil {
						return err
					}
					m.Any = append(m.Any, anyEl)
				}
			}
		case xml.EndElement:
			break lCT_OfficeArtExtension
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OfficeArtExtension and its children
func (m *CT_OfficeArtExtension) Validate() error {
	return m.ValidateWithPath("CT_OfficeArtExtension")
}

// ValidateWithPath validates the CT_OfficeArtExtension and its children, prefixing error messages with path
func (m *CT_OfficeArtExtension) ValidateWithPath(path string) error {
	return nil
}
