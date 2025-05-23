// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"

	"github.com/928799934/gooxml"
)

type CT_RPrDefault struct {
	// Run Properties
	RPr *CT_RPr
}

func NewCT_RPrDefault() *CT_RPrDefault {
	ret := &CT_RPrDefault{}
	return ret
}

func (m *CT_RPrDefault) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RPrDefault) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_RPrDefault:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rPr"}:
				m.RPr = NewCT_RPr()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_RPrDefault %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RPrDefault
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RPrDefault and its children
func (m *CT_RPrDefault) Validate() error {
	return m.ValidateWithPath("CT_RPrDefault")
}

// ValidateWithPath validates the CT_RPrDefault and its children, prefixing error messages with path
func (m *CT_RPrDefault) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	return nil
}
