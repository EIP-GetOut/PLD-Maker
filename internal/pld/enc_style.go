package pld

func (cli *Client) SetFont(familyStr string, styleStr string, size float64) {
	cli.Pdf.SetFont(familyStr, styleStr, size)
}
func (cli *Client) SetFillColor(r int, g int, b int) {
	cli.Pdf.SetFillColor(r, g, b)
}

func (cli *Client) SetDrawColor(r int, g int, b int) {
	cli.Pdf.SetDrawColor(r, g, b)
}

func (cli *Client) SetTextColor(r int, g int, b int) {
	cli.Pdf.SetTextColor(r, g, b)
}

// Curssor position
// X width
// Y height

// Get Cursor X
func (cli *Client) GetX() float64 {
	return cli.Pdf.GetX()
}

// Set Cursor X
func (cli *Client) SetX(x float64) {
	cli.Pdf.SetX(x)
}

// Get Cursor Y
func (cli *Client) GetY() float64 {
	return cli.Pdf.GetY()
}

// Set Cursor Y
func (cli *Client) SetY(y float64) {
	cli.Pdf.SetX(y)
}

// Set Cursor XY
func (cli *Client) SetXY(x, y float64) {
	cli.Pdf.SetXY(x, y)
}
