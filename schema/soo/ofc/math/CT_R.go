// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"fmt"

	"github.com/928799934/gooxml"
)

type CT_R struct {
	RPr    *CT_RPR
	Choice []*CT_RChoice
}

func NewCT_R() *CT_R {
	ret := &CT_R{}
	return ret
}

func (m *CT_R) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "m:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_R) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_R:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "rPr"}:
				m.RPr = NewCT_RPR()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "t"}:
				tmp := NewCT_RChoice()
				if err := d.DecodeElement(&tmp.T, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_R %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_R
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_R and its children
func (m *CT_R) Validate() error {
	return m.ValidateWithPath("CT_R")
}

// ValidateWithPath validates the CT_R and its children, prefixing error messages with path
func (m *CT_R) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
