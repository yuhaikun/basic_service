package main

import (
	basic "basic_service/kitex_gen/basic"
	"context"
)

// FoundationImpl implements the last service interface defined in the IDL.
type FoundationImpl struct {
}

// Feed implements the FoundationImpl interface.
func (s *FoundationImpl) Feed(ctx context.Context, req *basic.DouyinFeedRequest) (resp *basic.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the FoundationImpl interface.
func (s *FoundationImpl) PublishAction(ctx context.Context, req *basic.DouyinPublishActionRequest) (resp *basic.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPublishList implements the FoundationImpl interface.
func (s *FoundationImpl) GetPublishList(ctx context.Context, req *basic.DouyinPublishActionRequest) (resp *basic.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the FoundationImpl interface.
func (s *FoundationImpl) GetUserInfo(ctx context.Context, req *basic.DouyinUserRequest) (resp *basic.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the FoundationImpl interface.
func (s *FoundationImpl) UserLogin(ctx context.Context, req *basic.DouyinUserLoginRequest) (resp *basic.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRegister implements the FoundationImpl interface.
func (s *FoundationImpl) UserRegister(ctx context.Context, req *basic.DouyinUserRegisterRequest) (resp *basic.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...

	return
}
