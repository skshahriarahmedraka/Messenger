<script lang="ts">
    import {ChatOrDock} from "$lib/store"
    import Moji from "./profilePic/ot2.jpg"
    import Moji2 from "./profilePic/moji2.jpg"
    // your script goes here
    let ChatOrDockHelper
    ChatOrDock.subscribe(val=>{
        ChatOrDockHelper=val 
    }) 

    function ChatOrDockFunc(){
        if (ChatOrDockHelper != 0){

            ChatOrDock.update(n=> n=0)
        }else{
            ChatOrDock.update(n=> n=1)
        }
    }
    let searchIconEvent:boolean=false
    let inputValue:string=""
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div class=" min-w-[380px] h-screen  bg-[#2f3136] flex flex-col transition-all duration-500 ease-linear  ">
    <!-- scrollbar-thin  scrollbar-thumb-gray-500 scrollbar-track-gray-400 -->
    <div class="flex flex-row ">
        <!-- company logo -->
        <div class="avatar  py-1 px-2 " on:click="{ChatOrDockFunc}">
            <div class=" w-14 h-14 hover:rounded-xl   rounded-3xl active:rounded-md hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer {false ? "ring  ring-offset-base-100  ring-blue-500" : "" }">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16"><path fill="#FFC107" d="M10.376.985a1.426 1.426 0 0 0-2.711.881l3.685 11.339a1.426 1.426 0 0 0 1.734.884c.756-.218 1.218-1.021.976-1.764L10.376.985z"/><path fill="#4CAF50" d="M4.666 2.841a1.425 1.425 0 1 0-2.711.881L5.64 15.06a1.426 1.426 0 0 0 1.734.884c.756-.218 1.218-1.021.976-1.764L4.666 2.841z"/><path fill="#EC407A" d="M15.015 10.376a1.425 1.425 0 1 0-.881-2.711L2.795 11.351a1.426 1.426 0 0 0-.884 1.734c.218.756 1.021 1.218 1.764.976l11.34-3.685z"/><path fill="#472A49" d="M5.158 13.579l2.71-.881-.881-2.71-2.71.881.881 2.71z"/><path fill="#CC2027" d="M10.869 11.723l2.71-.881-.881-2.711-2.71.881.881 2.711z"/><path fill="#2196F3" d="M13.159 4.666a1.425 1.425 0 1 0-.881-2.711L.94 5.64a1.424 1.424 0 0 0-.884 1.734c.218.756 1.021 1.217 1.764.976l11.339-3.684z"/><path fill="#1A937D" d="M3.302 7.868l2.711-.881-.881-2.71-2.71.881.88 2.71z"/><path fill="#65863A" d="M9.013 6.013l2.711-.881-.881-2.711-2.711.881.881 2.711z"/></svg>
                <!-- <svg  enable-background="new 0 0 100 100" viewBox="0 0 100 100"><path fill="#ce3056" d="M3.63,62.22c0.563,12.911,18.923,12.908,19.484,0c0,0,0-9.73,0-9.73h-9.742 C8.011,52.489,3.63,56.864,3.63,62.22z"/><path fill="#5bb381" d="M62.234,47.511c5.362,0,9.742-4.375,9.742-9.73V13.417c-0.56-12.905-18.92-12.914-19.484,0 c0,0,0,24.364,0,24.364C52.492,43.136,56.872,47.511,62.234,47.511z"/><path fill="#ce3056" d="M37.765,52.489c-5.362,0-9.742,4.375-9.742,9.73v24.364c0.56,12.905,18.92,12.914,19.484,0 c0,0,0-24.364,0-24.364C47.508,56.864,43.128,52.489,37.765,52.489z"/><path fill="#5bb381" d="M96.37,37.78c-0.563-12.911-18.923-12.908-19.484,0c0,0,0,9.73,0,9.73h9.742 C91.989,47.511,96.37,43.136,96.37,37.78z"/><path fill="#61c3ec" d="M13.372,47.511h24.393c12.921-0.559,12.93-18.897,0-19.461H13.372C0.45,28.61,0.445,46.948,13.372,47.511z"/><path fill="#e4b34b" d="M62.234 76.853h-9.742v9.73c0 5.355 4.38 9.73 9.742 9.73C75.161 95.751 75.158 77.414 62.234 76.853zM86.628 52.489H62.234c-12.923.56-12.928 18.898 0 19.461 0 0 24.393 0 24.393 0C99.548 71.391 99.557 53.052 86.628 52.489z"/><path fill="#61c3ec" d="M37.765,3.686c-12.926,0.562-12.924,18.9,0,19.461c0,0,9.742,0,9.742,0v-9.73 C47.508,8.061,43.128,3.686,37.765,3.686z"/><path d="M62.182 75.355h-9.701c-.765 0-1.385.62-1.385 1.383v9.689c0 6.105 4.973 11.073 11.086 11.073C76.887 96.891 76.885 75.962 62.182 75.355zM62.182 94.733c-4.585 0-8.316-3.726-8.316-8.305v-8.305h8.316C73.204 78.558 73.203 94.298 62.182 94.733zM86.472 51.095h-24.29c-14.702.607-14.707 21.537 0 22.145 0 0 24.29 0 24.29 0C101.172 72.635 101.18 51.704 86.472 51.095zM86.472 70.473h-24.29c-4.585 0-8.316-3.726-8.316-8.305s3.731-8.305 8.316-8.305h24.29C97.489 54.296 97.498 70.036 86.472 70.473zM76.771 48.905h9.701c14.704-.607 14.705-21.538 0-22.145-6.112 0-11.086 4.967-11.086 11.073v9.689C75.386 48.285 76.006 48.905 76.771 48.905zM78.156 37.832c0-4.579 3.731-8.305 8.316-8.305 11.023.436 11.021 16.176 0 16.611h-8.316V37.832zM37.818 2.5c-14.706.608-14.704 21.538 0 22.145 0 0 9.701 0 9.701 0 .765 0 1.385-.62 1.385-1.383v-9.689C48.903 7.467 43.93 2.5 37.818 2.5zM46.133 21.878h-8.316c-4.585 0-8.316-3.726-8.316-8.305.437-11.01 16.196-11.008 16.631 0C46.133 13.572 46.133 21.878 46.133 21.878zM23.229 51.095h-9.701c-14.704.607-14.705 21.538 0 22.145 6.112 0 11.086-4.967 11.086-11.072v-9.689C24.614 51.715 23.994 51.095 23.229 51.095zM21.844 62.168c-.436 11.009-16.196 11.008-16.631 0 0-4.579 3.731-8.305 8.316-8.305h8.316V62.168z"/><path d="M37.818 51.095c-6.112 0-11.086 4.967-11.086 11.072v24.26c.606 14.683 21.562 14.691 22.172 0 0 0 0-24.26 0-24.26C48.903 56.063 43.93 51.095 37.818 51.095zM46.133 86.427c-.434 11.004-16.193 11.013-16.631 0 0 0 0-24.26 0-24.26 0-4.579 3.731-8.305 8.316-8.305s8.316 3.726 8.316 8.305V86.427zM13.529 48.905c0 0 24.289 0 24.289 0 14.702-.607 14.707-21.536 0-22.145-1.827.036-1.816 2.735 0 2.767 11.017.433 11.026 16.174 0 16.611 0 0-24.289 0-24.289 0-11.017-.433-11.027-16.174 0-16.611 0 0 18.216 0 18.216 0 1.817-.027 1.823-2.738 0-2.767 0 0-18.216 0-18.216 0C-1.171 27.365-1.182 48.295 13.529 48.905z"/><path d="M62.182,48.905c6.112,0,11.086-4.967,11.086-11.072c-0.036-1.824-2.738-1.814-2.77,0 c-0.434,11.004-16.194,11.013-16.631,0c0,0,0-24.26,0-24.26c0.434-11.003,16.193-11.014,16.631,0c0,0,0,18.714,0,18.714 c0.027,1.814,2.741,1.821,2.77,0c0,0,0-18.714,0-18.714c-0.605-14.681-21.562-14.692-22.172,0c0,0,0,24.26,0,24.26 C51.096,43.937,56.069,48.905,62.182,48.905z"/></svg> -->
            </div>
        </div>
        <!-- search -->
        {#if searchIconEvent}
             
             <div class=" flex-grow my-2 ml-1 mr-2 flex flex-row transition-all duration-300 ease-linear cursor-pointer text-lg text-gray-400">
                 <input placeholder="Search"  autocomplete="off" autofocus on:blur="{()=>{ inputValue==="" ? searchIconEvent=!searchIconEvent : searchIconEvent=searchIconEvent }}"  value="{inputValue}" type="text" class=" bg-slate-600 rounded-xl h-10 mt-2 w-full">
                <svg on:click={()=>searchIconEvent=!searchIconEvent} class="w-8 h-8 my-3 mr-0 rounded-xl hover:text-white text-slate-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
             </div>
        {:else}
             
             <div class=" flex-grow"></div>
             <!-- search -->
             <svg on:click="{()=>searchIconEvent=!searchIconEvent}" class="w-8 h-8 my-4 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
             <!-- notification -->
            <div class=" flex flex-row ">

                 <svg class="w-8 h-8 m-4 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
                <span class=" absolute mt-2  ml-7 h-5  px-1  border-2 border-[#202225] rounded-xl text-xs font-semibold  bg-red-500 text-white  ">
                        99+
                </span>
            </div>
             <div class=" my-1 mx-2 h-12 w-12 transition-all duration-300 ease-linear cursor-pointer">
               <img src="{ Moji2}" alt="" class=" rounded-2xl hover:rounded-xl active:rounded-md object-cover w-full h-full hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer  active:ring  active:ring-offset-base-50  active:ring-blue-600">
              
             </div>
        {/if}

    </div>
    <div class=" h-full w-full overflow-y-auto scrol  overflow-x-hidden ">

        <!-- user chat list -->
        {#each Array.from(Array(100+1).keys()).slice(1) as i }
         <div class=" min-w-full min-h-[72px] m-1 hover:bg-[#202225] flex flex-row hover:rounded-xl overflow-hidden ">
            <!-- user image -->
            <img src="{Moji}" alt="" class=" w-14 h-14 m-2 rounded-2xl">
            <div class="flex flex-col min-w-full min-h-full ">
                <div class="   h-8  mt-2 ml-0 flex flex-row">
                    <!-- User name -->
                    <p class=" w-56 font-bold text-lg text-ellipsis truncate   ">
                        skshahriarahmedrakarakarakaraka
                    </p>
                    <!-- mute icon -->
                    <svg class="w-5 h-5 m-1 flex-none " fill="currentColor" viewBox="0 0 20 20" ><path fill-rule="evenodd" d="M9.383 3.076A1 1 0 0110 4v12a1 1 0 01-1.707.707L4.586 13H2a1 1 0 01-1-1V8a1 1 0 011-1h2.586l3.707-3.707a1 1 0 011.09-.217zM12.293 7.293a1 1 0 011.414 0L15 8.586l1.293-1.293a1 1 0 111.414 1.414L16.414 10l1.293 1.293a1 1 0 01-1.414 1.414L15 11.414l-1.293 1.293a1 1 0 01-1.414-1.414L13.586 10l-1.293-1.293a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
                    <!-- last message time -->
                    <p class=" font-thin text-xs mt-3 mb-1 mx-0 flex-none ">
                        9:32 PM
                    </p>
                </div>
                <div class=" h-8 w-72 flex flex-row ">
                    <!-- last message -->
                    <p class="truncate w-64 h-full ">
                        a great idea, so the Russians can't interfere with Ukraine's communication
                    </p>
                    <!-- number of notification -->
                     <span class=" h-5  px-1  border-2 border-[#202225] rounded-xl text-xs font-semibold  bg-red-500 text-white  ">
                        99+
                    </span>
                </div>
            </div>
         </div>
    {:else}
    <!-- empty list -->
    {/each}
    
</div>
</div>