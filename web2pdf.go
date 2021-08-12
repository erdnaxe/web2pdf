// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"context"
	_ "embed"
	"flag"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

var (
	printBackground bool
	scale           float64
	paperWidth      float64
	paperHeight     float64
)

//go:embed cleanup.js
var cleanupScript string

func main() {
	// Parse optional command line arguments
	outputFile := *flag.String("output", "out.pdf", "write PDF document to `path`")
	printBackground = *flag.Bool("print-background", false, "print background graphics")
	scale = *flag.Float64("scale", 1, "`scale` of the webpage rendering")
	paperWidth = *flag.Float64("paper-width", 8.5, "paper `width` in inches, use 8.27 inch for A4")
	paperHeight = *flag.Float64("paper-height", 11, "paper `height` in inches, use 11.67 for A4")
	flag.Parse()

	// Parse positional command line arguments
	url := flag.Arg(0)
	if url == "" {
		flag.Usage()
		log.Fatal("Please specify an URL on the command line")
	}

	// Create Chrome DevTool context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Capture PDF document
	var buf []byte
	if err := chromedp.Run(ctx, printToPDFAction(url, &buf)); err != nil {
		log.Fatal(err)
	}

	// Save PDF document
	if err := ioutil.WriteFile(outputFile, buf, 0644); err != nil {
		log.Fatal(err)
	}
}

// printToPDFAction to print a specific page to PDF document in buffer res.
// It navigates to the specified web page, then inject JS and print.
func printToPDFAction(url string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Evaluate(cleanupScript, nil),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(printBackground).WithScale(scale).WithPaperWidth(paperWidth).WithPaperHeight(paperHeight).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
