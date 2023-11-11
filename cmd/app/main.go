package main

import (
	"cafe/cmd/app/handler"
	boardType3 "cafe/internal/cli/boardType"
	cafeRole2 "cafe/internal/cli/cafeRole"
	member3 "cafe/internal/cli/member"
	memberRole3 "cafe/internal/cli/memberRole"
	"cafe/internal/controller/ban"
	"cafe/internal/controller/boardType"
	cafe2 "cafe/internal/controller/cafe"
	"cafe/internal/controller/cafeRole"
	"cafe/internal/controller/member"
	"cafe/internal/controller/memberRole"
	handler2 "cafe/internal/deco/handler"
	ban3 "cafe/internal/repository/ban"
	cafe3 "cafe/internal/repository/cafe"
	"cafe/internal/repository/infla"
	ban2 "cafe/internal/service/ban"
	boardType2 "cafe/internal/service/boardType"
	"cafe/internal/service/cafe"
	member2 "cafe/internal/service/member"
	memberRole2 "cafe/internal/service/memberRole"
	"cafe/internal/service/role"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// 보드 액션
	//bAH := getBoardActionHandler()
	//r.PathPrefix("/cafes/{cafeId:[0-9]+}/board-actions").Handler(bAH)

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
var roleController = cafeRole.NewController(role.NewService(cafeRole2.NewRequester()))
var memberRoleController = memberRole.NewController(memberRole2.NewService(memberRole3.NewRequester()))

//func getBoardActionHandler() http.Handler {
//	return handler.NewBoardActionHandler(boardAction.NewController(boardAction2.NewService(boardAction3.NewRequester())))
//}

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
