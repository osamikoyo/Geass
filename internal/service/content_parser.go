package service

type Image struct{
	Src string
	Alt string
}

type Link struct{
	Text string
	Href string
}

type Content struct{
	FullText string
	Images []Image
}

type Technical struct{
	Code uint32
	ContentType string
}

type Metadata struct{
	Lang string
	Robots string
}

type PageInfo struct{
	Url string
	Title string
	MetadataDescription string
	Content Content
	CountKeyWord uint64
	Links []Link
	Technical string
	Metadata Metadata
}