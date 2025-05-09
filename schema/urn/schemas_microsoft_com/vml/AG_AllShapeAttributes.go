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
	"strconv"

	"github.com/928799934/gooxml/schema/soo/ofc/sharedTypes"
)

type AG_AllShapeAttributes struct {
	OpacityAttr        *string
	StrokedAttr        sharedTypes.ST_TrueFalse
	StrokecolorAttr    *string
	StrokeweightAttr   *string
	InsetpenAttr       sharedTypes.ST_TrueFalse
	ChromakeyAttr      *string
	FilledAttr         sharedTypes.ST_TrueFalse
	FillcolorAttr      *string
	SptAttr            *float32
	ConnectortypeAttr  OfcST_ConnectorType
	BwmodeAttr         OfcST_BWMode
	BwpureAttr         OfcST_BWMode
	BwnormalAttr       OfcST_BWMode
	ForcedashAttr      sharedTypes.ST_TrueFalse
	OleiconAttr        sharedTypes.ST_TrueFalse
	OleAttr            sharedTypes.ST_TrueFalseBlank
	PreferrelativeAttr sharedTypes.ST_TrueFalse
	CliptowrapAttr     sharedTypes.ST_TrueFalse
	ClipAttr           sharedTypes.ST_TrueFalse
}

func NewAG_AllShapeAttributes() *AG_AllShapeAttributes {
	ret := &AG_AllShapeAttributes{}
	return ret
}

func (m *AG_AllShapeAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.OpacityAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "opacity"},
			Value: fmt.Sprintf("%v", *m.OpacityAttr)})
	}
	if m.StrokedAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.StrokedAttr.MarshalXMLAttr(xml.Name{Local: "stroked"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StrokecolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "strokecolor"},
			Value: fmt.Sprintf("%v", *m.StrokecolorAttr)})
	}
	if m.StrokeweightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "strokeweight"},
			Value: fmt.Sprintf("%v", *m.StrokeweightAttr)})
	}
	if m.InsetpenAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.InsetpenAttr.MarshalXMLAttr(xml.Name{Local: "insetpen"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ChromakeyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "chromakey"},
			Value: fmt.Sprintf("%v", *m.ChromakeyAttr)})
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
	if m.SptAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:spt"},
			Value: fmt.Sprintf("%v", *m.SptAttr)})
	}
	if m.ConnectortypeAttr != OfcST_ConnectorTypeUnset {
		attr, err := m.ConnectortypeAttr.MarshalXMLAttr(xml.Name{Local: "connectortype"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
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
	if m.ForcedashAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ForcedashAttr.MarshalXMLAttr(xml.Name{Local: "forcedash"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OleiconAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.OleiconAttr.MarshalXMLAttr(xml.Name{Local: "oleicon"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OleAttr != sharedTypes.ST_TrueFalseBlankUnset {
		attr, err := m.OleAttr.MarshalXMLAttr(xml.Name{Local: "ole"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.PreferrelativeAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.PreferrelativeAttr.MarshalXMLAttr(xml.Name{Local: "preferrelative"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CliptowrapAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.CliptowrapAttr.MarshalXMLAttr(xml.Name{Local: "cliptowrap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ClipAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ClipAttr.MarshalXMLAttr(xml.Name{Local: "clip"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	return nil
}

func (m *AG_AllShapeAttributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "connectortype" {
			m.ConnectortypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwpure" {
			m.BwpureAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "cliptowrap" {
			m.CliptowrapAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "preferrelative" {
			m.PreferrelativeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "ole" {
			m.OleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "oleicon" {
			m.OleiconAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "forcedash" {
			m.ForcedashAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwnormal" {
			m.BwnormalAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "spt" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.SptAttr = &pt
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "clip" {
			m.ClipAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwmode" {
			m.BwmodeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "stroked" {
			m.StrokedAttr.UnmarshalXMLAttr(attr)
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
		if attr.Name.Local == "filled" {
			m.FilledAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "chromakey" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ChromakeyAttr = &parsed
			continue
		}
		if attr.Name.Local == "insetpen" {
			m.InsetpenAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "strokeweight" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StrokeweightAttr = &parsed
			continue
		}
		if attr.Name.Local == "strokecolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StrokecolorAttr = &parsed
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
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_AllShapeAttributes: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_AllShapeAttributes and its children
func (m *AG_AllShapeAttributes) Validate() error {
	return m.ValidateWithPath("AG_AllShapeAttributes")
}

// ValidateWithPath validates the AG_AllShapeAttributes and its children, prefixing error messages with path
func (m *AG_AllShapeAttributes) ValidateWithPath(path string) error {
	if err := m.StrokedAttr.ValidateWithPath(path + "/StrokedAttr"); err != nil {
		return err
	}
	if err := m.InsetpenAttr.ValidateWithPath(path + "/InsetpenAttr"); err != nil {
		return err
	}
	if err := m.FilledAttr.ValidateWithPath(path + "/FilledAttr"); err != nil {
		return err
	}
	if err := m.ConnectortypeAttr.ValidateWithPath(path + "/ConnectortypeAttr"); err != nil {
		return err
	}
	if err := m.BwmodeAttr.ValidateWithPath(path + "/BwmodeAttr"); err != nil {
		return err
	}
	if err := m.BwpureAttr.ValidateWithPath(path + "/BwpureAttr"); err != nil {
		return err
	}
	if err := m.BwnormalAttr.ValidateWithPath(path + "/BwnormalAttr"); err != nil {
		return err
	}
	if err := m.ForcedashAttr.ValidateWithPath(path + "/ForcedashAttr"); err != nil {
		return err
	}
	if err := m.OleiconAttr.ValidateWithPath(path + "/OleiconAttr"); err != nil {
		return err
	}
	if err := m.OleAttr.ValidateWithPath(path + "/OleAttr"); err != nil {
		return err
	}
	if err := m.PreferrelativeAttr.ValidateWithPath(path + "/PreferrelativeAttr"); err != nil {
		return err
	}
	if err := m.CliptowrapAttr.ValidateWithPath(path + "/CliptowrapAttr"); err != nil {
		return err
	}
	if err := m.ClipAttr.ValidateWithPath(path + "/ClipAttr"); err != nil {
		return err
	}
	return nil
}
