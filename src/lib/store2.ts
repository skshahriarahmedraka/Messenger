import { writable } from "svelte/store";
// friend images
import Moji from "./profilepic/moji2.jpg"
import Rafi from "./profilepic/r4.jpg"
import Ekbal from "./profilepic/e2.jpg"
import Ronok from "./profilepic/r5.jpg"
import Shakil from "./profilepic/s2.jpg"
import Hemel from "./profilepic/h1.jpg"
import Azim from "./profilepic/a1.jpg"
import Rafel from "./profilepic/r2.jpg"
import Rakib from "./profilepic/r.jpg"
import Istique from "./profilepic/ot2.jpg"
import Rakib2 from "./profilepic/r3.jpg"
import Toki from "./profilepic/t.jpg"
import Yurichi from "./profilepic/y.jpeg"
import Rengokhu from "./profilepic/rengokhu.png"
import Vegeta from "./profilepic/s2.png"
import Saad from "./profilepic/s3.jpg"
import Hankok from "./profilepic/h2.jpg"
import Nejuko from "./profilepic/n1.jpg"
import Shanks from "./profilepic/s4.jpg"
import Admiral from "./profilepic/s5.jpg"
import Law from "./profilepic/law.png"

// server images
import Postgres from "./serverpic/postgres.svg"
import Bd from "./serverpic/bd.svg"
import C from "./serverpic/c.svg"
import Chelsea from "./serverpic/chelsea.svg"
import Coffee from "./serverpic/coffee.png"
import Cpp from "./serverpic/cpp.svg"
import Docker from "./serverpic/docker.png"
import Go from "./serverpic/go.svg"
import Kubernetes from "./serverpic/kubernetes.svg"
import Mancity from "./serverpic/mancity.svg"
import Mazda from "./serverpic/mazda.png"
import Mongo from "./serverpic/mongo.svg"
import Pinterest from "./serverpic/pinterest.svg"
import Py from "./serverpic/py.svg"
import Real from "./serverpic/real.svg"
import Redis from "./serverpic/redis.svg"
import Sve from "./serverpic/svelte.svg"
import Tea from "./serverpic/tea2.jpg"
import Tensorflow from "./serverpic/tensorflow.svg"
import Linux from "./serverpic/linux.svg"
import Ubuntu from "./serverpic/ubuntu.svg"
// user profile pic
import Pro from "./profilepic/pro.jpeg"
// chat active





export const count = writable(0)
export const showPeopleList = writable(0)
export const ChatOrDock = writable(0)
export const MessageList = writable([])






export let FriendList = [
    {   "ProfileURL":"majibarrahman",
        "ProfileImage": Moji,
        "UserName": "Md Majibar Rahman",
        "LastMessage": "hope you are filling well",
        "LastMessageTime":"23 Aug 21",
        "SilentNotification": true,
        "NumberOfNotification": 43,
        "ActiveStatus": true,
        "LastActiveTime": "10:28 AM"
        
    },
    {   "ProfileURL":"rafihasan",
        "ProfileImage": Rafi ,
        "UserName": "md rafi hasan zihad",
        "LastMessage": "whats is going on bro",
        "LastMessageTime":"9:30 AM",
        "SilentNotification": false,
        "NumberOfNotification": 3,
        "ActiveStatus": true,
        "LastActiveTime": "6 Jul 17"
        
    },
    {   "ProfileURL":"ekbalhasan",
        "ProfileImage": Ekbal,
        "UserName": "Md ekbal Hasan",
        "LastMessage": "I was going home that day",
        "LastMessageTime":"9 jun 11",
        "SilentNotification": true,
        "NumberOfNotification": 4,
        "ActiveStatus": false,
        "LastActiveTime": "29 Jun 18"
        
    },
    {   "ProfileURL":"mdronok",
        "ProfileImage": Ronok,
        "UserName": "Md ronok shahriar",
        "LastMessage": "ethical hacking is my passion",
        "LastMessageTime":"10:34 PM",
        "SilentNotification": false,
        "NumberOfNotification": 1,
        "ActiveStatus": false,
        "LastActiveTime": "10:28 AM"
        
    },
    {   "ProfileURL":"samiulshakil",
        "ProfileImage": Shakil,
        "UserName": "md samiul alim shakil",
        "LastMessage": "I am feeling wanderful . wona go for a tour",
        "LastMessageTime":"12 mar 19",
        "SilentNotification": true,
        "NumberOfNotification": 7,
        "ActiveStatus": true,
        "LastActiveTime": "7 Aug 19"
        
    },
    {   "ProfileURL":"hemelakash",
        "ProfileImage": Hemel,
        "UserName": "Hemel sarkar akash",
        "LastMessage": "Got a job in Apple as a CEO",
        "LastMessageTime":"12 min",
        "SilentNotification": true,
        "NumberOfNotification": 6,
    "ActiveStatus": false,
        "LastActiveTime": "10:28 AM"
        },
    {   "ProfileURL":"hazem",
        "ProfileImage": Azim,
        "UserName": "Md Azim Mahfuz",
        "LastMessage": "what a great day . got a job on Uniliver",
        "LastMessageTime":"23 sec",
        "SilentNotification": true,
        "NumberOfNotification":7 ,
    "ActiveStatus": true,
        "LastActiveTime": "1:06 AM"
        },

    {   "ProfileURL":"rafel",
        "ProfileImage": Rafel,
        "UserName": "Md rafel",
        "LastMessage": "bought a new bike ,lets to on a ride",
        "LastMessageTime":"4:23 AM",
        "SilentNotification": false,
        "NumberOfNotification": 0,
    "ActiveStatus": false,
        "LastActiveTime": "10:28 AM"
        },
    {   "ProfileURL":"rakibulislam",
        "ProfileImage": Rakib,
        "UserName": "md Rakibul islam",
        "LastMessage": "whats's you doing , bro , see you later",
        "LastMessageTime":"23 dec 21",
        "SilentNotification": true,
        "NumberOfNotification": 2 ,
    "ActiveStatus": true,
        "LastActiveTime": "00:28 AM"
        },
    {   "ProfileURL":"istique",
        "ProfileImage": Istique ,
        "UserName": "istique mahmud pritom",
        "LastMessage": "i am on a train , going to Dhaka",
        "LastMessageTime":"1 jun 21",
        "SilentNotification": false,
        "NumberOfNotification": 0,
    "ActiveStatus": true,
        "LastActiveTime": "11:28 AM"
        },
    {   "ProfileURL":"mdrakib",
        "ProfileImage": Rakib2,
        "UserName": "Md Rakib Hasan",
        "LastMessage": "lets do party to night . Real madrid vs PSG . the game will be epic",
        "LastMessageTime":"9 mar",
        "SilentNotification": true,
        "NumberOfNotification": 2,
    "ActiveStatus": true,
        "LastActiveTime": "00:00 AM"
        },
    {   "ProfileURL":"toki",
        "ProfileImage": Toki,
        "UserName": "md toki",
        "LastMessage": "whats up bro what is happening there",
        "LastMessageTime":"21 jun",
        "SilentNotification": false,
        "NumberOfNotification": 0,
    "ActiveStatus": false,
        "LastActiveTime": "9 Feb 30"
        },
    {   "ProfileURL":"yurichi",
        "ProfileImage": Yurichi,
        "UserName": "Yurichi",
        "LastMessage": "are you enjoying my story",
        "LastMessageTime":"12 jan 16",
        "SilentNotification": false,
        "NumberOfNotification": 9,
    "ActiveStatus": false,
        "LastActiveTime": "23 Dec 21"
        },
    {   "ProfileURL":"rengokhu",
        "ProfileImage": Rengokhu,
        "UserName": "kyjiro Rengokhu",
        "LastMessage": "mugen the train was amayzing",
        "LastMessageTime":"19 feb 19",
        "SilentNotification": false,
        "NumberOfNotification": 5,
    "ActiveStatus": true,
        "LastActiveTime": "10:28 PM"
        },
    {   "ProfileURL":"vegeta",
        "ProfileImage": Vegeta,
        "UserName": "Prince Vegeta",
        "LastMessage": "the sayian pride will never parish",
        "LastMessageTime":"9 sep 20",
        "SilentNotification": false,
        "NumberOfNotification": 0 ,
    "ActiveStatus": false,
        "LastActiveTime": "10:28 AM"
        },

     {   "ProfileURL":"saad",
        "ProfileImage": Saad ,
        "UserName": "Hasan mahmud",
        "LastMessage": "what's your physical condition , come to my home ",
        "LastMessageTime":"yestarday",
        "SilentNotification": false,
        "NumberOfNotification": 5,
    "ActiveStatus": true,
        "LastActiveTime": "3:00 AM"
        },
      {   "ProfileURL":"hankok",
        "ProfileImage": Hankok,
        "UserName": "Boa Hankok",
        "LastMessage": "where is my luffy",
        "LastMessageTime":"12 jan 00",
        "SilentNotification": false,
        "NumberOfNotification": 99,
    "ActiveStatus": false,
        "LastActiveTime": "12:28 AM"
        },
       {   "ProfileURL":"nejuko",
        "ProfileImage": Nejuko,
        "UserName": "Nejuko",
        "LastMessage": "hmm",
        "LastMessageTime":"2 jan 34",
        "SilentNotification": true,
        "NumberOfNotification": 9,
    "ActiveStatus": true,
        "LastActiveTime": "Active"
        },
       {   "ProfileURL":"shanks",
        "ProfileImage": Shanks,
        "UserName": "D. Shanka",
        "LastMessage": "after shreding tears you will be a real man",
        "LastMessageTime":"1:00 AM",
        "SilentNotification": false,
        "NumberOfNotification": 45,
    "ActiveStatus": false,
        "LastActiveTime": "10:00 AM"
        },
       {   "ProfileURL":"admiral",
        "ProfileImage": Admiral ,
        "UserName": "Nevy Admiral",
        "LastMessage": "justice must be absolute",
        "LastMessageTime":"6:39 AM",
        "SilentNotification": false,
        "NumberOfNotification": 0,
    "ActiveStatus": false,
        "LastActiveTime": "9:28 AM"
        },
         {   "ProfileURL":"law",
        "ProfileImage": Law,
        "UserName": "Trafalger D. Water Law",
        "LastMessage": "     i must find the meaning of meaning of D.",
        "LastMessageTime":"yesterday",
        "SilentNotification": false,
        "NumberOfNotification": 34,
    "ActiveStatus": true,
        "LastActiveTime": "Active"
        }]

export let ServerList = [
    {   
        "ServerURL":"postgres",
        "Name": "Postgres",
        "NewMessage": true,
        "NumberOfNewMessage": 3,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": Postgres
        
    },
    {   
        "ServerURL":"mongodb",
        "Name": "MongoDB",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": true,
        "NumberOfNotification": 3,
        "ServerImage": Mongo
        
    },
    {   
        "ServerURL":"docker",
        "Name": "Docker",
        "NewMessage": true,
        "NumberOfNewMessage": 4,
        "Notification": true,
        "NumberOfNotification": 6,
        "ServerImage": Docker
        
    },
    {   
        "ServerURL":"kubernetes",
        "Name": "Kubernetes",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": true,
        "NumberOfNotification": 99,
        "ServerImage": Kubernetes
        
    },
    {   
        "ServerURL":"redis",
        "Name": "Redis",
        "NewMessage": true,
        "NumberOfNewMessage": 7,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": Redis
        
    },
    {   
        "ServerURL":"tensorflow",
        "Name": "Tensorflow",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": Tensorflow
        
    },
    {   
        "ServerURL":"macity",
        "Name": "Mancity",
        "NewMessage": true,
        "NumberOfNewMessage": 2,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": Mancity
        
    },
    {   
        "ServerURL":"cpp",
        "Name": "C++",
        "NewMessage": true,
        "NumberOfNewMessage": 4,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": Cpp
        
    },
    {   
        "ServerURL":"ubuntu",
        "Name": "Ubuntu",
        "NewMessage": true,
        "NumberOfNewMessage": 7,
        "Notification": true,
        "NumberOfNotification": 2,
        "ServerImage": Ubuntu
        
    },
    {   
        "ServerURL":"realmadrid",
        "Name": "Real Madrid",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": Real
        
    },
    {   
        "ServerURL":"bangladesh",
        "Name": "bangladesh",
        "NewMessage": true,
        "NumberOfNewMessage": 6,
        "Notification": true,
        "NumberOfNotification": 4,
        "ServerImage": Bd
        
    },
    {   
        "ServerURL":"c",
        "Name": "C",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": C
        
    },
    {   
        "ServerURL":"coffee",
        "Name": "Coffee",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": true,
        "NumberOfNotification": 45,
        "ServerImage": Coffee
        
    },
    {   
        "ServerURL":"chelsea",
        "Name": "Chelsea",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": Chelsea
        
    },
    {   
        "ServerURL":"go",
        "Name": "Go",
        "NewMessage": true,
        "NumberOfNewMessage": 9,
        "Notification": true,
        "NumberOfNotification": 65,
        "ServerImage": Go
        
    },
    {   
        "ServerURL":"mazda",
        "Name": "Mazda",
        "NewMessage": false,
        "NumberOfNewMessage": 2,
        "Notification": true,
        "NumberOfNotification": 5,
        "ServerImage": Mazda
        
    },
    {   
        "ServerURL":"pinterest",
        "Name": "Pinterest",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": Pinterest
        
    },
    {   
        "ServerURL":"svelte",
        "Name": "Svelte",
        "NewMessage": true,
        "NumberOfNewMessage": 3,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": Sve
        
    },
    {   
        "ServerURL":"tea",
        "Name": "Tea",
        "NewMessage": true,
        "NumberOfNewMessage": 345,
        "Notification": true,
        "NumberOfNotification": 999,
        "ServerImage": Tea
        
    },
    {   
        "ServerURL":"linux",
        "Name": "Linux",
        "NewMessage": true,
        "NumberOfNewMessage": 3,
        "Notification": true,
        "NumberOfNotification": 90,
        "ServerImage": Linux
        
    }
]

export let MyPro = writable({
    "Name": "Portgas D. Ace",
    "ProfileUrl": "ace",
    "ProfileImage": "http://127.0.0.1:9000/photos/p.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=admin%2F20220315%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20220315T083515Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3B%20filename%3D%22p.png%22&X-Amz-Signature=3b52ccf26069d09b2d15cc5b9219aaa6108e2779569430161fff002609914059",
    
})

export let UseFriendList = {
    5265758567567: {
        8846236503729: [
            [0, "ace", "hello brother , what's you doing"],
            [1, "law", "i am  fine , whats about you"],
            [2, "law", "i am  fine , whats about you"],
            [3, "ace", "hello brother , what's you doing"],
            [4, "law", "i am  fine , whats about you"],
            [5, "law", "i am  fine , whats about you"],
            [6, "ace", "hello brother , what's you doing"],
            [7, "law", "i am  fine , whats about you"],
            [8, "law", "i am  fine , whats about you"],
            [9, "ace", "hello brother , what's you doing"],
            [10, "law", "i am  fine , whats about you"],
            [11, "law", "i am  fine , whats about you"],
            [12, "ace", "hello brother , what's you doing"],
            [13, "law", "i am  fine , whats about you"],
            [14, "law", "i am  fine , whats about you"],
        ],
        8846235603729: [
            [0, "ace", "hello brother , what's you doing"],
            [1, "law", "i am  fine , whats about you"],
            [2, "law", "i am  fine , whats about you"],
            [3, "ace", "hello brother , what's you doing"],
            [4, "law", "i am  fine , whats about you"],
            [5, "law", "i am  fine , whats about you"],
            [6, "ace", "hello brother , what's you doing"],
            [7, "law", "i am  fine , whats about you"],
            [8, "law", "i am  fine , whats about you"],
            [9, "ace", "hello brother , what's you doing"],
            [10, "law", "i am  fine , whats about you"],
            [11, "law", "i am  fine , whats about you"],
            [12, "ace", "hello brother , what's you doing"],
            [13, "law", "i am  fine , whats about you"],
            [14, "law", "i am  fine , whats about you"],
        ],
        8846866503729: [
            [0, "ace", "hello brother , what's you doing"],
            [1, "law", "i am  fine , whats about you"],
            [2, "law", "i am  fine , whats about you"],
            [3, "ace", "hello brother , what's you doing"],
            [4, "law", "i am  fine , whats about you"],
            [5, "law", "i am  fine , whats about you"],
            [6, "ace", "hello brother , what's you doing"],
            [7, "law", "i am  fine , whats about you"],
            [8, "law", "i am  fine , whats about you"],
            [9, "ace", "hello brother , what's you doing"],
            [10, "law", "i am  fine , whats about you"],
            [11, "law", "i am  fine , whats about you"],
            [12, "ace", "hello brother , what's you doing"],
            [13, "law", "i am  fine , whats about you"],
            [14, "law", "i am  fine , whats about you"],
        ],
    }
}

