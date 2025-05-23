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

type CT_GroupChrPr struct {
	Chr    *CT_Char
	Pos    *CT_TopBot
	VertJc *CT_TopBot
	CtrlPr *CT_CtrlPr
}

func NewCT_GroupChrPr() *CT_GroupChrPr {
	ret := &CT_GroupChrPr{}
	return ret
}

func (m *CT_GroupChrPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Chr != nil {
		sechr := xml.StartElement{Name: xml.Name{Local: "m:chr"}}
		e.EncodeElement(m.Chr, sechr)
	}
	if m.Pos != nil {
		sepos := xml.StartElement{Name: xml.Name{Local: "m:pos"}}
		e.EncodeElement(m.Pos, sepos)
	}
	if m.VertJc != nil {
		severtJc := xml.StartElement{Name: xml.Name{Local: "m:vertJc"}}
		e.EncodeElement(m.VertJc, severtJc)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GroupChrPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GroupChrPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "chr"}:
				m.Chr = NewCT_Char()
				if err := d.DecodeElement(m.Chr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "pos"}:
				m.Pos = NewCT_TopBot()
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "vertJc"}:
				m.VertJc = NewCT_TopBot()
				if err := d.DecodeElement(m.VertJc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "ctrlPr"}:
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_GroupChrPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupChrPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupChrPr and its children
func (m *CT_GroupChrPr) Validate() error {
	return m.ValidateWithPath("CT_GroupChrPr")
}

// ValidateWithPath validates the CT_GroupChrPr and its children, prefixing error messages with path
func (m *CT_GroupChrPr) ValidateWithPath(path string) error {
	if m.Chr != nil {
		if err := m.Chr.ValidateWithPath(path + "/Chr"); err != nil {
			return err
		}
	}
	if m.Pos != nil {
		if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
			return err
		}
	}
	if m.VertJc != nil {
		if err := m.VertJc.ValidateWithPath(path + "/VertJc"); err != nil {
			return err
		}
	}
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
