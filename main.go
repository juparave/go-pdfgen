package main

import (
	"bytes"
	"fmt"
	"genpdf/model"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	date := time.Date(2023, time.March, 16, 0, 0, 0, 0, time.UTC)
	articles := []model.Article{
		{Name: "Item 1", Price: 10.0},
		{Name: "Item 2", Price: 25.0},
	}

	invoice := model.Invoice{
		ID:       "INV-123",
		Customer: model.Customer{Name: "John Doe"},
		Date:     &date,
		Articles: articles,
		Total:    35.0,
		Tax:      3.5,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Printf("Error parsing template: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, invoice)
		if err != nil {
			fmt.Printf("Error executing template: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		// if r.Method != http.MethodPost {
		// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		// 	return
		// }

		// Render the template to a string
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, invoice)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Generate the PDF
		pdfBytes, err := generatePDF(tpl.String())
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Serve the PDF
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=invoice.pdf")
		w.Write(pdfBytes)
	})

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func generatePDF(htmlContent string) ([]byte, error) {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	// pdfg.Grayscale.Set(true)

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeLetter)
	pdfg.MarginBottom.Set(10)
	pdfg.MarginTop.Set(10)
	pdfg.MarginLeft.Set(10)
	pdfg.MarginRight.Set(10)

	page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(htmlContent)))

	// Set options for this page
	page.FooterRight.Set("[page] of [topage]")
	page.FooterFontSize.Set(10)
	// page.Zoom.Set(0.95)

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	return pdfg.Bytes(), nil
}
