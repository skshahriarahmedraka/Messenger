<script lang="ts">
	import { showPeopleList } from '$lib/store2';
import Hash from './_profileImg/hash.svelte';
import AudioCall from './_profileImg/audioCall.svelte';
import Circle from './_profileImg/circle.svelte';
import Options from './_profileImg/options.svelte';
	import Taw from './_profileImg/t.jpg';
import VideoCall from './_profileImg/videoCall.svelte';
import Notification from './_profileImg/Notification.svelte';
import People from './_profileImg/people.svelte';
import Search from './_profileImg/search.svelte';
import Cross from './_profileImg/Cross.svelte';
import Attherate from './_profileImg/attherate.svelte';

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

	export let MyPro: any;
</script>

<div class="  bg-[#36393f] h-12 w-full flex flex-row items-center border-b-2 border-solid border-[#2e3136] ">
	<div
		class="flex flex-row items-center mx-2 space-x-1 hover:fill-white fill-slate-400 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer"
	>
		<Attherate class="w-6 h-6 " />
		<p class="my-3 text-lg font-semibold ">Gaming Chat</p>
	</div>
	<!-- gap -->
	<!-- search -->
	{#if searchIconEvent}
		<div
			class=" flex-grow my-2 ml-1 mr-2 flex flex-row transition-all duration-300 ease-linear cursor-pointer"
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
				class="     w-full    self-center h-10  p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 ml-4 mr-2 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
			/>
            <button on:click={() => {
                searchIconEvent = !searchIconEvent;
                inputValue = '';
            }} class="">
             <Cross class="w-8 h-8 mt-1 place-content-center  rounded-xl hover:fill-white fill-slate-400" />
        
        </button>
			
		</div>
	{:else}
		<div class=" flex-grow" />
		<div on:click={() => (searchIconEvent = !searchIconEvent)} class="">
			<Search class="w-7 h-7 my-4 hover:text-white text-slate-400  transition-all duration-200 ease-linear cursor-pointer" />
		</div>
	{/if}
	<!-- audio call -->
	<AudioCall class="w-6 h-6  ml-2 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer"/>
	<!-- video call -->
	<VideoCall class="w-7 h-7  ml-2 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer"/>
	<!-- hash -->
	<Hash class="w-7 h-7 my-3 mx-2 hover:text-white text-slate-400  transition-all duration-200 ease-linear cursor-pointer" />
	<!-- notification -->
	<div class=" flex flex-row ">
		<Notification class="w-7 h-7 my-4 mr-4 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" />
		<div
			class=" absolute mt-3 ml-2 h-4  px-1  border-1 border-[#202225] rounded-xl text-xs font-semibold  bg-red-500 text-white  "
		>
			99+
		</div>
	</div>
	<!-- show people pannel -->
	<div class=" w-6 h-6" on:click={showPeopleListFunc}>
		<People class="w-6 h-6  hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" />
	</div>
	<div class="flex flex-col ">
		<div class="mr-2 mt-1  w-60 text-right font-medium  truncate text-gray-300 ">
			Sk Shahriar Ahmed Raka
		</div>
		{#if true}
			<div class="  truncate text-right text-sm text-[#6e6a5c] mr-2 place-content-end ">
				Active 1 hour ago
			</div>
		{:else}
			<div class=" items-center gap-1 place-content-end  flex flex-row mr-2  ">
				<Circle class="fill-green-400 w-2 h-2  text-right   place-content-center" />
				<div class=" text-right text-sm text-gray-400">Active</div>
			</div>
		{/if}
	</div>
	<div class=" h-10 w-10 my-1 mx-2">
		<img
			src={MyPro['ProfileImage']}
			alt=""
			class=" rounded-2xl  object-cover w-full h-full   hover:rounded-xl active:rounded-md  transition-all duration-100  ease-linear cursor-pointer "
		/>
	</div>
	<div class=" flex items-center ">
		<Options class="w-7 h-7 mt-1  hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" />
	</div>
</div>

<style>
	/* your styles go here */
</style>
