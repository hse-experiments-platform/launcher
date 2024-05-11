package domain

import (
	"strconv"
)

func GetBucketName(userID int64) string {
	return "user-" + strconv.FormatInt(userID, 10)
}

func GetObjectName(datasetID int64) string {
	return strconv.FormatInt(datasetID, 10)
}
