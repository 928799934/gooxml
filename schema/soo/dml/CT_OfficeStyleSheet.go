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

type CT_OfficeStyleSheet struct {
	NameAttr          *string
	ThemeElements     *CT_BaseStyles
	ObjectDefaults    *CT_ObjectStyleDefaults
	ExtraClrSchemeLst *CT_ColorSchemeList
	CustClrLst        *CT_CustomColorList
	ExtLst            *CT_OfficeArtExtensionList
}

func NewCT_OfficeStyleSheet() *CT_OfficeStyleSheet {
	ret := &CT_OfficeStyleSheet{}
	ret.ThemeElements = NewCT_BaseStyles()
	return ret
}

func (m *CT_OfficeStyleSheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	e.EncodeToken(start)
	sethemeElements := xml.StartElement{Name: xml.Name{Local: "a:themeElements"}}
	e.EncodeElement(m.ThemeElements, sethemeElements)
	if m.ObjectDefaults != nil {
		seobjectDefaults := xml.StartElement{Name: xml.Name{Local: "a:objectDefaults"}}
		e.EncodeElement(m.ObjectDefaults, seobjectDefaults)
	}
	if m.ExtraClrSchemeLst != nil {
		seextraClrSchemeLst := xml.StartElement{Name: xml.Name{Local: "a:extraClrSchemeLst"}}
		e.EncodeElement(m.ExtraClrSchemeLst, seextraClrSchemeLst)
	}
	if m.CustClrLst != nil {
		secustClrLst := xml.StartElement{Name: xml.Name{Local: "a:custClrLst"}}
		e.EncodeElement(m.CustClrLst, secustClrLst)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OfficeStyleSheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ThemeElements = NewCT_BaseStyles()
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
	}
lCT_OfficeStyleSheet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "themeElements"}:
				if err := d.DecodeElement(m.ThemeElements, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "objectDefaults"}:
				m.ObjectDefaults = NewCT_ObjectStyleDefaults()
				if err := d.DecodeElement(m.ObjectDefaults, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extraClrSchemeLst"}:
				m.ExtraClrSchemeLst = NewCT_ColorSchemeList()
				if err := d.DecodeElement(m.ExtraClrSchemeLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "custClrLst"}:
				m.CustClrLst = NewCT_CustomColorList()
				if err := d.DecodeElement(m.CustClrLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_OfficeStyleSheet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OfficeStyleSheet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OfficeStyleSheet and its children
func (m *CT_OfficeStyleSheet) Validate() error {
	return m.ValidateWithPath("CT_OfficeStyleSheet")
}

// ValidateWithPath validates the CT_OfficeStyleSheet and its children, prefixing error messages with path
func (m *CT_OfficeStyleSheet) ValidateWithPath(path string) error {
	if err := m.ThemeElements.ValidateWithPath(path + "/ThemeElements"); err != nil {
		return err
	}
	if m.ObjectDefaults != nil {
		if err := m.ObjectDefaults.ValidateWithPath(path + "/ObjectDefaults"); err != nil {
			return err
		}
	}
	if m.ExtraClrSchemeLst != nil {
		if err := m.ExtraClrSchemeLst.ValidateWithPath(path + "/ExtraClrSchemeLst"); err != nil {
			return err
		}
	}
	if m.CustClrLst != nil {
		if err := m.CustClrLst.ValidateWithPath(path + "/CustClrLst"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
