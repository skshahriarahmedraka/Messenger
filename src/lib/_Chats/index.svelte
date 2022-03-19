<script context="module">
    // export async function load({fetch}){
    //     const res = await fetch("http.localhost:3000/")
    // }
</script>

<script lang="ts">
import { goto } from "$app/navigation";

import { url } from "$app/stores";

    // import {page} from "$app/stores"
    // $: console.log("$page.params ",$page.params)
    import {ChatOrDock} from "$lib/store2"

    import {FriendList ,MyPro} from "$lib/store2"



    
    let ChatOrDockHelper
    ChatOrDock.subscribe(val=>{
        ChatOrDockHelper=val 
    }) 

    function ChatOrDockFunc(){
        
        goto("/", {
			replaceState: true,
			noscroll: true,
			keepfocus: true
		});
        if (ChatOrDockHelper != 0){

            ChatOrDock.update(n=> n=0)
        }else{
            ChatOrDock.update(n=> n=1)
        }
    }
    let searchIconEvent:boolean=false
    let inputValue:string=""

    let ActiveChat:string=""

    function OnClickFriend(e){
        ActiveChat=e.ProfileURL
        goto(e.ProfileURL, {
			replaceState: true,
			noscroll: true,
			keepfocus: true
		});

    }


</script>

<style>
    /* your styles go here */
</style>


<div class=" w-[485px] h-full  bg-[#2f3136] flex flex-col transition-all duration-500 ease-linear  ">
    <div class="flex flex-row ">
        <!-- company logo -->
        <div class="avatar  py-1 px-2 " on:click={ChatOrDockFunc}>
            <div class=" w-14 h-14 hover:rounded-xl   rounded-3xl active:rounded-md hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer {false ? "ring  ring-offset-base-100  ring-blue-500" : "" }">
                <svg  viewBox="0 0 16 16"><path fill="#FFC107" d="M10.376.985a1.426 1.426 0 0 0-2.711.881l3.685 11.339a1.426 1.426 0 0 0 1.734.884c.756-.218 1.218-1.021.976-1.764L10.376.985z"/><path fill="#4CAF50" d="M4.666 2.841a1.425 1.425 0 1 0-2.711.881L5.64 15.06a1.426 1.426 0 0 0 1.734.884c.756-.218 1.218-1.021.976-1.764L4.666 2.841z"/><path fill="#EC407A" d="M15.015 10.376a1.425 1.425 0 1 0-.881-2.711L2.795 11.351a1.426 1.426 0 0 0-.884 1.734c.218.756 1.021 1.218 1.764.976l11.34-3.685z"/><path fill="#472A49" d="M5.158 13.579l2.71-.881-.881-2.71-2.71.881.881 2.71z"/><path fill="#CC2027" d="M10.869 11.723l2.71-.881-.881-2.711-2.71.881.881 2.711z"/><path fill="#2196F3" d="M13.159 4.666a1.425 1.425 0 1 0-.881-2.711L.94 5.64a1.424 1.424 0 0 0-.884 1.734c.218.756 1.021 1.217 1.764.976l11.339-3.684z"/><path fill="#1A937D" d="M3.302 7.868l2.711-.881-.881-2.71-2.71.881.88 2.71z"/><path fill="#65863A" d="M9.013 6.013l2.711-.881-.881-2.711-2.711.881.881 2.711z"/></svg>
            </div>
        </div>
        <!-- search -->
        {#if searchIconEvent}         
             <div class="  flex-grow my-2 ml-1 mr-2 flex flex-row transition-all duration-300 ease-linear cursor-pointer text-lg text-gray-400">  
                <input placeholder="Search"  autocomplete="off" autofocus on:blur="{()=>{ inputValue==="" ? searchIconEvent=!searchIconEvent : searchIconEvent=searchIconEvent }}"  bind:value="{inputValue}" type="text" class="    w-full  self-center h-12  p-2 text-lg font-medium text-[#98999e] outline-0 focus:border-sky-600  bg-[#303338] border-2 mx-1 my-0  border-[#24262b]  active:border-gray-800 rounded-2xl">
                <svg on:click={()=>{searchIconEvent=!searchIconEvent ; inputValue="" } } class="w-8 h-8 my-3 mr-0 rounded-xl hover:text-white text-slate-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
             </div>
        {:else}
             
            <div class=" flex-grow"></div>
             <!-- search -->
            <svg on:click="{()=>searchIconEvent=!searchIconEvent}" class="w-8 h-8 my-4 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
             <!-- notification -->
            <div class=" flex flex-row ">

                <svg class="w-8 h-8 m-4 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
                <span class="  absolute mt-2  ml-7 h-5  px-1  border-2 border-[#202225] rounded-xl text-xs font-semibold  bg-red-500 text-white  ">
                        99+
                </span>
            </div>
            <div class=" my-1 mx-2 h-12 w-12 transition-all duration-300 ease-linear cursor-pointer">
               <img src="{ $MyPro["ProfileImage"]}" alt="" class=" rounded-2xl hover:rounded-xl active:rounded-md object-cover w-full h-full hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer  active:ring  active:ring-offset-base-50  active:ring-blue-600">              
            </div>
        {/if}

    </div>
    <div class=" h-full w-full overflow-y-auto scrol overflow-x-hidden ">

        <!-- user chat list -->
        {#each FriendList as i }
         <div on:click={ ()=>{OnClickFriend(i)  }} class="  w-full h-[72px] m-1  { ActiveChat=== i.ProfileURL ? 'bg-slate-600 rounded-xl' : 'hover:bg-[#202225]' } flex flex-row hover:rounded-xl  ">
            <!-- active indecator -->
               
            <!-- user image -->
            <img src="{i.ProfileImage}" alt="" class=" w-14 h-14 m-2 aspect-square  rounded-3xl hover:rounded-xl active:rounded-md object-cover hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer  active:ring  active:ring-offset-base-50  active:ring-blue-600">
            <div class="flex flex-col w-full h-full mr-2">
                <div class="   h-8  mt-2 ml-0 flex flex-row">
                    <!-- User name -->
                    <p class=" w-56  font-bold text-lg text-ellipsis truncate   ">
                        {i.UserName}
                    </p>
                    <div class=" flex-grow"></div>
                    <!-- mute icon -->
                    {#if i.SilentNotification}
                         <svg class="w-5 h-5 m-1 flex-none " fill="currentColor" viewBox="0 0 20 20" ><path fill-rule="evenodd" d="M9.383 3.076A1 1 0 0110 4v12a1 1 0 01-1.707.707L4.586 13H2a1 1 0 01-1-1V8a1 1 0 011-1h2.586l3.707-3.707a1 1 0 011.09-.217zM12.293 7.293a1 1 0 011.414 0L15 8.586l1.293-1.293a1 1 0 111.414 1.414L16.414 10l1.293 1.293a1 1 0 01-1.414 1.414L15 11.414l-1.293 1.293a1 1 0 01-1.414-1.414L13.586 10l-1.293-1.293a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
                    {/if}
                    <!-- last message time -->
                    <p class="   font-thin text-xs mt-3 mb-1 mx-0 flex-none  ">
                        <!-- 34 aug21  -->
                        {i.LastMessageTime}
                    </p>
                </div>
                <div class=" h-8 w-72 flex flex-row ">
                    <!-- last message -->
                    <p class="truncate w-72  h-full   ">
                        {i.LastMessage}
                    </p>
                    <!-- number of notification -->
                   {#if i.SilentNotification && i.NumberOfNotification!=0 }
                        <span class="   h-5  px-1  border-2 border-[#202225] rounded-xl text-xs font-semibold bg-slate-500 text-white  ">
                            {i.NumberOfNotification}
                        </span>
                   {:else if i.NumberOfNotification != 0 }
                        <span class="   h-5  px-1  border-2 border-[#202225] rounded-xl text-xs font-semibold bg-red-500 text-white  ">
                            {i.NumberOfNotification}
                        </span>
                    {/if}
                </div>
            </div>
         </div>
    {:else}
    <!-- empty list -->
    {/each}
    
</div>
</div>