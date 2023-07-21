package nextgen

type NGFileType string

var NGFileTypes = struct {
	File   NGFileType
	Folder NGFileType
}{
	File:   "File",
	Folder: "Folder",
}

var NGFileTypeValues = []string{
	NGFileTypes.File.String(),
	NGFileTypes.Folder.String(),
}

func (e NGFileType) String() string {
	return string(e)
}
