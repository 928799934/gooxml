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

type CT_Cell3D struct {
	PrstMaterialAttr ST_PresetMaterialType
	Bevel            *CT_Bevel
	LightRig         *CT_LightRig
	ExtLst           *CT_OfficeArtExtensionList
}

func NewCT_Cell3D() *CT_Cell3D {
	ret := &CT_Cell3D{}
	ret.Bevel = NewCT_Bevel()
	return ret
}

func (m *CT_Cell3D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PrstMaterialAttr != ST_PresetMaterialTypeUnset {
		attr, err := m.PrstMaterialAttr.MarshalXMLAttr(xml.Name{Local: "prstMaterial"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	sebevel := xml.StartElement{Name: xml.Name{Local: "a:bevel"}}
	e.EncodeElement(m.Bevel, sebevel)
	if m.LightRig != nil {
		selightRig := xml.StartElement{Name: xml.Name{Local: "a:lightRig"}}
		e.EncodeElement(m.LightRig, selightRig)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Cell3D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Bevel = NewCT_Bevel()
	for _, attr := range start.Attr {
		if attr.Name.Local == "prstMaterial" {
			m.PrstMaterialAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_Cell3D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "bevel"}:
				if err := d.DecodeElement(m.Bevel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lightRig"}:
				m.LightRig = NewCT_LightRig()
				if err := d.DecodeElement(m.LightRig, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Cell3D %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Cell3D
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Cell3D and its children
func (m *CT_Cell3D) Validate() error {
	return m.ValidateWithPath("CT_Cell3D")
}

// ValidateWithPath validates the CT_Cell3D and its children, prefixing error messages with path
func (m *CT_Cell3D) ValidateWithPath(path string) error {
	if err := m.PrstMaterialAttr.ValidateWithPath(path + "/PrstMaterialAttr"); err != nil {
		return err
	}
	if err := m.Bevel.ValidateWithPath(path + "/Bevel"); err != nil {
		return err
	}
	if m.LightRig != nil {
		if err := m.LightRig.ValidateWithPath(path + "/LightRig"); err != nil {
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
