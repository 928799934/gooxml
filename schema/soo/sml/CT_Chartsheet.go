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

	"github.com/928799934/gooxml"
)

type CT_Chartsheet struct {
	// Chart Sheet Properties
	SheetPr *CT_ChartsheetPr
	// Chart Sheet Views
	SheetViews *CT_ChartsheetViews
	// Chart Sheet Protection
	SheetProtection *CT_ChartsheetProtection
	// Custom Chart Sheet Views
	CustomSheetViews *CT_CustomChartsheetViews
	PageMargins      *CT_PageMargins
	PageSetup        *CT_CsPageSetup
	HeaderFooter     *CT_HeaderFooter
	// Drawing
	Drawing       *CT_Drawing
	LegacyDrawing *CT_LegacyDrawing
	// Legacy Drawing Reference in Header Footer
	LegacyDrawingHF *CT_LegacyDrawing
	// Drawing Reference in Header Footer
	DrawingHF       *CT_DrawingHF
	Picture         *CT_SheetBackgroundPicture
	WebPublishItems *CT_WebPublishItems
	ExtLst          *CT_ExtensionList
}

func NewCT_Chartsheet() *CT_Chartsheet {
	ret := &CT_Chartsheet{}
	ret.SheetViews = NewCT_ChartsheetViews()
	ret.Drawing = NewCT_Drawing()
	return ret
}

func (m *CT_Chartsheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SheetPr != nil {
		sesheetPr := xml.StartElement{Name: xml.Name{Local: "ma:sheetPr"}}
		e.EncodeElement(m.SheetPr, sesheetPr)
	}
	sesheetViews := xml.StartElement{Name: xml.Name{Local: "ma:sheetViews"}}
	e.EncodeElement(m.SheetViews, sesheetViews)
	if m.SheetProtection != nil {
		sesheetProtection := xml.StartElement{Name: xml.Name{Local: "ma:sheetProtection"}}
		e.EncodeElement(m.SheetProtection, sesheetProtection)
	}
	if m.CustomSheetViews != nil {
		secustomSheetViews := xml.StartElement{Name: xml.Name{Local: "ma:customSheetViews"}}
		e.EncodeElement(m.CustomSheetViews, secustomSheetViews)
	}
	if m.PageMargins != nil {
		sepageMargins := xml.StartElement{Name: xml.Name{Local: "ma:pageMargins"}}
		e.EncodeElement(m.PageMargins, sepageMargins)
	}
	if m.PageSetup != nil {
		sepageSetup := xml.StartElement{Name: xml.Name{Local: "ma:pageSetup"}}
		e.EncodeElement(m.PageSetup, sepageSetup)
	}
	if m.HeaderFooter != nil {
		seheaderFooter := xml.StartElement{Name: xml.Name{Local: "ma:headerFooter"}}
		e.EncodeElement(m.HeaderFooter, seheaderFooter)
	}
	sedrawing := xml.StartElement{Name: xml.Name{Local: "ma:drawing"}}
	e.EncodeElement(m.Drawing, sedrawing)
	if m.LegacyDrawing != nil {
		selegacyDrawing := xml.StartElement{Name: xml.Name{Local: "ma:legacyDrawing"}}
		e.EncodeElement(m.LegacyDrawing, selegacyDrawing)
	}
	if m.LegacyDrawingHF != nil {
		selegacyDrawingHF := xml.StartElement{Name: xml.Name{Local: "ma:legacyDrawingHF"}}
		e.EncodeElement(m.LegacyDrawingHF, selegacyDrawingHF)
	}
	if m.DrawingHF != nil {
		sedrawingHF := xml.StartElement{Name: xml.Name{Local: "ma:drawingHF"}}
		e.EncodeElement(m.DrawingHF, sedrawingHF)
	}
	if m.Picture != nil {
		sepicture := xml.StartElement{Name: xml.Name{Local: "ma:picture"}}
		e.EncodeElement(m.Picture, sepicture)
	}
	if m.WebPublishItems != nil {
		sewebPublishItems := xml.StartElement{Name: xml.Name{Local: "ma:webPublishItems"}}
		e.EncodeElement(m.WebPublishItems, sewebPublishItems)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Chartsheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SheetViews = NewCT_ChartsheetViews()
	m.Drawing = NewCT_Drawing()
lCT_Chartsheet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetPr"}:
				m.SheetPr = NewCT_ChartsheetPr()
				if err := d.DecodeElement(m.SheetPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetViews"}:
				if err := d.DecodeElement(m.SheetViews, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetProtection"}:
				m.SheetProtection = NewCT_ChartsheetProtection()
				if err := d.DecodeElement(m.SheetProtection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "customSheetViews"}:
				m.CustomSheetViews = NewCT_CustomChartsheetViews()
				if err := d.DecodeElement(m.CustomSheetViews, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pageMargins"}:
				m.PageMargins = NewCT_PageMargins()
				if err := d.DecodeElement(m.PageMargins, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pageSetup"}:
				m.PageSetup = NewCT_CsPageSetup()
				if err := d.DecodeElement(m.PageSetup, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "headerFooter"}:
				m.HeaderFooter = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.HeaderFooter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "drawing"}:
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "legacyDrawing"}:
				m.LegacyDrawing = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "legacyDrawingHF"}:
				m.LegacyDrawingHF = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawingHF, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "drawingHF"}:
				m.DrawingHF = NewCT_DrawingHF()
				if err := d.DecodeElement(m.DrawingHF, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "picture"}:
				m.Picture = NewCT_SheetBackgroundPicture()
				if err := d.DecodeElement(m.Picture, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "webPublishItems"}:
				m.WebPublishItems = NewCT_WebPublishItems()
				if err := d.DecodeElement(m.WebPublishItems, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Chartsheet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Chartsheet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Chartsheet and its children
func (m *CT_Chartsheet) Validate() error {
	return m.ValidateWithPath("CT_Chartsheet")
}

// ValidateWithPath validates the CT_Chartsheet and its children, prefixing error messages with path
func (m *CT_Chartsheet) ValidateWithPath(path string) error {
	if m.SheetPr != nil {
		if err := m.SheetPr.ValidateWithPath(path + "/SheetPr"); err != nil {
			return err
		}
	}
	if err := m.SheetViews.ValidateWithPath(path + "/SheetViews"); err != nil {
		return err
	}
	if m.SheetProtection != nil {
		if err := m.SheetProtection.ValidateWithPath(path + "/SheetProtection"); err != nil {
			return err
		}
	}
	if m.CustomSheetViews != nil {
		if err := m.CustomSheetViews.ValidateWithPath(path + "/CustomSheetViews"); err != nil {
			return err
		}
	}
	if m.PageMargins != nil {
		if err := m.PageMargins.ValidateWithPath(path + "/PageMargins"); err != nil {
			return err
		}
	}
	if m.PageSetup != nil {
		if err := m.PageSetup.ValidateWithPath(path + "/PageSetup"); err != nil {
			return err
		}
	}
	if m.HeaderFooter != nil {
		if err := m.HeaderFooter.ValidateWithPath(path + "/HeaderFooter"); err != nil {
			return err
		}
	}
	if err := m.Drawing.ValidateWithPath(path + "/Drawing"); err != nil {
		return err
	}
	if m.LegacyDrawing != nil {
		if err := m.LegacyDrawing.ValidateWithPath(path + "/LegacyDrawing"); err != nil {
			return err
		}
	}
	if m.LegacyDrawingHF != nil {
		if err := m.LegacyDrawingHF.ValidateWithPath(path + "/LegacyDrawingHF"); err != nil {
			return err
		}
	}
	if m.DrawingHF != nil {
		if err := m.DrawingHF.ValidateWithPath(path + "/DrawingHF"); err != nil {
			return err
		}
	}
	if m.Picture != nil {
		if err := m.Picture.ValidateWithPath(path + "/Picture"); err != nil {
			return err
		}
	}
	if m.WebPublishItems != nil {
		if err := m.WebPublishItems.ValidateWithPath(path + "/WebPublishItems"); err != nil {
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
