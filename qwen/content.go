package qwen

// qwen(text-generation) and qwen-vl(multi-modal) have different data format
// so define generic interfaces for them.
type IQwenContent interface {
	*TextContent | *VLContentList | *AudioContentList | *FileContentList
	IQwenContentMethods
}

// TODO: langchaingo 中有使用这个 interface, 稍后看看是否需要重新设计.
type IQwenContentMethods interface {
	ToBytes() []byte
	ToString() string
	SetText(text string)
	AppendText(text string)
	SetBlob(url string)
}

// content with blob url: e.g. image, audio, file...
type IBlobContent interface {
	GetBlob() string
}

var _ IBlobContent = AudioContent{}
var _ IBlobContent = VLContent{}

// type IBlobContentList interface {
// 	[]AudioContent | []VLContent
// }

type IBlobListConvert interface {
	ConvertToBlobList() []IBlobContent
}

func popBlobContent(rawList IBlobListConvert) (IBlobContent, bool) {
	list := rawList.ConvertToBlobList()
	return innerGetBlob(&list)

}

func innerGetBlob(list *[]IBlobContent) (IBlobContent, bool) {
	hasBlob := false
	for i, v := range *list {
		if v.GetBlob() != "" {
			hasBlob = true
			preSlice := (*list)[:i]
			if i == len(*list)-1 {
				*list = preSlice
			} else {
				postSlice := (*list)[i+1:]
				*list = append(*list, preSlice...)
				*list = append(*list, postSlice...)
			}

			return v, hasBlob
		}
	}

	return nil, hasBlob
}
