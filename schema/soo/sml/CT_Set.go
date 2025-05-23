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

type CT_Set struct {
	// Number of Tuples
	CountAttr *uint32
	// Maximum Rank Requested
	MaxRankAttr int32
	// MDX Set Definition
	SetDefinitionAttr string
	// Set Sort Order
	SortTypeAttr ST_SortType
	// Query Failed
	QueryFailedAttr *bool
	// Tuples
	Tpls []*CT_Tuples
	// Sort By Tuple
	SortByTuple *CT_Tuples
}

func NewCT_Set() *CT_Set {
	ret := &CT_Set{}
	return ret
}

func (m *CT_Set) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxRank"},
		Value: fmt.Sprintf("%v", m.MaxRankAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "setDefinition"},
		Value: fmt.Sprintf("%v", m.SetDefinitionAttr)})
	if m.SortTypeAttr != ST_SortTypeUnset {
		attr, err := m.SortTypeAttr.MarshalXMLAttr(xml.Name{Local: "sortType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.QueryFailedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "queryFailed"},
			Value: fmt.Sprintf("%d", b2i(*m.QueryFailedAttr))})
	}
	e.EncodeToken(start)
	if m.Tpls != nil {
		setpls := xml.StartElement{Name: xml.Name{Local: "ma:tpls"}}
		for _, c := range m.Tpls {
			e.EncodeElement(c, setpls)
		}
	}
	if m.SortByTuple != nil {
		sesortByTuple := xml.StartElement{Name: xml.Name{Local: "ma:sortByTuple"}}
		e.EncodeElement(m.SortByTuple, sesortByTuple)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Set) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
		if attr.Name.Local == "maxRank" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.MaxRankAttr = int32(parsed)
			continue
		}
		if attr.Name.Local == "setDefinition" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SetDefinitionAttr = parsed
			continue
		}
		if attr.Name.Local == "sortType" {
			m.SortTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "queryFailed" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.QueryFailedAttr = &parsed
			continue
		}
	}
lCT_Set:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tpls"}:
				tmp := NewCT_Tuples()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tpls = append(m.Tpls, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sortByTuple"}:
				m.SortByTuple = NewCT_Tuples()
				if err := d.DecodeElement(m.SortByTuple, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Set %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Set
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Set and its children
func (m *CT_Set) Validate() error {
	return m.ValidateWithPath("CT_Set")
}

// ValidateWithPath validates the CT_Set and its children, prefixing error messages with path
func (m *CT_Set) ValidateWithPath(path string) error {
	if err := m.SortTypeAttr.ValidateWithPath(path + "/SortTypeAttr"); err != nil {
		return err
	}
	for i, v := range m.Tpls {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tpls[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.SortByTuple != nil {
		if err := m.SortByTuple.ValidateWithPath(path + "/SortByTuple"); err != nil {
			return err
		}
	}
	return nil
}
