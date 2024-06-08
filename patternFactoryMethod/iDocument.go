package main

type IDocument interface {
	getType(typeDoc string, maxSize int, isChangeable bool)
	getDocument(size int, name string)
	sendDocument(sendBy []string)
	printDocument()
}
