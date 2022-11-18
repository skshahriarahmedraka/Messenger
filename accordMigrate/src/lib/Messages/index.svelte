<script lang="ts">

    import {ActiveFrndData,ActiveFrndChatShort, showPeopleList, UserProData, UserActive,ActiveConversationData, ActiveConversationID} from "$lib/store2"
    import { fade, blur, fly, slide, scale } from "svelte/transition";

    import Message from "./message.svelte"
    export let FriendMsg:any
    // export let MyPro:any
    let showPeopleListValue: number

    // let ReactionStruct= {
    //     UserID : "" as  string,
    //     ReactionID :"" as number,
    // }
    let  MessageData= {
        ConversationID: "" as string ,
        SenderID : "" as string,
        SenderName : "" as string,
        Message :"" as  string,
        Reactions :[] as number[],
        UserReaction :  [] as  {
            UserID : "",
            ReactionID :"" ,
        },
        Timestamp :"" as string
    }

    showPeopleList.subscribe(val=>{
        showPeopleListValue=val 
    }) 
    let writeFocus:boolean=false
    let messengerValue:string=""

    async  function SendUserMsg (){
        // await  fetch("")
        await fetch(`/api/sendmessage/`, {
            method:"POST",
            mode: 'no-cors',
            body: JSON.stringify(MessageData)

        })
            .then((res) => {
                return res.json();
            })
            .then((d) => {
                console.log("Sent message response : ",d)
                //   console.log("🚀 ~ file: +page.server.ts ~ line 34 ~ constload:PageServerLoad= ~ resdata", resdata)
                //   console.log("🚀 ~ file: +layout.server.ts ~ line 24 ~ constload:PageServerLoad= ~ resdata", resdata)
                //   UserProData.set(resdata);
            });
    }
    function SendMessage(){
        messengerValue.trimStart()
        messengerValue.trimEnd()
        if (messengerValue != ""){
            MessageData.ConversationID=$ActiveFrndChatShort.ConversationID
            MessageData.SenderID=$UserProData.UserID
            MessageData.SenderName=$UserProData.UserName
            MessageData.Message=messengerValue
            let currentDate = new Date().toJSON().slice(0, 10);
            console.log(currentDate)
            MessageData.Timestamp=currentDate
            // $MessageList=[messengerValue,...$MessageList]
            // console.log($MessageList)
            SendUserMsg()
            messengerValue=""
        }
    }
    function keyHandler(e: { key: string; }){
        if (e.key ==="Enter"){
            // SendMessage()
            SendMyMsg()
        }
    }


        let req ={
            ConversationID : $ActiveFrndChatShort.ConversationID
        }
    console.log("Sent message  req : ",req)
    async  function GetConversationMsg (){
        // await  fetch("")
        MessageData.ConversationID=$ActiveFrndChatShort.ConversationID
        await fetch("/api/getconversationmsg/", {
            method:"POST",
            mode: 'no-cors',
            body: JSON.stringify(req)

        })
            .then((res) => {
                return res.json();
            })
            .then((d) => {
                console.log("Sent message response : ",d)
                //   console.log("🚀 ~ file: +page.server.ts ~ line 34 ~ constload:PageServerLoad= ~ resdata", resdata)
                //   console.log("🚀 ~ file: +layout.server.ts ~ line 24 ~ constload:PageServerLoad= ~ resdata", resdata)
                //   UserProData.set(resdata);
            });
    }

        // GetConversationMsg()

    let SendMoney={
        SenderUUID : "" as string,
        ReceiverUUID : "" as string ,
        Amount: 0 as number
    }
    async function SendMyMsg() {
        messengerValue.trimStart()
        messengerValue.trimEnd()
        if (messengerValue != ""){
            MessageData.ConversationID=$ActiveFrndChatShort.ConversationID
            MessageData.SenderID=$UserProData.UserID
            MessageData.SenderName=$UserProData.UserName
            MessageData.Message=messengerValue
            let currentDate = new Date().toJSON().slice(0, 10);
            console.log(currentDate)
            MessageData.Timestamp=currentDate
            // $MessageList=[messengerValue,...$MessageList]
            // console.log($MessageList)
            // SendUserMsg()
        }else {
            return
        }


        let messageInput = ""
        let socket = new WebSocket("ws://127.0.0.1:8889/gin/user/sendmessage")



        socket.onopen = () => {
            socket.send(JSON.stringify(MessageData))
        }

        socket.onmessage = (event) => {

            let data = JSON.parse(event.data)
            console.log("websocket data",data)
            // if (checkForId(data)) return

            // messages.update( msgs => {
            // 	if(Array.isArray(data)) return [...msgs, ...data]
            // 	return [...msgs, data]
            // })
        }
        messengerValue=""

        const sendMessage = () => {
            if(messageInput.length){
                socket.send(JSON.stringify({"UUID":"What a message"}))
            }
            messageInput = ""
        }
        GetAllConversationData()
        // console.log("ActiveConversationID",$ActiveConversationID)
    }

    async  function GetAllConversationData(){
        let messageInput = ""
        let socket = new WebSocket("ws://127.0.0.1:8889/gin/user/getconversationmsg")

        let req ={
            ConversationID : $ActiveConversationID
        }

        socket.onopen = () => {
            socket.send(JSON.stringify(req))
        }

        socket.onmessage = (event) => {

            let data = JSON.parse(event.data)
            console.log("websocket GetAllConversationData : ",data)
            ActiveConversationData.set(data)
            console.log("ActiveConversationData ",$ActiveConversationData)

        }
        messengerValue=""

        const sendMessage = () => {
            if(messageInput.length){
                socket.send(JSON.stringify({"UUID":"What a message"}))
            }
            messageInput = ""
        }

    }



    let Messagex = {
        SenderID: "" as string ,
        SenderName : "" as string,
        Message : "" as string,
        Reactions : [] as number[],
        UserReaction : [] as {
            UserID : ""  ,
            ReactionID : 0
        }[] ,
        Timestamp : "" as string,
    }
    let MsgArr : {SenderName: string, Message: string, UserReaction: {UserID: "", ReactionID: 0}[], Reactions: number[], SenderID: string, Timestamp: string}[]

    Messagex.Message="whats up"
    Messagex.SenderName="Sk shahriar"
    Messagex.Timestamp="12 july"
    Messagex.SenderID="skssar"
    // MsgArr.push(Messagex)

    // var intervalId = window.setInterval(function(){
    //     // call your function here
    //     // GetAllConversationData()
    // }, 5000);
</script>

<style>
    /* your styles go here */
</style>


<div class=" {showPeopleListValue!=0 ? "w-5/6": "w-full"} h-full overflow-hidden   bg-[#36393f] flex flex-col">
    <!-- all messaging -->
    <div class=" flex-grow bg-[#36393f]  z-10 flex flex-col-reverse overflow-y-scroll scrol3 ">
        <!--{#each [...FriendMsg[1]].reverse() as item, index (item.id) }-->
        <!--        <Message message={item} />-->
        <!--{/each}-->
    {#if $ActiveConversationData.length !=0 }
        {#each $ActiveConversationData.reverse() as message }
        <Message message={message} />
            {/each}
        {:else }
        <p class =""> NO Message found </p>
        {/if}

    <!--{#each MsgArr as item }-->
<!--    {/each}-->

    </div>
    <!-- write message -->
    <div class="  w-full  py-2 px-1 flex flex-row h-14 place-items-center  ">
        <!-- all -->
        <button class=" hover:bg-gray-800 hover:rounded-xl">
            <svg class=" w-7 h-7 m-2  fill-gray-400" version="1.1" id="Capa_1" x="0px" y="0px" viewBox="0 0 512 512" style="enable-background:new 0 0 512 512;" xml:space="preserve" width="512" height="512"><g>	<path d="M85.333,0h64c47.128,0,85.333,38.205,85.333,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   C38.205,234.667,0,196.462,0,149.333v-64C0,38.205,38.205,0,85.333,0z"/>	<path d="M362.667,0h64C473.795,0,512,38.205,512,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   c-47.128,0-85.333-38.205-85.333-85.333v-64C277.333,38.205,315.538,0,362.667,0z"/>	<path d="M85.333,277.333h64c47.128,0,85.333,38.205,85.333,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   C38.205,512,0,473.795,0,426.667v-64C0,315.538,38.205,277.333,85.333,277.333z"/><path d="M362.667,277.333h64c47.128,0,85.333,38.205,85.333,85.333v64C512,473.795,473.795,512,426.667,512h-64   c-47.128,0-85.333-38.205-85.333-85.333v-64C277.333,315.538,315.538,277.333,362.667,277.333z"/></g></svg>
        </button>
        <!-- camera -->
        <button class=" hover:bg-gray-800 hover:rounded-xl">
            <svg class="w-7 h-7 m-2  fill-gray-400" xmlns="http://www.w3.org/2000/svg" id="Filled" viewBox="0 0 24 24" width="512" height="512"><path d="M17.721,3,16.308,1.168A3.023,3.023,0,0,0,13.932,0H10.068A3.023,3.023,0,0,0,7.692,1.168L6.279,3Z"/><circle cx="12" cy="14" r="4"/><path d="M19,5H5a5.006,5.006,0,0,0-5,5v9a5.006,5.006,0,0,0,5,5H19a5.006,5.006,0,0,0,5-5V10A5.006,5.006,0,0,0,19,5ZM12,20a6,6,0,1,1,6-6A6.006,6.006,0,0,1,12,20Z"/></svg>
        </button>
        <!-- images -->
        <button class=" hover:bg-gray-800 hover:rounded-xl">
            <svg class="w-6 h-6 m-2  fill-gray-400" xmlns="http://www.w3.org/2000/svg" id="Filled" viewBox="0 0 24 24" width="512" height="512"><path d="M11.122,12.536a3,3,0,0,0-4.244,0l-6.84,6.84A4.991,4.991,0,0,0,5,24H19a4.969,4.969,0,0,0,2.753-.833Z"/><circle cx="18" cy="6" r="2"/><path d="M19,0H5A5.006,5.006,0,0,0,0,5V16.586l5.464-5.464a5,5,0,0,1,7.072,0L23.167,21.753A4.969,4.969,0,0,0,24,19V5A5.006,5.006,0,0,0,19,0ZM18,10a4,4,0,1,1,4-4A4,4,0,0,1,18,10Z"/></svg>
        </button>
        <!-- voice -->
        <button class=" hover:bg-gray-800 hover:rounded-xl ">
            <svg class="w-7 h-7 m-2  fill-gray-400" xmlns="http://www.w3.org/2000/svg" id="Layer_1" data-name="Layer 1" viewBox="0 0 24 24" width="512" height="512"><path d="M22,13a9.01,9.01,0,0,1-9,9H11a9.011,9.011,0,0,1-9-9H0A11.013,11.013,0,0,0,11,24h2A11.013,11.013,0,0,0,24,13Z"/><path d="M9,13H5.071a7,7,0,0,0,13.858,0H15V11h4V8H15V6h3.929A7,7,0,0,0,5.071,6H9V8H5v3H9Z"/></svg>
        </button>
            <input  autocomplete="off" bind:value="{messengerValue}" on:keypress="{keyHandler}" type="text" name="string" placeholder="Type Message" class="  min-w-fit grow   self-center h-10 max-w-full p-2 text-lg font-medium text-[#afb0b4] outline-none focus:border-sky-500  bg-[#303338] border-2 ml-2 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl  " on:focus="{()=>{writeFocus=true}}" on:blur="{()=>{writeFocus = false}}">
            <button  class="w-10 h-10 ml-2  hover:bg-gray-800 hover:rounded-xl ">
                <svg class="w-10 h-10" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM7 9a1 1 0 100-2 1 1 0 000 2zm7-1a1 1 0 11-2 0 1 1 0 012 0zm-.464 5.535a1 1 0 10-1.415-1.414 3 3 0 01-4.242 0 1 1 0 00-1.415 1.414 5 5 0 007.072 0z" clip-rule="evenodd"></path></svg>
            </button>
            {#if writeFocus || messengerValue.length>0  } 
            <button in:scale out:scale  on:click="{()=>{
                SendMyMsg()
                        GetAllConversationData()

            }}" type="submit" class=" w-10  h-10   ">
                <svg class=" fill-cyan-600" style="enable-background:new 0 0 24 24;" version="1.1" viewBox="0 0 24 24" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><g id="info"/><g id="icons"><path d="M21.5,11.1l-17.9-9C2.7,1.7,1.7,2.5,2.1,3.4l2.5,6.7L16,12L4.6,13.9l-2.5,6.7c-0.3,0.9,0.6,1.7,1.5,1.2l17.9-9   C22.2,12.5,22.2,11.5,21.5,11.1z" id="send"/></g></svg>
            </button>
            {/if}
        </div>

</div>