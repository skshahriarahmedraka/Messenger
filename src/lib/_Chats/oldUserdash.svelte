<script lang="ts">
	import { goto } from '$app/navigation';

	// import { url } from "$app/stores";

	// import {page} from "$app/stores"
	// $: console.log("$page.params ",$page.params)
	// import { ChatOrDock } from '$lib/store2';
	import Notification from '$lib/icons/notification.svelte';
	import Search from '$lib/_Navbar/_profileImg/search.svelte';
	import Cross from '$lib/_Navbar/_profileImg/Cross.svelte';
    // import UserDashPannel from "$lib/UserDashPanel/index.svelte"

	// import {FriendList ,MyPro} from "$lib/store2"
	// export let FriendList: any\
	export let MyPro: any;

	let searchIconEvent: boolean = false;
	let inputValue: string = '';

	
</script>

<div class="flex flex-row justify-center items-center   ">
    <!-- company logo -->
    <!-- <div class="avatar  py-1 px-2 " on:click={ChatOrDockFunc}>
        <div class=" w-14 h-14 hover:rounded-xl   rounded-3xl active:rounded-md hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer {false ? "ring  ring-offset-base-100  ring-blue-500" : "" }">
            <svg  viewBox="0 0 16 16"><path fill="#FFC107" d="M10.376.985a1.426 1.426 0 0 0-2.711.881l3.685 11.339a1.426 1.426 0 0 0 1.734.884c.756-.218 1.218-1.021.976-1.764L10.376.985z"/><path fill="#4CAF50" d="M4.666 2.841a1.425 1.425 0 1 0-2.711.881L5.64 15.06a1.426 1.426 0 0 0 1.734.884c.756-.218 1.218-1.021.976-1.764L4.666 2.841z"/><path fill="#EC407A" d="M15.015 10.376a1.425 1.425 0 1 0-.881-2.711L2.795 11.351a1.426 1.426 0 0 0-.884 1.734c.218.756 1.021 1.218 1.764.976l11.34-3.685z"/><path fill="#472A49" d="M5.158 13.579l2.71-.881-.881-2.71-2.71.881.881 2.71z"/><path fill="#CC2027" d="M10.869 11.723l2.71-.881-.881-2.711-2.71.881.881 2.711z"/><path fill="#2196F3" d="M13.159 4.666a1.425 1.425 0 1 0-.881-2.711L.94 5.64a1.424 1.424 0 0 0-.884 1.734c.218.756 1.021 1.217 1.764.976l11.339-3.684z"/><path fill="#1A937D" d="M3.302 7.868l2.711-.881-.881-2.71-2.71.881.88 2.71z"/><path fill="#65863A" d="M9.013 6.013l2.711-.881-.881-2.711-2.711.881.881 2.711z"/></svg>
        </div>
    </div> -->
    <!-- search -->
    {#if searchIconEvent}
        <div
            class="   flex-grow my-2 ml-1 mr-2 flex flex-row transition-all duration-300 ease-linear cursor-pointer text-lg text-gray-400"
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
                class="    w-full  self-center h-12  p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-1 my-0  border-[#24262b]  active:border-gray-800 rounded-2xl"
            />
            <div
                on:click={() => {
                    searchIconEvent = !searchIconEvent;
                    inputValue = '';
                }}
                class=""
            >
                <Cross
                    class="w-8 h-8 mt-1 place-content-center  rounded-xl hover:fill-white fill-slate-400"
                />

                <!-- <Search /> -->
            </div>
        </div>
    {:else}
        <div class=" flex-grow" />
        <!-- search -->
        <div on:click={() => (searchIconEvent = !searchIconEvent)} class="">
            <Search
                class="w-8 h-8 my-4  hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer"
            />
            <!-- <svg  class="w-8 h-8 my-4  hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg> -->
        </div>
        <!-- notification -->
        <div class=" flex flex-row ">
            <Notification />

            <span
                class="  absolute mt-2  ml-7 h-5  px-1  border-2 border-[#202225] rounded-xl text-xs font-semibold  bg-red-500 text-white  "
            >
                99+
            </span>
        </div>
        <div class="  my-1 mx-2 h-10 w-10 ">
            <img
                src={MyPro['ProfileImage']}
                alt=""
                class=" rounded-2xl  object-cover w-full h-full   hover:rounded-xl active:rounded-md  transition-all duration-100  ease-linear cursor-pointer"
            />
        </div>
    {/if}
</div>