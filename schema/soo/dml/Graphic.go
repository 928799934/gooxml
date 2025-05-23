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

type Graphic struct {
	CT_GraphicalObject
}

func NewGraphic() *Graphic {
	ret := &Graphic{}
	ret.CT_GraphicalObject = *NewCT_GraphicalObject()
	return ret
}

func (m *Graphic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_GraphicalObject.MarshalXML(e, start)
}

func (m *Graphic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_GraphicalObject = *NewCT_GraphicalObject()
lGraphic:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "graphicData"}:
				if err := d.DecodeElement(m.GraphicData, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on Graphic %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lGraphic
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Graphic and its children
func (m *Graphic) Validate() error {
	return m.ValidateWithPath("Graphic")
}

// ValidateWithPath validates the Graphic and its children, prefixing error messages with path
func (m *Graphic) ValidateWithPath(path string) error {
	if err := m.CT_GraphicalObject.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
