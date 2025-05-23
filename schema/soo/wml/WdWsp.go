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
	"strconv"

	"github.com/928799934/gooxml"
	"github.com/928799934/gooxml/schema/soo/dml"
)

type WdWsp struct {
	WdCT_WordprocessingShape
}

func NewWdWsp() *WdWsp {
	ret := &WdWsp{}
	ret.WdCT_WordprocessingShape = *NewWdCT_WordprocessingShape()
	return ret
}

func (m *WdWsp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.WdCT_WordprocessingShape.MarshalXML(e, start)
}

func (m *WdWsp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.WdCT_WordprocessingShape = *NewWdCT_WordprocessingShape()
	for _, attr := range start.Attr {
		if attr.Name.Local == "normalEastAsianFlow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NormalEastAsianFlowAttr = &parsed
			continue
		}
	}
lWdWsp:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvPr"}:
				m.CNvPr = dml.NewCT_NonVisualDrawingProps()
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvSpPr"}:
				m.Choice = NewWdCT_WordprocessingShapeChoice()
				if err := d.DecodeElement(&m.Choice.CNvSpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvCnPr"}:
				m.Choice = NewWdCT_WordprocessingShapeChoice()
				if err := d.DecodeElement(&m.Choice.CNvCnPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "spPr"}:
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "style"}:
				m.Style = dml.NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "txbx"}:
				m.WChoice = NewWdCT_WordprocessingShapeChoice1()
				if err := d.DecodeElement(&m.WChoice.Txbx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "linkedTxbx"}:
				m.WChoice = NewWdCT_WordprocessingShapeChoice1()
				if err := d.DecodeElement(&m.WChoice.LinkedTxbx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "bodyPr"}:
				if err := d.DecodeElement(m.BodyPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdWsp %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdWsp
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdWsp and its children
func (m *WdWsp) Validate() error {
	return m.ValidateWithPath("WdWsp")
}

// ValidateWithPath validates the WdWsp and its children, prefixing error messages with path
func (m *WdWsp) ValidateWithPath(path string) error {
	if err := m.WdCT_WordprocessingShape.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
