// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/928799934/gooxml"
)

type CT_VolType struct {
	// Type
	TypeAttr ST_VolDepType
	// Main
	Main []*CT_VolMain
}

func NewCT_VolType() *CT_VolType {
	ret := &CT_VolType{}
	ret.TypeAttr = ST_VolDepType(1)
	return ret
}

func (m *CT_VolType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	semain := xml.StartElement{Name: xml.Name{Local: "ma:main"}}
	for _, c := range m.Main {
		e.EncodeElement(c, semain)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_VolType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_VolDepType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_VolType:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "main"}:
				tmp := NewCT_VolMain()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Main = append(m.Main, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_VolType %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_VolType
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_VolType and its children
func (m *CT_VolType) Validate() error {
	return m.ValidateWithPath("CT_VolType")
}

// ValidateWithPath validates the CT_VolType and its children, prefixing error messages with path
func (m *CT_VolType) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_VolDepTypeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	for i, v := range m.Main {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Main[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
