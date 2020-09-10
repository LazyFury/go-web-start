package upload

// NewEchoUploader EchoUploader
func NewEchoUploader() *Uploader {
	u := &Uploader{
		BaseDir:      "./static/upload",
		UploadMethod: defaultUpload,
		GetFile:      defaultGetFile,
	}
	return u
}
