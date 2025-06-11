package inputs

import "mime/multipart"

// UploadAvatarRequest ...
type UploadAvatarRequest struct {
	Image *multipart.FileHeader `form:"image"`
}