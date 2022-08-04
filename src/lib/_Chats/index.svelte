<script context="module">
	// export async function load({fetch}){
	//     const res = await fetch("http.localhost:3000/")
	// }
</script>

<script lang="ts">
	import { goto } from '$app/navigation';

	// import { url } from "$app/stores";

	// import {page} from "$app/stores"
	// $: console.log("$page.params ",$page.params)
	import { ChatOrDock } from '$lib/store2';
	import Notification from '$lib/icons/notification.svelte';
	import Search from '$lib/_Navbar/_profileImg/search.svelte';
	import Cross from '$lib/_Navbar/_profileImg/Cross.svelte';
    import UserDashPannel from "$lib/UserDashPanel/index.svelte"

	// import {FriendList ,MyPro} from "$lib/store2"
	export let FriendList: any;
	export let MyPro: any;
    let SearchStyle:string
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
				class="  hover:text-white   w-full h-14 m-1  {ActiveChat === i['ProfileURL']
					? 'bg-slate-600 rounded-xl'
					: 'hover:bg-[#202225]'} flex flex-row hover:rounded-xl  "
			>
				<!-- active indecator -->

				<!-- user image -->
				<img
					src={i['ProfileImage']}
					alt=""
					class=" w-12  h-12 m-2 aspect-square  rounded-3xl hover:rounded-xl active:rounded-md object-cover hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer  active:ring  active:ring-offset-base-50  active:ring-blue-600"
				/>
				<div class="flex flex-col w-full h-full mr-2 ">
					<div class="   h-8  mt-2 ml-0 flex flex-row">
						<!-- User name -->
						<p class=" w-56  font-bold text-lg text-ellipsis truncate   ">
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
					<div class=" h-8 w-72 flex flex-row ">
						<!-- last message -->
						<p class="truncate w-72  h-full   ">
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
        <UserDashPannel MyPro={MyPro} bind:SearchStyle={SearchStyle} />
    </div>
	
</div>

<style>
	/* your styles go here */
</style>
