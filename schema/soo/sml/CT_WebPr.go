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
)

type CT_WebPr struct {
	// XML Source
	XmlAttr *bool
	// Import XML Source Data
	SourceDataAttr *bool
	// Parse PRE
	ParsePreAttr *bool
	// Consecutive Delimiters
	ConsecutiveAttr *bool
	// Use First Row
	FirstRowAttr *bool
	// Created in Excel 97
	Xl97Attr *bool
	// Dates as Text
	TextDatesAttr *bool
	// Refreshed in Excel 2000
	Xl2000Attr *bool
	// URL
	UrlAttr *string
	// Web Post
	PostAttr *string
	// HTML Tables Only
	HtmlTablesAttr *bool
	// HTML Formatting Handling
	HtmlFormatAttr ST_HtmlFmt
	// Edit Query URL
	EditPageAttr *string
	// Tables
	Tables *CT_Tables
}

func NewCT_WebPr() *CT_WebPr {
	ret := &CT_WebPr{}
	return ret
}

func (m *CT_WebPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.XmlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xml"},
			Value: fmt.Sprintf("%d", b2i(*m.XmlAttr))})
	}
	if m.SourceDataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sourceData"},
			Value: fmt.Sprintf("%d", b2i(*m.SourceDataAttr))})
	}
	if m.ParsePreAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "parsePre"},
			Value: fmt.Sprintf("%d", b2i(*m.ParsePreAttr))})
	}
	if m.ConsecutiveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "consecutive"},
			Value: fmt.Sprintf("%d", b2i(*m.ConsecutiveAttr))})
	}
	if m.FirstRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "firstRow"},
			Value: fmt.Sprintf("%d", b2i(*m.FirstRowAttr))})
	}
	if m.Xl97Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xl97"},
			Value: fmt.Sprintf("%d", b2i(*m.Xl97Attr))})
	}
	if m.TextDatesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "textDates"},
			Value: fmt.Sprintf("%d", b2i(*m.TextDatesAttr))})
	}
	if m.Xl2000Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xl2000"},
			Value: fmt.Sprintf("%d", b2i(*m.Xl2000Attr))})
	}
	if m.UrlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "url"},
			Value: fmt.Sprintf("%v", *m.UrlAttr)})
	}
	if m.PostAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "post"},
			Value: fmt.Sprintf("%v", *m.PostAttr)})
	}
	if m.HtmlTablesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "htmlTables"},
			Value: fmt.Sprintf("%d", b2i(*m.HtmlTablesAttr))})
	}
	if m.HtmlFormatAttr != ST_HtmlFmtUnset {
		attr, err := m.HtmlFormatAttr.MarshalXMLAttr(xml.Name{Local: "htmlFormat"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.EditPageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "editPage"},
			Value: fmt.Sprintf("%v", *m.EditPageAttr)})
	}
	e.EncodeToken(start)
	if m.Tables != nil {
		setables := xml.StartElement{Name: xml.Name{Local: "ma:tables"}}
		e.EncodeElement(m.Tables, setables)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WebPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "xl2000" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.Xl2000Attr = &parsed
			continue
		}
		if attr.Name.Local == "xml" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.XmlAttr = &parsed
			continue
		}
		if attr.Name.Local == "parsePre" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ParsePreAttr = &parsed
			continue
		}
		if attr.Name.Local == "consecutive" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ConsecutiveAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FirstRowAttr = &parsed
			continue
		}
		if attr.Name.Local == "xl97" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.Xl97Attr = &parsed
			continue
		}
		if attr.Name.Local == "textDates" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TextDatesAttr = &parsed
			continue
		}
		if attr.Name.Local == "sourceData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SourceDataAttr = &parsed
			continue
		}
		if attr.Name.Local == "url" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UrlAttr = &parsed
			continue
		}
		if attr.Name.Local == "post" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PostAttr = &parsed
			continue
		}
		if attr.Name.Local == "htmlTables" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HtmlTablesAttr = &parsed
			continue
		}
		if attr.Name.Local == "htmlFormat" {
			m.HtmlFormatAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "editPage" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EditPageAttr = &parsed
			continue
		}
	}
lCT_WebPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tables"}:
				m.Tables = NewCT_Tables()
				if err := d.DecodeElement(m.Tables, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_WebPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WebPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WebPr and its children
func (m *CT_WebPr) Validate() error {
	return m.ValidateWithPath("CT_WebPr")
}

// ValidateWithPath validates the CT_WebPr and its children, prefixing error messages with path
func (m *CT_WebPr) ValidateWithPath(path string) error {
	if err := m.HtmlFormatAttr.ValidateWithPath(path + "/HtmlFormatAttr"); err != nil {
		return err
	}
	if m.Tables != nil {
		if err := m.Tables.ValidateWithPath(path + "/Tables"); err != nil {
			return err
		}
	}
	return nil
}
