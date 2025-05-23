// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package format_test

import (
	"testing"

	"github.com/928799934/gooxml/spreadsheet/format"
)

func TestCellFormattingNumber(t *testing.T) {
	td := []struct {
		Inp float64
		Fmt string
		Exp string
	}{
		// General format, gathered by testing with Mac Excel 365
		{1.0, "", "1"},
		{0.0, "General", "0"},
		{1.0, "General", "1"},
		{1.23, "General", "1.23"},
		{123.456789, "", "123.456789"},
		{12341234.12341234, "", "12341234.12"},
		{12341234.125, "", "12341234.13"},
		{123412341234, "", "1.23412E+11"},
		{12345.12341234, "", "12345.12341"},
		{1e11, "", "1.00E+11"},
		{1e308, "", "1.00E+308"},
		{1234e15, "", "1.23E+18"},
		{1e-10, "", "1.00E-10"},
		{99999999.995, "", "100000000"},
		{.123400, "", "0.1234"},
		{.123412341234, "", "0.123412341"},
		{1.000000000001, "", "1"},
		{.9999999, "", "0.9999999"},
		{.99999999, "", "0.99999999"},
		{0.123999999, "", "0.124"},
		{10.19999999, "", "10.2"},
		{.999999999, "", "1"},
		{.9999999999, "", "1"},
		{42996.6996269676, "0.00", "42996.70"},

		{1.23, "0", "1"},
		{1.23, "0\\ ", "1 "},
		{3.14, "0", "3"},
		{1.23, "00", "01"},
		{1234, "00", "1234"},
		{1, "#,##0", "1"},
		{100, "#,##0", "100"},
		{1000, "#,##0", "1,000"},
		{10000, "#,##0", "10,000"},
		{100000, "#,##0", "100,000"},
		{1000000, "#,##0", "1,000,000"},
		{10000000, "#,##0", "10,000,000"},
		{12, "$0", "$12"},
		{12.234, "$0.00", "$12.23"},
		{12.235, "$0.00", "$12.24"},
		{0.999, "$0.00", "$1.00"},
		{0.123, "0.00", "0.12"},
		{0.123, "0.0", "0.1"},
		{0.123, "0.000", "0.123"},
		{0.123, "0.0000", "0.1230"},
		{1000, "#,##0.00", "1,000.00"},

		{100, "2.345", "2100.345"},
		{100, "20.345", "2100.345"},
		{100, "200.345", "2100.345"},
		{100, "2000.345", "2100.345"},
		{100, "20000.345", "20100.345"},
		{18175133, "#,##0", "18,175,133"},

		{1.23, "-", "-"},
		{1, "-.", "-1"},

		{-1.23, "0.00", "-1.23"},
		{-1.23, "$0.00", "-$1.23"},

		{1, "0.00%", "100.00%"},
		{1, "%0", "%100"},
		{.9512, "0.00%", "95.12%"},
		{1.23, "General", "1.23"},
		{1.23, "(#.##)", "(1.23)"},
		{4.23, "(#,###)", "(4)"},

		{123.456, `"foo"0"bar"`, "foo123bar"},

		// negative format
		{1234, "$#,##0_);($#,##0)", "$1,234"},
		{-1234, "$#,##0_);($#,##0)", "($1,234)"},
		{-4, "#,##0_);[Red](#,##0)", "(4)"},

		// fractions
		{1.5, `0/100`, "150/100"},
		{0.5, "0/1000", "500/1000"},
		{1.25, "0 0/100", "1 25/100"},
		{0.5, "0/10", "5/10"},
		{0.25, "0/10", "3/10"},
		{0.25, "?/?", "1/4"},
		{0.1, "?/?", "0/1"},
		{0.2, "?/?", "1/5"},
		{0.3, "?/?", "2/7"},
		{0.4, "?/?", "2/5"},
		{0.5, "?/?", "1/2"},
		{0.6, "?/?", "3/5"},
		{0.7, "?/?", "2/3"},
		{0.8, "?/?", "4/5"},
		{0.9, "?/?", "8/9"},
		{1, "?/?", "1/1"},
		{25.2, "?/?", "126/5"},
		{0.52, "??/??", "13/25"},
		{0.5, "# ?/?", "1/2"},
		{1.5, "# ?/?", "1 1/2"},

		// dates & times
		{42996.6996269676, "d-mmm-yy", "18-Sep-17"},
		{42996.6996269676, "m/d/yy", "9/18/17"},
		{42996.6996269676, "h:mm AM/PM", "4:47 PM"},
		{42996.6996269676, "h:mm:ss AM/PM", "4:47:28 PM"},
		{42996.6996269676, "mm:s.0", "47:27.8"},
		{42996.6996269676, "mm:ss.0", "47:27.8"},
		{42996.6996269676, "mm:ss.00", "47:27.77"},
		{42996.6996269676, "mm:ss.000", "47:27.770"},
		{42996.6996269676, "mm:ss", "47:28"},
		{42996.6996269676, "dddd, mmmm d, yyyy", "Monday, September 18, 2017"},
		{42996.6996269676, "[$-409]h:mm:ss AM/PM", "4:47:28 PM"},
		{42996.6996269676, " mm:ss ", " 47:28 "},
		{42996.6996269676, "mmmmm", "S"}, // first letter of month
		{1.2345, `[h]:mm:ss"s"`, "29:37:41s"},

		// absolute times
		{4, "[h]", "96"},
		{4.26, "[h]:mm:ss", "102:14:24"},
		{1, "[m]:ss", "1440:00"},
		{1, "[s]", "86400"},

		// exponential
		{10, "0.00E+00", "1.00E+01"},
		{.5, "0.00E+00", "5.00E-01"},
		{4e305, "0.00E+00", "4.00E+305"},
		{4e-305, "0.00E+00", "4.00E-305"},
		{4e25, "0.00E+00", "4.00E+25"},
		{4e-25, "0.00E+00", "4.00E-25"},
		{4e305, "0.00E+00000", "4.00E+00305"},
		{1, "##0.0E+0", "1.0E+0"},

		// special
		{0, `0;(0);"ZERO";-`, "ZERO"},
		{1, `0;(0);"ZERO";-`, "1"},
		{-1, `0;(0);"ZERO";-`, "(1)"},
	}
	for _, tc := range td {
		got := format.Number(tc.Inp, tc.Fmt)
		if got != tc.Exp {
			t.Errorf("expected %s, got %s for %g %s", tc.Exp, got, tc.Inp, tc.Fmt)
		}
	}
}

func TestCellFormattingValue(t *testing.T) {
	td := []struct {
		Inp string
		Fmt string
		Exp string
	}{
		{"0.0", "General", "0"},
		{"1.0", "General", "1"},
		{"1.23", "General", "1.23"},
		{"foo", `"bar"@"baz"`, "barfoobaz"},
	}
	for _, tc := range td {
		got := format.Value(tc.Inp, tc.Fmt)
		if got != tc.Exp {
			t.Errorf("expected %s, got '%s' for '%s'/%s", tc.Exp, got, tc.Inp, tc.Fmt)
		}
	}
}
