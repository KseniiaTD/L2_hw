package main

type Image struct {
	Document
}

func newImage(name string, size int) IDocument {
	return &Image{
		Document: Document{
			name:         name,
			size:         size,
			typeDoc:      "img",
			maxSize:      512,
			isChangeable: false,
		},
	}
}
