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
	"strconv"

	"github.com/928799934/gooxml"
)

type Stroke struct {
	CT_Stroke
}

func NewStroke() *Stroke {
	ret := &Stroke{}
	ret.CT_Stroke = *NewCT_Stroke()
	return ret
}

func (m *Stroke) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Stroke.MarshalXML(e, start)
}

func (m *Stroke) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Stroke = *NewCT_Stroke()
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "relid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RelidAttr = &parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RIdAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "forcedash" {
			m.ForcedashAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "althref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlthrefAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "href" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HrefAttr = &parsed
			continue
		}
		if attr.Name.Local == "imagesize" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ImagesizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "startarrowlength" {
			m.StartarrowlengthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "joinstyle" {
			m.JoinstyleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "endarrow" {
			m.EndarrowAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dashstyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DashstyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "endarrowwidth" {
			m.EndarrowwidthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "src" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SrcAttr = &parsed
			continue
		}
		if attr.Name.Local == "endarrowlength" {
			m.EndarrowlengthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "linestyle" {
			m.LinestyleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "opacity" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OpacityAttr = &parsed
			continue
		}
		if attr.Name.Local == "weight" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WeightAttr = &parsed
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
		if attr.Name.Local == "miterlimit" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.MiterlimitAttr = &parsed
			continue
		}
		if attr.Name.Local == "color" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ColorAttr = &parsed
			continue
		}
		if attr.Name.Local == "imageaspect" {
			m.ImageaspectAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "filltype" {
			m.FilltypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "startarrowwidth" {
			m.StartarrowwidthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "insetpen" {
			m.InsetpenAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "imagealignshape" {
			m.ImagealignshapeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "startarrow" {
			m.StartarrowAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "endcap" {
			m.EndcapAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "color2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Color2Attr = &parsed
			continue
		}
	}
lStroke:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "left"}:
				m.Left = NewOfcLeft()
				if err := d.DecodeElement(m.Left, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "top"}:
				m.Top = NewOfcTop()
				if err := d.DecodeElement(m.Top, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "right"}:
				m.Right = NewOfcRight()
				if err := d.DecodeElement(m.Right, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "bottom"}:
				m.Bottom = NewOfcBottom()
				if err := d.DecodeElement(m.Bottom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "column"}:
				m.Column = NewOfcColumn()
				if err := d.DecodeElement(m.Column, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on Stroke %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lStroke
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Stroke and its children
func (m *Stroke) Validate() error {
	return m.ValidateWithPath("Stroke")
}

// ValidateWithPath validates the Stroke and its children, prefixing error messages with path
func (m *Stroke) ValidateWithPath(path string) error {
	if err := m.CT_Stroke.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
