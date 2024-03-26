package qwen

type AudioContent struct {
	Audio string `json:"audio,omitempty"`
	Text  string `json:"text,omitempty"`
}

func (ac AudioContent) GetBlob() string {
	return ac.Audio
}

// AudioContentList is used for multi-modal generation.
type AudioContentList []AudioContent

var _ IQwenContentMethods = &AudioContentList{}

func NewAudioContentList() *AudioContentList {
	ac := AudioContentList(make([]AudioContent, 0))
	return &ac
}

func (acList *AudioContentList) ToBytes() []byte {
	if acList == nil || len(*acList) == 0 {
		return []byte("")
	}
	return []byte((*acList)[0].Text)
}

func (acList *AudioContentList) ToString() string {
	if acList == nil || len(*acList) == 0 {
		return ""
	}
	return (*acList)[0].Text
}

func (acList *AudioContentList) SetText(s string) {
	if acList == nil {
		panic("AudioContentList is nil")
	}
	*acList = append(*acList, AudioContent{Text: s})
}

func (acList *AudioContentList) AppendText(s string) {
	if acList == nil || len(*acList) == 0 {
		panic("AudioContentList is nil or empty")
	}
	(*acList)[0].Text += s
}

func (acList *AudioContentList) SetBlob(url string) {
	if acList == nil {
		panic("AudioContentList is nil or empty")
	}
	*acList = append(*acList, AudioContent{Audio: url})
}

func (acList *AudioContentList) PopAudioContent() (AudioContent, bool) {
	blobContent, hasAudio := popBlobContent(acList)

	if content, ok := blobContent.(AudioContent); ok {
		return content, hasAudio
	}

	return AudioContent{}, false
}

func (acList *AudioContentList) ConvertToBlobList() []IBlobContent {
	if acList == nil {
		panic("VLContentList is nil or empty")
	}

	list := make([]IBlobContent, len(*acList))
	for i, v := range *acList {
		list[i] = v
	}
	return list
}
