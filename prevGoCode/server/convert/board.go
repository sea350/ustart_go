package convert

import(
        "github.com/sea350/ustart/server/proto"
        "github.com/sea350/ustart/types"
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
