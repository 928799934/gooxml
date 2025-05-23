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

	"github.com/928799934/gooxml/schema/soo/ofc/sharedTypes"
)

type OfcCT_Ink struct {
	IAttr           *string
	AnnotationAttr  sharedTypes.ST_TrueFalse
	ContentTypeAttr *string
}

func NewOfcCT_Ink() *OfcCT_Ink {
	ret := &OfcCT_Ink{}
	return ret
}

func (m *OfcCT_Ink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "i"},
			Value: fmt.Sprintf("%v", *m.IAttr)})
	}
	if m.AnnotationAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AnnotationAttr.MarshalXMLAttr(xml.Name{Local: "annotation"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ContentTypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "contentType"},
			Value: fmt.Sprintf("%v", *m.ContentTypeAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_Ink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "i" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IAttr = &parsed
			continue
		}
		if attr.Name.Local == "annotation" {
			m.AnnotationAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "contentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentTypeAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_Ink: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_Ink and its children
func (m *OfcCT_Ink) Validate() error {
	return m.ValidateWithPath("OfcCT_Ink")
}

// ValidateWithPath validates the OfcCT_Ink and its children, prefixing error messages with path
func (m *OfcCT_Ink) ValidateWithPath(path string) error {
	if err := m.AnnotationAttr.ValidateWithPath(path + "/AnnotationAttr"); err != nil {
		return err
	}
	return nil
}
