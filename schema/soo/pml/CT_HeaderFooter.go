// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/928799934/gooxml"
)

type CT_HeaderFooter struct {
	// Slide Number Placeholder
	SldNumAttr *bool
	// Header Placeholder
	HdrAttr *bool
	// Footer Placeholder
	FtrAttr *bool
	// Date/Time Placeholder
	DtAttr *bool
	ExtLst *CT_ExtensionListModify
}

func NewCT_HeaderFooter() *CT_HeaderFooter {
	ret := &CT_HeaderFooter{}
	return ret
}

func (m *CT_HeaderFooter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SldNumAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sldNum"},
			Value: fmt.Sprintf("%d", b2i(*m.SldNumAttr))})
	}
	if m.HdrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hdr"},
			Value: fmt.Sprintf("%d", b2i(*m.HdrAttr))})
	}
	if m.FtrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ftr"},
			Value: fmt.Sprintf("%d", b2i(*m.FtrAttr))})
	}
	if m.DtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dt"},
			Value: fmt.Sprintf("%d", b2i(*m.DtAttr))})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_HeaderFooter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sldNum" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SldNumAttr = &parsed
			continue
		}
		if attr.Name.Local == "hdr" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HdrAttr = &parsed
			continue
		}
		if attr.Name.Local == "ftr" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FtrAttr = &parsed
			continue
		}
		if attr.Name.Local == "dt" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DtAttr = &parsed
			continue
		}
	}
lCT_HeaderFooter:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_HeaderFooter %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_HeaderFooter
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_HeaderFooter and its children
func (m *CT_HeaderFooter) Validate() error {
	return m.ValidateWithPath("CT_HeaderFooter")
}

// ValidateWithPath validates the CT_HeaderFooter and its children, prefixing error messages with path
func (m *CT_HeaderFooter) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
