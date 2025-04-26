package types

type IFSBussiness interface {
	OpenPDF(filePath string) (string, error)
	ListPDFInDirectory(directory string) ([]string, error)
	CopyPDF(src, dst string) error
	CreateDirectory(path string) error
}
