package route

import (
	"app/database"
	// "app/middleware"
	// "app/middleware"

	"github.com/gorilla/mux"
)

func Router(r *mux.Router){
	// DB1:=database.PostgresInit()
	DB2:= database.MongodbConnection()
	H:= database.DatabaseInitialization(DB2)

	// r.Use(middleware.ReqLog)
	r.HandleFunc("/login", H.Login).Methods("POST")
	r.HandleFunc("/sveltekit/login", H.SveltekitLogin).Methods("POST")
	r.HandleFunc("/register", H.Register).Methods("POST")
	r.HandleFunc("/sveltekit/register", H.SveltekitRegister).Methods("POST")
	
	// r.Use(middleware.AuthMiddleware)
	r.HandleFunc("/", H.Home).Methods("GET")

	r.HandleFunc("/{UID}", H.UserProfile).Methods("GET")
	r.HandleFunc("/{UID}/profile", H.UserPublicProfile).Methods("GET")
	r.HandleFunc("/{UID}/profile/edit", H.UserProfileEdit).Methods("POST")

	r.HandleFunc("/s/{SID}", H.ServerProfile).Methods("GET")
	r.HandleFunc("/s/{SID}/settings", H.ServerSettings).Methods("GET")
	r.HandleFunc("/s/{SID}/settings/edit", H.ServerSettingsEdit).Methods("POST")
	r.HandleFunc("/s/{SID}/{CID}", H.ServerChannel).Methods("GET")
	r.HandleFunc("/s/{SID}/{CID}/settings", H.ServerChannelSettings	).Methods("GET")
	r.HandleFunc("/s/{SID}/{CID}/settings/edit", H.ServerChannelSettingsEdit).Methods("POST")

	r.HandleFunc("/api/getimglink",H.ImgLink).Methods("POST")


	r.HandleFunc("/ws/{SendID}/{ReceiverID}",H.Message)
	// r.HandleFunc("/", handler.Home)
	// r.HandleFunc("/", handler.Home)
	// r.HandleFunc("/", handler.Home)
	// r.HandleFunc("/", handler.Home)
	
}
