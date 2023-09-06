package pdf

type Client interface {
	//Fundamental
	NewFile(filename string)
	CloseFile()

	//HeaderFooter
	Header(left, center, right string)
	Footer(left, center, right string, pageNo bool, firstPageNo bool)

	//Text
	Title(str string, params *TextParams) //30px
	SubTitle(str string, params *TextParams)
	Heading1(str string, params *TextParams)
	Heading2(str string, params *TextParams)
	Text(str string, params *TextParams)

	//More
	Image(y, w, h float64, filepath string)
	Table(data [][]string, tableParams *TableParams)

	//Miscellaneous
	UnicodeTranslator(str string) string // simplified from fpdf with codePageString
}
