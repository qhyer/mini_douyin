package data

import (
	"bytes"
	"context"
	"douyin/app/video/publish/common/constants"
	do "douyin/app/video/publish/common/entity"
	"douyin/app/video/publish/common/mapper"
	po "douyin/app/video/publish/common/model"
	"douyin/app/video/publish/job/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type videoRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoRepo(data *Data, logger log.Logger) biz.VideoRepo {
	return &videoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *videoRepo) BatchCreateVideo(ctx context.Context, videos []*do.Video) error {
	// TODO
	panic("implement me")
}

func (r *videoRepo) CreateVideo(ctx context.Context, video *do.Video) error {
	v, err := mapper.VideoToPO(video)
	if err != nil {
		r.log.Errorf("mapper video to po error: %v", err)
		return err
	}
	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(constants.PublishRecordTableName).Create(v).Error; err != nil {
			return err
		}
		pubCnt := po.PublishCount{
			UserID: v.AuthorID,
		}
		if err := tx.Table(constants.PublishCountTableName(v.AuthorID)).FirstOrInit(&pubCnt, pubCnt).UpdateColumn("video_count", gorm.Expr("video_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		r.log.Errorf("create video error: %v", err)
		return err
	}
	return nil
}

func (r *videoRepo) GetVideo(ctx context.Context, objectName string) ([]byte, error) {
	res, err := r.getObject(ctx, constants.VideoBucketName, objectName)
	if err != nil {
		r.log.Errorf("get video error: %v", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res)
	if err != nil {
		r.log.Errorf("read video error: %v", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

func (r *videoRepo) UploadCover(ctx context.Context, data []byte, objectName string) error {
	_, err := r.putObject(ctx, constants.CoverBucketName, objectName, data)
	if err != nil {
		r.log.Errorf("upload cover error: %v", err)
		return err
	}
	return nil
}

func (r *videoRepo) putObject(ctx context.Context, bucketName, objectName string, data []byte) (info minio.UploadInfo, err error) {
	reader := bytes.NewReader(data)
	return r.data.minio.PutObject(ctx, bucketName, objectName, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
}

func (r *videoRepo) getObject(ctx context.Context, bucketName, objectName string) (obj *minio.Object, err error) {
	return r.data.minio.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
}
