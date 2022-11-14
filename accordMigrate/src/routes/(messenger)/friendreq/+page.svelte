<script lang="ts">
	// your script goes here
	import Coin2 from '$lib/icons/coin2.svelte';
	import DiscordNoOne from '$lib/icons/discordNoOne.svelte';
	import { UserProData } from '$lib/store2';
	// import Nop from '$lib/icons/Errorsign.svelte';

	import type { PageData } from './$types';
	export let data: PageData;

	let { Userdata } = data;
	console.log('🚀 ~ file: +page.svelte ~ line 10 ~ Userdata', Userdata);

	UserProData.update((d) => (d = Userdata));

	let FrndSuggList: {
		UUID: string;
		UserID: string;
		UserName: string;
		ProfileImg: string;
		UserBio: string;
	}[] = [] as {
		UUID: string;
		UserID: string;
		UserName: string;
		ProfileImg: string;
		UserBio: string;
	}[];

	async function FetchSuggestionList() {
		await fetch(`http://localhost:8888/user/frndsuggestion`, {
			method: 'POST',
			mode: 'cors',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				UUID: $UserProData.UUID
			})
		})
			.then((res) => {
				return res.json();
			})
			.then((d) => {
				FrndSuggList = d;
				console.log('🚀 ~ file: index.svelte ~ line 296 ~ .then ~ FrndSuggList', FrndSuggList);
			})
			.catch((err) => {
				console.log(err);
			});
	}
	FetchSuggestionList();
	// let UserReqList: {
	// 	UUID: string;
	// 	Coin: number;
	// 	ReqList: {
	// 		Amount: number;
	// 		JWT: string;
	// 		Status: string;
	// 	}[]
	// }[]= []
	// async function FetchUserReqList() {

	// 	await fetch(`/api/profile/moneymanagement`, {
	// 			method: 'GET',
	// 			mode: 'no-cors',

	// 	}).then((res)=>{return res.json()}).
	// 	then((data)=>{
	// 	console.log("🚀 ~ file: +page.server.ts ~ line 112 ~ constload:PageServerLoad= ~ data", data)
	// 	UserReqList=data
	// 	})
	// }

	// FetchUserReqList()
	// let UserReqList = [
	// 	{
	// 		UserID: 'skraka',
	// 		Coin: 28.23,
	// 		ReqList: [
	// 			{ Amount: 245, JWT: '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', Status: 'pending' },
	// 			{ Amount: 235, JWT: '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', Status: 'accepted' },
	// 			{ Amount: 523, JWT: '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', Status: 'pending' }
	// 		]
	// 	},
	// 	{
	// 		UserID: 'linux',
	// 		Coin: 78.23,
	// 		ReqList: [
	// 			{ Amount: 245, JWT: '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', Status: 'accepted' },
	// 			{ Amount: 36, JWT: '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', Status: 'pending' },
	// 			{ Amount: 36, JWT: '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', Status: 'pending' }
	// 		]
	// 	},
	// 	{
	// 		UserID: 'whatsup',
	// 		Coin: 72.23,
	// 		ReqList: [
	// 			{ Amount: 36, JWT: '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', Status: 'pending' },
	// 			{ Amount: 36, JWT: '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', Status: 'pending' },
	// 			{ Amount: 36, JWT: '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', Status: 'pending' }
	// 		]
	// 	}
	// ];
	// async function AcceptReq(Amount: number, UserID: string) {
	// 	console.log('🚀 ~ file: +page.svelte ~ line 45 ~ AcceptReq ~ UserID', UserID);
	// 	console.log('🚀 ~ file: +page.svelte ~ line 44 ~ AcceptReq ~ Amount', Amount);
	// 	let res = await fetch('/api/profile/moneyreqaccept', {
	// 		method: 'POST',
	// 		headers: {
	// 			'Content-Type': 'application/json'
	// 		},
	// 		body: JSON.stringify({
	// 			UserID: UserID,
	// 			Amount: Amount
	// 		})
	// 	});
	// 	let data = await res.json();
	// 	console.log('🚀 ~ file: +page.svelte ~ line 45 ~ AcceptReq ~ data', data);
	// 	FetchUserReqList()
	// }

	// const DefaultUserReqList: {
	// 	UUID: string;
	// 	Coin: number;
	// 	ReqList: {
	// 		Amount: number;
	// 		JWT: string;
	// 		Status: string;
	// 	}[]
	// }[]= [] as  {
	// 	UUID: string;
	// 	Coin: number;
	// 	ReqList: {
	// 		Amount: number;
	// 		JWT: string;
	// 		Status: string;
	// 	}[]
	// }[]
</script>

<!-- markup (zero or more items) goes here -->
<div
	class=" flex xl:flex-col   h-screen w-full overflow-hidden bg-[#36393f] md:flex-col md:overflow-y-scroll lg:flex-row  "
>
	<!-- FRIEND req -->
	<div class="flex w-[80%]  flex-col p-5 transition-all  duration-200  ease-linear gap-4">
		<p class=" text-3xl text-white ">All Friend Request :</p>
		<div class="flex flex-row flex-wrap justify-evenly">
			{#if FrndSuggList.length > 0}
				<!-- content here -->
				{#each FrndSuggList as i}
					<!-- {#each Array.from(Array(10 + 1).keys()).slice(1) as i} -->
					<div class=" flex h-[350px] w-[200px]  flex-col    rounded-lg bg-gray-600 bg-opacity-80 ">
						<!--  -->
						<img
							src={i.ProfileImg}
							alt="UserImage"
							class=" h-[60%]  w-fit rounded-t-lg object-cover"
						/>
						<p class="mx-1 font-sans text-base font-semibold line-clamp-2  ">{i['UserName']}</p>
						<p
							class=" lamp-1 mx-1 text-center font-sans text-sm
				 font-light text-gray-300 "
						>
							{i['UserBio'] ? i['UserBio'] : 'No Bio'}
							<!-- Work at Google INCWork at Google INCWork at Google INC -->
						</p>
						<div class="my-2 flex flex-col items-center justify-center     gap-2 text-base">
							<button
								class="w-[90%] rounded-md bg-[#3982e4] px-2 py-1 font-semibold text-slate-200 hover:bg-blue-600 active:bg-blue-700"
								>Add Friend</button
							>
							<button
								class="w-[90%]  rounded-md border-0 bg-[#3a3b3c] px-2 py-1 text-white hover:bg-gray-700 active:bg-gray-600"
								>Remove</button
							>
						</div>
					</div>
				{/each}
			{:else}
				<DiscordNoOne class="self-center justify-self-center" />
			{/if}
		</div>
	</div>
	<!-- People you may know -->
	<div class="flex w-[80%] flex-col p-5 transition-all  duration-200  ease-linear gap-4">
		<p class=" text-3xl text-white w-full ">People you may know</p>
		<div class="flex flex-row flex-wrap justify-evenly">
			{#if FrndSuggList.length > 0}
				<!-- content here -->
				{#each FrndSuggList as i}
					<!-- {#each Array.from(Array(10 + 1).keys()).slice(1) as i} -->
					<div class=" flex h-[350px] w-[200px]  flex-col    rounded-lg bg-gray-600 bg-opacity-80 ">
						<!--  -->
						<img
							src={i.ProfileImg}
							alt="UserImage"
							class=" h-[60%]  w-fit rounded-t-lg object-cover"
						/>
						<p class="mx-1 font-sans text-base font-semibold line-clamp-2  ">{i['UserName']}</p>
						<p
							class=" lamp-1 mx-1 text-center font-sans text-sm
				 font-light text-gray-300 "
						>
							{i['UserBio'] ? i['UserBio'] : 'No Bio'}
							<!-- Work at Google INCWork at Google INCWork at Google INC -->
						</p>
						<div class="my-2 flex flex-col items-center justify-center     gap-2 text-base">
							<button
								class="w-[90%] rounded-md bg-[#3982e4] px-2 py-1 font-semibold text-slate-200 hover:bg-blue-600 active:bg-blue-700"
								>Add Friend</button
							>
							<button
								class="w-[90%]  rounded-md border-0 bg-[#3a3b3c] px-2 py-1  text-slate-200 hover:bg-gray-700 active:bg-gray-600"
								>Remove</button
							>
						</div>
					</div>
				{/each}
			{:else}
				<DiscordNoOne class="self-center justify-self-center" />
			{/if}
		</div>
	</div>
</div>
