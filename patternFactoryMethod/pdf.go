package main

type PDF struct {
	Document
}

func newPDF(name string, size int) IDocument {
	return &PDF{
		Document: Document{
			name:         name,
			size:         size,
			typeDoc:      "pdf",
			maxSize:      128,
			isChangeable: false,
		},
	}
}
