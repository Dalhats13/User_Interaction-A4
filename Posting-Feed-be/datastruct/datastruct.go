package datastruct

type HelloWorldRequest struct {
	NAME string `json:"name"`
}

type HelloWorldResponse struct {
	MESSAGE string `json:"message"`
}

type guidelinesActive struct {
	GuidelinesName string
	GuidelinesDesc string
	GuidelinesType string
	GuidelinesLink string
}

//feed create
type FeedCreateRequest struct {
	IMAGE_FEED   string `json:"image_feed"`
	CAPTION_FEED string `json:"caption_feed"`
}

type FeedCreateResponse struct {
	STATUS  int    `json:"code"`
	MESSAGE string `json:"message"`
}
