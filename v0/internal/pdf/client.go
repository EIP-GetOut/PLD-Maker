package db

type Client interface {
	Header(left, center, right string)
	Footer(left, center, right string, pageNo bool, firstPageNo bool)
	Image(y, w, h float64, filepath string)
	H1()
	H2()
	H3()
	Text(str string)
}
