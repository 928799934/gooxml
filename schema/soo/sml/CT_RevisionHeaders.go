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

type CT_RevisionHeaders struct {
	// Last Revision GUID
	GuidAttr string
	// Last GUID
	LastGuidAttr *string
	// Shared Workbook
	SharedAttr *bool
	// Disk Revisions
	DiskRevisionsAttr *bool
	// History
	HistoryAttr *bool
	// Track Revisions
	TrackRevisionsAttr *bool
	// Exclusive Mode
	ExclusiveAttr *bool
	// Revision Id
	RevisionIdAttr *uint32
	// Version
	VersionAttr *int32
	// Keep Change History
	KeepChangeHistoryAttr *bool
	// Protected
	ProtectedAttr *bool
	// Preserve History
	PreserveHistoryAttr *uint32
	// Header
	Header []*CT_RevisionHeader
}

func NewCT_RevisionHeaders() *CT_RevisionHeaders {
	ret := &CT_RevisionHeaders{}
	ret.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	return ret
}

func (m *CT_RevisionHeaders) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "guid"},
		Value: fmt.Sprintf("%v", m.GuidAttr)})
	if m.LastGuidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lastGuid"},
			Value: fmt.Sprintf("%v", *m.LastGuidAttr)})
	}
	if m.SharedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "shared"},
			Value: fmt.Sprintf("%d", b2i(*m.SharedAttr))})
	}
	if m.DiskRevisionsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "diskRevisions"},
			Value: fmt.Sprintf("%d", b2i(*m.DiskRevisionsAttr))})
	}
	if m.HistoryAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "history"},
			Value: fmt.Sprintf("%d", b2i(*m.HistoryAttr))})
	}
	if m.TrackRevisionsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "trackRevisions"},
			Value: fmt.Sprintf("%d", b2i(*m.TrackRevisionsAttr))})
	}
	if m.ExclusiveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "exclusive"},
			Value: fmt.Sprintf("%d", b2i(*m.ExclusiveAttr))})
	}
	if m.RevisionIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "revisionId"},
			Value: fmt.Sprintf("%v", *m.RevisionIdAttr)})
	}
	if m.VersionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "version"},
			Value: fmt.Sprintf("%v", *m.VersionAttr)})
	}
	if m.KeepChangeHistoryAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "keepChangeHistory"},
			Value: fmt.Sprintf("%d", b2i(*m.KeepChangeHistoryAttr))})
	}
	if m.ProtectedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "protected"},
			Value: fmt.Sprintf("%d", b2i(*m.ProtectedAttr))})
	}
	if m.PreserveHistoryAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preserveHistory"},
			Value: fmt.Sprintf("%v", *m.PreserveHistoryAttr)})
	}
	e.EncodeToken(start)
	seheader := xml.StartElement{Name: xml.Name{Local: "ma:header"}}
	for _, c := range m.Header {
		e.EncodeElement(c, seheader)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RevisionHeaders) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	for _, attr := range start.Attr {
		if attr.Name.Local == "exclusive" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ExclusiveAttr = &parsed
			continue
		}
		if attr.Name.Local == "lastGuid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LastGuidAttr = &parsed
			continue
		}
		if attr.Name.Local == "shared" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SharedAttr = &parsed
			continue
		}
		if attr.Name.Local == "diskRevisions" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DiskRevisionsAttr = &parsed
			continue
		}
		if attr.Name.Local == "history" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HistoryAttr = &parsed
			continue
		}
		if attr.Name.Local == "trackRevisions" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TrackRevisionsAttr = &parsed
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
		if attr.Name.Local == "revisionId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RevisionIdAttr = &pt
			continue
		}
		if attr.Name.Local == "version" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.VersionAttr = &pt
			continue
		}
		if attr.Name.Local == "keepChangeHistory" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.KeepChangeHistoryAttr = &parsed
			continue
		}
		if attr.Name.Local == "protected" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ProtectedAttr = &parsed
			continue
		}
		if attr.Name.Local == "preserveHistory" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.PreserveHistoryAttr = &pt
			continue
		}
	}
lCT_RevisionHeaders:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "header"}:
				tmp := NewCT_RevisionHeader()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Header = append(m.Header, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_RevisionHeaders %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RevisionHeaders
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RevisionHeaders and its children
func (m *CT_RevisionHeaders) Validate() error {
	return m.ValidateWithPath("CT_RevisionHeaders")
}

// ValidateWithPath validates the CT_RevisionHeaders and its children, prefixing error messages with path
func (m *CT_RevisionHeaders) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.GuidAttr) {
		return fmt.Errorf(`%s/m.GuidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.GuidAttr)
	}
	if m.LastGuidAttr != nil {
		if !sharedTypes.ST_GuidPatternRe.MatchString(*m.LastGuidAttr) {
			return fmt.Errorf(`%s/m.LastGuidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, *m.LastGuidAttr)
		}
	}
	for i, v := range m.Header {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Header[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
