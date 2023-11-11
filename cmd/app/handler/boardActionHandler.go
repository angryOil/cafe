package handler

//
//import (
//	"github.com/gorilla/mux"
//	"net/http"
//)
//
//// 2가지 역활을 합니다
//// 1. 보드 타입에 따른 crud 권한 저장
//// 2. 권한 비교결과 알려주는 기능 (단 작성자는 무조건 read,delete 가능)
//
//type BoardActionHandler struct {
//	c boardactionCon
//}
//
//func NewBoardActionHandler(c boardAction.Controller) http.Handler {
//	r := mux.NewRouter()
//	h := BoardActionHandler{c: c}
//
//	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-actions/{typeId:[0-9]+}", h.getList).Methods(http.MethodGet)
//	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-actions/{typeId:[0-9]+}", h.create).Methods(http.MethodPost)
//	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-actions/{typeId:[0-9]+}/{id:[0-9]+}", h.patch).Methods(http.MethodPatch)
//	r.HandleFunc("/cafes/{cafeId:[0-9]+}/board-actions/{typeId:[0-9]+}/{id:[0-9]+}", h.delete).Methods(http.MethodDelete)
//	return r
//}
//func (h BoardActionHandler) getList(w http.ResponseWriter, r *http.Request) {
//	//vars := mux.Vars(r)
//	//cafeId, err := strconv.Atoi(vars["cafeId"])
//	//if err != nil {
//	//	http.Error(w, "invalid cafe id", http.StatusBadRequest)
//	//	return
//	//}
//	//typeId, err := strconv.Atoi(vars["typeId"])
//}
//
//func (h BoardActionHandler) create(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func (h BoardActionHandler) patch(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func (h BoardActionHandler) delete(w http.ResponseWriter, r *http.Request) {
//
//}
