<script lang="ts">
	import AddReaction from './reactions/addReaction.svelte';
import Angryx from './reactions/angry.svelte';

	import Carex from './reactions/carex.svelte';
	import Cry from './reactions/cry.svelte';
	import Hahax from './reactions/hahay.svelte';

	import Happy from './reactions/happy.svelte';

	import Likex from './reactions/likey.svelte';
	import Lovex from './reactions/love.svelte';
	import Sadx from './reactions/sadx.svelte';

	import Unlike from './reactions/unlike.svelte';
	import Wowx from './reactions/wowx.svelte';

	import Ek2 from './profilePic/e2.jpg';
import Party from './reactions/party.svelte';
import Angrytalk from './reactions/angrytalk.svelte';
import Fillinglove from './reactions/fillinglove.svelte';
	import {UserProData, ActiveFrndData, ActiveConversationID, ActiveConversationData} from '$lib/store2';

	import Person from '$lib/icons/person.svelte';

	function GetAllConversationData(){
	
		let socket = new WebSocket("ws://127.0.0.1:8889/gin/user/getconversationmsg")
	
		let req ={
			ConversationID : $ActiveConversationID
		}

		socket.onopen = () => {
			socket.send(JSON.stringify(req))
		}

		socket.onmessage = (event) => {

			let data = JSON.parse(event.data)
			
			ActiveConversationData.set(data)
		

		}
		


	}

	let MsgData=$ActiveConversationData
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
		MsgData=data
		ActiveConversationData.set(data)
		console.log("ActiveConversationData ",$ActiveConversationData)

	}
	const GetMessage = () => {
		req ={
			ConversationID : $ActiveConversationID
		}
		socket.send(JSON.stringify(req))

	}


	function ReactCount(react: { [key: string]: number }) {
		let r = [
			[0, 0],
			[1, 0],
			[2, 0],
			[3, 0],
			[4, 0],
			[5, 0],
			[6, 0],
			[7, 0],
			[8, 0],
			[9, 0]
		];
		for (let x in react) {
			r[react[x]][1] += 1;
		}

		r.sort(compareSecondColumn);

		function compareSecondColumn(a: any, b: any) {
			if (a[1] === b[1]) {
				return 0;
			} else {
				return a[1] > b[1] ? -1 : 1;
			}
		}

		
		return r;
	}

	let reactions = [Likex, Lovex, Happy, Fillinglove, Wowx, Hahax, Unlike, Sadx, Cry, Angryx];

	// like 1 , love 2 ,happy 3 , care 4 ,wow 5,haha 6,Unlike 7 ,sad 8, cry 9,angry 10
	let GivingReact=false
	var intervalId = window.setInterval(function(){
		GetMessage()
	}, 2000);

	function  GetMsgImg(SenderID:string,UserID :string) {
		if (SenderID === UserID) {
			return $UserProData.ProfileImg
		}else {
			return $ActiveFrndData.ProfileImg
		}
	}
</script>



{#each $ActiveConversationData.reverse() as Message }

	<div class=" flex flex-col ?">
		<div
			class=" flex  {Message.SenderID === $UserProData.UserID
				? 'flex-row-reverse'
				: 'flex-row'}    mr-20 mt-4 ml-5 max-w-full   overflow-ellipsis text-base text-white "
		>
			<!-- <div class="w-[80%] ?"> -->
			{#if GetMsgImg(Message.SenderID,$UserProData.UserID) != "" }
			<img src={GetMsgImg(Message.SenderID,$UserProData.UserID)} alt="" class=" m-2 h-10 w-10 rounded-2xl object-cover " />
				{:else }
				<Person
						class="m-2 h-10 w-10 rounded-2xl cursor-pointer border-2  border-gray-400  fill-gray-400 object-cover p-2    transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md"
				/>
				{/if}
			<div class="flex max-w-[80%] flex-col flex-wrap ">
				<div
					class="flex flex-row  {Message.SenderID != $UserProData.UserID ? 'self-start' : 'self-end'} gap-1"
				>
					<button class=" h-6 text-[#04b1aa] line-clamp-1 hover:text-[#47faf4]">
						{Message.SenderName}
					</button>
					<div class=" mt-2 ml-2 text-xs text-[#6e6a5c]">
						{Message.Timestamp}
					</div>
				</div>
				<!-- MESSAGES -->
				<div
					class=" text-[15px] {Message.SenderID === $UserProData.UserID
						? 'bg-blue-400 bg-opacity-20 p-2 '
						: 'bg-gray-700'}  rounded-lg p-2 "
				>
					{Message.Message}
				</div>
				<!-- REACTIONS -->
				<div class="no-scrollbar flex h-6 w-full flex-row overflow-x-scroll gap-1 ">
					<!-- {#each ReactCount(message.react) as x} -->
					<!--{#each message.numberOfReact as x,id}-->
					<!--	-->
					<!--	{#if x != 0}-->
					<!--		-->
					<!--		<div class=" border-0 bg-[#2f3136] rounded-xl w-fit  flex flex-row justify-center items-center gap-1">-->
					<!--			<svelte:component this={reactions[id]} class=" h-[18px] w-[18px]" />-->
					<!--			<p class=" h-6 self-end  text-xs p-1 mr-1">{x}</p>-->
					<!--		</div>-->
					<!--	{/if}-->
					<!--{:else}-->
					<!--	-->
					<!--{/each}-->
					<div class="relative" on:click="{()=>{GivingReact=!GivingReact}}">
						{#if GivingReact}
							<div class=" fixed  z-50  bg-white  h-6 w-36 rounded-lg ">
								<!--  -->
							</div>
						{/if}
						<AddReaction  class="h-6 w-6" />
					</div>



				</div>
			</div>
		</div>
	</div>
{/each}

