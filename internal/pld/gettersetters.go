package pld

func (cli *Client) SetFont(familyStr string, styleStr string, size float64) {
	cli.Pdf.SetFont(familyStr, styleStr, size)
}
func (cli *Client) SetFillColor(r int, g int, b int) {
	cli.Pdf.SetFillColor(r, g, b)
}

func (cli *Client) SetAutoPageBreak(auto bool, margin float64) {
	cli.Pdf.SetAutoPageBreak(auto, margin)
}

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
