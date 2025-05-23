// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"

	"github.com/928799934/gooxml"
	"github.com/928799934/gooxml/schema/soo/dml"
)

type CT_UpDownBar struct {
	SpPr *dml.CT_ShapeProperties
}

func NewCT_UpDownBar() *CT_UpDownBar {
	ret := &CT_UpDownBar{}
	return ret
}

func (m *CT_UpDownBar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "c:spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_UpDownBar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_UpDownBar:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"}:
				m.SpPr = dml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_UpDownBar %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_UpDownBar
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_UpDownBar and its children
func (m *CT_UpDownBar) Validate() error {
	return m.ValidateWithPath("CT_UpDownBar")
}

// ValidateWithPath validates the CT_UpDownBar and its children, prefixing error messages with path
func (m *CT_UpDownBar) ValidateWithPath(path string) error {
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	return nil
}
