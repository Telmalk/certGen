package pdf

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
	path2 "path"
	"training.go/gencert/cert"
)

type PdfSaver struct {
	OutputDir string
}

func background(pdf *gofpdf.Fpdf)  {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageW, pageH := pdf.GetPageSize()
	pdf.ImageOptions("img/background.png",
		0, 0,
		pageW, pageH,
		false, opts, 0, "")
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert)  {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	margin := 30.0
	x := 0.0
	imgW := 30.0
	filename := "img/gopher.png"
	pdf.ImageOptions(filename,
		x + margin, 20,
		imgW, 0,
		false, opts, 0, "")
	pageW, _ := pdf.GetPageSize()
	x  = pageW - imgW
	pdf.ImageOptions(filename,
		x - margin, 20,
		imgW, 0,
		false, opts, 0, "")
	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}

func footer(pdf *gofpdf.Fpdf)  {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	margin := 20.0
	imgW := 50.0
	filename := "img/stamp.png"
	pageW, pageH := pdf.GetPageSize()
	x  := pageW - imgW - margin
	y := pageH - imgW - 10.0
	pdf.ImageOptions(filename,
		x - margin, y,
		imgW, 40,
		false, opts, 0, "")
}

func (p *PdfSaver) Save(cert cert.Cert) error  {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	background(pdf)
	header(pdf, &cert)
	pdf.Ln(30)
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")
	pdf.Ln(30)

	footer(pdf)

	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path2.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to '%v'\n", path)
	return nil
}

func New(outputDir string) (*PdfSaver, error)  {
	var p *PdfSaver
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return p, err
	}
	p = &PdfSaver{
		OutputDir: outputDir,
	}
	return p, nil
}