package convert

import(
        "github.com/nicolaifsf/TheTandon-Server/server/proto"
        "github.com/nicolaifsf/TheTandon-Server/types"
)

func ToBoardProto(b types.Board) *proto.Board {
    board := proto.Board {
            BoardID: b.BoardID,
            BoardName: b.BoardName,
            ThreadIDs: b.ThreadIDs,
            DocIDs: b.DocIDs,
    }
    return &board
}
