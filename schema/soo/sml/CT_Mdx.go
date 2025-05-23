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

type CT_Mdx struct {
	// Connection Name Index
	NAttr uint32
	// Cube Function Tag
	FAttr ST_MdxFunctionType
	// Tuple MDX Metadata
	T *CT_MdxTuple
	// Set MDX Metadata
	Ms *CT_MdxSet
	// Member Property MDX Metadata
	P *CT_MdxMemeberProp
	// KPI MDX Metadata
	K *CT_MdxKPI
}

func NewCT_Mdx() *CT_Mdx {
	ret := &CT_Mdx{}
	ret.FAttr = ST_MdxFunctionType(1)
	return ret
}

func (m *CT_Mdx) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "n"},
		Value: fmt.Sprintf("%v", m.NAttr)})
	attr, err := m.FAttr.MarshalXMLAttr(xml.Name{Local: "f"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	if m.T != nil {
		set := xml.StartElement{Name: xml.Name{Local: "ma:t"}}
		e.EncodeElement(m.T, set)
	}
	if m.Ms != nil {
		sems := xml.StartElement{Name: xml.Name{Local: "ma:ms"}}
		e.EncodeElement(m.Ms, sems)
	}
	if m.P != nil {
		sep := xml.StartElement{Name: xml.Name{Local: "ma:p"}}
		e.EncodeElement(m.P, sep)
	}
	if m.K != nil {
		sek := xml.StartElement{Name: xml.Name{Local: "ma:k"}}
		e.EncodeElement(m.K, sek)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Mdx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.FAttr = ST_MdxFunctionType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "n" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.NAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "f" {
			m.FAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_Mdx:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "t"}:
				m.T = NewCT_MdxTuple()
				if err := d.DecodeElement(m.T, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "ms"}:
				m.Ms = NewCT_MdxSet()
				if err := d.DecodeElement(m.Ms, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "p"}:
				m.P = NewCT_MdxMemeberProp()
				if err := d.DecodeElement(m.P, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "k"}:
				m.K = NewCT_MdxKPI()
				if err := d.DecodeElement(m.K, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Mdx %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Mdx
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Mdx and its children
func (m *CT_Mdx) Validate() error {
	return m.ValidateWithPath("CT_Mdx")
}

// ValidateWithPath validates the CT_Mdx and its children, prefixing error messages with path
func (m *CT_Mdx) ValidateWithPath(path string) error {
	if m.FAttr == ST_MdxFunctionTypeUnset {
		return fmt.Errorf("%s/FAttr is a mandatory field", path)
	}
	if err := m.FAttr.ValidateWithPath(path + "/FAttr"); err != nil {
		return err
	}
	if m.T != nil {
		if err := m.T.ValidateWithPath(path + "/T"); err != nil {
			return err
		}
	}
	if m.Ms != nil {
		if err := m.Ms.ValidateWithPath(path + "/Ms"); err != nil {
			return err
		}
	}
	if m.P != nil {
		if err := m.P.ValidateWithPath(path + "/P"); err != nil {
			return err
		}
	}
	if m.K != nil {
		if err := m.K.ValidateWithPath(path + "/K"); err != nil {
			return err
		}
	}
	return nil
}
