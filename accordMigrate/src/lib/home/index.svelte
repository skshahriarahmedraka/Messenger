<script lang="ts">
	// let suggestedFrineds = [
	// 	{
	// 		name: 'md abdullah ',
	// 		bio: 'work at Google ',
	// 		img: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1660151599/samples/about_dmmlnj.jpg'
	// 	},
	// 	{
	// 		name: '',
	// 		bio: '',
	// 		img: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1660151599/samples/about_dmmlnj.jpg'
	// 	}
	// ];
	// import Moji from '$lib/profilepic/moji2.jpg';
	// import Rafi from '$lib/profilepic/r4.jpg';
	// import Ekbal from '$lib/profilepic/e2.jpg';
	// import Ronok from '$lib/profilepic/r5.jpg';
	// import Shakil from '$lib/profilepic/s2.jpg';
	// import Hemel from '$lib/profilepic/h1.jpg';
	// import Azim from '$lib/profilepic/a1.jpg';
	// import Rafel from '$lib/profilepic/r2.jpg';
	// import Rakib from '$lib/profilepic/r.jpg';
	// import Istique from '$lib/profilepic/ot2.jpg';
	// import Rakib2 from '$lib/profilepic/r3.jpg';
	// import Toki from '$lib/profilepic/t.jpg';
	// import Yurichi from '$lib/profilepic/y.jpeg';
	// import Rengokhu from '$lib/profilepic/rengokhu.png';
	// import Vegeta from '$lib/profilepic/s2.png';
	// import Saad from '$lib/profilepic/s3.jpg';
	// import Hankok from '$lib/profilepic/h2.jpg';
	// import Nejuko from '$lib/profilepic/n1.jpg';
	// import Shanks from '$lib/profilepic/s4.jpg';
	// import Admiral from '$lib/profilepic/s5.jpg';
	// import Law from '$lib/profilepic/law.png';

	import { UserProData } from '$lib/store2';

	// let x = [
	// 	{
	// 		ProfileURL: 'majibarrahman',
	// 		ProfileImage: Moji,
	// 		UserName: 'Md Majibar Rahman',
	// 		LastMessage: 'hope you are filling well',
	// 		LastMessageTime: '23 Aug 21',
	// 		SilentNotification: true,
	// 		NumberOfNotification: 43,
	// 		ActiveStatus: true,
	// 		LastActiveTime: '10:28 AM'
	// 	},
	// 	{
	// 		ProfileURL: 'rafihasan',
	// 		ProfileImage: Rafi,
	// 		UserName: 'md rafi hasan zihad',
	// 		LastMessage: 'whats is going on bro',
	// 		LastMessageTime: '9:30 AM',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 3,
	// 		ActiveStatus: true,
	// 		LastActiveTime: '6 Jul 17'
	// 	},
	// 	{
	// 		ProfileURL: 'ekbalhasan',
	// 		ProfileImage: Ekbal,
	// 		UserName: 'Md ekbal Hasan',
	// 		LastMessage: 'I was going home that day',
	// 		LastMessageTime: '9 jun 11',
	// 		SilentNotification: true,
	// 		NumberOfNotification: 4,
	// 		ActiveStatus: false,
	// 		LastActiveTime: '29 Jun 18'
	// 	},
	// 	{
	// 		ProfileURL: 'mdronok',
	// 		ProfileImage: Ronok,
	// 		UserName: 'Md ronok shahriar',
	// 		LastMessage: 'ethical hacking is my passion',
	// 		LastMessageTime: '10:34 PM',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 1,
	// 		ActiveStatus: false,
	// 		LastActiveTime: '10:28 AM'
	// 	},
	// 	{
	// 		ProfileURL: 'samiulshakil',
	// 		ProfileImage: Shakil,
	// 		UserName: 'md samiul alim shakil',
	// 		LastMessage: 'I am feeling wanderful . wona go for a tour',
	// 		LastMessageTime: '12 mar 19',
	// 		SilentNotification: true,
	// 		NumberOfNotification: 7,
	// 		ActiveStatus: true,
	// 		LastActiveTime: '7 Aug 19'
	// 	},
	// 	{
	// 		ProfileURL: 'hemelakash',
	// 		ProfileImage: Hemel,
	// 		UserName: 'Hemel sarkar akash',
	// 		LastMessage: 'Got a job in Apple as a CEO',
	// 		LastMessageTime: '12 min',
	// 		SilentNotification: true,
	// 		NumberOfNotification: 6,
	// 		ActiveStatus: false,
	// 		LastActiveTime: '10:28 AM'
	// 	},
	// 	{
	// 		ProfileURL: 'hazem',
	// 		ProfileImage: Azim,
	// 		UserName: 'Md Azim Mahfuz',
	// 		LastMessage: 'what a great day . got a job on Uniliver',
	// 		LastMessageTime: '23 sec',
	// 		SilentNotification: true,
	// 		NumberOfNotification: 7,
	// 		ActiveStatus: true,
	// 		LastActiveTime: '1:06 AM'
	// 	},

	// 	{
	// 		ProfileURL: 'rafel',
	// 		ProfileImage: Rakib,
	// 		UserName: 'Md rafel',
	// 		LastMessage: 'bought a new bike ,lets to on a ride',
	// 		LastMessageTime: '4:23 AM',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 0,
	// 		ActiveStatus: false,
	// 		LastActiveTime: '10:28 AM'
	// 	},
	// 	{
	// 		ProfileURL: 'rakibulislam',
	// 		ProfileImage: Rafel,
	// 		UserName: 'md Rakibul islam',
	// 		LastMessage: "whats's you doing , bro , see you later",
	// 		LastMessageTime: '23 dec 21',
	// 		SilentNotification: true,
	// 		NumberOfNotification: 2,
	// 		ActiveStatus: true,
	// 		LastActiveTime: '00:28 AM'
	// 	},
	// 	{
	// 		ProfileURL: 'istique',
	// 		ProfileImage: Istique,
	// 		UserName: 'istique mahmud pritom',
	// 		LastMessage: 'i am on a train , going to Dhaka',
	// 		LastMessageTime: '1 jun 21',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 0,
	// 		ActiveStatus: true,
	// 		LastActiveTime: '11:28 AM'
	// 	},
	// 	{
	// 		ProfileURL: 'mdrakib',
	// 		ProfileImage: Rakib2,
	// 		UserName: 'Md Rakib Hasan',
	// 		LastMessage: 'lets do party to night . Real madrid vs PSG . the game will be epic',
	// 		LastMessageTime: '9 mar',
	// 		SilentNotification: true,
	// 		NumberOfNotification: 2,
	// 		ActiveStatus: true,
	// 		LastActiveTime: '00:00 AM'
	// 	},
	// 	{
	// 		ProfileURL: 'toki',
	// 		ProfileImage: Yurichi,
	// 		UserName: 'md toki',
	// 		LastMessage: 'whats up bro what is happening there',
	// 		LastMessageTime: '21 jun',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 0,
	// 		ActiveStatus: false,
	// 		LastActiveTime: '9 Feb 30'
	// 	},
	// 	{
	// 		ProfileURL: 'yurichi',
	// 		ProfileImage: Toki,
	// 		UserName: 'Yurichi',
	// 		LastMessage: 'are you enjoying my story',
	// 		LastMessageTime: '12 jan 16',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 9,
	// 		ActiveStatus: false,
	// 		LastActiveTime: '23 Dec 21'
	// 	},
	// 	{
	// 		ProfileURL: 'rengokhu',
	// 		ProfileImage: Rengokhu,
	// 		UserName: 'kyjiro Rengokhu',
	// 		LastMessage: 'mugen the train was amayzing',
	// 		LastMessageTime: '19 feb 19',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 5,
	// 		ActiveStatus: true,
	// 		LastActiveTime: '10:28 PM'
	// 	},
	// 	{
	// 		ProfileURL: 'vegeta',
	// 		ProfileImage: Vegeta,
	// 		UserName: 'Prince Vegeta',
	// 		LastMessage: 'the sayian pride will never parish',
	// 		LastMessageTime: '9 sep 20',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 0,
	// 		ActiveStatus: false,
	// 		LastActiveTime: '10:28 AM'
	// 	},

	// 	{
	// 		ProfileURL: 'saad',
	// 		ProfileImage: Saad,
	// 		UserName: 'Hasan mahmud',
	// 		LastMessage: "what's your physical condition , come to my home ",
	// 		LastMessageTime: 'yestarday',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 5,
	// 		ActiveStatus: true,
	// 		LastActiveTime: '3:00 AM'
	// 	},
	// 	{
	// 		ProfileURL: 'hankok',
	// 		ProfileImage: Hankok,
	// 		UserName: 'Boa Hankok',
	// 		LastMessage: 'where is my luffy',
	// 		LastMessageTime: '12 jan 00',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 99,
	// 		ActiveStatus: false,
	// 		LastActiveTime: '12:28 AM'
	// 	},
	// 	{
	// 		ProfileURL: 'nejuko',
	// 		ProfileImage: Nejuko,
	// 		UserName: 'Nejuko',
	// 		LastMessage: 'hmm',
	// 		LastMessageTime: '2 jan 34',
	// 		SilentNotification: true,
	// 		NumberOfNotification: 9,
	// 		ActiveStatus: true,
	// 		LastActiveTime: 'Active'
	// 	},
	// 	{
	// 		ProfileURL: 'shanks',
	// 		ProfileImage: Shanks,
	// 		UserName: 'D. Shanka',
	// 		LastMessage: 'after shreding tears you will be a real man',
	// 		LastMessageTime: '1:00 AM',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 45,
	// 		ActiveStatus: false,
	// 		LastActiveTime: '10:00 AM'
	// 	},
	// 	{
	// 		ProfileURL: 'admiral',
	// 		ProfileImage:
	// 			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
	// 		UserName: 'Nevy Admiral',
	// 		LastMessage: 'justice must be absolute',
	// 		LastMessageTime: '6:39 AM',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 0,
	// 		ActiveStatus: false,
	// 		LastActiveTime: '9:28 AM'
	// 	},
	// 	{
	// 		ProfileURL: 'law',
	// 		ProfileImage: Law,
	// 		UserName: 'Trafalger D. Water Law',
	// 		LastMessage: '     i must find the meaning of meaning of D.',
	// 		LastMessageTime: 'yesterday',
	// 		SilentNotification: false,
	// 		NumberOfNotification: 34,
	// 		ActiveStatus: true,
	// 		LastActiveTime: 'Active'
	// 	}
	// ];
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

	async function SendFriendReq() {
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
	SendFriendReq();
	import Error2 from '$lib/Error/error2.svelte';
	import { page } from '$app/stores';
	import Loading from '$lib/Loading/solarAccord.svelte';
	import { afterUpdate } from 'svelte';
	import DiscordNoOne from '$lib/icons/discordNoOne.svelte';
	import Accord from '$lib/Dock/icons/accord.svelte';

	let loadingState: boolean = true;
	afterUpdate(() => {
		loadingState = false;
	});
</script>

{#if loadingState}
	<Loading class=" scale-150" />
{:else if $page.error}
	<!-- <Error1/> -->
	<Error2 />
{:else}
	<div class=" m-4 flex flex-row flex-wrap justify-evenly  gap-2">
		{#if FrndSuggList.length > 0}
			<!-- content here -->
			{#each FrndSuggList as i}
				<div class=" flex h-[350px] w-[200px]  flex-col    rounded-lg bg-gray-600 bg-opacity-80 ">
					{#if i.ProfileImg.trim()=== ""}
						 <!-- content here -->
						 <Accord  class="h-[60%] self-center bg-slate-600 w-full rounded-t-lg object-cover" />
					{:else}
						 <!-- else content here -->
						 <img
							 src={i.ProfileImg}
							 alt="UserImage"
							 class=" h-[60%]  w-fit rounded-t-lg object-cover"
						 />
					{/if}
					<p class="mx-1 font-sans text-base font-semibold line-clamp-2 self-center   ">{i['UserName']}</p>
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
			<!-- else content here -->
			<DiscordNoOne class="self-center justify-self-center" />
		{/if}
	</div>
{/if}
