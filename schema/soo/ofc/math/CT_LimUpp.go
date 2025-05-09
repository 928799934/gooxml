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

	"github.com/928799934/gooxml"
)

type CT_LimUpp struct {
	LimUppPr *CT_LimUppPr
	E        *CT_OMathArg
	Lim      *CT_OMathArg
}

func NewCT_LimUpp() *CT_LimUpp {
	ret := &CT_LimUpp{}
	ret.E = NewCT_OMathArg()
	ret.Lim = NewCT_OMathArg()
	return ret
}

func (m *CT_LimUpp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.LimUppPr != nil {
		selimUppPr := xml.StartElement{Name: xml.Name{Local: "m:limUppPr"}}
		e.EncodeElement(m.LimUppPr, selimUppPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	selim := xml.StartElement{Name: xml.Name{Local: "m:lim"}}
	e.EncodeElement(m.Lim, selim)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LimUpp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
	m.Lim = NewCT_OMathArg()
lCT_LimUpp:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "limUppPr"}:
				m.LimUppPr = NewCT_LimUppPr()
				if err := d.DecodeElement(m.LimUppPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "e"}:
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "lim"}:
				if err := d.DecodeElement(m.Lim, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_LimUpp %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LimUpp
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LimUpp and its children
func (m *CT_LimUpp) Validate() error {
	return m.ValidateWithPath("CT_LimUpp")
}

// ValidateWithPath validates the CT_LimUpp and its children, prefixing error messages with path
func (m *CT_LimUpp) ValidateWithPath(path string) error {
	if m.LimUppPr != nil {
		if err := m.LimUppPr.ValidateWithPath(path + "/LimUppPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	if err := m.Lim.ValidateWithPath(path + "/Lim"); err != nil {
		return err
	}
	return nil
}
