package route

import (
	"app/database"
	// "app/middleware"
	"app/handler"

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
	r.HandleFunc("/", H.ServerSettings).Methods("GET")
	r.HandleFunc("/api/home", H.Home).Methods("GET")// test route

	r.HandleFunc("/user/{UUID}", H.UserProData).Methods("GET")//complete
	r.HandleFunc("/user/frndsuggestion", H.FrndSuggestion).Methods("POST")// complete
	r.HandleFunc("/user/sendfrndreq", H.SendFriendReq).Methods("POST")// not complete
	r.HandleFunc("/user/userfriendreqlist/{UUID}", H.UserFriendReqList).Methods("GET")// complete
	r.HandleFunc("/user/userfriendreqpendinglist/{UUID}", H.UserFriendReqPendingList).Methods("GET")// not complete

	r.HandleFunc("/user/useracceptfrndreq/", H.UserAcceptFrndReq).Methods("POST")// complete : have to check
	r.HandleFunc("/user/frndlist/{UUID}", H.FrndList).Methods("GET")// complete : get your friend list
	r.HandleFunc("/user/chatfrndlist", H.ChatFrndList).Methods("POST")// complete : get your friend list
	r.HandleFunc("/user/frienddata/{UUID}", H.UserFrndData).Methods("GET")// complete : get your friend list
	r.HandleFunc("/user/getconversationid/", H.GetConversationID).Methods("POST")// incomplete : get your friend list
	
	r.HandleFunc("/user/sendmessage/", H.UserSendMessage).Methods("POST")// complete : have to check
	r.HandleFunc("/user/getconversationmsg/", H.UserConversationMsg).Methods("POST")// not complete : have to check
	
	r.HandleFunc("/user/sendmoney/", H.UserSendMoney).Methods("POST")// complete : have to check

	
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
	r.HandleFunc("/user/rechargewallet", H.UserRechargeWallet).Methods("POST")// complete
	
	
	r.HandleFunc("/admin/moneymanagement", H.MoneyManagement).Methods("GET")//Complete
	r.HandleFunc("/admin/moneyreqaccept", H.AdminMoneyReqAccept).Methods("POST")//complete
	
	//////////////
	r.HandleFunc("/api/imgupload", H.ImgLink).Methods("POST")
	r.HandleFunc("/api/frienddata/{UUID}", H.UserFrndData).Methods("GET")// complete : get your friend list
	r.HandleFunc("/api/getconversationid/", H.GetConversationID).Methods("POST")// incomplete : get your friend list
	r.HandleFunc("/api/getconversationmsg/", H.UserConversationMsg).Methods("POST")// not complete : have to check
	
	// r.HandleFunc("/api/imgupload", H.AdminUserList).Methods("GET")//complete

	go handler.Main(&H)

	// r.HandleFunc("/ws/{SendID}/{ReceiverID}", H.Message)
	// r.Handle("/socketio", H.MsgSocketIO())
	// r.HandleFunc("/ws/{SendID}/{ReceiverID}",)
	// r.HandleFunc("/ws/{SendID}/{ReceiverID}", websocket.Handler(H.MsgNetSocket))
	// r.HandleFunc("/ws/raka", H.MsgGobwas).Methods("GET")
	// r.HandleFunc("/", handler.Home)
	// r.HandleFunc("/", handler.Home)
	// r.HandleFunc("/", handler.Home)
	// r.HandleFunc("/", handler.Home)

}
