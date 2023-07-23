package pld

type Report struct {
	Name  string
	Notes string
}

type ProgressReport struct {
	Global   string
	Reports  []Report
	Problems string
	Comments string
}

func (cli *Client) AddProgressReport(progressReport ProgressReport) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	cli.Pdf.Ln(5)

	//Global
	cli.Pdf.SetDrawColor(0, 0, 0)
	cli.Pdf.SetFillColor(164, 194, 244)
	cli.Pdf.SetTextColor(0, 0, 0)
	cli.Pdf.SetFont("Arial", "B", 15)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith, 7, tr("Avancement global pour ce rendez-vous"), "1", "C", true)

	cli.Pdf.SetFont("Arial", "", 10)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith, 7, tr(WrapText(progressReport.Global, 50)), "1", "", false)
	cli.Pdf.Ln(-1)

	//Avancement individuel
	cli.Pdf.SetFillColor(164, 194, 244)
	cli.Pdf.SetFont("Arial", "B", 15)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith, 7, tr("Avancement individuel"), "1", "C", true)

	for _, report := range progressReport.Reports {
		cli.Pdf.SetFillColor(201, 218, 248)
		cli.Pdf.SetFont("Arial", "B", 10)
		cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
		cli.Pdf.MultiCell(cli.CardWith, 7, tr(report.Name), "1", "", true)

		cli.Pdf.SetFont("Arial", "", 10)
		cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
		cli.Pdf.MultiCell(cli.CardWith, 7, tr(WrapText(report.Notes, 50)), "1", "", false)
	}
	cli.Pdf.Ln(-1)

	//Problems
	cli.Pdf.SetFillColor(164, 194, 244)
	cli.Pdf.SetFont("Arial", "B", 15)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith, 7, tr("Points bloquants"), "1", "C", true)

	cli.Pdf.SetFont("Arial", "", 10)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith, 7, tr(WrapText(progressReport.Problems, 50)), "1", "", false)
	cli.Pdf.Ln(-1)

	//Problems
	cli.Pdf.SetFillColor(164, 194, 244)
	cli.Pdf.SetFont("Arial", "B", 15)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith, 7, tr("Commentaire général"), "1", "C", true)

	cli.Pdf.SetFont("Arial", "", 10)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith, 7, tr(WrapText(progressReport.Comments, 50)), "1", "", false)
}
