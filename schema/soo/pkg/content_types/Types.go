// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package content_types

import (
	"encoding/xml"

	"github.com/928799934/gooxml"
)

type Types struct {
	CT_Types
}

func NewTypes() *Types {
	ret := &Types{}
	ret.CT_Types = *NewCT_Types()
	return ret
}

func (m *Types) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/package/2006/content-types"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "Types"
	return m.CT_Types.MarshalXML(e, start)
}

func (m *Types) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Types = *NewCT_Types()
lTypes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/content-types", Local: "Default"}:
				tmp := NewDefault()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Default = append(m.Default, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/content-types", Local: "Override"}:
				tmp := NewOverride()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Override = append(m.Override, tmp)
			default:
				gooxml.Log("skipping unsupported element on Types %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lTypes
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Types and its children
func (m *Types) Validate() error {
	return m.ValidateWithPath("Types")
}

// ValidateWithPath validates the Types and its children, prefixing error messages with path
func (m *Types) ValidateWithPath(path string) error {
	if err := m.CT_Types.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
