package main

type Excel struct {
	Document
}

func newExcel(name string, size int) IDocument {
	return &Excel{
		Document: Document{
			name:         name,
			size:         size,
			typeDoc:      "xlsx",
			maxSize:      1024,
			isChangeable: true,
		},
	}
}
