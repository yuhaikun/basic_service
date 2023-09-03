package mysql

import (
	"basic_service/kitex_gen/basic"
	"log"
	"time"
)

func GetVideoList(p *basic.DouyinFeedRequest) ([]*basic.Video, error) {

	var videos []*basic.Video

	// 在查询语句后使用 Find 方法来获取结果
	if err := DB.Table("videos").Where("create_at < ?", time.Unix(p.LatestTime, 0)).Order("created_at desc").Limit(30).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}

func SaveVideoList(p *basic.DouyinPublishActionRequest, userID int64, videoURL string) error {
	//timeStamp := time.Now()
	video := &basic.Video{
		//Id:            nil,
		//Author:        nil,
		PlayUrl:       videoURL,
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         p.Title,
		//CreatedAt:     &timeStamp,
		//AuthorID:      &userID,
		//AuthorName:    nil,
	}
	err := DB.Table("videos").Create(video).Error
	if err != nil {
		log.Printf("save video failed:%v", err)
		return err
	}
	return nil
}

func GetPublishList(userID int64) ([]*basic.Video, error) {
	var videos []*basic.Video
	err := DB.Table("videos").Where("author_id = ?", userID).Find(&videos).Error
	if err != nil {
		log.Printf("%d get publishList failed : %v", userID, err)
		return nil, err
	}
	return videos, nil
}
