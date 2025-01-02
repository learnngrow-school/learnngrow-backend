package utils

// todo:
// import (
// 	"io"
// 	"mime/multipart"
// 	"regexp"
//
// 	"github.com/gin-gonic/gin"
// 	"learn-n-grow.dev/m/db/repository"
// 	"learn-n-grow.dev/m/internal"
// )
//
// func sanitizeFilename(filename string) string {
// 	re := regexp.MustCompile(`[^a-zA-Zа-яА-ЯёЁ0-9._-]`)
// 	sanitized := re.ReplaceAllString(filename, "_")
//
// 	if len(sanitized) > 100 {
// 		sanitized = sanitized[:100]
// 	}
//
// 	return sanitized
// }
//
// func Upload(c *gin.Context, file multipart.FileHeader) error {
// 	data, err := file.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer data.Close()
//
// 	fname := sanitizeFilename(file.Filename)
// 	fsize := file.Size
// 	fdata, err := io.ReadAll(data)
// 	if err != nil {
// 		return err
// 	}
//
// 	params := repository.UploadFileParams{
// 		Fname: fname,
// 		Fsize: int32(fsize),
// 		Fdata: fdata,
// 		Slug:  GetSlug(10),
// 	}
//
// 	internal.Server.Repo.UploadFile(c, params)
// 	return nil
// }
