<script lang="ts">
    throw new Error("@migration task: Add data prop (https://github.com/sveltejs/kit/discussions/5774#discussioncomment-3292707)");

    import Servers from "$lib/Servers/index.svelte"
	import NavBar from '$lib/Navbar/index.svelte';
    import Peoples from '$lib/Peoples/index.svelte';
	import Messages from '$lib/Messages/index.svelte';
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