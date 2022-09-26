
<script lang="ts">
    import Servers from "$lib/Servers/index.svelte"
	import NavBar from '$lib/Navbar/index.svelte';
    import Peoples from '$lib/Peoples/index.svelte';
	import Messages from '$lib/Messages/index.svelte';
	import { showPeopleList } from '$lib/store2';

  
    let MyPro= {
			Name: "Sk Shahriar Ahmed Raka",
			Userid: 'skraka',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
			CoverImage:"https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png"
		};
    let ServerPro={
    "Name": "Competitive Programming for Pro",
    "ProfileUrl": "ace",
    "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1642489744/samples/1007550_t0uscy.jpg"
}
let  ServerMsg = [
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

    let showPeopleListValue: number;
	
	showPeopleList.subscribe((val) => {
		showPeopleListValue = val;
	});
    let PeopleData={
        "image":ServerPro.ProfileImage,
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
    
</script>

<style>
    /* your styles go here */
</style>

<!-- <svelte:head>
	<title> Accord server </title>
</svelte:head> -->

<Servers MyPro={MyPro}/>

<div class=" w-full h-screen overflow-hidden bg-stone-600 flex flex-col flex-nowrap ">
    <!-- nav bar -->
    <NavBar  MyPro={ServerPro} />

    <!-- message and people -->
    <div class="  w-full h-full  bg-[#2f3136] flex flex-row overflow-hidden">
        <!-- messaging 830px -->
        <Messages FriendMsg={ServerMsg} MyPro={MyPro} />
        <!-- people -->
        {#if showPeopleListValue != 0}
            <Peoples PeopleData={PeopleData}/>
        {/if}
    </div>
</div>