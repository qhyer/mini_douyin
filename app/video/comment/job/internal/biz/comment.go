package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"douyin/app/video/comment/common/event"
)

type CommentRepo interface {
	UpdateVideoCommentCount(ctx context.Context, videoId int64, incr int64) error
	BatchUpdateVideoCommentCount(ctx context.Context, stats map[int64]int64) error
	UpdateVideoCommentTempCountCache(ctx context.Context, procId int, stat *event.CommentStat) error
	GetVideoCommentTempCountFromCache(ctx context.Context, procId int) (map[int64]int64, error)
	PurgeVideoCommentTempCountCache(ctx context.Context, procId int) error
	CreateComment(ctx context.Context, comment *event.CommentAction) error
	BatchCreateComment(ctx context.Context, comments []*event.CommentAction) error
	DeleteComment(ctx context.Context, comment *event.CommentAction) error
	BatchDeleteComment(ctx context.Context, comments []*event.CommentAction) error
}

type CommentUsecase struct {
	repo CommentRepo
	log  *log.Helper
}

func NewCommentUsecase(repo CommentRepo, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentUsecase) CreateComment(ctx context.Context, comment *event.CommentAction) error {
	err := uc.repo.CreateComment(ctx, comment)
	if err != nil {
		uc.log.Errorf("CreateComment error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) BatchCreateComment(ctx context.Context, commentActs []*event.CommentAction) error {
	err := uc.repo.BatchCreateComment(ctx, commentActs)
	if err != nil {
		uc.log.Errorf("BatchCreateComment error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) DeleteComment(ctx context.Context, comment *event.CommentAction) error {
	err := uc.repo.DeleteComment(ctx, comment)
	if err != nil {
		uc.log.Errorf("DeleteComment error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) BatchDeleteComment(ctx context.Context, commentActs []*event.CommentAction) error {
	err := uc.repo.BatchDeleteComment(ctx, commentActs)
	if err != nil {
		uc.log.Errorf("BatchDeleteComment error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) UpdateVideoCommentCount(ctx context.Context, videoId int64, incr int64) error {
	err := uc.repo.UpdateVideoCommentCount(ctx, videoId, incr)
	if err != nil {
		uc.log.Errorf("UpdateVideoCommentCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) BatchUpdateVideoCommentCount(ctx context.Context, stats map[int64]int64) error {
	err := uc.repo.BatchUpdateVideoCommentCount(ctx, stats)
	if err != nil {
		uc.log.Errorf("BatchUpdateVideoCommentCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) UpdateVideoCommentTempCountCache(ctx context.Context, procId int, stat *event.CommentStat) error {
	err := uc.repo.UpdateVideoCommentTempCountCache(ctx, procId, stat)
	if err != nil {
		uc.log.Errorf("UpdateVideoCommentTempCountCache error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) GetVideoCommentTempCountCache(ctx context.Context, procId int) (map[int64]int64, error) {
	stat, err := uc.repo.GetVideoCommentTempCountFromCache(ctx, procId)
	if err != nil {
		uc.log.Errorf("GetVideoCommentTempCountCache error(%v)", err)
		return nil, err
	}
	return stat, nil
}

func (uc *CommentUsecase) PurgeVideoCommentTempCountCache(ctx context.Context, procId int) error {
	err := uc.repo.PurgeVideoCommentTempCountCache(ctx, procId)
	if err != nil {
		uc.log.Errorf("PurgeVideoCommentTempCountCache error(%v)", err)
		return err
	}
	return nil
}
