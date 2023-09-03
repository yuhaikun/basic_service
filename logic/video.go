package logic

import (
	"basic_service/dao/mysql"
	"basic_service/kitex_gen/basic"
)

func GetVideoStream(p *basic.DouyinFeedRequest) ([]*basic.Video, error) {

	return mysql.GetVideoList(p)
}

//func PublishVideo(p *basic.DouyinPublishActionRequest, fd *bytes.Reader, userID int64) error {
//
//	videoURL, _, err := obs_object.PutOBSFile("api-test-file", *p.Title, fd)
//	if err != nil {
//		return err
//	}
//	err = mysql.SaveVideoList(p, userID, videoURL)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func GetPublishList(userID int64) ([]*basic.Video, error) {

	return mysql.GetPublishList(userID)
}
