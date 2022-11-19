<script lang="ts">
	import { goto } from '$app/navigation';

	// import { url } from "$app/stores";

	// import {page} from "$app/stores"
	// $: console.log("$page.params ",$page.params)
	import {ChatOrDock, ActiveFrndChatShort, ActiveConversationID, ActiveConversationData} from '$lib/store2';
	import Notification from '$lib/icons/notification.svelte';
	// import Search from '$lib/Navbar/profileImg/search.svelte';
	// import Cross from '$lib/Navbar/profileImg/Cross.svelte';
	import UserDashPannel from '$lib/UserDashPanel/index.svelte';
	import { fade, blur, fly, slide, scale } from 'svelte/transition';

	import { UserProData,UserActive, ActiveFrndData } from '$lib/store2';
	import Person from '$lib/icons/person.svelte';
	import DiscordNoOne3 from '$lib/icons/discordNoOne3.svelte';
	
	// import {FriendList ,MyPro} from "$lib/store2"
	let FriendList: {
		UUID: string;
		UserID: string;
		UserImg: string;
		UserName: string;
		LastMessage: string;
		LastMessageTime: string;
		SilentNotification: boolean;
		NumberOfNotification: number;
		ActiveStatus: boolean;
		LastActiveTime: string;
		ConversationID:string
	}[] = [] as {
		UUID: string;
		UserID: string;
		UserImg: string;
		UserName: string;
		LastMessage: string;
		LastMessageTime: string;
		SilentNotification: boolean;
		NumberOfNotification: number;
		ActiveStatus: boolean;
		LastActiveTime: string;
		ConversationID:string
	}[];
	// export let MyPro: any;

	let ChatOrDockHelper: number;
	ChatOrDock.subscribe((val) => {
		ChatOrDockHelper = val;
	});

	function ChatOrDockFunc() {
		goto('/', {
			replaceState: true,
			noscroll: true,
			keepfocus: true
		});
		if (ChatOrDockHelper != 0) {
			ChatOrDock.update((n) => (n = 0));
		} else {
			ChatOrDock.update((n) => (n = 1));
		}
	}
	let searchIconEvent = false;
	let inputValue = '';

	let title: string | undefined = 'Accord';

	function OnClickFriend(e: {
		ProfileURL: any;
		ProfileImage?: string;
		UserName?: string;
		LastMessage?: string;
		LastMessageTime?: string;
		SilentNotification?: boolean;
		NumberOfNotification?: number;
		ActiveStatus?: boolean;
		LastActiveTime?: string;
	}) {
		title = e.UserName;
		ActiveChat = e.ProfileURL;
		goto(e.ProfileURL, {
			replaceState: true,
			noscroll: true,
			keepfocus: true
		});
	}
	let showUsername = true;

	//{
	//	ProfileURL: 'majibarrahman',
	//	ProfileImage: Moji,//UserImg
	//	UserName: 'Md Majibar Rahman',
	//	LastMessage: 'hope you are filling well',
	//	LastMessageTime: '23 Aug 21',
	//	SilentNotification: true,
	//	NumberOfNotification: 43,
	//	ActiveStatus: true,
	//	LastActiveTime: '10:28 AM'
	//},

	async function GetUserChatFrndList() {
		await fetch(`/api/profile/chatfrndlist`, {
			method: 'POST',
			mode: 'no-cors',
			body: JSON.stringify({
				UUID: $UserProData.UUID
			})
		})
			.then((res) => {
				return res.json();
			})
			.then((d) => {
				FriendList = d;
				console.log('🚀 ~ file: index.svelte ~ line 109 ~ .then ~ FriendList', FriendList);
			});
	}
	GetUserChatFrndList();

	async function GetFrndData(PathUrl:string){
		await fetch(`/api/frienddata/`,{
			method:"POST",
			mode: 'no-cors',
			body: JSON.stringify({"UserID":PathUrl})
		}).then((res)=>{
			return res.json()
		}).then((data)=>{
			console.log(" post Frnddata ",data)
			ActiveFrndData.set(data)
			console.log("ActiveFrndData",$ActiveFrndData)
			// Frnddata=d
		})
		// return Frnddata
	}
	// let ActiveChat: string = '';

	//         return res.json()

	let ConversationID: string ="" as string
	// async function GetConversationID(){
	// 	let req ={
	// 		ReqUUID : $UserProData.UUID as string ,
	// 		FrndUUID: $ActiveFrndChatShort.UUID as string
	// 	}
	//
	// 	console.log("req :" ,req)
	// 	await fetch(`/api/getconversationid/`,{
	// 		method: "POST",
	// 		body: JSON.stringify(req)
	// 	}).then((res)=>{
	// 		return res.json()
	// 	}).then((d)=>{
	// 		console.log(" post ConversationID  ConversationID  v ConversationID ConversationID  ConversationID  ",d)
	// 		ConversationID=d.ConversationID
	// 		ActiveConversationID.set(ConversationID)
	// 	})
	//
	// }
	//  function GetConversationID(){
	// 	 let getUserProData
	// 	 let getfrndProData
	// 	 UserProData.subscribe((d)=>{getUserProData=d})
	// 	 ActiveFrndChatShort.subscribe((d)=>{getfrndProData=d})
	// 	 socket.close();
	// }
		let req ={
			ReqUUID : $UserProData.UUID as string ,
			FrndUUID: $ActiveFrndChatShort.UUID as string
		}
		// let messageInput = ""
		let socket = new WebSocket("ws://127.0.0.1:8889/gin/user/getconversationid/")

		console.log("GetConversationID send req ",req)

		socket.onopen = () => {
			socket.send(JSON.stringify(req))
		}

		socket.onmessage = (event) => {

			let data = JSON.parse(event.data)
			console.log("websocket Active ConversationID : ",data)
			ConversationID=data.ConversationID
			ActiveConversationID.set(ConversationID)
			console.log("websocket Get active ConversationID : ",data)

		}
	const GetConversationID = () => {
		// if(messageInput.length){
		req ={
			ReqUUID : $UserProData.UUID as string ,
			FrndUUID: $ActiveFrndChatShort.UUID as string
		}
		socket.send(JSON.stringify(req))
		// }
		// messageInput = ""
	}
	// GetConversationID()

	//  function GetAllConversationData(){
	// 	// let messageInput = ""
	// 	// messengerValue=""
	// 	//  socket.close();
	//
	//
	// }
		let socket2 = new WebSocket("ws://127.0.0.1:8889/gin/user/getconversationmsg")
		let get
		 ActiveConversationID.subscribe((d)=>{get=d})
		// let req2 ={
		// 	ConversationID : get
		// }

		socket2.onopen = () => {
			socket2.send(JSON.stringify({
				ConversationID : $ActiveConversationID
			}))
		}

		socket2.onmessage = (event) => {

			let data = JSON.parse(event.data)
			console.log("websocket GetAllConversationData : ",data)
			ActiveConversationData.set(data)
			console.log("ActiveConversationData ",$ActiveConversationData)

		}
	const GetAllConversationData = () => {
		let req ={
			ConversationID : $ActiveConversationID
		}
		socket2.send(JSON.stringify(req))

	}
	// socket2.close();

	function  ResetGetAllConversationData(){
		ActiveConversationData.set([{
			SenderID: "" as string ,
			SenderName : "" as string,
			Message : "" as string,
			Reactions : [] as number[],
			UserReaction : [] as {
				UserID : ""  ,
				ReactionID : 0
			}[] ,
			Timestamp : "" as string,
		}])

		ActiveConversationID.set("")
	}


	// GetFrndData()
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<div
	class="  flex h-screen min-w-[404px] max-w-[404px] flex-col overflow-hidden bg-[#2f3136] transition-all duration-500 ease-linear  "
>
	<div class=" scrol h-full w-full  overflow-y-auto overflow-x-hidden  text-gray-300  ">
		<!-- user chat list -->
		{#if FriendList.length > 0}
			{#each FriendList as i}
				<button
					on:click={() => {
						//ActiveChat=i.UUID 
						UserActive.set(i.UserID);
						GetFrndData(i.UserID)
						ActiveFrndChatShort.set(i)
						console.log("ActiveFrndChatShort" ,$ActiveFrndChatShort)
						//ResetGetAllConversationData()

						GetConversationID()
						//GetAllConversationData()
						GetAllConversationData()
						socket2.close()
						goto(`/${i.UserID}`)}
					}
					class="  m-1   h-16 w-full hover:text-white    {$UserActive === i.UserID
						? 'bg-slate-600 rounded-xl'
						: 'hover:bg-[#202225]'} flex flex-row hover:rounded-xl  "
				>
					<!-- active indecator -->

					<!-- user image -->
					<!-- <img
					src={i['ProfileImage']}
					alt=""
					class=" w-14 h-14 m-2   rounded-3xl hover:rounded-xl active:rounded-md object-cover hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer  "
				/> -->
					<!-- USER IMAGE -->
					<button
						class=" group my-2   ml-1  flex  w-16 flex-row transition-all duration-150 ease-linear hover:rounded-lg"
					>
						<div class=" relative flex ">
							<!-- <svg  class="fill-white  -left-[63px] -top-1  absolute  h-16 w-16 " viewBox="0 0 448 512" xmlns="http://www.w3.org/2000/svg"><path d="M448 95.1v320c0 35.35-28.65 64-64 64H64c-35.35 0-64-28.65-64-64v-320c0-35.35 28.65-63.1 64-63.1h320C419.3 31.1 448 60.65 448 95.1z"/></svg> -->
							<!-- <div in:scale="{{  duration: 200 }}" out:scale class="bg-white -left-[58px]  absolute transition-all duration-200 ease-linear    {ActiveServer===u["ServerURL"] ? "rounded-md h-[50px] w-[50px]" : "rounded-xl h-[10px] w-[50px] group-hover:rounded-md group-hover:h-[30px] group-hover:w-[50px] " }"></div> -->
							<!-- <svg in:scale="{{  duration: 200 }}" out:scale class="fill-white   -left-[55px] -top-[25px]   absolute  h-[50px] w-[50px] " viewBox="0 0 448 512" xmlns="http://www.w3.org/2000/svg"><path d="M384 32H64C28.65 32 0 60.66 0 96v320c0 35.34 28.65 64 64 64h320c35.35 0 64-28.66 64-64V96C448 60.66 419.3 32 384 32zM319.1 280h-192C114.8 280 103.1 269.2 103.1 256c0-13.2 10.8-24 24-24h192c13.2 0 23.1 10.8 23.1 24C343.1 269.2 333.2 280 319.1 280z"/></svg> -->
							{#if i.UserImg === ''}
								<Person
									class="  h-12 w-12 cursor-pointer rounded-xl border-2  border-gray-400  fill-gray-400 object-cover p-2    transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md"
								/>
							{:else}
								<!-- <img
					src={$UserProData.ProfileImg}
					alt="coverImage"
					class="fixed -mt-7 ml-5 h-12 w-12 rounded-xl object-cover"
						/> -->
								<img
									src={i.UserImg}
									alt=""
									class="  h-[48px] w-[48px] cursor-pointer rounded-[1.25rem] object-cover transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md  "
								/>
							{/if}
						</div>
					</button>
					<!-- NAME , LAST MESSAGE -->
					<div class="mx-2 flex h-full w-full  flex-col justify-center  ">
						<div class=" mt-2 ml-0 flex flex-row text-left">
							<!-- User name -->
							<p class=" w-56  text-base font-bold    ">
								{i.UserName}
							</p>
							<div class=" flex-grow" />
							<!-- mute icon -->
							{#if i.SilentNotification}
								<svg class="m-1 h-5 w-5 flex-none " fill="currentColor" viewBox="0 0 20 20"
									><path
										fill-rule="evenodd"
										d="M9.383 3.076A1 1 0 0110 4v12a1 1 0 01-1.707.707L4.586 13H2a1 1 0 01-1-1V8a1 1 0 011-1h2.586l3.707-3.707a1 1 0 011.09-.217zM12.293 7.293a1 1 0 011.414 0L15 8.586l1.293-1.293a1 1 0 111.414 1.414L16.414 10l1.293 1.293a1 1 0 01-1.414 1.414L15 11.414l-1.293 1.293a1 1 0 01-1.414-1.414L13.586 10l-1.293-1.293a1 1 0 010-1.414z"
										clip-rule="evenodd"
									/></svg
								>
							{/if}
							<!-- last message time -->
							<p class="   mx-0 mt-3 mb-1 flex-none text-xs font-thin  ">
								<!-- 34 aug21  -->
								{i.LastMessageTime.split(/(\s+)/)[0]}
							</p>
						</div>
						<div class="  mb-2 flex w-72 flex-row">
							<!-- last message -->
							<p class=" h-full w-72  text-left  text-sm  line-clamp-1">
								 {i.LastMessage}
<!--								LastMessage-->
							</p>
							<!-- number of notification -->
							{#if i.SilentNotification && i.NumberOfNotification != 0}
								<span
									class="   h-5  rounded-xl  border-2 border-[#202225] bg-slate-500 px-1 text-xs font-semibold text-white  "
								>
									{i.NumberOfNotification}

								</span>
							{:else if i.NumberOfNotification != 0}
								<span
									class="   h-5  rounded-xl  border-2 border-[#202225] bg-red-500 px-1 text-xs font-semibold text-white  "
								>
									{i.NumberOfNotification}
								</span>
							{/if}
						</div>
					</div>
				</button>
			<!-- {:else} -->
				<!-- empty list -->
			{/each}
		{:else}
			<!-- else content here -->
			<div class="flex h-full w-full flex-col items-center justify-center">
				<DiscordNoOne3 class=" w-[80%]  self-center" />

				<p class=" font-Poppins text-lg font-semibold ">No Friends</p>
			</div>
		{/if}
	</div>
	<div class="">
		<UserDashPannel {showUsername} />
	</div>
</div>

<style>
	/* your styles go here */
</style>
