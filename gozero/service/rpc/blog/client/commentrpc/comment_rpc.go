// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package commentrpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/commentrpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BatchResp                  = commentrpc.BatchResp
	CommentDetails             = commentrpc.CommentDetails
	CommentNewReq              = commentrpc.CommentNewReq
	CountResp                  = commentrpc.CountResp
	EmptyReq                   = commentrpc.EmptyReq
	EmptyResp                  = commentrpc.EmptyResp
	FindCommentListReq         = commentrpc.FindCommentListReq
	FindCommentListResp        = commentrpc.FindCommentListResp
	FindCommentReplyListReq    = commentrpc.FindCommentReplyListReq
	FindCommentReplyListResp   = commentrpc.FindCommentReplyListResp
	FindLikeCommentResp        = commentrpc.FindLikeCommentResp
	FindTopicCommentCountsReq  = commentrpc.FindTopicCommentCountsReq
	FindTopicCommentCountsResp = commentrpc.FindTopicCommentCountsResp
	IdReq                      = commentrpc.IdReq
	IdsReq                     = commentrpc.IdsReq
	UpdateCommentContentReq    = commentrpc.UpdateCommentContentReq
	UpdateCommentReviewReq     = commentrpc.UpdateCommentReviewReq
	UserIdReq                  = commentrpc.UserIdReq

	CommentRpc interface {
		// 创建评论
		AddComment(ctx context.Context, in *CommentNewReq, opts ...grpc.CallOption) (*CommentDetails, error)
		// 删除评论
		DeleteComment(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询评论
		GetComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*CommentDetails, error)
		// 查询评论列表
		FindCommentList(ctx context.Context, in *FindCommentListReq, opts ...grpc.CallOption) (*FindCommentListResp, error)
		// 查询评论回复列表
		FindCommentReplyList(ctx context.Context, in *FindCommentReplyListReq, opts ...grpc.CallOption) (*FindCommentReplyListResp, error)
		// 查询评论回复数量
		FindTopicCommentCounts(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*FindTopicCommentCountsResp, error)
		// 更新评论审核状态
		UpdateCommentReview(ctx context.Context, in *UpdateCommentReviewReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 更新评论
		UpdateCommentContent(ctx context.Context, in *UpdateCommentContentReq, opts ...grpc.CallOption) (*CommentDetails, error)
		// 点赞评论
		LikeComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*EmptyResp, error)
		// 用户点赞的评论
		FindUserLikeComment(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindLikeCommentResp, error)
	}

	defaultCommentRpc struct {
		cli zrpc.Client
	}
)

func NewCommentRpc(cli zrpc.Client) CommentRpc {
	return &defaultCommentRpc{
		cli: cli,
	}
}

// 创建评论
func (m *defaultCommentRpc) AddComment(ctx context.Context, in *CommentNewReq, opts ...grpc.CallOption) (*CommentDetails, error) {
	client := commentrpc.NewCommentRpcClient(m.cli.Conn())
	return client.AddComment(ctx, in, opts...)
}

// 删除评论
func (m *defaultCommentRpc) DeleteComment(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := commentrpc.NewCommentRpcClient(m.cli.Conn())
	return client.DeleteComment(ctx, in, opts...)
}

// 查询评论
func (m *defaultCommentRpc) GetComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*CommentDetails, error) {
	client := commentrpc.NewCommentRpcClient(m.cli.Conn())
	return client.GetComment(ctx, in, opts...)
}

// 查询评论列表
func (m *defaultCommentRpc) FindCommentList(ctx context.Context, in *FindCommentListReq, opts ...grpc.CallOption) (*FindCommentListResp, error) {
	client := commentrpc.NewCommentRpcClient(m.cli.Conn())
	return client.FindCommentList(ctx, in, opts...)
}

// 查询评论回复列表
func (m *defaultCommentRpc) FindCommentReplyList(ctx context.Context, in *FindCommentReplyListReq, opts ...grpc.CallOption) (*FindCommentReplyListResp, error) {
	client := commentrpc.NewCommentRpcClient(m.cli.Conn())
	return client.FindCommentReplyList(ctx, in, opts...)
}

// 查询评论回复数量
func (m *defaultCommentRpc) FindTopicCommentCounts(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*FindTopicCommentCountsResp, error) {
	client := commentrpc.NewCommentRpcClient(m.cli.Conn())
	return client.FindTopicCommentCounts(ctx, in, opts...)
}

// 更新评论审核状态
func (m *defaultCommentRpc) UpdateCommentReview(ctx context.Context, in *UpdateCommentReviewReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := commentrpc.NewCommentRpcClient(m.cli.Conn())
	return client.UpdateCommentReview(ctx, in, opts...)
}

// 更新评论
func (m *defaultCommentRpc) UpdateCommentContent(ctx context.Context, in *UpdateCommentContentReq, opts ...grpc.CallOption) (*CommentDetails, error) {
	client := commentrpc.NewCommentRpcClient(m.cli.Conn())
	return client.UpdateCommentContent(ctx, in, opts...)
}

// 点赞评论
func (m *defaultCommentRpc) LikeComment(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := commentrpc.NewCommentRpcClient(m.cli.Conn())
	return client.LikeComment(ctx, in, opts...)
}

// 用户点赞的评论
func (m *defaultCommentRpc) FindUserLikeComment(ctx context.Context, in *UserIdReq, opts ...grpc.CallOption) (*FindLikeCommentResp, error) {
	client := commentrpc.NewCommentRpcClient(m.cli.Conn())
	return client.FindUserLikeComment(ctx, in, opts...)
}