<script lang="ts">
	// import userBanner from './_images/whiteBeard.jpg';
	// import Moji2 from "./_images/moji2.jpg"
	// import { MyPro} from "$lib/store2"
	// import Messaging from './_icons/message.svelte';
	import Microphone from './_icons/microphone.svelte';
	import Video from './_icons/video.svelte';
	// import Announcement from './_icons/announcement.svelte';
	// import Hash from './_icons/hash.svelte';
	// import Rule from './_icons/rule.svelte';
	// import InvitePeople from './_icons/invitePeople.svelte';
	// import AeroUp from './_icons/aeroUp.svelte';
	// import AeroDown from './_icons/aeroDown.svelte';
	// import AudioCall from '$lib/_Navbar/_profileImg/audioCall.svelte';
	import RedOff from './_icons/redOff.svelte';
	import Speaker from './_icons/Speaker.svelte';
	import Notification from './_icons/notification.svelte';
	import Settings from './_icons/settings.svelte';
	import Search from './_icons/search.svelte';

	import Cross from '$lib/_Navbar/_profileImg/Cross.svelte';
	export let MyPro: any;
	let searchIconEvent: boolean = false;
	let inputValue: string = '';
    export let showUsername:boolean
	let userOptions = { video: false, microphone: false, notificationSound: false, sound: false };
	let SearchStyle: string="lg" // sm,md,lg
</script>

<div class="w-full h-14 bg-[#292b2f] flex flex-row items-center ">
	{#if SearchStyle === 'sm'}
    <!-- WRITE SEARCH -->
		<div
			class="  flex-grow justify-center items-center flex flex-row transition-all duration-300 ease-linear cursor-pointer text-lg text-gray-400"
		>
			<input
				placeholder="Search"
				autocomplete="off"
				on:blur={() => {
					inputValue.trim() === '' ? (SearchStyle = 'lg') : (searchIconEvent = searchIconEvent);
				}}
				bind:value={inputValue}
				type="text"
				class="    w-full  self-center h-10  p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-1 my-0  border-[#24262b]  active:border-gray-800 rounded-2xl"
			/>
			<div
				on:click={() => {
					SearchStyle = 'lg';
					inputValue = '';
				}}
				class=""
			>
				<Cross
					class="w-8 h-8 mr-2 place-content-center  rounded-xl hover:fill-white fill-slate-400"
				/>

			</div>
		</div>
	{:else if SearchStyle === 'md'}
		<div class=" hover:bg-gray-700 rounded-lg  mr-2">
			<Search
				class="w-6 h-6 m-1  fill-slate-400 hover:fill-white transition-all duration-100 ease-linear cursor-pointer"
			/>
		</div>
		<div class=" hover:bg-gray-700 rounded-lg  mr-2">
			<Search
				class="w-6 h-6 m-1  fill-slate-400 hover:fill-white transition-all duration-100 ease-linear cursor-pointer"
			/>
		</div>
	{:else if SearchStyle === 'lg'}
		<!-- USER PROFILE -->
        {#if showUsername}
             <div class="   mx-2 h-10 min-w-[2.5rem] ">
                <img
                    src={MyPro.ProfileImage}
                    alt=""
                    class=" rounded-2xl  object-cover w-full h-full   hover:rounded-xl active:rounded-md  transition-all duration-100  ease-linear cursor-pointer"
                />

            </div>
            <div class="flex flex-col  ">
                <p class=" overflow-hidden line-clamp-1 text-base text-slate-300 ">{MyPro.Name}</p>
                <p on:click={()=>{navigator.clipboard.writeText( "http://localhost:3000/"+MyPro.ProfileUrl)}} class=" cursor-pointer line-clamp-1 text-sm text-gray-500 hover:text-gray-400">@{MyPro.ProfileUrl}</p>

            </div>
            <div class=" flex-grow" />

        {:else}
             <div class="   mx-2 h-10 w-10 ">
                 <img
                     src={MyPro['ProfileImage']}
                     alt=""
                     class=" rounded-2xl  object-cover w-full h-full   hover:rounded-xl active:rounded-md  transition-all duration-100  ease-linear cursor-pointer"
                 />
             </div>
             
             <div class=" flex-grow" />
        {/if}
		<!-- video -->
		<div
			class=" hover:bg-gray-700 rounded-lg"
			on:click={() => {
				userOptions['video'] = !userOptions['video'];
			}}
		>
			{#if !userOptions['video']}
				<RedOff class=" absolute  fill-[#a4393c]  h-6 w-6  mt-1 mb-2 ml-1 " />
			{/if}
			<!-- <svg class=" absolute  fill-[#a4393c] stroke-2 h-4 w-4 my-4" viewBox="0 0 640 512" xmlns="http://www.w3.org/2000/svg"><path d="M24.03 0c5.156 0 10.37 1.672 14.78 5.109l591.1 463.1c10.44 8.172 12.25 23.27 4.062 33.7c-8.125 10.41-23.19 12.28-33.69 4.078L9.189 42.89c-10.44-8.172-12.25-23.27-4.062-33.7C9.845 3.156 16.91 0 24.03 0z"/></svg> -->
			<Video />
		</div>
		<!-- microphone -->
		<div
			class="hover:bg-gray-700 rounded-lg "
			on:click={() => {
				userOptions['microphone'] = !userOptions['microphone']
			}}
		>
			{#if !userOptions['microphone']}
				<RedOff class=" absolute  fill-[#a4393c]  h-6 w-6  mt-1 mb-2 ml-1 " />
			{/if}
			<Microphone
				class="w-6 h-6 m-1  fill-slate-400 hover:fill-white transition-all duration-100 ease-linear cursor-pointer"
			/>
		</div>
		<!-- sound -->
		<div
			class="hover:bg-gray-700 rounded-lg "
			on:click={() => {
				userOptions['sound'] = !userOptions['sound'];
			}}
		>
			{#if !userOptions['sound']}
				<RedOff class=" absolute  fill-[#a4393c]  h-6 w-6  mt-1 mb-2 ml-1 " />
			{/if}
			<Speaker />
		</div>
		<!-- notification sound -->
		<div
			class="hover:bg-gray-700 rounded-lg "
			on:click={() => {
				userOptions['notificationSound'] = !userOptions['notificationSound'];
			}}
		>
			{#if !userOptions['notificationSound']}
				<RedOff class=" absolute  fill-[#a4393c]  h-6 w-6  mt-1 mb-2 ml-1 " />
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
			class=" hover:bg-gray-700 rounded-lg  mr-2"
		>
			<Search
				class="w-6 h-6 m-1  fill-slate-400 hover:fill-white transition-all duration-100 ease-linear cursor-pointer"
			/>
		</div>
	{/if}
</div>

<style>
	/* your styles go here */
</style>
