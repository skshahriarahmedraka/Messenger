<script context="module">
     export async function load ({}){
       let  FriendData={
    "Name": "Portgas D. Ace",
    "ProfileUrl": "ace",
    "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1642489744/samples/1007550_t0uscy.jpg"
}

    let FriendMsg= [
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
        ]

        let FrinedStat={
            
        }

        return {
            props:{
                FriendData:FriendData,
                FriendMsg:FriendMsg
            }
        }
     }
</script>


<script lang="ts">
    	import Dock from '$lib/_Dock/index.svelte';
	import Servers from '$lib/_Servers/index.svelte';
	import NavBar from '$lib/_Navbar/index.svelte';
	import Peoples from '$lib/_Peoples/index.svelte';
	import Messages from '$lib/_Messages/index.svelte';
    import Chat from "$lib/_Chats/index.svelte"
    import WebSocket, { WebSocketServer } from 'ws'
    

	import { showPeopleList } from '$lib/store2';
    let showPeopleListValue: number;
	// let ChatOrDockValue: number;
    export let FriendData:any 
    export let FriendMsg:any 
	showPeopleList.subscribe((val) => {
		showPeopleListValue = val;
	});
    let x = [
    {   "ProfileURL":"majibarrahman",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Md Majibar Rahman",
        "LastMessage": "hope you are filling well",
        "LastMessageTime":"23 Aug 21",
        "SilentNotification": true,
        "NumberOfNotification": 43,
        "ActiveStatus": true,
        "LastActiveTime": "10:28 AM"
        
    },
    {   "ProfileURL":"rafihasan",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png" ,
        "UserName": "md rafi hasan zihad",
        "LastMessage": "whats is going on bro",
        "LastMessageTime":"9:30 AM",
        "SilentNotification": false,
        "NumberOfNotification": 3,
        "ActiveStatus": true,
        "LastActiveTime": "6 Jul 17"
        
    },
    {   "ProfileURL":"ekbalhasan",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Md ekbal Hasan",
        "LastMessage": "I was going home that day",
        "LastMessageTime":"9 jun 11",
        "SilentNotification": true,
        "NumberOfNotification": 4,
        "ActiveStatus": false,
        "LastActiveTime": "29 Jun 18"
        
    },
    {   "ProfileURL":"mdronok",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Md ronok shahriar",
        "LastMessage": "ethical hacking is my passion",
        "LastMessageTime":"10:34 PM",
        "SilentNotification": false,
        "NumberOfNotification": 1,
        "ActiveStatus": false,
        "LastActiveTime": "10:28 AM"
        
    },
    {   "ProfileURL":"samiulshakil",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "md samiul alim shakil",
        "LastMessage": "I am feeling wanderful . wona go for a tour",
        "LastMessageTime":"12 mar 19",
        "SilentNotification": true,
        "NumberOfNotification": 7,
        "ActiveStatus": true,
        "LastActiveTime": "7 Aug 19"
        
    },
    {   "ProfileURL":"hemelakash",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Hemel sarkar akash",
        "LastMessage": "Got a job in Apple as a CEO",
        "LastMessageTime":"12 min",
        "SilentNotification": true,
        "NumberOfNotification": 6,
    "ActiveStatus": false,
        "LastActiveTime": "10:28 AM"
        },
    {   "ProfileURL":"hazem",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Md Azim Mahfuz",
        "LastMessage": "what a great day . got a job on Uniliver",
        "LastMessageTime":"23 sec",
        "SilentNotification": true,
        "NumberOfNotification":7 ,
    "ActiveStatus": true,
        "LastActiveTime": "1:06 AM"
        },

    {   "ProfileURL":"rafel",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Md rafel",
        "LastMessage": "bought a new bike ,lets to on a ride",
        "LastMessageTime":"4:23 AM",
        "SilentNotification": false,
        "NumberOfNotification": 0,
    "ActiveStatus": false,
        "LastActiveTime": "10:28 AM"
        },
    {   "ProfileURL":"rakibulislam",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "md Rakibul islam",
        "LastMessage": "whats's you doing , bro , see you later",
        "LastMessageTime":"23 dec 21",
        "SilentNotification": true,
        "NumberOfNotification": 2 ,
    "ActiveStatus": true,
        "LastActiveTime": "00:28 AM"
        },
    {   "ProfileURL":"istique",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png" ,
        "UserName": "istique mahmud pritom",
        "LastMessage": "i am on a train , going to Dhaka",
        "LastMessageTime":"1 jun 21",
        "SilentNotification": false,
        "NumberOfNotification": 0,
    "ActiveStatus": true,
        "LastActiveTime": "11:28 AM"
        },
    {   "ProfileURL":"mdrakib",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Md Rakib Hasan",
        "LastMessage": "lets do party to night . Real madrid vs PSG . the game will be epic",
        "LastMessageTime":"9 mar",
        "SilentNotification": true,
        "NumberOfNotification": 2,
    "ActiveStatus": true,
        "LastActiveTime": "00:00 AM"
        },
    {   "ProfileURL":"toki",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "md toki",
        "LastMessage": "whats up bro what is happening there",
        "LastMessageTime":"21 jun",
        "SilentNotification": false,
        "NumberOfNotification": 0,
    "ActiveStatus": false,
        "LastActiveTime": "9 Feb 30"
        },
    {   "ProfileURL":"yurichi",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Yurichi",
        "LastMessage": "are you enjoying my story",
        "LastMessageTime":"12 jan 16",
        "SilentNotification": false,
        "NumberOfNotification": 9,
    "ActiveStatus": false,
        "LastActiveTime": "23 Dec 21"
        },
    {   "ProfileURL":"rengokhu",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "kyjiro Rengokhu",
        "LastMessage": "mugen the train was amayzing",
        "LastMessageTime":"19 feb 19",
        "SilentNotification": false,
        "NumberOfNotification": 5,
    "ActiveStatus": true,
        "LastActiveTime": "10:28 PM"
        },
    {   "ProfileURL":"vegeta",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Prince Vegeta",
        "LastMessage": "the sayian pride will never parish",
        "LastMessageTime":"9 sep 20",
        "SilentNotification": false,
        "NumberOfNotification": 0 ,
    "ActiveStatus": false,
        "LastActiveTime": "10:28 AM"
        },

     {   "ProfileURL":"saad",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png" ,
        "UserName": "Hasan mahmud",
        "LastMessage": "what's your physical condition , come to my home ",
        "LastMessageTime":"yestarday",
        "SilentNotification": false,
        "NumberOfNotification": 5,
    "ActiveStatus": true,
        "LastActiveTime": "3:00 AM"
        },
      {   "ProfileURL":"hankok",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Boa Hankok",
        "LastMessage": "where is my luffy",
        "LastMessageTime":"12 jan 00",
        "SilentNotification": false,
        "NumberOfNotification": 99,
    "ActiveStatus": false,
        "LastActiveTime": "12:28 AM"
        },
       {   "ProfileURL":"nejuko",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Nejuko",
        "LastMessage": "hmm",
        "LastMessageTime":"2 jan 34",
        "SilentNotification": true,
        "NumberOfNotification": 9,
    "ActiveStatus": true,
        "LastActiveTime": "Active"
        },
       {   "ProfileURL":"shanks",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "D. Shanka",
        "LastMessage": "after shreding tears you will be a real man",
        "LastMessageTime":"1:00 AM",
        "SilentNotification": false,
        "NumberOfNotification": 45,
    "ActiveStatus": false,
        "LastActiveTime": "10:00 AM"
        },
       {   "ProfileURL":"admiral",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png" ,
        "UserName": "Nevy Admiral",
        "LastMessage": "justice must be absolute",
        "LastMessageTime":"6:39 AM",
        "SilentNotification": false,
        "NumberOfNotification": 0,
    "ActiveStatus": false,
        "LastActiveTime": "9:28 AM"
        },
         {   "ProfileURL":"law",
        "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png",
        "UserName": "Trafalger D. Water Law",
        "LastMessage": "     i must find the meaning of meaning of D.",
        "LastMessageTime":"yesterday",
        "SilentNotification": false,
        "NumberOfNotification": 34,
    "ActiveStatus": true,
        "LastActiveTime": "Active"
        }]
    
    let x3 ={
    "Name": "Portgas D. Ace",
    "ProfileUrl": "ace",
    "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg"
}

let PeopleData={
        "image":FriendData["ProfileImage"],
        "username":"Sk Shahriar Ahmed Raka",
        "lastseen":"9 jun22",
        "mobile":"+8801838810189",
        "bio":"Lazy is such an ugly word, I prefer to selectively participate. I already want to take a nap tomorrow. ",
        "url":"shahriarraka",
        "notification":true ,
        "contents":[
            [213,"images"],
            [41,"videos"],
            [5,"files"],
            [6,"shared links"],
            [41,"audios"],
            [64,"gifs"],
            [23,"recomendedserver"],
            [67 ,"customizeChat"],
            [ 67,"ChangeTheme"],

            [4 ,"Changeimoji"],
            [675,"EditNickname"],
            [63 ,"shareThisContact"],
            [54 ,"DelteContact"],
            [436 ,"DeleteHistory"],
            [75 ,"DeleteChat"],
            [865 ,"BlockUser"]
        ],
        "recomendationlist":[
            []
        ],
        "cutomize":{
            "color":"blue",
            "emoji":"smile",
            "nickname":"ssar",
        },
        

    }
    // export let server;

    // let socket = new WebSocket(server)
    // console.log("🚀 ~ file: index.svelte ~ line 320 ~ socket", socket)
  
  
    // let server:string="ws://localhost:8888/ws/raka/ssar" 
    // let socket =new WebSocket(server)

    // socket.onopen=()=>{
    //     socket.send(JSON.stringify({name:"raka",message:"whats your name??"}))

    // }
    // socket.onmessage=(event)=>{
    //     let data =JSON.parse(event.data)
    //     console.log("🚀 ~ file: index.svelte ~ line 332 ~ data", data)
        
    // }
    // let demoSend:string=""
    // const sendMessage =()=>{
    //     if (true  ){
    //         socket.send(JSON.stringify({name:"sendmessage",message:"i am send message"}))
    //     }

    // }
    
</script>

<!-- <svelte:head>
	<title> Accord Friend </title>
</svelte:head> -->
<style>
    /* your styles go here */
</style>

<Chat FriendList={x} MyPro={x3} />

<div class=" w-full h-screen overflow-hidden bg-stone-600 flex flex-col flex-nowrap ">
    <!-- nav bar -->
    <NavBar  MyPro={FriendData} />

    <!-- message and people -->
    <div class="  w-full h-full  bg-[#2f3136] flex flex-row overflow-hidden">
        <!-- messaging 830px -->
        <Messages FriendMsg={FriendMsg} />
        <!-- people -->
        {#if showPeopleListValue != 0}
            <Peoples PeopleData={PeopleData}/>
        {/if}
    </div>
</div>