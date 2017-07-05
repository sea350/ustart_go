package convert

import(
        "github.com/sea350/ustart/server/proto"
        "github.com/sea350/ustart/types"
)

func ToThreadProto(t types.Thread) *proto.Thread {
        var comments []*proto.Comment
        for _, element := range t.Comments {
            comments = append(comments, ToCommentProto(element)) 
        }
        thread := proto.Thread{
                ThreadName: t.ThreadName,
                ThreadID: t.ThreadID,
                PosterID: t.PosterID,
                PosterUsername: t.PosterUsername,
                Content: t.Content,
                Comments: comments,
                Timestamp: t.Timestamp,
        }
        return &thread
}
