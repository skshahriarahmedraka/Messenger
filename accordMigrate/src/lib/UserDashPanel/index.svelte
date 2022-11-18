<script lang="ts">
	// import userBanner from './_images/whiteBeard.jpg';
	// import Moji2 from "./_images/moji2.jpg"
	// import { $UserProData} from "$lib/store2"
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
	import WalpaperLogo from '$lib/icons/walpaperLogo.svelte';
	import Cross from '$lib/Navbar/profileImg/Cross.svelte';
	import { goto } from '$app/navigation';
	import { UserProData ,UserSettingSelect } from '$lib/store2';
	// export let $UserProData: any;
	let searchIconEvent = false;
	let inputValue = '';
	export let showUsername: boolean;
	let userOptions = { video: false, microphone: false, notificationSound: false, sound: false };
	let SearchStyle = 'lg'; // sm,md,lg
	let showUserid = true;
	let showUserOptions = true;
	import Coin from '$lib/icons/coin2.svelte';
	import Person from '$lib/icons/person.svelte';

	async function LogoutReq() {
		await fetch('/api/logout', {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			}
		})
			.then((res) => {
				return res.json();
			})
			.then((d) => {
				console.log('🚀 ~ file: index.svelte ~ line 43 ~ LogoutReq ~ d', d);
				goto('/login');
			});
	}
</script>

{#if showUserOptions}
	<div
		class=" mb-0 flex h-fit w-full flex-row flex-wrap items-center justify-center  rounded-xl bg-[#18191c] bg-opacity-80 text-gray-300 "
		on:blur={() => {
			showUserOptions = false;
		}}
	>
		<div class="h-40 w-full">
			{#if $UserProData.BannerImg.trim() === ''}
				<div
					class="flex h-40 w-full items-center justify-center rounded-xl border-2 border-slate-400 border-opacity-30 "
				>
					<WalpaperLogo class=" h-20 w-20" />
				</div>
			{:else}
				<img
					src={$UserProData.BannerImg}
					alt="coverImage"
					class="h-40 w-full rounded-xl object-cover "
				/>
			{/if}
			{#if $UserProData.ProfileImg === ''}
				<Person
					class=" fixed -mt-7 ml-5 h-12 w-12 cursor-pointer rounded-xl border-2  border-gray-400  fill-gray-400 object-cover p-2    transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md"
				/>
			{:else}
				<!-- <img
				src={$UserProData.ProfileImg}
				alt="ProfileImg"
				class=" h-full  w-full cursor-pointer rounded-2xl   object-cover transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md "
			/> -->
				<img
					src={$UserProData.ProfileImg}
					alt="coverImage"
					class="fixed -mt-7 ml-5 h-12 w-12 rounded-xl object-cover"
				/>
			{/if}
		</div>
		<p class="  mt-5 self-center font-Poppins text-2xl ">{$UserProData.UserName}</p>
		<!-- <div class=" flex flex-col ">

			<div
			class=" flex mt-12  ml-20   w-full flex-row items-center gap-2 rounded-2xl border-0 border-blue-500 bg-sky-300 bg-opacity-30 pr-2  sm:hidden md:hidden lg:hidden xl:hidden xs:hidden xxl:contents"
		>
			<Coin class="storke-[1px] h-8 w-8 stroke-yellow-300" />
			<p class="text-slate-200 ">{826753}</p>
		</div>
		</div> -->
		<div class="mt-2 inline-flex h-16 w-full ">
			<button
				on:click={() => {
					goto(`/${$UserProData.UserID}/profile`);
				}}
				class=" flex h-16 w-1/2 flex-row items-center justify-center  gap-2 rounded-lg hover:bg-gray-900 "
			>
				<Coin class="storke-[1px] h-8 w-8 stroke-yellow-300" />
				<p class="font-Poppins text-slate-200 ">{$UserProData.Coin}</p>
			</button>
			<button
				on:click={() => {
					goto(`/${$UserProData.UserID}/profile/settings`);
				}}
				class=" h-16  w-1/2 rounded-lg  font-Poppins hover:bg-gray-900 ">Settings</button
			>
		</div>
		<div class="inline-flex h-16 w-full">
			<button
				on:click={() => {
					UserSettingSelect.set("Request Token")
			goto(`/${$UserProData.UserID}/profile/settings`);

				}}
				class=" h-16 w-1/2 rounded-lg font-Poppins hover:bg-gray-900">Recharge Wallet</button
			>
			<button
				on:click={() => {
					LogoutReq();
				}}
				class=" h-16 w-1/2 rounded-lg  font-Poppins text-red-500 hover:bg-gray-900 ">Logout</button
			>
		</div>
	</div>
{/if}
<div class="flex h-14 w-full flex-row items-center justify-center bg-[#292b2f]  ">
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
				class=" mx-1  my-0 h-10  w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e]  outline-none  focus:border-sky-500 active:border-gray-800"
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
			<div
				on:click={() => {
					showUserOptions = !showUserOptions;
				}}
				class="   mx-2 h-10 min-w-[2.5rem]  "
			>
				{#if $UserProData.ProfileImg === ''}
					<Person
						class="h-10 w-10  cursor-pointer rounded-2xl border-2  border-gray-400 fill-gray-400 p-2    transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md"
					/>
				{:else}
					<img
						src={$UserProData.ProfileImg}
						alt=""
						class="  h-10  w-10  cursor-pointer rounded-2xl   object-cover transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md"
					/>
				{/if}
			</div>
			<div class="flex flex-col  ">
				<p class=" overflow-hidden text-base text-slate-300 line-clamp-1 ">
					{$UserProData.UserName}
				</p>
				{#if showUserid}
					<p
						on:click={() => {
							navigator.clipboard.writeText(`http://${process.env.HOST}/`+$UserProData.UserID);
							showUserid = false;
						}}
						class=" cursor-pointer text-sm text-gray-500 line-clamp-1 hover:text-gray-400"
					>
						@{$UserProData.UserID}
					</p>
				{:else}
					<p
						on:click={() => {
							navigator.clipboard.writeText(`http://${process.env.HOST}/` + $UserProData.UserID);
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
				{#if $UserProData.ProfileImg === ''}
					<Person
						class="h-10 w-10  cursor-pointer rounded-2xl border-2  border-gray-400 fill-gray-400 p-2    transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md"
					/>
				{:else}
					<img
						src={$UserProData.ProfileImg}
						alt=""
						class=" h-full  w-full cursor-pointer rounded-2xl   object-cover transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md"
					/>
				{/if}
			</div>

			<div class=" flex-grow" />
		{/if}
		<!-- video -->
<!--		<div-->
<!--			class=" rounded-lg hover:bg-gray-700"-->
<!--			on:click={() => {-->
<!--				userOptions['video'] = !userOptions['video'];-->
<!--			}}-->
<!--		>-->
<!--			{#if !userOptions['video']}-->
<!--				<RedOff class=" absolute  mt-1  mb-2 ml-1  h-6 w-6 fill-[#a4393c] " />-->
<!--			{/if}-->
<!--			<Video />-->
<!--		</div>-->
		<!-- microphone -->
<!--		<div-->
<!--			class="rounded-lg hover:bg-gray-700 "-->
<!--			on:click={() => {-->
<!--				userOptions['microphone'] = !userOptions['microphone'];-->
<!--			}}-->
<!--		>-->
<!--			{#if !userOptions['microphone']}-->
<!--				<RedOff class=" absolute  mt-1  mb-2 ml-1  h-6 w-6 fill-[#a4393c] " />-->
<!--			{/if}-->
<!--			<Microphone-->
<!--				class="m-1 h-6 w-6  cursor-pointer fill-slate-400 transition-all duration-100 ease-linear hover:fill-white"-->
<!--			/>-->
<!--		</div>-->
		<!-- sound -->
<!--		<div-->
<!--			class="rounded-lg hover:bg-gray-700 "-->
<!--			on:click={() => {-->
<!--				userOptions['sound'] = !userOptions['sound'];-->
<!--			}}-->
<!--		>-->
<!--			{#if !userOptions['sound']}-->
<!--				<RedOff class=" absolute  mt-1  mb-2 ml-1  h-6 w-6 fill-[#a4393c] " />-->
<!--			{/if}-->
<!--			<Speaker />-->
<!--		</div>-->
		<!-- notification sound -->
<!--		<div-->
<!--			class="rounded-lg hover:bg-gray-700 "-->
<!--			on:click={() => {-->
<!--				userOptions['notificationSound'] = !userOptions['notificationSound'];-->
<!--			}}-->
<!--		>-->
<!--			{#if !userOptions['notificationSound']}-->
<!--				<RedOff class=" absolute  mt-1  mb-2 ml-1  h-6 w-6 fill-[#a4393c] " />-->
<!--			{/if}-->
<!--			<Notification />-->
<!--		</div>-->

<!--		<div-->
<!--			on:click={() => {-->
<!--				SearchStyle = 'sm';-->
<!--			}}-->
<!--			class=" mr-2 rounded-lg  hover:bg-gray-700"-->
<!--		>-->
<!--			<Search-->
<!--				class="m-1 h-6 w-6  cursor-pointer fill-slate-400 transition-all duration-100 ease-linear hover:fill-white"-->
<!--			/>-->
<!--		</div>-->
	{/if}
</div>

<style>
	/* your styles go here */
</style>
