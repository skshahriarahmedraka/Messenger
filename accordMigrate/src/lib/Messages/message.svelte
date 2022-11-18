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
	export let message: {SenderName: string, Message: string, UserReaction: {UserID: "", ReactionID: 0}[], Reactions: number[], SenderID: string, Timestamp: string};
	import {UserProData , ActiveFrndData} from '$lib/store2';
	console.log("my message :" ,message)
	// var person = { fname: 'Nick', lname: 'Jonas', age: 26 };
	// for (let x in person) {
	// 	console.log(x + ': ' + person[x]);
	// }
	import Person from '$lib/icons/person.svelte';



	let Message = {
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

	Message.Message="whats up"
	Message.SenderName="Sk shahriar"
	Message.Timestamp="12 july"
	Message.SenderID="skssar"
	// message =Message

	Message=message
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
			// console.log(x + ': ' + react[x]);
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

		// r = r.sort(function (a, b) {
		// 	return a[1] - b[1];
		// });

		// r.sort();
		// r.reverse();
		return r;
	}

	let reactions = [Likex, Lovex, Happy, Fillinglove, Wowx, Hahax, Unlike, Sadx, Cry, Angryx];

	// like 1 , love 2 ,happy 3 , care 4 ,wow 5,haha 6,Unlike 7 ,sad 8, cry 9,angry 10
	let GivingReact:boolean=false
</script>

<!-- content here -->
<!-- { message.writer=== UserProData.Name  ? "flex-row-reverse" : "flex-row" }  -->
<!-- {#if true} -->
<div class=" flex flex-col">
	<div
		class=" flex  {Message.SenderID === $UserProData.UserID
			? 'flex-row-reverse'
			: 'flex-row'}    mr-20 mt-4 ml-5 max-w-full   overflow-ellipsis text-base text-white "
	>
		<!-- <div class="w-[80%] ?"> -->
		{#if $ActiveFrndData.ProfileImg != "" }
		<img src={$ActiveFrndData.ProfileImg} alt="" class=" m-2 h-10 w-10 rounded-2xl " />
			{:else }
<!--			<img src={Ek2} alt="" class=" m-2 h-10 w-10 rounded-2xl " />-->
			<Person
					class="m-2 h-10 w-10 rounded-2xl cursor-pointer rounded-xl border-2  border-gray-400  fill-gray-400 object-cover p-2    transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md"
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

						</div>
					{/if}
					<AddReaction  class="h-6 w-6" />
				</div>
				

			
			</div>
		</div>
	</div>
</div>

<!-- {/if} -->
<style>
	/* your styles go here */
</style>
