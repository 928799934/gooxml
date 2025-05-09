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

type CT_TrackChangesView struct {
	// Display Visual Indicator Of Markup Area
	MarkupAttr *sharedTypes.ST_OnOff
	// Display Comments
	CommentsAttr *sharedTypes.ST_OnOff
	// Display Content Revisions
	InsDelAttr *sharedTypes.ST_OnOff
	// Display Formatting Revisions
	FormattingAttr *sharedTypes.ST_OnOff
	// Display Ink Annotations
	InkAnnotationsAttr *sharedTypes.ST_OnOff
}

func NewCT_TrackChangesView() *CT_TrackChangesView {
	ret := &CT_TrackChangesView{}
	return ret
}

func (m *CT_TrackChangesView) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.MarkupAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:markup"},
			Value: fmt.Sprintf("%v", *m.MarkupAttr)})
	}
	if m.CommentsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:comments"},
			Value: fmt.Sprintf("%v", *m.CommentsAttr)})
	}
	if m.InsDelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:insDel"},
			Value: fmt.Sprintf("%v", *m.InsDelAttr)})
	}
	if m.FormattingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:formatting"},
			Value: fmt.Sprintf("%v", *m.FormattingAttr)})
	}
	if m.InkAnnotationsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:inkAnnotations"},
			Value: fmt.Sprintf("%v", *m.InkAnnotationsAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TrackChangesView) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "markup" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.MarkupAttr = &parsed
			continue
		}
		if attr.Name.Local == "comments" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.CommentsAttr = &parsed
			continue
		}
		if attr.Name.Local == "insDel" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.InsDelAttr = &parsed
			continue
		}
		if attr.Name.Local == "formatting" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.FormattingAttr = &parsed
			continue
		}
		if attr.Name.Local == "inkAnnotations" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.InkAnnotationsAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TrackChangesView: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TrackChangesView and its children
func (m *CT_TrackChangesView) Validate() error {
	return m.ValidateWithPath("CT_TrackChangesView")
}

// ValidateWithPath validates the CT_TrackChangesView and its children, prefixing error messages with path
func (m *CT_TrackChangesView) ValidateWithPath(path string) error {
	if m.MarkupAttr != nil {
		if err := m.MarkupAttr.ValidateWithPath(path + "/MarkupAttr"); err != nil {
			return err
		}
	}
	if m.CommentsAttr != nil {
		if err := m.CommentsAttr.ValidateWithPath(path + "/CommentsAttr"); err != nil {
			return err
		}
	}
	if m.InsDelAttr != nil {
		if err := m.InsDelAttr.ValidateWithPath(path + "/InsDelAttr"); err != nil {
			return err
		}
	}
	if m.FormattingAttr != nil {
		if err := m.FormattingAttr.ValidateWithPath(path + "/FormattingAttr"); err != nil {
			return err
		}
	}
	if m.InkAnnotationsAttr != nil {
		if err := m.InkAnnotationsAttr.ValidateWithPath(path + "/InkAnnotationsAttr"); err != nil {
			return err
		}
	}
	return nil
}
