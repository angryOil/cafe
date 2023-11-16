package main

import (
	"cafe/cmd/app/handler"
	board4 "cafe/cmd/app/handler/board"
	boardAction4 "cafe/cmd/app/handler/boardAction"
	board3 "cafe/internal/cli/board"
	boardAction3 "cafe/internal/cli/boardAction"
	boardType3 "cafe/internal/cli/boardType"
	cafeRole2 "cafe/internal/cli/cafeRole"
	member3 "cafe/internal/cli/member"
	memberRole3 "cafe/internal/cli/memberRole"
	"cafe/internal/controller/ban"
	"cafe/internal/controller/board"
	"cafe/internal/controller/boardAction"
	"cafe/internal/controller/boardType"
	cafe2 "cafe/internal/controller/cafe"
	cafeRole3 "cafe/internal/controller/cafeRole"
	"cafe/internal/controller/member"
	"cafe/internal/controller/memberRole"
	handler2 "cafe/internal/deco/handler"
	ban3 "cafe/internal/repository/ban"
	cafe3 "cafe/internal/repository/cafe"
	"cafe/internal/repository/infla"
	ban2 "cafe/internal/service/ban"
	board2 "cafe/internal/service/board"
	boardAction2 "cafe/internal/service/boardAction"
	boardType2 "cafe/internal/service/boardType"
	"cafe/internal/service/cafe"
	"cafe/internal/service/cafeRole"
	member2 "cafe/internal/service/member"
	memberRole2 "cafe/internal/service/memberRole"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	boardH := getBoardHandler()
	r.PathPrefix("/cafes/{cafeId:[0-9]+}/boards").Handler(boardH)

	//보드 액션
	bAH := getBoardActionHandler()
	r.PathPrefix("/cafes/{cafeId:[0-9]+}/board-actions").Handler(bAH)

	// 멤버 룰
	mrH := getMemberRoleHandler()
	r.PathPrefix("/cafes/{cafeId:[0-9]+}/member-roles").Handler(mrH)

	// 카페 룰
	crH := getRoleHandler()
	r.PathPrefix("/cafes/{cafeId:[0-9]+}/roles").Handler(crH)

	// 보드 타입
	btH := getBoardTypeHandler()
	r.PathPrefix("/cafes/{cafeId:[0-9]+}/board-types").Handler(btH)
	r.PathPrefix("/cafes/board-types").Handler(btH)

	// 벤
	bH := getBandHandler()
	r.PathPrefix("/cafes/{cafeId:[0-9]+}/ban").Handler(bH)
	r.PathPrefix("/cafes/ban").Handler(bH)

	// 멤버
	mh := getMemberHandler()
	r.PathPrefix("/cafes/{cafeId:[0-9]+}/members").Handler(mh)
	r.PathPrefix("/cafes/members").Handler(mh)
	// 카페
	cafeH := getCafeHandler()
	r.PathPrefix("/cafes").Handler(cafeH)

	wrappedRouter := handler2.NewDecoHandler(r, handler2.JwtMiddleware)
	http.ListenAndServe(":8083", wrappedRouter)
}

var boardTypeController = boardType.NewController(boardType2.NewService(boardType3.NewRequester()))
var banController = ban.NewController(ban2.NewService(ban3.NewBanRepository(infla.NewDB())))
var cafeController = cafe2.NewCafeController(cafe.NewService(cafe3.NewRepository(infla.NewDB())))
var memberController = member.NewController(member2.NewService(member3.NewRequester()))
var roleController = cafeRole3.NewController(cafeRole.NewService(cafeRole2.NewRequester()))
var memberRoleController = memberRole.NewController(memberRole2.NewService(memberRole3.NewRequester()))
var boardActionController = boardAction.NewController(boardAction2.NewService(boardAction3.NewRequester()))
var boardController = board.NewController(board2.NewService(board3.NewRequester()))

func getBoardHandler() http.Handler {
	return board4.NewHandler(boardController)
}

func getBoardActionHandler() http.Handler {
	return boardAction4.NewBoardActionHandler(boardActionController)
}

func getMemberRoleHandler() http.Handler {
	return handler.NewMemberRoleHandler(cafeController, memberController, roleController, memberRoleController)
}

func getRoleHandler() http.Handler {
	return handler.NewCafeRoleHandler(cafeController, roleController)
}

func getBoardTypeHandler() http.Handler {
	return handler.NewBoardTypeHandler(boardTypeController, cafeController)
}

func getBandHandler() http.Handler {
	return handler.NewBanHandler(banController, memberController, cafeController)
}
func getMemberHandler() http.Handler {
	return handler.NewMemberHandler(memberController, cafeController)
}

func getCafeHandler() http.Handler {
	return handler.NewCafeHandler(cafeController)
}
