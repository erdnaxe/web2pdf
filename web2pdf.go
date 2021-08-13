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

//go:embed cleanup.js
var cleanupScript string

func main() {
	// Parse optional command line arguments
	outputFile := *flag.String("output", "out.pdf", "write PDF document to `path`")
	printToPDFParams := page.PrintToPDFParams{}
	flag.BoolVar(&printToPDFParams.Landscape, "landscape", false, "set paper orientation to landscape")
	flag.BoolVar(&printToPDFParams.DisplayHeaderFooter, "display-header-footer", false, "display header and footer")
	flag.BoolVar(&printToPDFParams.PrintBackground, "print-background", false, "print background graphics")
	flag.Float64Var(&printToPDFParams.Scale, "scale", 1, "`scale` of the webpage rendering")
	flag.Float64Var(&printToPDFParams.PaperWidth, "paper-width", 8.5, "paper `width` in inches, use 8.27 inch for A4")
	flag.Float64Var(&printToPDFParams.PaperHeight, "paper-height", 11, "paper `height` in inches, use 11.67 inch for A4")
	flag.Float64Var(&printToPDFParams.MarginTop, "margin-top", 0.4, "top `margin` in inch")
	flag.Float64Var(&printToPDFParams.MarginBottom, "margin-bottom", 0.4, "bottom `margin` in inch")
	flag.Float64Var(&printToPDFParams.MarginLeft, "margin-left", 0.4, "left `margin` in inch")
	flag.Float64Var(&printToPDFParams.MarginRight, "margin-right", 0.4, "right `margin` in inch")
	flag.StringVar(&printToPDFParams.PageRanges, "page-ranges", "", "paper `ranges` to print, e.g., '1-5, 8, 11-13' (defaults to all pages)")
	flag.StringVar(&printToPDFParams.HeaderTemplate, "header-template", "", "HTML template for the print header, use these CSS classes to inject values: date, title, url, pageNumber, totalPages.")
	flag.StringVar(&printToPDFParams.FooterTemplate, "footer-template", "", "HTML template for the print footer, same format as --header-template")
	flag.Parse()

	// Parse positional command line arguments
	url := flag.Arg(0)
	if url == "" {
		flag.Usage()
		log.Fatal("Please specify an URL on the command line")
	}

	// Print
	log.Printf("Printing %s to %s", url, outputFile)
	PrintToPDF(url, printToPDFParams, outputFile)
}

// PrintToPDF prints the specified `url` to `outputFile` PDF document. Printing
// options can be specified within `printToPDFParams`.
func PrintToPDF(
	url string,
	printToPDFParams page.PrintToPDFParams,
	outputFile string,
) {
	// Create Chrome DevTool context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Capture PDF document
	var buf []byte
	if err := chromedp.Run(ctx, printToPDFAction(url, printToPDFParams, &buf)); err != nil {
		log.Fatal(err)
	}

	// Save PDF document
	if err := ioutil.WriteFile(outputFile, buf, 0644); err != nil {
		log.Fatal(err)
	}
}

// printToPDFAction is a Chrome DevTool Protocol action that navigates to the
// specified web page, then inject JS and print it to a PDF document in `res`
// buffer.
func printToPDFAction(
	url string,
	printToPDFParams page.PrintToPDFParams,
	res *[]byte,
) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Evaluate(cleanupScript, nil),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := printToPDFParams.Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
