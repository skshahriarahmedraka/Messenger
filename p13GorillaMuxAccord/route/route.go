package route

import (
	"app/database"
	// "app/middleware"
	// "app/middleware"

	"github.com/gorilla/mux"
	// "golang.org/x/net/websocket"
)

func Router(r *mux.Router) {
	// DB1:=database.PostgresInit()
	DB1 := database.MongoUserDBConn()
	DB2 := database.MongoMsgDBConn()
	H := database.DatabaseInitialization(DB1, DB2)

	// r.Use(middleware.ReqLog)
	// r.HandleFunc("/login", H.Login).Methods("POST")
	r.HandleFunc("/sveltekit/login", H.SveltekitLogin).Methods("POST")
	// r.HandleFunc("/register", H.Register).Methods("POST")
	r.HandleFunc("/sveltekit/register", H.SveltekitRegister).Methods("POST")

	// r.Use(middleware.AuthMiddleware)
	r.HandleFunc("/", H.Home).Methods("GET")

	// r.HandleFunc("/{UserID}", H.UserProfile).Methods("GET")
	r.HandleFunc("/user/{UUID}", H.UserProData).Methods("GET")//complete
	r.HandleFunc("/user/{UUID}/frndsuggestion", H.FrndSuggestion).Methods("GET")// incomplete
	
	// r.HandleFunc("/{UID}/profile", H.UserPublicProfile).Methods("GET")
	r.HandleFunc("/user/profile/update", H.UserProfileUpdate).Methods("POST")//complete
	
	
	
	r.HandleFunc("/s/{SID}", H.ServerProfile).Methods("GET")
	r.HandleFunc("/s/{SID}/settings", H.ServerSettings).Methods("GET")
	r.HandleFunc("/s/{SID}/settings/edit", H.ServerSettingsEdit).Methods("POST")
	r.HandleFunc("/s/{SID}/{CID}", H.ServerChannel).Methods("GET")
	r.HandleFunc("/s/{SID}/{CID}/settings", H.ServerChannelSettings).Methods("GET")
	r.HandleFunc("/s/{SID}/{CID}/settings/edit", H.ServerChannelSettingsEdit).Methods("POST")
	
	// Admin
	r.HandleFunc("/api/getimglink", H.ImgLink).Methods("POST")
	
	r.HandleFunc("/user/moneytokenreq", H.UserTokenReq).Methods("POST")// complete
	r.HandleFunc("/user/moneytokenreqlist", H.UserTokenReqList).Methods("POST")//complete
	r.HandleFunc("/user/rechargewallet", H.UserRechargeWallet).Methods("POST")


	r.HandleFunc("/admin/moneymanagement", H.MoneyManagement).Methods("GET")//Complete
	r.HandleFunc("/admin/moneyreqaccept", H.AdminMoneyReqAccept).Methods("POST")//complete


	// r.HandleFunc("/ws/{SendID}/{ReceiverID}", H.Message).Methods("GET")
	r.Handle("/socketio", H.MsgSocketIO())
	// r.HandleFunc("/ws/{SendID}/{ReceiverID}",)
	// r.HandleFunc("/ws/{SendID}/{ReceiverID}", websocket.Handler(H.MsgNetSocket))
	// r.HandleFunc("/ws/{SendID}/{ReceiverID}", H.MsgGobwas)
	// r.HandleFunc("/", handler.Home)
	// r.HandleFunc("/", handler.Home)
	// r.HandleFunc("/", handler.Home)
	// r.HandleFunc("/", handler.Home)

}
