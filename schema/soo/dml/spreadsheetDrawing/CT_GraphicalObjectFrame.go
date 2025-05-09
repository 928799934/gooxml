// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/928799934/gooxml"
	"github.com/928799934/gooxml/schema/soo/dml"
)

type CT_GraphicalObjectFrame struct {
	MacroAttr        *string
	FPublishedAttr   *bool
	NvGraphicFramePr *CT_GraphicalObjectFrameNonVisual
	Xfrm             *dml.CT_Transform2D
	Graphic          *dml.Graphic
}

func NewCT_GraphicalObjectFrame() *CT_GraphicalObjectFrame {
	ret := &CT_GraphicalObjectFrame{}
	ret.NvGraphicFramePr = NewCT_GraphicalObjectFrameNonVisual()
	ret.Xfrm = dml.NewCT_Transform2D()
	ret.Graphic = dml.NewGraphic()
	return ret
}

func (m *CT_GraphicalObjectFrame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.MacroAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "macro"},
			Value: fmt.Sprintf("%v", *m.MacroAttr)})
	}
	if m.FPublishedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fPublished"},
			Value: fmt.Sprintf("%d", b2i(*m.FPublishedAttr))})
	}
	e.EncodeToken(start)
	senvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "xdr:nvGraphicFramePr"}}
	e.EncodeElement(m.NvGraphicFramePr, senvGraphicFramePr)
	sexfrm := xml.StartElement{Name: xml.Name{Local: "xdr:xfrm"}}
	e.EncodeElement(m.Xfrm, sexfrm)
	segraphic := xml.StartElement{Name: xml.Name{Local: "a:graphic"}}
	e.EncodeElement(m.Graphic, segraphic)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GraphicalObjectFrame) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvGraphicFramePr = NewCT_GraphicalObjectFrameNonVisual()
	m.Xfrm = dml.NewCT_Transform2D()
	m.Graphic = dml.NewGraphic()
	for _, attr := range start.Attr {
		if attr.Name.Local == "macro" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MacroAttr = &parsed
			continue
		}
		if attr.Name.Local == "fPublished" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FPublishedAttr = &parsed
			continue
		}
	}
lCT_GraphicalObjectFrame:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "nvGraphicFramePr"}:
				if err := d.DecodeElement(m.NvGraphicFramePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "xfrm"}:
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "graphic"}:
				if err := d.DecodeElement(m.Graphic, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_GraphicalObjectFrame %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GraphicalObjectFrame
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GraphicalObjectFrame and its children
func (m *CT_GraphicalObjectFrame) Validate() error {
	return m.ValidateWithPath("CT_GraphicalObjectFrame")
}

// ValidateWithPath validates the CT_GraphicalObjectFrame and its children, prefixing error messages with path
func (m *CT_GraphicalObjectFrame) ValidateWithPath(path string) error {
	if err := m.NvGraphicFramePr.ValidateWithPath(path + "/NvGraphicFramePr"); err != nil {
		return err
	}
	if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
		return err
	}
	if err := m.Graphic.ValidateWithPath(path + "/Graphic"); err != nil {
		return err
	}
	return nil
}
