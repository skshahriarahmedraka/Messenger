<script lang="ts" context="module">
	export async function load({}){
		let x3 ={
    "Name": "Portgas D. Ace",
    "ProfileUrl": "ace",
    "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg"
}
	return {
		props:{
			
		},
		stuff:{
			MyPro:x3
		}
	}
	}
</script>

<script lang="ts">
	import '../app.css';
    
	// import peopleIcon from "./_icons/users.svg"
	import Dock from '$lib/_Dock/index.svelte';
	import Servers from '$lib/_Servers/index.svelte';
	import NavBar from '$lib/_Navbar/index.svelte';
	import Peoples from '$lib/_Peoples/index.svelte';
	import Messages from '$lib/_Messages/index.svelte';
	import { showPeopleList, ChatOrDock } from '$lib/store2';

	import Chat from '$lib/_Chats/index.svelte';
	// export let MyPro:any 

	let showPeopleListValue: number;
	// let ChatOrDockValue: number;

	showPeopleList.subscribe((val) => {
		showPeopleListValue = val;
	});

	// ChatOrDock.subscribe((val) => {
	// 	ChatOrDockValue = val;
	// });

	let x2:any = [
    {   
        "ServerURL":"postgres",
        "Name": "Postgres",
        "NewMessage": true,
        "NumberOfNewMessage": 3,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"mongodb",
        "Name": "MongoDB",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": true,
        "NumberOfNotification": 3,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"docker",
        "Name": "Docker",
        "NewMessage": true,
        "NumberOfNewMessage": 4,
        "Notification": true,
        "NumberOfNotification": 6,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"kubernetes",
        "Name": "Kubernetes",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": true,
        "NumberOfNotification": 99,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"redis",
        "Name": "Redis",
        "NewMessage": true,
        "NumberOfNewMessage": 7,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"tensorflow",
        "Name": "Tensorflow",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"macity",
        "Name": "Mancity",
        "NewMessage": true,
        "NumberOfNewMessage": 2,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"cpp",
        "Name": "C++",
        "NewMessage": true,
        "NumberOfNewMessage": 4,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"ubuntu",
        "Name": "Ubuntu",
        "NewMessage": true,
        "NumberOfNewMessage": 7,
        "Notification": true,
        "NumberOfNotification": 2,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"realmadrid",
        "Name": "Real Madrid",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"bangladesh",
        "Name": "bangladesh",
        "NewMessage": true,
        "NumberOfNewMessage": 6,
        "Notification": true,
        "NumberOfNotification": 4,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"c",
        "Name": "C",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"coffee",
        "Name": "Coffee",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": true,
        "NumberOfNotification": 45,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"chelsea",
        "Name": "Chelsea",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"go",
        "Name": "Go",
        "NewMessage": true,
        "NumberOfNewMessage": 9,
        "Notification": true,
        "NumberOfNotification": 65,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"mazda",
        "Name": "Mazda",
        "NewMessage": false,
        "NumberOfNewMessage": 2,
        "Notification": true,
        "NumberOfNotification": 5,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"pinterest",
        "Name": "Pinterest",
        "NewMessage": false,
        "NumberOfNewMessage": 0,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"svelte",
        "Name": "Svelte",
        "NewMessage": true,
        "NumberOfNewMessage": 3,
        "Notification": false,
        "NumberOfNotification": 0,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"tea",
        "Name": "Tea",
        "NewMessage": true,
        "NumberOfNewMessage": 345,
        "Notification": true,
        "NumberOfNotification": 999,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    },
    {   
        "ServerURL":"linux",
        "Name": "Linux",
        "NewMessage": true,
        "NumberOfNewMessage": 3,
        "Notification": true,
        "NumberOfNotification": 90,
        "ServerImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg"
        
    }
]
</script>

<svelte:head>
	<title> Accord </title>
</svelte:head>

<div class="flex flex-row h-screen w-full transition-all duration-500 ease-linear ">
	<!-- left avater bar -->

	<Dock ServerList={x2} />

	<slot />

	<!-- everything -->
	
</div>

<style>
</style>
