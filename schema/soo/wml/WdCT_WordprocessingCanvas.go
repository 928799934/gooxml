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
	"github.com/928799934/gooxml/schema/soo/dml"
)

type WdCT_WordprocessingCanvas struct {
	Bg     *dml.CT_BackgroundFormatting
	Whole  *dml.CT_WholeE2oFormatting
	Choice []*WdCT_WordprocessingCanvasChoice
	ExtLst *dml.CT_OfficeArtExtensionList
}

func NewWdCT_WordprocessingCanvas() *WdCT_WordprocessingCanvas {
	ret := &WdCT_WordprocessingCanvas{}
	return ret
}

func (m *WdCT_WordprocessingCanvas) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Bg != nil {
		sebg := xml.StartElement{Name: xml.Name{Local: "wp:bg"}}
		e.EncodeElement(m.Bg, sebg)
	}
	if m.Whole != nil {
		sewhole := xml.StartElement{Name: xml.Name{Local: "wp:whole"}}
		e.EncodeElement(m.Whole, sewhole)
	}
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_WordprocessingCanvas) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lWdCT_WordprocessingCanvas:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "bg"}:
				m.Bg = dml.NewCT_BackgroundFormatting()
				if err := d.DecodeElement(m.Bg, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "whole"}:
				m.Whole = dml.NewCT_WholeE2oFormatting()
				if err := d.DecodeElement(m.Whole, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wsp"}:
				tmp := NewWdCT_WordprocessingCanvasChoice()
				if err := d.DecodeElement(&tmp.Wsp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/picture", Local: "pic"}:
				tmp := NewWdCT_WordprocessingCanvasChoice()
				if err := d.DecodeElement(&tmp.Pic, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "contentPart"}:
				tmp := NewWdCT_WordprocessingCanvasChoice()
				if err := d.DecodeElement(&tmp.ContentPart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wgp"}:
				tmp := NewWdCT_WordprocessingCanvasChoice()
				if err := d.DecodeElement(&tmp.Wgp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "graphicFrame"}:
				tmp := NewWdCT_WordprocessingCanvasChoice()
				if err := d.DecodeElement(&tmp.GraphicFrame, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdCT_WordprocessingCanvas %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_WordprocessingCanvas
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_WordprocessingCanvas and its children
func (m *WdCT_WordprocessingCanvas) Validate() error {
	return m.ValidateWithPath("WdCT_WordprocessingCanvas")
}

// ValidateWithPath validates the WdCT_WordprocessingCanvas and its children, prefixing error messages with path
func (m *WdCT_WordprocessingCanvas) ValidateWithPath(path string) error {
	if m.Bg != nil {
		if err := m.Bg.ValidateWithPath(path + "/Bg"); err != nil {
			return err
		}
	}
	if m.Whole != nil {
		if err := m.Whole.ValidateWithPath(path + "/Whole"); err != nil {
			return err
		}
	}
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
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
