import * as cookie from "cookie"
import * as jwt from "jsonwebtoken"
import dotenv from "dotenv"

dotenv.config()

export async function handle({event,resolve}) {
    const respose = resolve(event)
    
    return respose
}


// export function getSession(event ){
//     const MyCookie=cookie.parse(event.request.headers.get("cookie") || "")
//     console.log("🚀 ~ file: hooks.ts ~ line 21 ~ jwt.verify ~ MyCookie[\"Auth1\"] : ", MyCookie["Auth1"])
//     if (!MyCookie["Auth1"]){
//         return {user:{authenticated:false}}
//     }
//     const value:string|undefined =process.env.VITE_AUTHKEY
//     jwt.verify(MyCookie["Auth1"],value,(err)=>{
//     console.log("🚀 ~ file: hooks.ts ~ line 21 ~ jwt.verify ~ err : ", err)
//         if (err){

//             return {user:{authenticated:false}}
            
//         }

//     })
//     return {
//         user: {
//             authenticated:true 
//         }
//     }
// }