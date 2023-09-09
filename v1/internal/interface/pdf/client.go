package pdf

type Client interface {
	//Fundamental
	NewFile(filename string)
	CloseFile()
	NewPage()

	//HeaderFooter
	Header(left, center, right string)
	Footer(left, center, right string, footerParams *FooterParams)

	//Text
	Title(str string, params *TextParams) //30px
	SubTitle(str string, params *TextParams)
	Heading1(str string, params *TextParams)
	Heading2(str string, params *TextParams)
	Text(str string, params *TextParams)

	//More
	Image(filepath string, w, h float64, params *ImageParams)
	Table(data Table)

	//Miscellaneous
	UnicodeTranslator(str string) string // simplified from fpdf with codePageString
}
