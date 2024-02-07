package service

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/jinzhu/copier"

	v1 "douyin/api/bff"
	"douyin/app/interface/bff/service/internal/biz"
	"douyin/common/ecode"
)

type BFFService struct {
	v1.UnimplementedBFFServer

	accountUsecase  *biz.AccountUsecase
	feedUsecase     *biz.FeedUsecase
	relationUsecase *biz.RelationUsecase
	commentUsecase  *biz.CommentUsecase
	favoriteUsecase *biz.FavoriteUsecase
	publishUsecase  *biz.PublishUsecase

	log *log.Helper
}

func NewBFFService(au *biz.AccountUsecase, fu *biz.FeedUsecase, ru *biz.RelationUsecase, cu *biz.CommentUsecase,
	favu *biz.FavoriteUsecase, pu *biz.PublishUsecase, logger log.Logger,
) *BFFService {
	return &BFFService{
		accountUsecase:  au,
		feedUsecase:     fu,
		relationUsecase: ru,
		commentUsecase:  cu,
		favoriteUsecase: favu,
		publishUsecase:  pu,
		log:             log.NewHelper(logger),
	}
}

func (s *BFFService) UserRegister(ctx context.Context, req *v1.UserRegisterRequest) (*v1.UserRegisterReply, error) {
	res, err := s.accountUsecase.Register(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.UserRegisterReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}
	return &v1.UserRegisterReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
		UserId:     res.GetUserId(),
		Token:      res.GetToken(),
	}, nil
}

func (s *BFFService) UserLogin(ctx context.Context, req *v1.UserLoginRequest) (*v1.UserLoginReply, error) {
	res, err := s.accountUsecase.Login(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.UserLoginReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}
	return &v1.UserLoginReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
		UserId:     res.GetUserId(),
		Token:      res.GetToken(),
	}, nil
}

func (s *BFFService) GetUserInfo(ctx context.Context, req *v1.GetUserInfoRequest) (*v1.GetUserInfoReply, error) {
	userId := ctx.Value("userId").(int64)
	toUserId := req.GetUserId()

	res, err := s.accountUsecase.GetUserInfo(ctx, userId, toUserId)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetUserInfoReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	user := res.GetUser()

	reply := &v1.GetUserInfoReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
		User:       &v1.User{},
	}
	err = copier.Copy(&reply.User, &user)
	return reply, nil
}

func (s *BFFService) GetPublishList(ctx context.Context, req *v1.GetPublishListRequest) (*v1.GetPublishListReply, error) {
	userId := ctx.Value("userId").(int64)
	toUserId := req.GetUserId()
	res, err := s.publishUsecase.GetUserPublishedVideoList(ctx, userId, toUserId)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetPublishListReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	publishList := res.GetVideoList()

	publishListReply := make([]*v1.Video, 0, len(publishList))
	for _, video := range publishList {
		user := video.GetAuthor()
		v := &v1.Video{Author: &v1.User{}}
		err := copier.Copy(&v, &video)
		if err != nil {
			s.log.Errorf("copier.Copy error(%v)", err)
		}
		err = copier.Copy(&v.Author, &user)
		if err != nil {
			s.log.Errorf("copier.Copy error(%v)", err)
		}
		publishListReply = append(publishListReply, v)
	}

	return &v1.GetPublishListReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
		VideoList:  publishListReply,
	}, nil
}

func (s *BFFService) PublishAction(ctx context.Context, req *v1.PublishActionRequest) (*v1.PublishActionReply, error) {
	videoData := req.GetData()
	videoTitle := req.GetTitle()
	userId := ctx.Value("userId").(int64)
	res, err := s.publishUsecase.PublishVideo(ctx, videoData, userId, videoTitle)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.PublishActionReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	return &v1.PublishActionReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
	}, nil
}

func (s *BFFService) Feed(ctx context.Context, req *v1.FeedRequest) (*v1.FeedReply, error) {
	userId := ctx.Value("userId").(int64)
	latestTimeStr := req.GetLatestTime()

	latestTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
	if err != nil {
		return &v1.FeedReply{
			StatusCode: ecode.ParamErr.ErrCode,
			StatusMsg:  ecode.ParamErr.ErrMsg,
		}, nil
	}

	res, err := s.feedUsecase.Feed(ctx, userId, latestTime)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.FeedReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	videoList := res.GetVideoList()

	videoListReply := make([]*v1.Video, 0, len(videoList))
	for _, video := range videoList {
		user := video.GetAuthor()
		v := &v1.Video{Author: &v1.User{}}
		err := copier.Copy(&v, &video)
		if err != nil {
			s.log.Errorf("copier.Copy error(%v)", err)
		}
		err = copier.Copy(&v.Author, &user)
		if err != nil {
			s.log.Errorf("copier.Copy error(%v)", err)
		}
		videoListReply = append(videoListReply, v)
	}
	return &v1.FeedReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
		VideoList:  videoListReply,
	}, nil
}

func (s *BFFService) GetFollowerList(ctx context.Context, req *v1.GetFollowerListRequest) (*v1.GetFollowerListReply, error) {
	userId := ctx.Value("userId").(int64)
	toUserId := req.GetUserId()
	res, err := s.relationUsecase.GetUserFollowerList(ctx, userId, toUserId)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetFollowerListReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	followerList := res.GetUsers()

	followerListReply := make([]*v1.User, 0, len(followerList))
	for _, user := range followerList {
		u := &v1.User{}
		err := copier.Copy(&u, &user)
		if err != nil {
			s.log.Errorf("copier.Copy error(%v)", err)
		}
		followerListReply = append(followerListReply, u)
	}

	return &v1.GetFollowerListReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
		UserList:   followerListReply,
	}, nil
}

func (s *BFFService) GetFollowList(ctx context.Context, req *v1.GetFollowListRequest) (*v1.GetFollowListReply, error) {
	userId := ctx.Value("userId").(int64)
	toUserId := req.GetUserId()
	res, err := s.relationUsecase.GetUserFollowList(ctx, userId, toUserId)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetFollowListReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	followList := res.GetUsers()

	followListReply := make([]*v1.User, 0, len(followList))
	for _, user := range followList {
		u := &v1.User{}
		err := copier.Copy(&u, &user)
		if err != nil {
			s.log.Errorf("copier.Copy error(%v)", err)
		}
		followListReply = append(followListReply, u)
	}

	return &v1.GetFollowListReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
		UserList:   followListReply,
	}, nil
}

func (s *BFFService) RelationAction(ctx context.Context, req *v1.RelationActionRequest) (*v1.RelationActionReply, error) {
	userId := ctx.Value("userId").(int64)
	toUserId := req.GetToUserId()
	action := req.GetActionType()
	res, err := s.relationUsecase.RelationAction(ctx, userId, toUserId, int32(action))
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.RelationActionReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	return &v1.RelationActionReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
	}, nil
}

func (s *BFFService) GetFriendList(ctx context.Context, req *v1.GetFriendListRequest) (*v1.GetFriendListReply, error) {
	userId := ctx.Value("userId").(int64)
	res, err := s.relationUsecase.GetUserFriendList(ctx, userId)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetFriendListReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	friendList := res.GetUsers()

	friendListReply := make([]*v1.User, 0, len(friendList))
	for _, user := range friendList {
		u := &v1.User{}
		err := copier.Copy(&u, &user)
		if err != nil {
			s.log.Errorf("copier.Copy error(%v)", err)
		}
		friendListReply = append(friendListReply, u)
	}

	return &v1.GetFriendListReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
		UserList:   friendListReply,
	}, nil
}

func (s *BFFService) GetMessageList(ctx context.Context, req *v1.GetMessageListRequest) (*v1.GetMessageListReply, error) {
	// todo
	return &v1.GetMessageListReply{}, nil
}

func (s *BFFService) MessageAction(ctx context.Context, req *v1.MessageActionRequest) (*v1.MessageActionReply, error) {
	// todo
	return &v1.MessageActionReply{}, nil
}

func (s *BFFService) GetFavoriteVideoList(ctx context.Context, req *v1.GetFavoriteVideoListRequest) (*v1.GetFavoriteVideoListReply, error) {
	userId := ctx.Value("userId").(int64)
	toUserId := req.GetUserId()
	res, err := s.favoriteUsecase.GetUserFavoriteVideoList(ctx, userId, toUserId)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetFavoriteVideoListReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	favoriteList := res.GetVideoList()

	favoriteListReply := make([]*v1.Video, 0, len(favoriteList))
	for _, video := range favoriteList {
		user := video.GetAuthor()
		v := &v1.Video{Author: &v1.User{}}
		err := copier.Copy(&v, &video)
		if err != nil {
			s.log.Errorf("copier.Copy error(%v)", err)
		}
		err = copier.Copy(&v.Author, &user)
		if err != nil {
			s.log.Errorf("copier.Copy error(%v)", err)
		}
		favoriteListReply = append(favoriteListReply, v)
	}

	return &v1.GetFavoriteVideoListReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
		VideoList:  favoriteListReply,
	}, nil
}

func (s *BFFService) FavoriteAction(ctx context.Context, req *v1.FavoriteActionRequest) (*v1.FavoriteActionReply, error) {
	userId := ctx.Value("userId").(int64)
	videoId := req.GetVideoId()
	actionType := req.GetActionType()
	res, err := s.favoriteUsecase.FavoriteAction(ctx, userId, videoId, int32(actionType))
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.FavoriteActionReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	return &v1.FavoriteActionReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
	}, nil
}

func (s *BFFService) GetCommentList(ctx context.Context, req *v1.CommentListRequest) (*v1.CommentListReply, error) {
	videoId := req.GetVideoId()
	res, err := s.commentUsecase.GetCommentList(ctx, videoId)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CommentListReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	commentList := res.GetCommentList()

	commentListReply := make([]*v1.Comment, 0, len(commentList))
	for _, comment := range commentList {
		user := comment.GetUser()
		c := &v1.Comment{User: &v1.User{}}
		err := copier.Copy(&c, &comment)
		if err != nil {
			s.log.Errorf("copier.Copy error(%v)", err)
		}
		err = copier.Copy(&c.User, &user)
		if err != nil {
			s.log.Errorf("copier.Copy error(%v)", err)
		}
		commentListReply = append(commentListReply, c)
	}

	return &v1.CommentListReply{
		StatusCode:  res.GetStatusCode(),
		StatusMsg:   res.GetStatusMsg(),
		CommentList: commentListReply,
	}, nil
}

func (s *BFFService) CommentAction(ctx context.Context, req *v1.CommentActionRequest) (*v1.CommentActionReply, error) {
	userId := ctx.Value("userId").(int64)
	videoId := req.GetVideoId()
	content := req.GetCommentText()
	action := req.GetActionType()
	commentId := req.GetCommentId()
	res, err := s.commentUsecase.CommentAction(ctx, userId, videoId, commentId, int32(action), content)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CommentActionReply{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	return &v1.CommentActionReply{
		StatusCode: res.GetStatusCode(),
		StatusMsg:  res.GetStatusMsg(),
	}, nil
}
