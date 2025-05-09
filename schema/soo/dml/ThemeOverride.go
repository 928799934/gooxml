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

type ThemeOverride struct {
	CT_BaseStylesOverride
}

func NewThemeOverride() *ThemeOverride {
	ret := &ThemeOverride{}
	ret.CT_BaseStylesOverride = *NewCT_BaseStylesOverride()
	return ret
}

func (m *ThemeOverride) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "a:themeOverride"
	return m.CT_BaseStylesOverride.MarshalXML(e, start)
}

func (m *ThemeOverride) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_BaseStylesOverride = *NewCT_BaseStylesOverride()
lThemeOverride:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "clrScheme"}:
				m.ClrScheme = NewCT_ColorScheme()
				if err := d.DecodeElement(m.ClrScheme, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fontScheme"}:
				m.FontScheme = NewCT_FontScheme()
				if err := d.DecodeElement(m.FontScheme, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fmtScheme"}:
				m.FmtScheme = NewCT_StyleMatrix()
				if err := d.DecodeElement(m.FmtScheme, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on ThemeOverride %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lThemeOverride
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ThemeOverride and its children
func (m *ThemeOverride) Validate() error {
	return m.ValidateWithPath("ThemeOverride")
}

// ValidateWithPath validates the ThemeOverride and its children, prefixing error messages with path
func (m *ThemeOverride) ValidateWithPath(path string) error {
	if err := m.CT_BaseStylesOverride.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
