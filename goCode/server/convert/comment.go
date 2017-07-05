package convert

import(
        "github.com/sea350/ustart/server/proto"
        "github.com/sea350/ustart/types"
)

func ToCommentProto(c types.Comment) *proto.Comment {
        comment := proto.Comment {
                CommentID: c.CommentID,
                PosterID: c.PosterID,
                PosterUsername: c.PosterUsername,
                Timestamp: c.Timestamp,
                Content: c.Content,
        }
        return &comment
}
