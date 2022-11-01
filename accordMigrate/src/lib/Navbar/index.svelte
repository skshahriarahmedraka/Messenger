<script lang="ts">
	import { showPeopleList } from '$lib/store2';
	import Hash from './profileImg/hash.svelte';
	import AudioCall from './profileImg/audioCall.svelte';
	import Circle from './profileImg/circle.svelte';
	import Options from './profileImg/options.svelte';
	import Taw from './profileImg/t.jpg';
	import VideoCall from './profileImg/videoCall.svelte';
	import Notification from './profileImg/Notification.svelte';
	import People from './profileImg/people.svelte';
	import Search from './profileImg/search.svelte';
	import Cross from './profileImg/Cross.svelte';
	import Attherate from './profileImg/attherate.svelte';
	import { UserProData } from '$lib/store2';
	import Person from '$lib/icons/person.svelte';

	let showPeopleList1: number;
	showPeopleList.subscribe((val) => {
		showPeopleList1 = val;
	});

	function showPeopleListFunc() {
		if (showPeopleList1 != 0) {
			showPeopleList.update((n) => (n = 0));
		} else {
			showPeopleList.update((n) => (n = 1));
		}
	}
	let inputValue: string = '';
	let searchIconEvent: boolean = false;

	// export let $UserProData: any;
</script>

<div
	class="  flex h-12 w-full flex-row items-center border-b-2 border-solid border-[#2e3136] bg-[#36393f] "
>
	<div
		class="mx-2 flex cursor-pointer flex-row items-center space-x-1 fill-slate-400 text-slate-400 transition-all duration-200 ease-linear hover:fill-white hover:text-white"
	>
		<Attherate class="h-6 w-6 " />
		<p class="my-3 text-lg font-semibold ">Gaming Chat</p>
	</div>
	<!-- gap -->
	<!-- search -->
	{#if searchIconEvent}
		<div
			class=" my-2 ml-1 mr-2 flex flex-grow cursor-pointer flex-row transition-all duration-300 ease-linear"
		>
			<input
				placeholder="Search"
				autocomplete="off"
				on:blur={() => {
					inputValue === ''
						? (searchIconEvent = !searchIconEvent)
						: (searchIconEvent = searchIconEvent);
				}}
				bind:value={inputValue}
				type="text"
				class="     my-2    ml-4 mr-2  h-10 w-full self-center rounded-2xl border-2 border-[#24262b]  bg-[#303338] p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
			<button
				on:click={() => {
					searchIconEvent = !searchIconEvent;
					inputValue = '';
				}}
				class=""
			>
				<Cross
					class="mt-1 h-8 w-8 place-content-center  rounded-xl fill-slate-400 hover:fill-white"
				/>
			</button>
		</div>
	{:else}
		<div class=" flex-grow" />
		<div on:click={() => (searchIconEvent = !searchIconEvent)} class="">
			<Search
				class="my-4 h-7 w-7 cursor-pointer text-slate-400  transition-all duration-200 ease-linear hover:text-white"
			/>
		</div>
	{/if}
	<!-- audio call -->
	<AudioCall
		class="ml-2 h-6  w-6 cursor-pointer text-slate-400 transition-all duration-200 ease-linear hover:text-white"
	/>
	<!-- video call -->
	<VideoCall
		class="ml-2 h-7  w-7 cursor-pointer text-slate-400 transition-all duration-200 ease-linear hover:text-white"
	/>
	<!-- hash -->
	<Hash
		class="my-3 mx-2 h-7 w-7 cursor-pointer text-slate-400  transition-all duration-200 ease-linear hover:text-white"
	/>
	<!-- notification -->
	<div class=" flex flex-row ">
		<Notification
			class="my-4 mr-4 h-7 w-7 cursor-pointer text-slate-400 transition-all duration-200 ease-linear hover:text-white"
		/>
		<div
			class=" border-1 absolute mt-3 ml-2  h-4  rounded-xl border-[#202225] bg-red-500 px-1 text-xs  font-semibold text-white  "
		>
			99+
		</div>
	</div>
	<!-- show people pannel -->
	<div class=" h-6 w-6" on:click={showPeopleListFunc}>
		<People
			class="h-6 w-6  cursor-pointer text-slate-400 transition-all duration-200 ease-linear hover:text-white"
		/>
	</div>
	<div class="flex flex-col ">
		<div class="mr-2 mt-1  w-60 truncate text-right  font-medium text-gray-300 ">
			Sk Shahriar Ahmed Raka
		</div>
		{#if true}
			<div class="  mr-2 place-content-end truncate text-right text-sm text-[#6e6a5c] ">
				Active 1 hour ago
			</div>
		{:else}
			<div class=" mr-2 flex flex-row  place-content-end items-center gap-1  ">
				<Circle class="h-2 w-2 place-content-center  fill-green-400   text-right" />
				<div class=" text-right text-sm text-gray-400">Active</div>
			</div>
		{/if}
	</div>
	<div class=" my-1 mx-2 h-10 w-10">
		{#if $UserProData.ProfileImg === ''}
			<Person
				class="h-10 w-10  cursor-pointer rounded-2xl border-2  border-gray-400 fill-gray-400 p-2    transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md"
			/>
		{:else}
			<img
				src={$UserProData.ProfileImg}
				alt="ProfileImg"
				class=" h-full  w-full cursor-pointer rounded-2xl   object-cover transition-all  duration-100 ease-linear  hover:rounded-xl active:rounded-md "
			/>
		{/if}
	</div>
	<div class=" flex items-center ">
		<Options
			class="mt-1 h-7 w-7  cursor-pointer text-slate-400 transition-all duration-200 ease-linear hover:text-white"
		/>
	</div>
</div>

<style>
	/* your styles go here */
</style>
