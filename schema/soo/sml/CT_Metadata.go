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

type CT_Metadata struct {
	// Metadata Types Collection
	MetadataTypes *CT_MetadataTypes
	// Metadata String Store
	MetadataStrings *CT_MetadataStrings
	// MDX Metadata Information
	MdxMetadata *CT_MdxMetadata
	// Future Metadata
	FutureMetadata []*CT_FutureMetadata
	// Cell Metadata
	CellMetadata *CT_MetadataBlocks
	// Value Metadata
	ValueMetadata *CT_MetadataBlocks
	// Future Feature Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Metadata() *CT_Metadata {
	ret := &CT_Metadata{}
	return ret
}

func (m *CT_Metadata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.MetadataTypes != nil {
		semetadataTypes := xml.StartElement{Name: xml.Name{Local: "ma:metadataTypes"}}
		e.EncodeElement(m.MetadataTypes, semetadataTypes)
	}
	if m.MetadataStrings != nil {
		semetadataStrings := xml.StartElement{Name: xml.Name{Local: "ma:metadataStrings"}}
		e.EncodeElement(m.MetadataStrings, semetadataStrings)
	}
	if m.MdxMetadata != nil {
		semdxMetadata := xml.StartElement{Name: xml.Name{Local: "ma:mdxMetadata"}}
		e.EncodeElement(m.MdxMetadata, semdxMetadata)
	}
	if m.FutureMetadata != nil {
		sefutureMetadata := xml.StartElement{Name: xml.Name{Local: "ma:futureMetadata"}}
		for _, c := range m.FutureMetadata {
			e.EncodeElement(c, sefutureMetadata)
		}
	}
	if m.CellMetadata != nil {
		secellMetadata := xml.StartElement{Name: xml.Name{Local: "ma:cellMetadata"}}
		e.EncodeElement(m.CellMetadata, secellMetadata)
	}
	if m.ValueMetadata != nil {
		sevalueMetadata := xml.StartElement{Name: xml.Name{Local: "ma:valueMetadata"}}
		e.EncodeElement(m.ValueMetadata, sevalueMetadata)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Metadata) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Metadata:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "metadataTypes"}:
				m.MetadataTypes = NewCT_MetadataTypes()
				if err := d.DecodeElement(m.MetadataTypes, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "metadataStrings"}:
				m.MetadataStrings = NewCT_MetadataStrings()
				if err := d.DecodeElement(m.MetadataStrings, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "mdxMetadata"}:
				m.MdxMetadata = NewCT_MdxMetadata()
				if err := d.DecodeElement(m.MdxMetadata, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "futureMetadata"}:
				tmp := NewCT_FutureMetadata()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FutureMetadata = append(m.FutureMetadata, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cellMetadata"}:
				m.CellMetadata = NewCT_MetadataBlocks()
				if err := d.DecodeElement(m.CellMetadata, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "valueMetadata"}:
				m.ValueMetadata = NewCT_MetadataBlocks()
				if err := d.DecodeElement(m.ValueMetadata, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Metadata %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Metadata
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Metadata and its children
func (m *CT_Metadata) Validate() error {
	return m.ValidateWithPath("CT_Metadata")
}

// ValidateWithPath validates the CT_Metadata and its children, prefixing error messages with path
func (m *CT_Metadata) ValidateWithPath(path string) error {
	if m.MetadataTypes != nil {
		if err := m.MetadataTypes.ValidateWithPath(path + "/MetadataTypes"); err != nil {
			return err
		}
	}
	if m.MetadataStrings != nil {
		if err := m.MetadataStrings.ValidateWithPath(path + "/MetadataStrings"); err != nil {
			return err
		}
	}
	if m.MdxMetadata != nil {
		if err := m.MdxMetadata.ValidateWithPath(path + "/MdxMetadata"); err != nil {
			return err
		}
	}
	for i, v := range m.FutureMetadata {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FutureMetadata[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.CellMetadata != nil {
		if err := m.CellMetadata.ValidateWithPath(path + "/CellMetadata"); err != nil {
			return err
		}
	}
	if m.ValueMetadata != nil {
		if err := m.ValueMetadata.ValidateWithPath(path + "/ValueMetadata"); err != nil {
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
