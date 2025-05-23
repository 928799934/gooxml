// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/928799934/gooxml"
	"github.com/928799934/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_Filters struct {
	// Filter by Blank
	BlankAttr *bool
	// Calendar Type
	CalendarTypeAttr sharedTypes.ST_CalendarType
	// Filter
	Filter []*CT_Filter
	// Date Grouping
	DateGroupItem []*CT_DateGroupItem
}

func NewCT_Filters() *CT_Filters {
	ret := &CT_Filters{}
	return ret
}

func (m *CT_Filters) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BlankAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "blank"},
			Value: fmt.Sprintf("%d", b2i(*m.BlankAttr))})
	}
	if m.CalendarTypeAttr != sharedTypes.ST_CalendarTypeUnset {
		attr, err := m.CalendarTypeAttr.MarshalXMLAttr(xml.Name{Local: "calendarType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.Filter != nil {
		sefilter := xml.StartElement{Name: xml.Name{Local: "ma:filter"}}
		for _, c := range m.Filter {
			e.EncodeElement(c, sefilter)
		}
	}
	if m.DateGroupItem != nil {
		sedateGroupItem := xml.StartElement{Name: xml.Name{Local: "ma:dateGroupItem"}}
		for _, c := range m.DateGroupItem {
			e.EncodeElement(c, sedateGroupItem)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Filters) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "blank" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BlankAttr = &parsed
			continue
		}
		if attr.Name.Local == "calendarType" {
			m.CalendarTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_Filters:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "filter"}:
				tmp := NewCT_Filter()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Filter = append(m.Filter, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "dateGroupItem"}:
				tmp := NewCT_DateGroupItem()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DateGroupItem = append(m.DateGroupItem, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_Filters %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Filters
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Filters and its children
func (m *CT_Filters) Validate() error {
	return m.ValidateWithPath("CT_Filters")
}

// ValidateWithPath validates the CT_Filters and its children, prefixing error messages with path
func (m *CT_Filters) ValidateWithPath(path string) error {
	if err := m.CalendarTypeAttr.ValidateWithPath(path + "/CalendarTypeAttr"); err != nil {
		return err
	}
	for i, v := range m.Filter {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Filter[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.DateGroupItem {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DateGroupItem[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
