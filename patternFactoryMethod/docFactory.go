package main

import (
	"fmt"
	"strings"
)

func getDocument(docType string, size int, name string) (IDocument, error) {
	if strings.EqualFold(docType, "pdf") {
		return newPDF(name, size), nil
	} else if strings.EqualFold(docType, "image") {
		return newImage(name, size), nil
	} else if strings.EqualFold(docType, "excel") {
		return newExcel(name, size), nil
	}

	return nil, fmt.Errorf(" Document type is not found")
}
