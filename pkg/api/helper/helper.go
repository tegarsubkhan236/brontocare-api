package helper

import (
	"net/http"
)

func UniqueInt(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

//func ImageUploadHelper(input interface{}) (string, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	//create cloudinary instance
//	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
//	if err != nil {
//		return "", err
//	}
//
//	//Upload File
//	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: os.Getenv("CLOUDINARY_UPLOAD_FOLDER")})
//	if err != nil {
//		return "", err
//	}
//
//	return uploadParam.SecureURL, nil
//}

func SuccessResponse(message any) (int, map[string]any) {
	response := map[string]any{
		"data":   message,
		"status": "success",
	}
	return http.StatusOK, response
}

func BadResponse(message any) (int, map[string]any) {
	response := map[string]any{
		"data":   message,
		"status": "failed",
	}
	return http.StatusBadRequest, response
}

func InternalErrorResponse(message any) (int, map[string]any) {
	response := map[string]any{
		"data":   message,
		"status": "failed",
	}
	return http.StatusBadRequest, response
}
