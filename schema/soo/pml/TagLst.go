// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"

	"github.com/928799934/gooxml"
)

type TagLst struct {
	CT_TagList
}

func NewTagLst() *TagLst {
	ret := &TagLst{}
	ret.CT_TagList = *NewCT_TagList()
	return ret
}

func (m *TagLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "p:tagLst"
	return m.CT_TagList.MarshalXML(e, start)
}

func (m *TagLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_TagList = *NewCT_TagList()
lTagLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "tag"}:
				tmp := NewCT_StringTag()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tag = append(m.Tag, tmp)
			default:
				gooxml.Log("skipping unsupported element on TagLst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lTagLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the TagLst and its children
func (m *TagLst) Validate() error {
	return m.ValidateWithPath("TagLst")
}

// ValidateWithPath validates the TagLst and its children, prefixing error messages with path
func (m *TagLst) ValidateWithPath(path string) error {
	if err := m.CT_TagList.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
