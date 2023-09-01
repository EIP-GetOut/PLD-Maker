package db

type Client interface {
	Header(left, center, right string)
	Footer(left, center, right string, pageNo bool, firstPageNo bool)
}
