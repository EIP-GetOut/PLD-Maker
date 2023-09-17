package pdf

type ImageParams struct {
	X        float64
	XPercent bool
}

type Image struct {
	Filepath string
	Width    float64
	Height   float64
	Params   *ImageParams
}
