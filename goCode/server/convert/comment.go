package convert

import(
        "github.com/nicolaifsf/TheTandon-Server/server/proto"
        "github.com/nicolaifsf/TheTandon-Server/types"
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
