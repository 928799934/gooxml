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
	"fmt"
	"strconv"

	"github.com/928799934/gooxml"
	"github.com/928799934/gooxml/schema/soo/dml"
)

type CT_TLCommonMediaNodeData struct {
	// Volume
	VolAttr *dml.ST_PositiveFixedPercentage
	// Mute
	MuteAttr *bool
	// Number of Slides
	NumSldAttr *uint32
	// Show When Stopped
	ShowWhenStoppedAttr *bool
	// Common Time Node Properties
	CTn   *CT_TLCommonTimeNodeData
	TgtEl *CT_TLTimeTargetElement
}

func NewCT_TLCommonMediaNodeData() *CT_TLCommonMediaNodeData {
	ret := &CT_TLCommonMediaNodeData{}
	ret.CTn = NewCT_TLCommonTimeNodeData()
	ret.TgtEl = NewCT_TLTimeTargetElement()
	return ret
}

func (m *CT_TLCommonMediaNodeData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.VolAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "vol"},
			Value: fmt.Sprintf("%v", *m.VolAttr)})
	}
	if m.MuteAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "mute"},
			Value: fmt.Sprintf("%d", b2i(*m.MuteAttr))})
	}
	if m.NumSldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "numSld"},
			Value: fmt.Sprintf("%v", *m.NumSldAttr)})
	}
	if m.ShowWhenStoppedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showWhenStopped"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowWhenStoppedAttr))})
	}
	e.EncodeToken(start)
	secTn := xml.StartElement{Name: xml.Name{Local: "p:cTn"}}
	e.EncodeElement(m.CTn, secTn)
	setgtEl := xml.StartElement{Name: xml.Name{Local: "p:tgtEl"}}
	e.EncodeElement(m.TgtEl, setgtEl)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLCommonMediaNodeData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CTn = NewCT_TLCommonTimeNodeData()
	m.TgtEl = NewCT_TLTimeTargetElement()
	for _, attr := range start.Attr {
		if attr.Name.Local == "vol" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.VolAttr = &parsed
			continue
		}
		if attr.Name.Local == "mute" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MuteAttr = &parsed
			continue
		}
		if attr.Name.Local == "numSld" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.NumSldAttr = &pt
			continue
		}
		if attr.Name.Local == "showWhenStopped" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowWhenStoppedAttr = &parsed
			continue
		}
	}
lCT_TLCommonMediaNodeData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cTn"}:
				if err := d.DecodeElement(m.CTn, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "tgtEl"}:
				if err := d.DecodeElement(m.TgtEl, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TLCommonMediaNodeData %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLCommonMediaNodeData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLCommonMediaNodeData and its children
func (m *CT_TLCommonMediaNodeData) Validate() error {
	return m.ValidateWithPath("CT_TLCommonMediaNodeData")
}

// ValidateWithPath validates the CT_TLCommonMediaNodeData and its children, prefixing error messages with path
func (m *CT_TLCommonMediaNodeData) ValidateWithPath(path string) error {
	if m.VolAttr != nil {
		if err := m.VolAttr.ValidateWithPath(path + "/VolAttr"); err != nil {
			return err
		}
	}
	if err := m.CTn.ValidateWithPath(path + "/CTn"); err != nil {
		return err
	}
	if err := m.TgtEl.ValidateWithPath(path + "/TgtEl"); err != nil {
		return err
	}
	return nil
}
