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
	"fmt"

	"github.com/928799934/gooxml"
	"github.com/928799934/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_DocPartTypes struct {
	// Entry Is Of All Types
	AllAttr *sharedTypes.ST_OnOff
	// Entry Type
	Type []*CT_DocPartType
}

func NewCT_DocPartTypes() *CT_DocPartTypes {
	ret := &CT_DocPartTypes{}
	return ret
}

func (m *CT_DocPartTypes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AllAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:all"},
			Value: fmt.Sprintf("%v", *m.AllAttr)})
	}
	e.EncodeToken(start)
	if m.Type != nil {
		setype := xml.StartElement{Name: xml.Name{Local: "w:type"}}
		for _, c := range m.Type {
			e.EncodeElement(c, setype)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocPartTypes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "all" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.AllAttr = &parsed
			continue
		}
	}
lCT_DocPartTypes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "type"}:
				tmp := NewCT_DocPartType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Type = append(m.Type, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_DocPartTypes %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocPartTypes
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DocPartTypes and its children
func (m *CT_DocPartTypes) Validate() error {
	return m.ValidateWithPath("CT_DocPartTypes")
}

// ValidateWithPath validates the CT_DocPartTypes and its children, prefixing error messages with path
func (m *CT_DocPartTypes) ValidateWithPath(path string) error {
	if m.AllAttr != nil {
		if err := m.AllAttr.ValidateWithPath(path + "/AllAttr"); err != nil {
			return err
		}
	}
	for i, v := range m.Type {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Type[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
