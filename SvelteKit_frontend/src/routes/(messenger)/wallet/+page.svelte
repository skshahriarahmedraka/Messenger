

<script lang="ts">
	import Dock from '$lib/Dock/index.svelte';
	import Servers from '$lib/Servers/index.svelte';
	import NavBar from '$lib/Navbar/index.svelte';
	import Peoples from '$lib/Peoples/index.svelte';
	import Messages from '$lib/Messages/index.svelte';
	import Chat from '$lib/Chats/index.svelte';

	import { showPeopleList } from '$lib/store2';
	let showPeopleListValue: number;
	// let ChatOrDockValue: number;
	// export let FriendData: any;
	// export let FriendMsg: any;
    // let MyPro= {
	// 		Name: "Sk Shahriar Ahmed Raka",
	// 		Userid: 'skraka',
	// 		ProfileImage:
	// 			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
	// 		CoverImage:"https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png"
	// 	};

    let FriendData = {
			Name: 'Portgas D. Ace',
			ProfileUrl: 'ace',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1642489744/samples/1007550_t0uscy.jpg'
		};

		let FriendMsg = [
			{
				Conversation: {
					id: 123,
					members: ['skraka', 'ace']
				}
			},
			[
				{
					id: 0,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',
					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 1,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',
					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 2,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 3,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 4,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 5,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 6,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 7,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 8,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 9,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 10,
					time: '20 jun22',
					writer: 'skraka',
					writerName: 'Sk Shahriar Ahmed Raka',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				// {id:10.1, time:"20 jun22", writer: 'Sk Shahriar Ahmed Raka', message: "hello brother , what's you doing" },
				{
					id: 11,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 12,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 13,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message:
						"hello brother , what's you doingSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed Raka",

					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 14,
					time: '20 jun22',
					writer: 'skraka',
					writerName: 'Sk Shahriar Ahmed Raka',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 15,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 16,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },

					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 17,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 18,
					time: '20 jun22',
					writer: 'skraka',
					writerName: 'Sk Shahriar Ahmed Raka',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 19,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 1, skraka: 2 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 20,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 7, skraka: 1 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 21,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 5, skraka: 4 },
					numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 22,
					time: '20 jun22',
					writer: 'skraka',
					writerName: 'Sk Shahriar Ahmed Raka',

					message:
						"hello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doing",
					react: { ace: 2, skraka: 9, fr: 3, dfz: 2, sg: 3 },
					numberOfReact: [1, 1, 2, 3, 4, 5, 6, 7, 8, 9]
				},
				{
					id: 23,
					time: '20 jun22',
					writer: 'ace',
					writerName: 'Portgas D. Ace',

					message: "hello brother , what's you doing",
					react: { ace: 0, skraka: 1, sk: 2, f: 3, g: 4, p: 5, o: 6, i: 7, u: 8, w: 9 },
					numberOfReact: [1, 1, 2, 3, 4, 5, 67435, 7, 8, 99]
				}
			]
		];
	showPeopleList.subscribe((val) => {
		showPeopleListValue = val;
	});
	let x = [
		{
			ProfileURL: 'majibarrahman',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Md Majibar Rahman',
			LastMessage: 'hope you are filling well',
			LastMessageTime: '23 Aug 21',
			SilentNotification: true,
			NumberOfNotification: 43,
			ActiveStatus: true,
			LastActiveTime: '10:28 AM'
		},
		{
			ProfileURL: 'rafihasan',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'md rafi hasan zihad',
			LastMessage: 'whats is going on bro',
			LastMessageTime: '9:30 AM',
			SilentNotification: false,
			NumberOfNotification: 3,
			ActiveStatus: true,
			LastActiveTime: '6 Jul 17'
		},
		{
			ProfileURL: 'ekbalhasan',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Md ekbal Hasan',
			LastMessage: 'I was going home that day',
			LastMessageTime: '9 jun 11',
			SilentNotification: true,
			NumberOfNotification: 4,
			ActiveStatus: false,
			LastActiveTime: '29 Jun 18'
		},
		{
			ProfileURL: 'mdronok',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Md ronok shahriar',
			LastMessage: 'ethical hacking is my passion',
			LastMessageTime: '10:34 PM',
			SilentNotification: false,
			NumberOfNotification: 1,
			ActiveStatus: false,
			LastActiveTime: '10:28 AM'
		},
		{
			ProfileURL: 'samiulshakil',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'md samiul alim shakil',
			LastMessage: 'I am feeling wanderful . wona go for a tour',
			LastMessageTime: '12 mar 19',
			SilentNotification: true,
			NumberOfNotification: 7,
			ActiveStatus: true,
			LastActiveTime: '7 Aug 19'
		},
		{
			ProfileURL: 'hemelakash',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Hemel sarkar akash',
			LastMessage: 'Got a job in Apple as a CEO',
			LastMessageTime: '12 min',
			SilentNotification: true,
			NumberOfNotification: 6,
			ActiveStatus: false,
			LastActiveTime: '10:28 AM'
		},
		{
			ProfileURL: 'hazem',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Md Azim Mahfuz',
			LastMessage: 'what a great day . got a job on Uniliver',
			LastMessageTime: '23 sec',
			SilentNotification: true,
			NumberOfNotification: 7,
			ActiveStatus: true,
			LastActiveTime: '1:06 AM'
		},

		{
			ProfileURL: 'rafel',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Md rafel',
			LastMessage: 'bought a new bike ,lets to on a ride',
			LastMessageTime: '4:23 AM',
			SilentNotification: false,
			NumberOfNotification: 0,
			ActiveStatus: false,
			LastActiveTime: '10:28 AM'
		},
		{
			ProfileURL: 'rakibulislam',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'md Rakibul islam',
			LastMessage: "whats's you doing , bro , see you later",
			LastMessageTime: '23 dec 21',
			SilentNotification: true,
			NumberOfNotification: 2,
			ActiveStatus: true,
			LastActiveTime: '00:28 AM'
		},
		{
			ProfileURL: 'istique',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'istique mahmud pritom',
			LastMessage: 'i am on a train , going to Dhaka',
			LastMessageTime: '1 jun 21',
			SilentNotification: false,
			NumberOfNotification: 0,
			ActiveStatus: true,
			LastActiveTime: '11:28 AM'
		},
		{
			ProfileURL: 'mdrakib',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Md Rakib Hasan',
			LastMessage: 'lets do party to night . Real madrid vs PSG . the game will be epic',
			LastMessageTime: '9 mar',
			SilentNotification: true,
			NumberOfNotification: 2,
			ActiveStatus: true,
			LastActiveTime: '00:00 AM'
		},
		{
			ProfileURL: 'toki',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'md toki',
			LastMessage: 'whats up bro what is happening there',
			LastMessageTime: '21 jun',
			SilentNotification: false,
			NumberOfNotification: 0,
			ActiveStatus: false,
			LastActiveTime: '9 Feb 30'
		},
		{
			ProfileURL: 'yurichi',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Yurichi',
			LastMessage: 'are you enjoying my story',
			LastMessageTime: '12 jan 16',
			SilentNotification: false,
			NumberOfNotification: 9,
			ActiveStatus: false,
			LastActiveTime: '23 Dec 21'
		},
		{
			ProfileURL: 'rengokhu',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'kyjiro Rengokhu',
			LastMessage: 'mugen the train was amayzing',
			LastMessageTime: '19 feb 19',
			SilentNotification: false,
			NumberOfNotification: 5,
			ActiveStatus: true,
			LastActiveTime: '10:28 PM'
		},
		{
			ProfileURL: 'vegeta',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Prince Vegeta',
			LastMessage: 'the sayian pride will never parish',
			LastMessageTime: '9 sep 20',
			SilentNotification: false,
			NumberOfNotification: 0,
			ActiveStatus: false,
			LastActiveTime: '10:28 AM'
		},

		{
			ProfileURL: 'saad',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Hasan mahmud',
			LastMessage: "what's your physical condition , come to my home ",
			LastMessageTime: 'yestarday',
			SilentNotification: false,
			NumberOfNotification: 5,
			ActiveStatus: true,
			LastActiveTime: '3:00 AM'
		},
		{
			ProfileURL: 'hankok',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Boa Hankok',
			LastMessage: 'where is my luffy',
			LastMessageTime: '12 jan 00',
			SilentNotification: false,
			NumberOfNotification: 99,
			ActiveStatus: false,
			LastActiveTime: '12:28 AM'
		},
		{
			ProfileURL: 'nejuko',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Nejuko',
			LastMessage: 'hmm',
			LastMessageTime: '2 jan 34',
			SilentNotification: true,
			NumberOfNotification: 9,
			ActiveStatus: true,
			LastActiveTime: 'Active'
		},
		{
			ProfileURL: 'shanks',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'D. Shanka',
			LastMessage: 'after shreding tears you will be a real man',
			LastMessageTime: '1:00 AM',
			SilentNotification: false,
			NumberOfNotification: 45,
			ActiveStatus: false,
			LastActiveTime: '10:00 AM'
		},
		{
			ProfileURL: 'admiral',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Nevy Admiral',
			LastMessage: 'justice must be absolute',
			LastMessageTime: '6:39 AM',
			SilentNotification: false,
			NumberOfNotification: 0,
			ActiveStatus: false,
			LastActiveTime: '9:28 AM'
		},
		{
			ProfileURL: 'law',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060642/samples/law_hc6hfb.png',
			UserName: 'Trafalger D. Water Law',
			LastMessage: '     i must find the meaning of meaning of D.',
			LastMessageTime: 'yesterday',
			SilentNotification: false,
			NumberOfNotification: 34,
			ActiveStatus: true,
			LastActiveTime: 'Active'
		}
	];

	let x3 = {
		Name: 'Portgas D. Ace',
		ProfileUrl: 'ace',
		ProfileImage:
			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg'
	};

	let PeopleData = {
		image: FriendData['ProfileImage'],
		username: 'Sk Shahriar Ahmed Raka',
		lastseen: '9 jun22',
		mobile: '+8801838810189',
		bio: 'Lazy is such an ugly word, I prefer to selectively participate. I already want to take a nap tomorrow. ',
		url: 'shahriarraka',
		notification: true,
		contents: [
			[213, 'images'],
			[41, 'videos'],
			[5, 'files'],
			[6, 'shared links'],
			[41, 'audios'],
			[64, 'gifs'],
			[23, 'recomendedserver'],
			[67, 'customizeChat'],
			[67, 'ChangeTheme'],

			[4, 'Changeimoji'],
			[675, 'EditNickname'],
			[63, 'shareThisContact'],
			[54, 'DelteContact'],
			[436, 'DeleteHistory'],
			[75, 'DeleteChat'],
			[865, 'BlockUser']
		],
		recomendationlist: [[]],
		cutomize: {
			color: 'blue',
			emoji: 'smile',
			nickname: 'ssar'
		}
	};
	
	// import { io } from "socket.io-client";

	// const socket = io("ws://127.0.0.1:8888/socketio/");
	// console.log("🚀 ~ file: +page.svelte ~ line 589 ~ socket", socket)

</script>

<!-- <Chat FriendList={x}  /> -->

<div
	class=" flex  h-screen w-full overflow-hidden bg-[#36393f] md:flex-col md:overflow-y-scroll lg:flex-row "
>
	<div class=" no-scrollbar h-fit  justify-center md:w-full lg:w-1/2 lg:overflow-y-scroll mb-6  ">
		Wallet 
	</div>
</div>


<style>
	
</style>
