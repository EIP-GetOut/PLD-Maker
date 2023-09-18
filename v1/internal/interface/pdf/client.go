package pdf

type Client interface {
	//Fundamental
	NewFile(filename string)
	CloseFile()
	NewPage()
	NewLine()

	//HeaderFooter
	Header(left, center, right string)
	Footer(left, center, right string, footerParams *FooterParams)
	//Text
	Title(text Text) //30px
	SubTitle(text Text)
	Heading1(text Text)
	Heading2(text Text)
	Text(text Text)

	//More
	Image(image Image)
	Table(data Table)

	//Miscellaneous
	UnicodeTranslator(str string) string // simplified from fpdf with codePageString
}
