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

	"github.com/928799934/gooxml"
)

type EG_ShadeProperties struct {
	Lin  *CT_LinearShadeProperties
	Path *CT_PathShadeProperties
}

func NewEG_ShadeProperties() *EG_ShadeProperties {
	ret := &EG_ShadeProperties{}
	return ret
}

func (m *EG_ShadeProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Lin != nil {
		selin := xml.StartElement{Name: xml.Name{Local: "a:lin"}}
		e.EncodeElement(m.Lin, selin)
	}
	if m.Path != nil {
		sepath := xml.StartElement{Name: xml.Name{Local: "a:path"}}
		e.EncodeElement(m.Path, sepath)
	}
	return nil
}

func (m *EG_ShadeProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ShadeProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lin"}:
				m.Lin = NewCT_LinearShadeProperties()
				if err := d.DecodeElement(m.Lin, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "path"}:
				m.Path = NewCT_PathShadeProperties()
				if err := d.DecodeElement(m.Path, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_ShadeProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ShadeProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ShadeProperties and its children
func (m *EG_ShadeProperties) Validate() error {
	return m.ValidateWithPath("EG_ShadeProperties")
}

// ValidateWithPath validates the EG_ShadeProperties and its children, prefixing error messages with path
func (m *EG_ShadeProperties) ValidateWithPath(path string) error {
	if m.Lin != nil {
		if err := m.Lin.ValidateWithPath(path + "/Lin"); err != nil {
			return err
		}
	}
	if m.Path != nil {
		if err := m.Path.ValidateWithPath(path + "/Path"); err != nil {
			return err
		}
	}
	return nil
}
