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

type CT_DuotoneEffect struct {
	EG_ColorChoice []*EG_ColorChoice
}

func NewCT_DuotoneEffect() *CT_DuotoneEffect {
	ret := &CT_DuotoneEffect{}
	return ret
}

func (m *CT_DuotoneEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	for _, c := range m.EG_ColorChoice {
		c.MarshalXML(e, xml.StartElement{})
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DuotoneEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DuotoneEffect:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "scrgbClr"}:
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.ScrgbClr = NewCT_ScRgbColor()
				if err := d.DecodeElement(tmpcolorchoice.ScrgbClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "srgbClr"}:
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.SrgbClr = NewCT_SRgbColor()
				if err := d.DecodeElement(tmpcolorchoice.SrgbClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "hslClr"}:
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.HslClr = NewCT_HslColor()
				if err := d.DecodeElement(tmpcolorchoice.HslClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "sysClr"}:
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.SysClr = NewCT_SystemColor()
				if err := d.DecodeElement(tmpcolorchoice.SysClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "schemeClr"}:
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.SchemeClr = NewCT_SchemeColor()
				if err := d.DecodeElement(tmpcolorchoice.SchemeClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "prstClr"}:
				tmpcolorchoice := NewEG_ColorChoice()
				tmpcolorchoice.PrstClr = NewCT_PresetColor()
				if err := d.DecodeElement(tmpcolorchoice.PrstClr, &el); err != nil {
					return err
				}
				m.EG_ColorChoice = append(m.EG_ColorChoice, tmpcolorchoice)
			default:
				gooxml.Log("skipping unsupported element on CT_DuotoneEffect %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DuotoneEffect
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DuotoneEffect and its children
func (m *CT_DuotoneEffect) Validate() error {
	return m.ValidateWithPath("CT_DuotoneEffect")
}

// ValidateWithPath validates the CT_DuotoneEffect and its children, prefixing error messages with path
func (m *CT_DuotoneEffect) ValidateWithPath(path string) error {
	for i, v := range m.EG_ColorChoice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ColorChoice[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
