<script lang="ts">
	// import userBanner from './_images/whiteBeard.jpg';
	// import Moji2 from "./_images/moji2.jpg"
	// import { MyPro} from "$lib/store2"
	// import Messaging from './icons/message.svelte';
	import Microphone from './icons/microphone.svelte';
	import Video from './icons/video.svelte';
	// import Announcement from './icons/announcement.svelte';
	// import Hash from './icons/hash.svelte';
	// import Rule from './icons/rule.svelte';
	// import InvitePeople from './icons/invitePeople.svelte';
	// import AeroUp from './icons/aeroUp.svelte';
	// import AeroDown from './icons/aeroDown.svelte';
	// import AudioCall from '$lib/Navbar/profileImg/audioCall.svelte';
	import RedOff from './icons/redOff.svelte';
	import Speaker from './icons/Speaker.svelte';
	import Notification from './icons/notification.svelte';
	import Settings from './icons/settings.svelte';
	import Search from './icons/search.svelte';

	import Cross from '$lib/Navbar/profileImg/Cross.svelte';
import { goto } from '$app/navigation';
	export let MyPro: any;
	let searchIconEvent: boolean = false;
	let inputValue: string = '';
	export let showUsername: boolean;
	let userOptions = { video: false, microphone: false, notificationSound: false, sound: false };
	let SearchStyle: string = 'lg'; // sm,md,lg
	let showUserid: boolean = true;
	let showUserOptions = false;
	import Coin from '$lib/icons/coin2.svelte';
</script>

{#if showUserOptions}
	<div
		class=" mb-5 flex h-fit w-full flex-row flex-wrap   rounded-xl bg-[#18191c] bg-opacity-80 text-gray-300 "
		on:blur={() => {
			showUserOptions = false;
		}}
	>
		<div class="h-40 w-full">
			<img src={MyPro.CoverImage} alt="coverImage" class="h-40 w-full rounded-xl object-cover " />
			<img
				src={MyPro.ProfileImage}
				alt="coverImage"
				class="fixed -mt-7 ml-5 h-12 w-12 rounded-xl object-cover"
			/>
		</div>
		<p class="  ml-20 mt-6 self-center text-lg">{MyPro.Name}</p>
		<!-- <div class=" flex flex-col ">

			<div
			class=" flex mt-12  ml-20   w-full flex-row items-center gap-2 rounded-2xl border-0 border-blue-500 bg-sky-300 bg-opacity-30 pr-2  sm:hidden md:hidden lg:hidden xl:hidden xs:hidden xxl:contents"
		>
			<Coin class="storke-[1px] h-8 w-8 stroke-yellow-300" />
			<p class="text-slate-200 ">{826753}</p>
		</div>
		</div> -->
		<div class="w-full h-16 inline-flex mt-5">
			<button on:click={() => {goto(`/${MyPro.Userid}/profile`)}} class=" h-16 w-1/2 rounded-lg flex justify-center items-center  hover:bg-gray-900 "
				>
				<Coin class="storke-[1px] h-8 w-8 stroke-yellow-300" />
			<p class="text-slate-200 ">{826753}</p>
				</button
			>
			<button on:click={() => {goto(`/${MyPro.Userid}/settings`)}} class=" h-16  w-1/2 rounded-lg  hover:bg-gray-900 "
				>Settings</button
			>
		</div>
		<div class="w-full h-16 inline-flex">
			<button on:click={() => {goto(`/${MyPro.Userid}/profile/edit`)}} class=" h-16 w-1/2 rounded-lg hover:bg-gray-900 "
				>Edit Profile</button
			>
			<button on:click={() => {goto(`/${MyPro.Userid}/settings/edit`)}} class=" h-16 w-1/2 rounded-lg  hover:bg-gray-900 "
				>Edit Profile</button
			>

		</div>
	
	</div>
{/if}
<div class="flex h-14 w-full flex-row items-center bg-[#292b2f]  ">
	{#if SearchStyle === 'sm'}
		<!-- WRITE SEARCH -->
		<div
			class="  flex flex-grow cursor-pointer flex-row items-center justify-center text-lg text-gray-400 transition-all duration-300 ease-linear"
		>
			<input
				placeholder="Search"
				autocomplete="off"
				on:blur={() => {
					inputValue.trim() === '' ? (SearchStyle = 'lg') : (searchIconEvent = searchIconEvent);
				}}
				bind:value={inputValue}
				type="text"
				class="    mx-1  my-0 h-10  w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e]  outline-none  focus:border-sky-500 active:border-gray-800"
			/>
			<div
				on:click={() => {
					SearchStyle = 'lg';
					inputValue = '';
				}}
				class=""
			>
				<Cross
					class="mr-2 h-8 w-8 place-content-center  rounded-xl fill-slate-400 hover:fill-white"
				/>
			</div>
		</div>
	{:else if SearchStyle === 'md'}
		<div class=" mr-2 rounded-lg  hover:bg-gray-700">
			<Search
				class="m-1 h-6 w-6  cursor-pointer fill-slate-400 transition-all duration-100 ease-linear hover:fill-white"
			/>
		</div>
		<div class=" mr-2 rounded-lg  hover:bg-gray-700">
			<Search
				class="m-1 h-6 w-6  cursor-pointer fill-slate-400 transition-all duration-100 ease-linear hover:fill-white"
			/>
		</div>
	{:else if SearchStyle === 'lg'}
		<!-- USER PROFILE -->
		{#if showUsername}
			<div class="   mx-2 h-10 min-w-[2.5rem] ">
				<img
					on:click={() => {
						showUserOptions = !showUserOptions;
					}}
					src={MyPro.ProfileImage}
					alt=""
					class=" h-full  w-full cursor-pointer rounded-2xl   object-cover transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md"
				/>
			</div>
			<div class="flex flex-col  ">
				<p class=" overflow-hidden text-base text-slate-300 line-clamp-1 ">{MyPro.Name}</p>
				{#if showUserid}
					<!-- content here -->
					<p
						on:click={() => {
							navigator.clipboard.writeText('http://localhost:3000/' + MyPro.Userid);
							showUserid = false;
						}}
						class=" cursor-pointer text-sm text-gray-500 line-clamp-1 hover:text-gray-400"
					>
						@{MyPro.Userid}
					</p>
				{:else}
					<p
						on:click={() => {
							navigator.clipboard.writeText('http://localhost:3000/' + MyPro.Userid);
						}}
						class="{setTimeout(() => {
							showUserid = true;
						}, 1000)} cursor-pointer text-sm text-blue-500 line-clamp-1 "
					>
						Copied !!
					</p>
				{/if}
			</div>
			<div class=" flex-grow" />
		{:else}
			<div class="   mx-2 h-10 w-10 ">
				<img
					src={MyPro['ProfileImage']}
					alt=""
					class=" h-full  w-full cursor-pointer rounded-2xl   object-cover transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md"
				/>
			</div>

			<div class=" flex-grow" />
		{/if}
		<!-- video -->
		<div
			class=" rounded-lg hover:bg-gray-700"
			on:click={() => {
				userOptions['video'] = !userOptions['video'];
			}}
		>
			{#if !userOptions['video']}
				<RedOff class=" absolute  mt-1  mb-2 ml-1  h-6 w-6 fill-[#a4393c] " />
			{/if}
			<!-- <svg class=" absolute  fill-[#a4393c] stroke-2 h-4 w-4 my-4" viewBox="0 0 640 512" xmlns="http://www.w3.org/2000/svg"><path d="M24.03 0c5.156 0 10.37 1.672 14.78 5.109l591.1 463.1c10.44 8.172 12.25 23.27 4.062 33.7c-8.125 10.41-23.19 12.28-33.69 4.078L9.189 42.89c-10.44-8.172-12.25-23.27-4.062-33.7C9.845 3.156 16.91 0 24.03 0z"/></svg> -->
			<Video />
		</div>
		<!-- microphone -->
		<div
			class="rounded-lg hover:bg-gray-700 "
			on:click={() => {
				userOptions['microphone'] = !userOptions['microphone'];
			}}
		>
			{#if !userOptions['microphone']}
				<RedOff class=" absolute  mt-1  mb-2 ml-1  h-6 w-6 fill-[#a4393c] " />
			{/if}
			<Microphone
				class="m-1 h-6 w-6  cursor-pointer fill-slate-400 transition-all duration-100 ease-linear hover:fill-white"
			/>
		</div>
		<!-- sound -->
		<div
			class="rounded-lg hover:bg-gray-700 "
			on:click={() => {
				userOptions['sound'] = !userOptions['sound'];
			}}
		>
			{#if !userOptions['sound']}
				<RedOff class=" absolute  mt-1  mb-2 ml-1  h-6 w-6 fill-[#a4393c] " />
			{/if}
			<Speaker />
		</div>
		<!-- notification sound -->
		<div
			class="rounded-lg hover:bg-gray-700 "
			on:click={() => {
				userOptions['notificationSound'] = !userOptions['notificationSound'];
			}}
		>
			{#if !userOptions['notificationSound']}
				<RedOff class=" absolute  mt-1  mb-2 ml-1  h-6 w-6 fill-[#a4393c] " />
			{/if}
			<Notification />
			<!-- <svg class="w-8 h-8 my-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path></svg> -->
		</div>
		<!-- settings -->
		<!-- <div class=" hover:bg-gray-700 rounded-lg ">
        <Settings />
    </div> -->
		<div
			on:click={() => {
				SearchStyle = 'sm';
			}}
			class=" mr-2 rounded-lg  hover:bg-gray-700"
		>
			<Search
				class="m-1 h-6 w-6  cursor-pointer fill-slate-400 transition-all duration-100 ease-linear hover:fill-white"
			/>
		</div>
	{/if}
</div>

<style>
	/* your styles go here */
</style>
