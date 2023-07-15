package pld

// SetFont sets the font used to print character strings.
// It is mandatory to call this method at least once before printing text or the resulting document will not be valid.
//
// The font can be either a standard one or a font added via the AddFont() method or AddFontFromReader() method.
// Standard fonts use the Windows encoding cp1252 (Western Europe).
//
// The method can be called before the first page is created and the font is kept from page to page.
// If you just wish to change the current font size, it is simpler to call SetFontSize().
//
// Note: the font definition file must be accessible. An error is set if the file cannot be read.
//
// familyStr specifies the font family. It can be either a name defined by AddFont(), AddFontFromReader()
// or one of the standard families (case insensitive): "Courier" for fixed-width, "Helvetica" or "Arial"
// for sans serif, "Times" for serif, "Symbol" or "ZapfDingbats" for symbolic.
//
// styleStr can be "B" (bold), "I" (italic), "U" (underscore), "S" (strike-out) or any combination.
// The default value (specified with an empty string) is regular. Bold and italic styles do not apply to Symbol and ZapfDingbats.
//
// size is the font size measured in points. The default value is the current size.
// If no size has been specified since the beginning of the document, the value taken is 12.
func (cli *Client) SetFont(familyStr string, styleStr string, size float64) {
	cli.Pdf.SetFont(familyStr, styleStr, size)
}

// SetFillColor defines the color used for all filling operations (filled rectangles and cell backgrounds).
// It is expressed in RGB components (0 -255).
// The method can be called before the first page is created and the value is retained from page to page.
func (cli *Client) SetFillColor(r int, g int, b int) {
	cli.Pdf.SetFillColor(r, g, b)
}

// SetDrawColor defines the color used for all drawing operations (lines, rectangles and cell borders).
// It is expressed in RGB components (0 - 255). The method can be called before the first page is created.
// The value is retained from page to page.
func (cli *Client) SetDrawColor(r int, g int, b int) {
	cli.Pdf.SetDrawColor(r, g, b)
}

// SetTextColor defines the color used for text. It is expressed in RGB components (0 - 255).
// The method can be called before the first page is created. The value is retained from page to page.
func (cli *Client) SetTextColor(r int, g int, b int) {
	cli.Pdf.SetTextColor(r, g, b)
}

// styleStr can be "B" (bold), "I" (italic), "U" (underscore), "S" (strike-out) or any combination.
// The default value (specified with an empty string) is regular.
// Bold and italic styles do not apply to Symbol and ZapfDingbats.
func (cli *Client) SetFontStyle(styleStr string) {
	cli.Pdf.SetFontStyle(styleStr)
}

// SetAutoPageBreak enables or disables the automatic page breaking mode.
// When enabling, the second parameter is the distance from the bottom of the page that defines the triggering limit.
// By default, the mode is on and the margin is 2 cm.
func (cli *Client) SetAutoPageBreak(auto bool, margin float64) {
	cli.Pdf.SetAutoPageBreak(auto, margin)
}

//GetX returns the abscissa of the current position.

// Note: the value returned will be affected by the current cell margin.
// To account for this, you may need to either add the value returned by GetCellMargin()
// to it or call SetCellMargin(0) to remove the cell margin.
func (cli *Client) GetX() float64 {
	return cli.Pdf.GetX()
}

// SetX defines the abscissa of the current position.
// If the passed value is negative, it is relative to the right of the page.
func (cli *Client) SetX(x float64) {
	cli.Pdf.SetX(x)
}

// GetY returns the ordinate of the current position.
func (cli *Client) GetY() float64 {
	return cli.Pdf.GetY()
}

// SetX defines the abscissa of the current position.
// If the passed value is negative, it is relative to the right of the page.
func (cli *Client) SetY(y float64) {
	cli.Pdf.SetX(y)
}

// SetXY defines the abscissa and ordinate of the current position.
// If the passed values are negative, they are relative respectively to the right and bottom of the page.
func (cli *Client) SetXY(x, y float64) {
	cli.Pdf.SetXY(x, y)
}

func (cli *Client) PageNo() int {
	return cli.Pdf.PageNo()
}
