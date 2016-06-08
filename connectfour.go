package connectfour

import (
	"fmt"

	"github.com/moul/bolosseum/bots"
)

var Rows = 6
var Cols = 7

func NewConnectfourBot() *ConnectfourBot {
	return &ConnectfourBot{}
}

type ConnectfourBot struct{}

func (b *ConnectfourBot) Init(message bots.QuestionMessage) *bots.ReplyMessage {
	// FIXME: init ttt here
	return &bots.ReplyMessage{
		Name: "moul-connectfour",
	}
}

func (b *ConnectfourBot) PlayTurn(question bots.QuestionMessage) *bots.ReplyMessage {
	bot := NewConnectFour()

	bot.Board = question.Board.([][]string)

	fmt.Println(bot)

	return &bots.ReplyMessage{
		Play: 3,
	}
}

type ConnectFour struct {
	Board [][]string
}

func NewConnectFour() ConnectFour {
	cf := ConnectFour{
		Board: make([][]string, Rows),
	}
	for i := 0; i < Rows; i++ {
		cf.Board[i] = make([]string, Cols)
	}
	return cf
}
