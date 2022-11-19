<script lang="ts">
	import { goto } from '$app/navigation';
	import { ChatOrDock, UserProData, ActiveFrndData,UserSettingSelect } from '$lib/store2';
	import Accord from '$lib/Dock/icons/accord.svelte';
	import AccordLogo from '$lib/Dock/icons/accord.svelte';
	// import {ServerList} from "$lib/store2"

	import { fade, blur, fly, slide, scale } from 'svelte/transition';
	import Plus from '$lib/icons/plus.svelte';
	import Wallet3 from '$lib/icons/wallet3.svelte';
	import Wallet1 from '$lib/icons/wallet1.svelte';
	import FriendReq from '$lib/icons/FriendReq.svelte';
	import Frnds from '$lib/icons/Frnds.svelte';
	import Addfriends from '$lib/icons/addfriends.svelte';
	import AddManyFrnd from '$lib/icons/addManyFrnd.svelte';

	export let ServerList: any;
	let ChatOrDockHelper: number;
	ChatOrDock.subscribe((val) => {
		ChatOrDockHelper = val;
	});
	function ChatOrDockFunc() {
		ActiveFrndData.set($UserProData)
		if (ChatOrDockHelper != 0) {
			ChatOrDock.update((n) => (n = 0));
		} else {
			ChatOrDock.update((n) => (n = 1));
		}
		goto('/', {
			replaceState: true,
			noscroll: true,
			keepfocus: true
		});
	}
	let ActiveServer = '';
	let HoverServer = '';
	let title: string | undefined = 'Accord';

	function OnClickServer(e: {
		ServerURL: any;
		Name?: string;
		NewMessage?: boolean;
		NumberOfNewMessage?: number;
		Notification?: boolean;
		NumberOfNotification?: number;
		ServerImage?: string;
	}) {
		title = e.Name;

		ActiveServer = e.ServerURL;

		goto('/s/' + e.ServerURL, {
			replaceState: true,
			noscroll: true,
			keepfocus: true
		});
	}

	let showMenu = false;
	async function onRightClick(e: any) {
		if (showMenu) {
			showMenu = false;
			await new Promise((res) => setTimeout(res, 100));
		}
	}
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<!-- dock folder 48x48 and small icons 16x16 -->
<div
	class=" no-scrollbar flex   min-w-[72px] max-w-[72px] flex-col items-center  overflow-y-scroll bg-[#202225]  "
>
	<!-- company icon button -->
	<div
		class="avatar   my-1  flex h-[48px]  w-[48px] cursor-pointer   justify-center    object-cover  "
		on:click={ChatOrDockFunc}
	>
		<AccordLogo class="h-10  w-10  transition-all duration-100 ease-linear  hover:scale-125 " />
		<!-- <div class=" w-14 h-14 hover:rounded-xl   rounded-3xl active:rounded-md hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer {false ? "ring  ring-offset-base-100  ring-blue-500" : "" }">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16"><path fill="#FFC107" d="M10.376.985a1.426 1.426 0 0 0-2.711.881l3.685 11.339a1.426 1.426 0 0 0 1.734.884c.756-.218 1.218-1.021.976-1.764L10.376.985z"/><path fill="#4CAF50" d="M4.666 2.841a1.425 1.425 0 1 0-2.711.881L5.64 15.06a1.426 1.426 0 0 0 1.734.884c.756-.218 1.218-1.021.976-1.764L4.666 2.841z"/><path fill="#EC407A" d="M15.015 10.376a1.425 1.425 0 1 0-.881-2.711L2.795 11.351a1.426 1.426 0 0 0-.884 1.734c.218.756 1.021 1.218 1.764.976l11.34-3.685z"/><path fill="#472A49" d="M5.158 13.579l2.71-.881-.881-2.71-2.71.881.881 2.71z"/><path fill="#CC2027" d="M10.869 11.723l2.71-.881-.881-2.711-2.71.881.881 2.711z"/><path fill="#2196F3" d="M13.159 4.666a1.425 1.425 0 1 0-.881-2.711L.94 5.64a1.424 1.424 0 0 0-.884 1.734c.218.756 1.021 1.217 1.764.976l11.339-3.684z"/><path fill="#1A937D" d="M3.302 7.868l2.711-.881-.881-2.71-2.71.881.88 2.71z"/><path fill="#65863A" d="M9.013 6.013l2.711-.881-.881-2.711-2.711.881.881 2.711z"/></svg>
        </div> -->
	</div>
	<!-- user list -->
	<!--{#each ServerList as u}-->
	<!--	<button-->
	<!--		on:contextmenu|preventDefault={onRightClick}-->
	<!--		on:click={() => {-->
	<!--			ActiveServer = u['ServerURL'];-->
	<!--			OnClickServer(u);-->
	<!--		}}-->
	<!--		class="  group flex w-full  flex-row  items-center justify-center py-[5px]  transition-all duration-150 ease-linear hover:rounded-lg"-->
	<!--	>-->
	<!--		<div class=" relative flex items-center justify-center">-->
	<!--			<div-->
	<!--				in:scale={{ duration: 200 }}-->
	<!--				out:scale-->
	<!--				class="absolute -left-[58px]  bg-white transition-all duration-200 ease-linear    {ActiveServer ===-->
	<!--				u['ServerURL']-->
	<!--					? 'rounded-md h-[50px] w-[50px]'-->
	<!--					: 'rounded-xl h-[10px] w-[50px] group-hover:rounded-md group-hover:h-[30px] group-hover:w-[50px] '}"-->
	<!--			/>-->

	<!--			<img-->
	<!--				src={u['ServerImage']}-->
	<!--				alt=""-->
	<!--				class="  h-[48px] w-[48px] cursor-pointer rounded-[1.25rem] object-cover transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md  "-->
	<!--			/>-->
	<!--			{#if u['Notification']}-->
	<!--				<span-->
	<!--					class=" absolute -bottom-2  -right-2 rounded-full border-2 border-[#202225] bg-red-500  px-1 text-xs font-semibold text-white "-->
	<!--				>-->
	<!--					{u['NumberOfNotification']}-->
	<!--				</span>-->
	<!--			{/if}-->
	<!--		</div>-->
	<!--	</button>-->
	<!--{/each}-->

	<!-- create button -->
	<button
		on:click={() => {
			goto('/s/create');
		}}
		class="  group flex w-full  flex-row  items-center justify-center py-[5px]  transition-all duration-150 ease-linear hover:rounded-lg"
	>
		<div class=" relative flex items-center justify-center">
			<!-- <svg  class="fill-white  -left-[63px] -top-1  absolute  h-16 w-16 " viewBox="0 0 448 512" xmlns="http://www.w3.org/2000/svg"><path d="M448 95.1v320c0 35.35-28.65 64-64 64H64c-35.35 0-64-28.65-64-64v-320c0-35.35 28.65-63.1 64-63.1h320C419.3 31.1 448 60.65 448 95.1z"/></svg> -->
			<div
				in:scale={{ duration: 200 }}
				out:scale
				class="absolute -left-[58px] bg-transparent  transition-all duration-200 ease-linear group-hover:bg-white    {ActiveServer ===
				'create'
					? 'rounded-md h-[50px] w-[50px]'
					: 'rounded-xl h-[10px] w-[50px] group-hover:rounded-md group-hover:h-[30px] group-hover:w-[50px] '}"
			/>

			<Plus />
		</div>
	</button>
	<!-- Wallet -->
	<button
		on:click={() => {
			UserSettingSelect.set("Payment Mangement")
			goto(`/${$UserProData.UserID}/profile/settings`);
		}}
		class="  group flex w-full  flex-row  items-center justify-center py-[5px]  transition-all duration-150 ease-linear hover:rounded-lg"
	>
		<div class=" relative flex items-center justify-center">
			<!-- <svg  class="fill-white  -left-[63px] -top-1  absolute  h-16 w-16 " viewBox="0 0 448 512" xmlns="http://www.w3.org/2000/svg"><path d="M448 95.1v320c0 35.35-28.65 64-64 64H64c-35.35 0-64-28.65-64-64v-320c0-35.35 28.65-63.1 64-63.1h320C419.3 31.1 448 60.65 448 95.1z"/></svg> -->
			<div
				in:scale={{ duration: 200 }}
				out:scale
				class="absolute -left-[58px] bg-transparent  transition-all duration-200 ease-linear group-hover:bg-white    {ActiveServer ===
				'create'
					? 'rounded-md h-[50px] w-[50px]'
					: 'rounded-xl h-[10px] w-[50px] group-hover:rounded-md group-hover:h-[30px] group-hover:w-[50px] '}"
			/>

			<Wallet1 class="h-[50px] w-[50px]" />
		</div>
	</button>
	<!-- friend Req -->
	<button
		on:click={() => {
			goto('/friendreq');
		}}
		class="  group flex w-full  flex-row   items-center justify-center py-[5px]  transition-all duration-150 ease-linear hover:rounded-lg"
	>
		<div class=" relative flex items-center justify-center">
			<!-- <svg  class="fill-white  -left-[63px] -top-1  absolute  h-16 w-16 " viewBox="0 0 448 512" xmlns="http://www.w3.org/2000/svg"><path d="M448 95.1v320c0 35.35-28.65 64-64 64H64c-35.35 0-64-28.65-64-64v-320c0-35.35 28.65-63.1 64-63.1h320C419.3 31.1 448 60.65 448 95.1z"/></svg> -->
			<div
				in:scale={{ duration: 200 }}
				out:scale
				class="absolute -left-[58px] bg-transparent  transition-all duration-200 ease-linear group-hover:bg-white    {ActiveServer ===
				'create'
					? 'rounded-md h-[50px] w-[50px]'
					: 'rounded-xl h-[10px] w-[50px] group-hover:rounded-md group-hover:h-[30px] group-hover:w-[50px] '}"
			/>
			<AddManyFrnd class="h-[50px] w-[50px] p-1   " />
			<!-- <Wallet3 class="h-[50px] w-[50px] p-1" /> -->
		</div>
	</button>
	<!-- friend List -->
	<button
		on:click={() => {
			goto('/friendlist');
		}}
		class="  group flex w-full  flex-row   items-center justify-center py-[5px]  transition-all duration-150 ease-linear hover:rounded-lg"
	>
		<div class=" relative flex items-center justify-center">
			<!-- <svg  class="fill-white  -left-[63px] -top-1  absolute  h-16 w-16 " viewBox="0 0 448 512" xmlns="http://www.w3.org/2000/svg"><path d="M448 95.1v320c0 35.35-28.65 64-64 64H64c-35.35 0-64-28.65-64-64v-320c0-35.35 28.65-63.1 64-63.1h320C419.3 31.1 448 60.65 448 95.1z"/></svg> -->
			<div
				in:scale={{ duration: 200 }}
				out:scale
				class="absolute -left-[58px] bg-transparent  transition-all duration-200 ease-linear group-hover:bg-white    {ActiveServer ===
				'create'
					? 'rounded-md h-[50px] w-[50px]'
					: 'rounded-xl h-[10px] w-[50px] group-hover:rounded-md group-hover:h-[30px] group-hover:w-[50px] '}"
			/>
			<Frnds class="h-[50px] w-[50px] p-1   " />
			<!-- <Wallet3 class="h-[50px] w-[50px] p-1" /> -->
		</div>
	</button>
	<!-- AdminMoneyManagement -->
	{#if $UserProData.Accounttype === 'admin'}
		<!-- content here -->
		<button
			on:click={() => {
				goto('/moneymanagement');
			}}
			class="  group flex w-full  flex-row  items-center justify-center py-[5px]  transition-all duration-150 ease-linear hover:rounded-lg"
		>
			<div class=" relative flex items-center justify-center">
				<!-- <svg  class="fill-white  -left-[63px] -top-1  absolute  h-16 w-16 " viewBox="0 0 448 512" xmlns="http://www.w3.org/2000/svg"><path d="M448 95.1v320c0 35.35-28.65 64-64 64H64c-35.35 0-64-28.65-64-64v-320c0-35.35 28.65-63.1 64-63.1h320C419.3 31.1 448 60.65 448 95.1z"/></svg> -->
				<div
					in:scale={{ duration: 200 }}
					out:scale
					class="absolute -left-[58px] bg-transparent  transition-all duration-200 ease-linear group-hover:bg-white    {ActiveServer ===
					'create'
						? 'rounded-md h-[50px] w-[50px]'
						: 'rounded-xl h-[10px] w-[50px] group-hover:rounded-md group-hover:h-[30px] group-hover:w-[50px] '}"
				/>

				<Wallet3 class="h-[50px] w-[50px] p-1" />
			</div>
		</button>
	{/if}

	<!-- download button -->
	<div class="avatar  py-1 px-2">
		<div
			class=" flex h-[50px] w-[50px] cursor-pointer items-center justify-center rounded-3xl bg-gray-700 transition-all duration-150  ease-linear hover:rounded-xl hover:border-[1px]  hover:border-[#3ba55d] hover:transition-all hover:duration-150 hover:ease-linear active:rounded-md "
		>
			<svg
				class="h-8 w-8 stroke-[#3ba55d] stroke-[1px]"
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
				xmlns="http://www.w3.org/2000/svg"
				><path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"
				/></svg
			>
		</div>
	</div>
</div>

<style>
</style>
