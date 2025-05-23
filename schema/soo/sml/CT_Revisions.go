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

	"github.com/928799934/gooxml"
)

type CT_Revisions struct {
	// Revision Row Column Insert Delete
	Rrc []*CT_RevisionRowColumn
	// Revision Cell Move
	Rm []*CT_RevisionMove
	// Revision Custom View
	Rcv []*CT_RevisionCustomView
	// Revision Sheet Name
	Rsnm []*CT_RevisionSheetRename
	// Revision Insert Sheet
	Ris []*CT_RevisionInsertSheet
	// Revision Cell Change
	Rcc []*CT_RevisionCellChange
	// Revision Format
	Rfmt []*CT_RevisionFormatting
	// Revision AutoFormat
	Raf []*CT_RevisionAutoFormatting
	// Revision Defined Name
	Rdn []*CT_RevisionDefinedName
	// Revision Cell Comment
	Rcmt []*CT_RevisionComment
	// Revision Query Table
	Rqt []*CT_RevisionQueryTableField
	// Revision Merge Conflict
	Rcft []*CT_RevisionConflict
}

func NewCT_Revisions() *CT_Revisions {
	ret := &CT_Revisions{}
	return ret
}

func (m *CT_Revisions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Rrc != nil {
		serrc := xml.StartElement{Name: xml.Name{Local: "ma:rrc"}}
		for _, c := range m.Rrc {
			e.EncodeElement(c, serrc)
		}
	}
	if m.Rm != nil {
		serm := xml.StartElement{Name: xml.Name{Local: "ma:rm"}}
		for _, c := range m.Rm {
			e.EncodeElement(c, serm)
		}
	}
	if m.Rcv != nil {
		sercv := xml.StartElement{Name: xml.Name{Local: "ma:rcv"}}
		for _, c := range m.Rcv {
			e.EncodeElement(c, sercv)
		}
	}
	if m.Rsnm != nil {
		sersnm := xml.StartElement{Name: xml.Name{Local: "ma:rsnm"}}
		for _, c := range m.Rsnm {
			e.EncodeElement(c, sersnm)
		}
	}
	if m.Ris != nil {
		seris := xml.StartElement{Name: xml.Name{Local: "ma:ris"}}
		for _, c := range m.Ris {
			e.EncodeElement(c, seris)
		}
	}
	if m.Rcc != nil {
		sercc := xml.StartElement{Name: xml.Name{Local: "ma:rcc"}}
		for _, c := range m.Rcc {
			e.EncodeElement(c, sercc)
		}
	}
	if m.Rfmt != nil {
		serfmt := xml.StartElement{Name: xml.Name{Local: "ma:rfmt"}}
		for _, c := range m.Rfmt {
			e.EncodeElement(c, serfmt)
		}
	}
	if m.Raf != nil {
		seraf := xml.StartElement{Name: xml.Name{Local: "ma:raf"}}
		for _, c := range m.Raf {
			e.EncodeElement(c, seraf)
		}
	}
	if m.Rdn != nil {
		serdn := xml.StartElement{Name: xml.Name{Local: "ma:rdn"}}
		for _, c := range m.Rdn {
			e.EncodeElement(c, serdn)
		}
	}
	if m.Rcmt != nil {
		sercmt := xml.StartElement{Name: xml.Name{Local: "ma:rcmt"}}
		for _, c := range m.Rcmt {
			e.EncodeElement(c, sercmt)
		}
	}
	if m.Rqt != nil {
		serqt := xml.StartElement{Name: xml.Name{Local: "ma:rqt"}}
		for _, c := range m.Rqt {
			e.EncodeElement(c, serqt)
		}
	}
	if m.Rcft != nil {
		sercft := xml.StartElement{Name: xml.Name{Local: "ma:rcft"}}
		for _, c := range m.Rcft {
			e.EncodeElement(c, sercft)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Revisions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Revisions:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rrc"}:
				tmp := NewCT_RevisionRowColumn()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rrc = append(m.Rrc, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rm"}:
				tmp := NewCT_RevisionMove()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rm = append(m.Rm, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rcv"}:
				tmp := NewCT_RevisionCustomView()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcv = append(m.Rcv, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rsnm"}:
				tmp := NewCT_RevisionSheetRename()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rsnm = append(m.Rsnm, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "ris"}:
				tmp := NewCT_RevisionInsertSheet()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ris = append(m.Ris, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rcc"}:
				tmp := NewCT_RevisionCellChange()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcc = append(m.Rcc, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rfmt"}:
				tmp := NewCT_RevisionFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rfmt = append(m.Rfmt, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "raf"}:
				tmp := NewCT_RevisionAutoFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Raf = append(m.Raf, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rdn"}:
				tmp := NewCT_RevisionDefinedName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rdn = append(m.Rdn, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rcmt"}:
				tmp := NewCT_RevisionComment()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcmt = append(m.Rcmt, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rqt"}:
				tmp := NewCT_RevisionQueryTableField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rqt = append(m.Rqt, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rcft"}:
				tmp := NewCT_RevisionConflict()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcft = append(m.Rcft, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_Revisions %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Revisions
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Revisions and its children
func (m *CT_Revisions) Validate() error {
	return m.ValidateWithPath("CT_Revisions")
}

// ValidateWithPath validates the CT_Revisions and its children, prefixing error messages with path
func (m *CT_Revisions) ValidateWithPath(path string) error {
	for i, v := range m.Rrc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rrc[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rm {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rm[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rcv {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rcv[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rsnm {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rsnm[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Ris {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ris[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rcc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rcc[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rfmt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rfmt[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Raf {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Raf[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rdn {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rdn[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rcmt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rcmt[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rqt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rqt[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rcft {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rcft[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
