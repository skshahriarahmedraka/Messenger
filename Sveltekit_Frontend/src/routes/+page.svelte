<script lang="ts">
	// import '../app.css';

	// import { onMount, beforeUpdate, afterUpdate } from 'svelte';
	// let counter = 0;
	// 	beforeUpdate(() => {
	// 	console.log('the component is going to be updated');
	// 	});
	// 	afterUpdate(() => {
	// 	console.log('the component has been updated');
	// 	});
	import type { PageData } from './$types';
	import { UserProData, ActiveFrndData } from '$lib/store2';


	export let data: PageData;

	let { Userdata } = data;
	console.log('🚀 ~ file: +layout.svelte ~ line 26 ~ Userdata', Userdata);
	UserProData.update((d) => (d = Userdata));

	ActiveFrndData.set(Userdata)

	import Dock from '$lib/Dock/index.svelte';
	// import Servers from "$lib/_Servers/index.svelte"
	import NavBar from '$lib/Navbar/index.svelte';
	// import Peoples from "$lib/_Peoples/index.svelte"
	// import Messages from "$lib/_Messages/index.svelte"
	// import {showPeopleList,ChatOrDock} from "$lib/store2"
	import Home from '$lib/home/index.svelte';

	import Chat from '$lib/Chats/index.svelte';

	// let MyPro = {
	// 	Name: 'Sk Shahriar Ahmed Raka',
	// 	Userid: 'skraka',
	// 	ProfileImage:
	// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
	// 	CoverImage:
	// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png'
	// };

	// let showPeopleListValue: number
	// let ChatOrDockValue: number

	// showPeopleList.subscribe(val=>{
	//     showPeopleListValue=val
	// })

	// ChatOrDock.subscribe(val=>{
	//     ChatOrDockValue=val
	// })
	// friend images
	import Moji from '$lib/profilepic/moji2.jpg';
	import Rafi from '$lib/profilepic/r4.jpg';
	import Ekbal from '$lib/profilepic/e2.jpg';
	import Ronok from '$lib/profilepic/r5.jpg';
	import Shakil from '$lib/profilepic/s2.jpg';
	import Hemel from '$lib/profilepic/h1.jpg';
	import Azim from '$lib/profilepic/a1.jpg';
	import Rafel from '$lib/profilepic/r2.jpg';
	import Rakib from '$lib/profilepic/r.jpg';
	import Istique from '$lib/profilepic/ot2.jpg';
	import Rakib2 from '$lib/profilepic/r3.jpg';
	import Toki from '$lib/profilepic/t.jpg';
	import Yurichi from '$lib/profilepic/y.jpeg';
	import Rengokhu from '$lib/profilepic/rengokhu.png';
	import Vegeta from '$lib/profilepic/s2.png';
	import Saad from '$lib/profilepic/s3.jpg';
	import Hankok from '$lib/profilepic/h2.jpg';
	import Nejuko from '$lib/profilepic/n1.jpg';
	import Shanks from '$lib/profilepic/s4.jpg';
	import Admiral from '$lib/profilepic/s5.jpg';
	import Law from '$lib/profilepic/law.png';

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

	// let x3 = {
	// 	Name: 'Portgas D. Ace',
	// 	ProfileUrl: 'ace',
	// 	ProfileImage:
	// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg'
	// };

	let x2: any = [
		// {
		// 	ServerURL: 'postgres',
		// 	Name: 'Postgres',
		// 	NewMessage: true,
		// 	NumberOfNewMessage: 3,
		// 	Notification: false,
		// 	NumberOfNotification: 0,
		// 	ServerImage: Rafel
		// },
		// {
		// 	ServerURL: 'mongodb',
		// 	Name: 'MongoDB',
		// 	NewMessage: false,
		// 	NumberOfNewMessage: 0,
		// 	Notification: true,
		// 	NumberOfNotification: 3,
		// 	ServerImage:
		// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		// },
		// {
		// 	ServerURL: 'docker',
		// 	Name: 'Docker',
		// 	NewMessage: true,
		// 	NumberOfNewMessage: 4,
		// 	Notification: true,
		// 	NumberOfNotification: 6,
		// 	ServerImage:
		// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		// },
		// {
		// 	ServerURL: 'kubernetes',
		// 	Name: 'Kubernetes',
		// 	NewMessage: false,
		// 	NumberOfNewMessage: 0,
		// 	Notification: true,
		// 	NumberOfNotification: 99,
		// 	ServerImage:
		// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		// },
		// {
		// 	ServerURL: 'redis',
		// 	Name: 'Redis',
		// 	NewMessage: true,
		// 	NumberOfNewMessage: 7,
		// 	Notification: false,
		// 	NumberOfNotification: 0,
		// 	ServerImage:
		// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		// },
		// {
		// 	ServerURL: 'tensorflow',
		// 	Name: 'Tensorflow',
		// 	NewMessage: false,
		// 	NumberOfNewMessage: 0,
		// 	Notification: false,
		// 	NumberOfNotification: 0,
		// 	ServerImage:
		// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		// },
		// {
		// 	ServerURL: 'macity',
		// 	Name: 'Mancity',
		// 	NewMessage: true,
		// 	NumberOfNewMessage: 2,
		// 	Notification: false,
		// 	NumberOfNotification: 0,
		// 	ServerImage:
		// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		// },
		// {
		// 	ServerURL: 'cpp',
		// 	Name: 'C++',
		// 	NewMessage: true,
		// 	NumberOfNewMessage: 4,
		// 	Notification: false,
		// 	NumberOfNotification: 0,
		// 	ServerImage:
		// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		// },
		// {
		// 	ServerURL: 'ubuntu',
		// 	Name: 'Ubuntu',
		// 	NewMessage: true,
		// 	NumberOfNewMessage: 7,
		// 	Notification: true,
		// 	NumberOfNotification: 2,
		// 	ServerImage:
		// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		// },
		// {
		// 	ServerURL: 'realmadrid',
		// 	Name: 'Real Madrid',
		// 	NewMessage: false,
		// 	NumberOfNewMessage: 0,
		// 	Notification: false,
		// 	NumberOfNotification: 0,
		// 	ServerImage:
		// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		// },
		{
			ServerURL: 'bangladesh',
			Name: 'bangladesh',
			NewMessage: true,
			NumberOfNewMessage: 6,
			Notification: true,
			NumberOfNotification: 4,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'c',
			Name: 'C',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'coffee',
			Name: 'Coffee',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: true,
			NumberOfNotification: 45,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'chelsea',
			Name: 'Chelsea',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'go',
			Name: 'Go',
			NewMessage: true,
			NumberOfNewMessage: 9,
			Notification: true,
			NumberOfNotification: 65,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'mazda',
			Name: 'Mazda',
			NewMessage: false,
			NumberOfNewMessage: 2,
			Notification: true,
			NumberOfNotification: 5,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'pinterest',
			Name: 'Pinterest',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'svelte',
			Name: 'Svelte',
			NewMessage: true,
			NumberOfNewMessage: 3,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'tea',
			Name: 'Tea',
			NewMessage: true,
			NumberOfNewMessage: 345,
			Notification: true,
			NumberOfNotification: 999,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'linux',
			Name: 'Linux',
			NewMessage: true,
			NumberOfNewMessage: 3,
			Notification: true,
			NumberOfNotification: 90,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		}
	];
	const AccordImg = new URL('../lib/icons/accord.svg', import.meta.url).href;
	
</script>

<svelte:head>
	<title>Accord</title>
</svelte:head>

<!-- markup (zero or more items) goes here -->

<!-- {#if ChatOrDockValue} -->
<!-- content here -->

<!-- hash bar -->
<!-- <Servers/> -->
<!-- {:else} -->
<div class="flex h-screen w-full flex-row  ">
	<Dock ServerList={x2} />

	<Chat />
	<!-- else content here -->
	<!-- {/if} -->

	<!-- <Dock2/> -->

	<div class=" flex h-screen w-full flex-col   flex-nowrap overflow-hidden bg-[#36393f] ">
		<!-- nav bar -->
		<NavBar FrndPro={Userdata} />
		
			<!-- <div class="flex flex-row h-screen w-full transition-all duration-500 ease-linear ">
		<slot />
	</div> -->

			<!-- message and people -->
			<div class=" h-full w-full overflow-y-scroll overflow-x-hidden text-3xl text-white scrol6">
				<Home />
			</div>
	</div>
</div>
