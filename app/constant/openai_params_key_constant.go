package constant

const OpenaiParamKeyType = "type"
const OpenaiParamKeyDescription = "description"

var OpenaiParamKeyProperties = struct {
	PropertyName string
	Filepath     string
	Content      string
}{
	PropertyName: "properties",
	Filepath:     "file_path",
	Content:      "content",
}

const OpenaiParamKeyRequired = "required"
