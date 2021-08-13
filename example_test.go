// Copyright (c) 2021 Alexandre Iooss
// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"testing"

	"github.com/chromedp/cdproto/page"
)

func TestPrintToPDF(t *testing.T) {
	printToPDFParams := page.PrintToPDFParams{}
	if err := PrintToPDF("https://github.com/erdnaxe/web2pdf", printToPDFParams, "out.pdf"); err != nil {
		t.Error(err)
	}
}
