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

type CT_DdeItem struct {
	// DDE Name
	NameAttr *string
	// Object Linking TechnologyE
	OleAttr *bool
	// Advise
	AdviseAttr *bool
	// Data is an Image
	PreferPicAttr *bool
	// DDE Name Values
	Values *CT_DdeValues
}

func NewCT_DdeItem() *CT_DdeItem {
	ret := &CT_DdeItem{}
	return ret
}

func (m *CT_DdeItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.OleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ole"},
			Value: fmt.Sprintf("%d", b2i(*m.OleAttr))})
	}
	if m.AdviseAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "advise"},
			Value: fmt.Sprintf("%d", b2i(*m.AdviseAttr))})
	}
	if m.PreferPicAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preferPic"},
			Value: fmt.Sprintf("%d", b2i(*m.PreferPicAttr))})
	}
	e.EncodeToken(start)
	if m.Values != nil {
		sevalues := xml.StartElement{Name: xml.Name{Local: "ma:values"}}
		e.EncodeElement(m.Values, sevalues)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DdeItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
		if attr.Name.Local == "ole" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OleAttr = &parsed
			continue
		}
		if attr.Name.Local == "advise" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AdviseAttr = &parsed
			continue
		}
		if attr.Name.Local == "preferPic" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreferPicAttr = &parsed
			continue
		}
	}
lCT_DdeItem:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "values"}:
				m.Values = NewCT_DdeValues()
				if err := d.DecodeElement(m.Values, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_DdeItem %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DdeItem
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DdeItem and its children
func (m *CT_DdeItem) Validate() error {
	return m.ValidateWithPath("CT_DdeItem")
}

// ValidateWithPath validates the CT_DdeItem and its children, prefixing error messages with path
func (m *CT_DdeItem) ValidateWithPath(path string) error {
	if m.Values != nil {
		if err := m.Values.ValidateWithPath(path + "/Values"); err != nil {
			return err
		}
	}
	return nil
}
