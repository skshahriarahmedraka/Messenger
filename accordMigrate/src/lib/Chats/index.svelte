

<script lang="ts">
	import { goto } from '$app/navigation';

	// import { url } from "$app/stores";

	// import {page} from "$app/stores"
	// $: console.log("$page.params ",$page.params)
	import { ChatOrDock } from '$lib/store2';
	import Notification from '$lib/icons/notification.svelte';
	// import Search from '$lib/Navbar/profileImg/search.svelte';
	// import Cross from '$lib/Navbar/profileImg/Cross.svelte';
    import UserDashPannel from "$lib/UserDashPanel/index.svelte"
    import { fade, blur, fly, slide, scale } from "svelte/transition";


	// import {FriendList ,MyPro} from "$lib/store2"
	export let FriendList: any;
	export let MyPro: any;

	let ChatOrDockHelper: number;
	ChatOrDock.subscribe((val) => {
		ChatOrDockHelper = val;
	});

	function ChatOrDockFunc() {
		goto('/', {
			replaceState: true,
			noscroll: true,
			keepfocus: true
		});
		if (ChatOrDockHelper != 0) {
			ChatOrDock.update((n) => (n = 0));
		} else {
			ChatOrDock.update((n) => (n = 1));
		}
	}
	let searchIconEvent: boolean = false;
	let inputValue: string = '';

	let ActiveChat: string = '';

	let title: string | undefined = 'Accord';

	function OnClickFriend(e: {
		ProfileURL: any;
		ProfileImage?: string;
		UserName?: string;
		LastMessage?: string;
		LastMessageTime?: string;
		SilentNotification?: boolean;
		NumberOfNotification?: number;
		ActiveStatus?: boolean;
		LastActiveTime?: string;
	}) {
		title = e.UserName;
		ActiveChat = e.ProfileURL;
		goto(e.ProfileURL, {
			replaceState: true,
			noscroll: true,
			keepfocus: true
		});
	}
    let showUsername=true 
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<div
	class=" min-w-[404px] max-w-[404px] h-screen overflow-hidden bg-[#2f3136] flex flex-col transition-all duration-500 ease-linear  "
>
	<div class=" h-full w-full overflow-y-auto scrol overflow-x-hidden  text-gray-300  ">
		<!-- user chat list -->
		{#each FriendList as i}
			<button
				on:click={() => {
					OnClickFriend(i);
				}}
				class="  hover:text-white   w-full h-16 m-1    {ActiveChat === i['ProfileURL']
					? 'bg-slate-600 rounded-xl'
					: 'hover:bg-[#202225]'} flex flex-row hover:rounded-xl  "
			>
				<!-- active indecator -->

				<!-- user image -->
				<!-- <img
					src={i['ProfileImage']}
					alt=""
					class=" w-14 h-14 m-2   rounded-3xl hover:rounded-xl active:rounded-md object-cover hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer  "
				/> -->
                <!-- USER IMAGE -->
                <button   class=" group w-16   flex  flex-row  my-2 ml-1 hover:rounded-lg transition-all duration-150 ease-linear">
            
                    <div class=" relative flex ">
                        <!-- <svg  class="fill-white  -left-[63px] -top-1  absolute  h-16 w-16 " viewBox="0 0 448 512" xmlns="http://www.w3.org/2000/svg"><path d="M448 95.1v320c0 35.35-28.65 64-64 64H64c-35.35 0-64-28.65-64-64v-320c0-35.35 28.65-63.1 64-63.1h320C419.3 31.1 448 60.65 448 95.1z"/></svg> -->
                        <!-- <div in:scale="{{  duration: 200 }}" out:scale class="bg-white -left-[58px]  absolute transition-all duration-200 ease-linear    {ActiveServer===u["ServerURL"] ? "rounded-md h-[50px] w-[50px]" : "rounded-xl h-[10px] w-[50px] group-hover:rounded-md group-hover:h-[30px] group-hover:w-[50px] " }"></div> -->
                        <!-- <svg in:scale="{{  duration: 200 }}" out:scale class="fill-white   -left-[55px] -top-[25px]   absolute  h-[50px] w-[50px] " viewBox="0 0 448 512" xmlns="http://www.w3.org/2000/svg"><path d="M384 32H64C28.65 32 0 60.66 0 96v320c0 35.34 28.65 64 64 64h320c35.35 0 64-28.66 64-64V96C448 60.66 419.3 32 384 32zM319.1 280h-192C114.8 280 103.1 269.2 103.1 256c0-13.2 10.8-24 24-24h192c13.2 0 23.1 10.8 23.1 24C343.1 269.2 333.2 280 319.1 280z"/></svg> -->
                        
                        <img src="{i.ProfileImage}" alt="" class="  rounded-[1.25rem] hover:rounded-xl active:rounded-md object-cover h-[48px] w-[48px]  transition-all duration-100  ease-linear cursor-pointer  ">
                        
                    </div>
            
                </button>
                <!-- NAME , LAST MESSAGE -->
				<div class="flex flex-col w-full h-full  mx-2 justify-center  ">
					<div class="     mt-2 ml-0 flex flex-row text-left">
						<!-- User name -->
						<p class=" w-56  font-bold text-base    ">
							{i['UserName']}
						</p>
						<div class=" flex-grow" />
						<!-- mute icon -->
						{#if i['SilentNotification']}
							<svg class="w-5 h-5 m-1 flex-none " fill="currentColor" viewBox="0 0 20 20"
								><path
									fill-rule="evenodd"
									d="M9.383 3.076A1 1 0 0110 4v12a1 1 0 01-1.707.707L4.586 13H2a1 1 0 01-1-1V8a1 1 0 011-1h2.586l3.707-3.707a1 1 0 011.09-.217zM12.293 7.293a1 1 0 011.414 0L15 8.586l1.293-1.293a1 1 0 111.414 1.414L16.414 10l1.293 1.293a1 1 0 01-1.414 1.414L15 11.414l-1.293 1.293a1 1 0 01-1.414-1.414L13.586 10l-1.293-1.293a1 1 0 010-1.414z"
									clip-rule="evenodd"
								/></svg
							>
						{/if}
						<!-- last message time -->
						<p class="   font-thin text-xs mt-3 mb-1 mx-0 flex-none  ">
							<!-- 34 aug21  -->
							{i['LastMessageTime']}
						</p>
					</div>
					<div class="  w-72 flex flex-row mb-2">
						<!-- last message -->
						<p class=" line-clamp-1 w-72  h-full  text-left text-sm ">
							{i['LastMessage']}
						</p>
						<!-- number of notification -->
						{#if i['SilentNotification'] && i['NumberOfNotification'] != 0}
							<span
								class="   h-5  px-1  border-2 border-[#202225] rounded-xl text-xs font-semibold bg-slate-500 text-white  "
							>
								{i['NumberOfNotification']}
							</span>
						{:else if i['NumberOfNotification'] != 0}
							<span
								class="   h-5  px-1  border-2 border-[#202225] rounded-xl text-xs font-semibold bg-red-500 text-white  "
							>
								{i['NumberOfNotification']}
							</span>
						{/if}
					</div>
				</div>
			</button>
		{:else}
			<!-- empty list -->
		{/each}
	</div>
    <div class="">
        <UserDashPannel MyPro={MyPro}  {showUsername}/>
    </div>
	
</div>

<style>
	/* your styles go here */
</style>
