// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/928799934/gooxml"
	"github.com/928799934/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_Background struct {
	BwmodeAttr           OfcST_BWMode
	BwpureAttr           OfcST_BWMode
	BwnormalAttr         OfcST_BWMode
	TargetscreensizeAttr OfcST_ScreenSize
	Fill                 *Fill
	IdAttr               *string
	FilledAttr           sharedTypes.ST_TrueFalse
	FillcolorAttr        *string
}

func NewCT_Background() *CT_Background {
	ret := &CT_Background{}
	return ret
}

func (m *CT_Background) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BwmodeAttr != OfcST_BWModeUnset {
		attr, err := m.BwmodeAttr.MarshalXMLAttr(xml.Name{Local: "bwmode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BwpureAttr != OfcST_BWModeUnset {
		attr, err := m.BwpureAttr.MarshalXMLAttr(xml.Name{Local: "bwpure"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BwnormalAttr != OfcST_BWModeUnset {
		attr, err := m.BwnormalAttr.MarshalXMLAttr(xml.Name{Local: "bwnormal"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TargetscreensizeAttr != OfcST_ScreenSizeUnset {
		attr, err := m.TargetscreensizeAttr.MarshalXMLAttr(xml.Name{Local: "targetscreensize"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.FilledAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.FilledAttr.MarshalXMLAttr(xml.Name{Local: "filled"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FillcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillcolor"},
			Value: fmt.Sprintf("%v", *m.FillcolorAttr)})
	}
	e.EncodeToken(start)
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "v:fill"}}
		e.EncodeElement(m.Fill, sefill)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Background) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwmode" {
			m.BwmodeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwpure" {
			m.BwpureAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwnormal" {
			m.BwnormalAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "targetscreensize" {
			m.TargetscreensizeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "filled" {
			m.FilledAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "fillcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FillcolorAttr = &parsed
			continue
		}
	}
lCT_Background:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "fill"}:
				m.Fill = NewFill()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Background %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Background
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Background and its children
func (m *CT_Background) Validate() error {
	return m.ValidateWithPath("CT_Background")
}

// ValidateWithPath validates the CT_Background and its children, prefixing error messages with path
func (m *CT_Background) ValidateWithPath(path string) error {
	if err := m.BwmodeAttr.ValidateWithPath(path + "/BwmodeAttr"); err != nil {
		return err
	}
	if err := m.BwpureAttr.ValidateWithPath(path + "/BwpureAttr"); err != nil {
		return err
	}
	if err := m.BwnormalAttr.ValidateWithPath(path + "/BwnormalAttr"); err != nil {
		return err
	}
	if err := m.TargetscreensizeAttr.ValidateWithPath(path + "/TargetscreensizeAttr"); err != nil {
		return err
	}
	if m.Fill != nil {
		if err := m.Fill.ValidateWithPath(path + "/Fill"); err != nil {
			return err
		}
	}
	if err := m.FilledAttr.ValidateWithPath(path + "/FilledAttr"); err != nil {
		return err
	}
	return nil
}
