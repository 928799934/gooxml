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

type CT_ProtectedRange struct {
	// Legacy Password
	PasswordAttr *string
	// Sequence of References
	SqrefAttr ST_Sqref
	// Name
	NameAttr string
	// Security Descriptor
	SecurityDescriptorAttr *string
	// Cryptographic Algorithm Name
	AlgorithmNameAttr *string
	// Password Hash Value
	HashValueAttr *string
	// Salt Value for Password Verifier
	SaltValueAttr *string
	// Iterations to Run Hashing Algorithm
	SpinCountAttr *uint32
	// Security Descriptor
	SecurityDescriptor []string
}

func NewCT_ProtectedRange() *CT_ProtectedRange {
	ret := &CT_ProtectedRange{}
	return ret
}

func (m *CT_ProtectedRange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PasswordAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "password"},
			Value: fmt.Sprintf("%v", *m.PasswordAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sqref"},
		Value: fmt.Sprintf("%v", m.SqrefAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.SecurityDescriptorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "securityDescriptor"},
			Value: fmt.Sprintf("%v", *m.SecurityDescriptorAttr)})
	}
	if m.AlgorithmNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "algorithmName"},
			Value: fmt.Sprintf("%v", *m.AlgorithmNameAttr)})
	}
	if m.HashValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hashValue"},
			Value: fmt.Sprintf("%v", *m.HashValueAttr)})
	}
	if m.SaltValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "saltValue"},
			Value: fmt.Sprintf("%v", *m.SaltValueAttr)})
	}
	if m.SpinCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spinCount"},
			Value: fmt.Sprintf("%v", *m.SpinCountAttr)})
	}
	e.EncodeToken(start)
	if m.SecurityDescriptor != nil {
		sesecurityDescriptor := xml.StartElement{Name: xml.Name{Local: "ma:securityDescriptor"}}
		for _, c := range m.SecurityDescriptor {
			e.EncodeElement(c, sesecurityDescriptor)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ProtectedRange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "password" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PasswordAttr = &parsed
			continue
		}
		if attr.Name.Local == "sqref" {
			parsed, err := ParseSliceST_Sqref(attr.Value)
			if err != nil {
				return err
			}
			m.SqrefAttr = parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "securityDescriptor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SecurityDescriptorAttr = &parsed
			continue
		}
		if attr.Name.Local == "algorithmName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgorithmNameAttr = &parsed
			continue
		}
		if attr.Name.Local == "hashValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HashValueAttr = &parsed
			continue
		}
		if attr.Name.Local == "saltValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SaltValueAttr = &parsed
			continue
		}
		if attr.Name.Local == "spinCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SpinCountAttr = &pt
			continue
		}
	}
lCT_ProtectedRange:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "securityDescriptor"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.SecurityDescriptor = append(m.SecurityDescriptor, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_ProtectedRange %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ProtectedRange
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ProtectedRange and its children
func (m *CT_ProtectedRange) Validate() error {
	return m.ValidateWithPath("CT_ProtectedRange")
}

// ValidateWithPath validates the CT_ProtectedRange and its children, prefixing error messages with path
func (m *CT_ProtectedRange) ValidateWithPath(path string) error {
	return nil
}
