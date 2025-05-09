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
	"time"

	"github.com/928799934/gooxml"
	"github.com/928799934/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_RevisionHeader struct {
	// GUID
	GuidAttr string
	// Date Time
	DateTimeAttr time.Time
	// Last Sheet Id
	MaxSheetIdAttr uint32
	// User Name
	UserNameAttr string
	IdAttr       string
	// Minimum Revision Id
	MinRIdAttr *uint32
	// Max Revision Id
	MaxRIdAttr *uint32
	// Sheet Id Map
	SheetIdMap *CT_SheetIdMap
	// Reviewed List
	ReviewedList *CT_ReviewedRevisions
	ExtLst       *CT_ExtensionList
}

func NewCT_RevisionHeader() *CT_RevisionHeader {
	ret := &CT_RevisionHeader{}
	ret.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	ret.SheetIdMap = NewCT_SheetIdMap()
	return ret
}

func (m *CT_RevisionHeader) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "guid"},
		Value: fmt.Sprintf("%v", m.GuidAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dateTime"},
		Value: fmt.Sprintf("%v", m.DateTimeAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxSheetId"},
		Value: fmt.Sprintf("%v", m.MaxSheetIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "userName"},
		Value: fmt.Sprintf("%v", m.UserNameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.MinRIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minRId"},
			Value: fmt.Sprintf("%v", *m.MinRIdAttr)})
	}
	if m.MaxRIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxRId"},
			Value: fmt.Sprintf("%v", *m.MaxRIdAttr)})
	}
	e.EncodeToken(start)
	sesheetIdMap := xml.StartElement{Name: xml.Name{Local: "ma:sheetIdMap"}}
	e.EncodeElement(m.SheetIdMap, sesheetIdMap)
	if m.ReviewedList != nil {
		sereviewedList := xml.StartElement{Name: xml.Name{Local: "ma:reviewedList"}}
		e.EncodeElement(m.ReviewedList, sereviewedList)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RevisionHeader) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	m.SheetIdMap = NewCT_SheetIdMap()
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
		if attr.Name.Local == "guid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GuidAttr = parsed
			continue
		}
		if attr.Name.Local == "dateTime" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DateTimeAttr = parsed
			continue
		}
		if attr.Name.Local == "maxSheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.MaxSheetIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "userName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UserNameAttr = parsed
			continue
		}
		if attr.Name.Local == "minRId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MinRIdAttr = &pt
			continue
		}
		if attr.Name.Local == "maxRId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MaxRIdAttr = &pt
			continue
		}
	}
lCT_RevisionHeader:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetIdMap"}:
				if err := d.DecodeElement(m.SheetIdMap, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "reviewedList"}:
				m.ReviewedList = NewCT_ReviewedRevisions()
				if err := d.DecodeElement(m.ReviewedList, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_RevisionHeader %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RevisionHeader
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RevisionHeader and its children
func (m *CT_RevisionHeader) Validate() error {
	return m.ValidateWithPath("CT_RevisionHeader")
}

// ValidateWithPath validates the CT_RevisionHeader and its children, prefixing error messages with path
func (m *CT_RevisionHeader) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.GuidAttr) {
		return fmt.Errorf(`%s/m.GuidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.GuidAttr)
	}
	if err := m.SheetIdMap.ValidateWithPath(path + "/SheetIdMap"); err != nil {
		return err
	}
	if m.ReviewedList != nil {
		if err := m.ReviewedList.ValidateWithPath(path + "/ReviewedList"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
