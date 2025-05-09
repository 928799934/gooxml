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
	"github.com/928799934/gooxml/schema/soo/dml/picture"
)

type WdCT_WordprocessingGroupChoice struct {
	Wsp          []*WdWsp
	GrpSp        []*WdCT_WordprocessingGroup
	GraphicFrame []*WdCT_GraphicFrame
	Pic          []*picture.Pic
	ContentPart  []*WdCT_WordprocessingContentPart
}

func NewWdCT_WordprocessingGroupChoice() *WdCT_WordprocessingGroupChoice {
	ret := &WdCT_WordprocessingGroupChoice{}
	return ret
}

func (m *WdCT_WordprocessingGroupChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Wsp != nil {
		sewsp := xml.StartElement{Name: xml.Name{Local: "wp:wsp"}}
		for _, c := range m.Wsp {
			e.EncodeElement(c, sewsp)
		}
	}
	if m.GrpSp != nil {
		segrpSp := xml.StartElement{Name: xml.Name{Local: "wp:grpSp"}}
		for _, c := range m.GrpSp {
			e.EncodeElement(c, segrpSp)
		}
	}
	if m.GraphicFrame != nil {
		segraphicFrame := xml.StartElement{Name: xml.Name{Local: "wp:graphicFrame"}}
		for _, c := range m.GraphicFrame {
			e.EncodeElement(c, segraphicFrame)
		}
	}
	if m.Pic != nil {
		sepic := xml.StartElement{Name: xml.Name{Local: "pic:pic"}}
		for _, c := range m.Pic {
			e.EncodeElement(c, sepic)
		}
	}
	if m.ContentPart != nil {
		secontentPart := xml.StartElement{Name: xml.Name{Local: "wp:contentPart"}}
		for _, c := range m.ContentPart {
			e.EncodeElement(c, secontentPart)
		}
	}
	return nil
}

func (m *WdCT_WordprocessingGroupChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lWdCT_WordprocessingGroupChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wsp"}:
				tmp := NewWdWsp()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Wsp = append(m.Wsp, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "grpSp"}:
				tmp := NewWdCT_WordprocessingGroup()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GrpSp = append(m.GrpSp, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "graphicFrame"}:
				tmp := NewWdCT_GraphicFrame()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GraphicFrame = append(m.GraphicFrame, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/picture", Local: "pic"}:
				tmp := picture.NewPic()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Pic = append(m.Pic, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "contentPart"}:
				tmp := NewWdCT_WordprocessingContentPart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ContentPart = append(m.ContentPart, tmp)
			default:
				gooxml.Log("skipping unsupported element on WdCT_WordprocessingGroupChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_WordprocessingGroupChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_WordprocessingGroupChoice and its children
func (m *WdCT_WordprocessingGroupChoice) Validate() error {
	return m.ValidateWithPath("WdCT_WordprocessingGroupChoice")
}

// ValidateWithPath validates the WdCT_WordprocessingGroupChoice and its children, prefixing error messages with path
func (m *WdCT_WordprocessingGroupChoice) ValidateWithPath(path string) error {
	for i, v := range m.Wsp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Wsp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GrpSp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GrpSp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GraphicFrame {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GraphicFrame[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Pic {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Pic[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ContentPart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ContentPart[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
