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
	"strconv"

	"github.com/928799934/gooxml"
)

type CT_XmlColumnPr struct {
	// XML Map Id
	MapIdAttr uint32
	// XPath
	XpathAttr string
	// Denormalized
	DenormalizedAttr *bool
	// XML Data Type
	XmlDataTypeAttr string
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_XmlColumnPr() *CT_XmlColumnPr {
	ret := &CT_XmlColumnPr{}
	return ret
}

func (m *CT_XmlColumnPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "mapId"},
		Value: fmt.Sprintf("%v", m.MapIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xpath"},
		Value: fmt.Sprintf("%v", m.XpathAttr)})
	if m.DenormalizedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "denormalized"},
			Value: fmt.Sprintf("%d", b2i(*m.DenormalizedAttr))})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlDataType"},
		Value: fmt.Sprintf("%v", m.XmlDataTypeAttr)})
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_XmlColumnPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "mapId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.MapIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "xpath" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.XpathAttr = parsed
			continue
		}
		if attr.Name.Local == "denormalized" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DenormalizedAttr = &parsed
			continue
		}
		if attr.Name.Local == "xmlDataType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.XmlDataTypeAttr = parsed
			continue
		}
	}
lCT_XmlColumnPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_XmlColumnPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_XmlColumnPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_XmlColumnPr and its children
func (m *CT_XmlColumnPr) Validate() error {
	return m.ValidateWithPath("CT_XmlColumnPr")
}

// ValidateWithPath validates the CT_XmlColumnPr and its children, prefixing error messages with path
func (m *CT_XmlColumnPr) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
