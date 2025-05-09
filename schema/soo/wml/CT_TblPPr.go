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

	"github.com/928799934/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_TblPPr struct {
	// Distance From Left of Table to Text
	LeftFromTextAttr *sharedTypes.ST_TwipsMeasure
	// (Distance From Right of Table to Text
	RightFromTextAttr *sharedTypes.ST_TwipsMeasure
	// Distance From Top of Table to Text
	TopFromTextAttr *sharedTypes.ST_TwipsMeasure
	// Distance From Bottom of Table to Text
	BottomFromTextAttr *sharedTypes.ST_TwipsMeasure
	// Table Vertical Anchor
	VertAnchorAttr ST_VAnchor
	// Table Horizontal Anchor
	HorzAnchorAttr ST_HAnchor
	// Relative Horizontal Alignment From Anchor
	TblpXSpecAttr sharedTypes.ST_XAlign
	// Absolute Horizontal Distance From Anchor
	TblpXAttr *ST_SignedTwipsMeasure
	// Relative Vertical Alignment from Anchor
	TblpYSpecAttr sharedTypes.ST_YAlign
	// Absolute Vertical Distance From Anchor
	TblpYAttr *ST_SignedTwipsMeasure
}

func NewCT_TblPPr() *CT_TblPPr {
	ret := &CT_TblPPr{}
	return ret
}

func (m *CT_TblPPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LeftFromTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:leftFromText"},
			Value: fmt.Sprintf("%v", *m.LeftFromTextAttr)})
	}
	if m.RightFromTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rightFromText"},
			Value: fmt.Sprintf("%v", *m.RightFromTextAttr)})
	}
	if m.TopFromTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:topFromText"},
			Value: fmt.Sprintf("%v", *m.TopFromTextAttr)})
	}
	if m.BottomFromTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:bottomFromText"},
			Value: fmt.Sprintf("%v", *m.BottomFromTextAttr)})
	}
	if m.VertAnchorAttr != ST_VAnchorUnset {
		attr, err := m.VertAnchorAttr.MarshalXMLAttr(xml.Name{Local: "w:vertAnchor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HorzAnchorAttr != ST_HAnchorUnset {
		attr, err := m.HorzAnchorAttr.MarshalXMLAttr(xml.Name{Local: "w:horzAnchor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TblpXSpecAttr != sharedTypes.ST_XAlignUnset {
		attr, err := m.TblpXSpecAttr.MarshalXMLAttr(xml.Name{Local: "w:tblpXSpec"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TblpXAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tblpX"},
			Value: fmt.Sprintf("%v", *m.TblpXAttr)})
	}
	if m.TblpYSpecAttr != sharedTypes.ST_YAlignUnset {
		attr, err := m.TblpYSpecAttr.MarshalXMLAttr(xml.Name{Local: "w:tblpYSpec"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TblpYAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tblpY"},
			Value: fmt.Sprintf("%v", *m.TblpYAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblPPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "leftFromText" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.LeftFromTextAttr = &parsed
			continue
		}
		if attr.Name.Local == "rightFromText" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.RightFromTextAttr = &parsed
			continue
		}
		if attr.Name.Local == "topFromText" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.TopFromTextAttr = &parsed
			continue
		}
		if attr.Name.Local == "bottomFromText" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.BottomFromTextAttr = &parsed
			continue
		}
		if attr.Name.Local == "vertAnchor" {
			m.VertAnchorAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "horzAnchor" {
			m.HorzAnchorAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "tblpXSpec" {
			m.TblpXSpecAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "tblpX" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.TblpXAttr = &parsed
			continue
		}
		if attr.Name.Local == "tblpYSpec" {
			m.TblpYSpecAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "tblpY" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.TblpYAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TblPPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TblPPr and its children
func (m *CT_TblPPr) Validate() error {
	return m.ValidateWithPath("CT_TblPPr")
}

// ValidateWithPath validates the CT_TblPPr and its children, prefixing error messages with path
func (m *CT_TblPPr) ValidateWithPath(path string) error {
	if m.LeftFromTextAttr != nil {
		if err := m.LeftFromTextAttr.ValidateWithPath(path + "/LeftFromTextAttr"); err != nil {
			return err
		}
	}
	if m.RightFromTextAttr != nil {
		if err := m.RightFromTextAttr.ValidateWithPath(path + "/RightFromTextAttr"); err != nil {
			return err
		}
	}
	if m.TopFromTextAttr != nil {
		if err := m.TopFromTextAttr.ValidateWithPath(path + "/TopFromTextAttr"); err != nil {
			return err
		}
	}
	if m.BottomFromTextAttr != nil {
		if err := m.BottomFromTextAttr.ValidateWithPath(path + "/BottomFromTextAttr"); err != nil {
			return err
		}
	}
	if err := m.VertAnchorAttr.ValidateWithPath(path + "/VertAnchorAttr"); err != nil {
		return err
	}
	if err := m.HorzAnchorAttr.ValidateWithPath(path + "/HorzAnchorAttr"); err != nil {
		return err
	}
	if err := m.TblpXSpecAttr.ValidateWithPath(path + "/TblpXSpecAttr"); err != nil {
		return err
	}
	if m.TblpXAttr != nil {
		if err := m.TblpXAttr.ValidateWithPath(path + "/TblpXAttr"); err != nil {
			return err
		}
	}
	if err := m.TblpYSpecAttr.ValidateWithPath(path + "/TblpYSpecAttr"); err != nil {
		return err
	}
	if m.TblpYAttr != nil {
		if err := m.TblpYAttr.ValidateWithPath(path + "/TblpYAttr"); err != nil {
			return err
		}
	}
	return nil
}
