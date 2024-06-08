package main

func main() {
	book1 := &Book{name: "Hobbit", author: "J. R. R. Tolkien", publishingHouse: "asti", binding: "hard", pageCnt: 312, price: 80.99}
	jeans1 := &Jeans{model: "H&M skiny", size: 38, color: "blue", length: 32, price: 34.95}
	laptop1 := &Laptop{model: "Lenovo Yoga", cpu: "Intel Celeron N4020", diagonal: 15.6, color: "pink", diskType: "SSD", diskGb: 128, os: "Windows 11 Home", price: 415}

	info := &InfoSbj{}

	book1.accept(info)
	jeans1.accept(info)
	laptop1.accept(info)

	packageSubject := &PackageSbj{}

	book1.accept(packageSubject)
	jeans1.accept(packageSubject)
	laptop1.accept(packageSubject)
}
