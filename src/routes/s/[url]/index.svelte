<script lang="ts" context="module">
    export async function load ({stuff}:any){
        let ServerPro={
    "Name": "Competitive Programming for Pro",
    "ProfileUrl": "ace",
    "ProfileImage": "https://res.cloudinary.com/dqo0ssnti/image/upload/v1642489744/samples/1007550_t0uscy.jpg"
}
let  ServerMsg = [
			{id:0, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:1, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:2, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:3, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:4, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:5, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:6, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:7, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:8, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:9, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:10, time:"20 jun22", writer: 'Sk Shahriar Ahmed Raka', message: "hello brother , what's you doing" },
			// {id:10.1, time:"20 jun22", writer: 'Sk Shahriar Ahmed Raka', message: "hello brother , what's you doing" },
			{id:11, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:12, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:13, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doingSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed Raka" },
			{id:14, time:"20 jun22", writer: 'Sk Shahriar Ahmed Raka', message: "hello brother , what's you doing" },
			{id:15, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:16, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:17, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:18, time:"20 jun22", writer: 'Sk Shahriar Ahmed Raka', message: "hello brother , what's you doing" },
			{id:19, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:20, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:21, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			{id:22, time:"20 jun22", writer: 'Sk Shahriar Ahmed Raka', message: "hello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doing" },
			{id:23, time:"20 jun22", writer: 'ace', message: "hello brother , what's you doing" },
			
		]
        

        
        return {
            props:{
                
                ServerPro:ServerPro,
                ServerMsg:ServerMsg,
                MyPro:stuff.MyPro
            }
        }
    }
</script>
<script lang="ts">
    import Servers from "$lib/_Servers/index.svelte"
	import NavBar from '$lib/_Navbar/index.svelte';
    import Peoples from '$lib/_Peoples/index.svelte';
	import Messages from '$lib/_Messages/index.svelte';
	import { showPeopleList } from '$lib/store2';

    export let ServerPro:any 
    export let ServerMsg:any 
    export let MyPro:any

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