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

type CT_TblCellMar struct {
	// Table Cell Top Margin Default
	Top *CT_TblWidth
	// Table Cell Leading Margin Default
	Start *CT_TblWidth
	// Table Cell Leading Margin Default
	Left *CT_TblWidth
	// Table Cell Bottom Margin Default
	Bottom *CT_TblWidth
	// Table Cell Trailing Margin Default
	End *CT_TblWidth
	// Table Cell Trailing Margin Default
	Right *CT_TblWidth
}

func NewCT_TblCellMar() *CT_TblCellMar {
	ret := &CT_TblCellMar{}
	return ret
}

func (m *CT_TblCellMar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Top != nil {
		setop := xml.StartElement{Name: xml.Name{Local: "w:top"}}
		e.EncodeElement(m.Top, setop)
	}
	if m.Start != nil {
		sestart := xml.StartElement{Name: xml.Name{Local: "w:start"}}
		e.EncodeElement(m.Start, sestart)
	}
	if m.Left != nil {
		seleft := xml.StartElement{Name: xml.Name{Local: "w:left"}}
		e.EncodeElement(m.Left, seleft)
	}
	if m.Bottom != nil {
		sebottom := xml.StartElement{Name: xml.Name{Local: "w:bottom"}}
		e.EncodeElement(m.Bottom, sebottom)
	}
	if m.End != nil {
		seend := xml.StartElement{Name: xml.Name{Local: "w:end"}}
		e.EncodeElement(m.End, seend)
	}
	if m.Right != nil {
		seright := xml.StartElement{Name: xml.Name{Local: "w:right"}}
		e.EncodeElement(m.Right, seright)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblCellMar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TblCellMar:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "top"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "top"}:
				m.Top = NewCT_TblWidth()
				if err := d.DecodeElement(m.Top, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "start"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "start"}:
				m.Start = NewCT_TblWidth()
				if err := d.DecodeElement(m.Start, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "left"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "left"}:
				m.Left = NewCT_TblWidth()
				if err := d.DecodeElement(m.Left, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bottom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bottom"}:
				m.Bottom = NewCT_TblWidth()
				if err := d.DecodeElement(m.Bottom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "end"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "end"}:
				m.End = NewCT_TblWidth()
				if err := d.DecodeElement(m.End, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "right"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "right"}:
				m.Right = NewCT_TblWidth()
				if err := d.DecodeElement(m.Right, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TblCellMar %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TblCellMar
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TblCellMar and its children
func (m *CT_TblCellMar) Validate() error {
	return m.ValidateWithPath("CT_TblCellMar")
}

// ValidateWithPath validates the CT_TblCellMar and its children, prefixing error messages with path
func (m *CT_TblCellMar) ValidateWithPath(path string) error {
	if m.Top != nil {
		if err := m.Top.ValidateWithPath(path + "/Top"); err != nil {
			return err
		}
	}
	if m.Start != nil {
		if err := m.Start.ValidateWithPath(path + "/Start"); err != nil {
			return err
		}
	}
	if m.Left != nil {
		if err := m.Left.ValidateWithPath(path + "/Left"); err != nil {
			return err
		}
	}
	if m.Bottom != nil {
		if err := m.Bottom.ValidateWithPath(path + "/Bottom"); err != nil {
			return err
		}
	}
	if m.End != nil {
		if err := m.End.ValidateWithPath(path + "/End"); err != nil {
			return err
		}
	}
	if m.Right != nil {
		if err := m.Right.ValidateWithPath(path + "/Right"); err != nil {
			return err
		}
	}
	return nil
}
