<script lang="ts"> 
    import {goto } from "$app/navigation"
    import {ChatOrDock} from "$lib/store2"
import Accord from "$lib/_Dock/_icons/accord.svelte";
    import AccordLogo from "$lib/_Dock/_icons/accord.svelte"
    // import {ServerList} from "$lib/store2"
    
    import { fade, blur, fly, slide, scale } from "svelte/transition";

    export let ServerList:any
    let ChatOrDockHelper:number 
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
    let ActiveServer:string=""
    let HoverServer:string=""
    function OnClickServer(e: { ServerURL: any; Name?: string; NewMessage?: boolean; NumberOfNewMessage?: number; Notification?: boolean; NumberOfNotification?: number; ServerImage?: string; }){
        ActiveServer=e.ServerURL
        
        goto("/s/"+e.ServerURL,{
           replaceState: true,
			noscroll: true,
			keepfocus: true 
        })
    }


    
</script>

<style>
</style>

<!-- dock folder 48x48 and small icons 16x16 -->
<div class=" flex flex-col   min-w-[72px] max-w-[72px] bg-[#202225] items-center  overflow-y-scroll no-scrollbar transition-all duration-300 ease-linear ">

    <!-- company icon button -->
    <div class="avatar   my-1  flex justify-center  h-[48px] w-[48px]   object-cover    cursor-pointer  " on:click="{ChatOrDockFunc}">
        <AccordLogo class="h-10  w-10  hover:scale-125 transition-all duration-100  ease-linear " />
        <!-- <div class=" w-14 h-14 hover:rounded-xl   rounded-3xl active:rounded-md hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer {false ? "ring  ring-offset-base-100  ring-blue-500" : "" }">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16"><path fill="#FFC107" d="M10.376.985a1.426 1.426 0 0 0-2.711.881l3.685 11.339a1.426 1.426 0 0 0 1.734.884c.756-.218 1.218-1.021.976-1.764L10.376.985z"/><path fill="#4CAF50" d="M4.666 2.841a1.425 1.425 0 1 0-2.711.881L5.64 15.06a1.426 1.426 0 0 0 1.734.884c.756-.218 1.218-1.021.976-1.764L4.666 2.841z"/><path fill="#EC407A" d="M15.015 10.376a1.425 1.425 0 1 0-.881-2.711L2.795 11.351a1.426 1.426 0 0 0-.884 1.734c.218.756 1.021 1.218 1.764.976l11.34-3.685z"/><path fill="#472A49" d="M5.158 13.579l2.71-.881-.881-2.71-2.71.881.881 2.71z"/><path fill="#CC2027" d="M10.869 11.723l2.71-.881-.881-2.711-2.71.881.881 2.711z"/><path fill="#2196F3" d="M13.159 4.666a1.425 1.425 0 1 0-.881-2.711L.94 5.64a1.424 1.424 0 0 0-.884 1.734c.218.756 1.021 1.217 1.764.976l11.339-3.684z"/><path fill="#1A937D" d="M3.302 7.868l2.711-.881-.881-2.71-2.71.881.88 2.71z"/><path fill="#65863A" d="M9.013 6.013l2.711-.881-.881-2.711-2.711.881.881 2.711z"/></svg>
        </div> -->
    </div>
    <!-- user list -->
    {#each ServerList as u}
    <button  on:click="{()=>{ ActiveServer=u["ServerURL"]  ; OnClickServer(u) }}" class=" group w-full py-[5px]  flex  flex-row items-center justify-center  hover:rounded-lg transition-all duration-150 ease-linear">
            
        <div class=" relative flex justify-center items-center">
            <!-- <svg  class="fill-white  -left-[63px] -top-1  absolute  h-16 w-16 " viewBox="0 0 448 512" xmlns="http://www.w3.org/2000/svg"><path d="M448 95.1v320c0 35.35-28.65 64-64 64H64c-35.35 0-64-28.65-64-64v-320c0-35.35 28.65-63.1 64-63.1h320C419.3 31.1 448 60.65 448 95.1z"/></svg> -->
            <div in:scale="{{  duration: 200 }}" out:scale class="bg-white -left-[58px]  absolute transition-all duration-200 ease-linear    {ActiveServer===u["ServerURL"] ? "rounded-md h-[50px] w-[50px]" : "rounded-xl h-[10px] w-[50px] group-hover:rounded-md group-hover:h-[30px] group-hover:w-[50px] " }"></div>
            <!-- <svg in:scale="{{  duration: 200 }}" out:scale class="fill-white   -left-[55px] -top-[25px]   absolute  h-[50px] w-[50px] " viewBox="0 0 448 512" xmlns="http://www.w3.org/2000/svg"><path d="M384 32H64C28.65 32 0 60.66 0 96v320c0 35.34 28.65 64 64 64h320c35.35 0 64-28.66 64-64V96C448 60.66 419.3 32 384 32zM319.1 280h-192C114.8 280 103.1 269.2 103.1 256c0-13.2 10.8-24 24-24h192c13.2 0 23.1 10.8 23.1 24C343.1 269.2 333.2 280 319.1 280z"/></svg> -->
            
            <img src="{u["ServerImage"]}" alt="" class="  rounded-[1.25rem] hover:rounded-xl active:rounded-md object-cover h-[48px] w-[48px]  transition-all duration-100  ease-linear cursor-pointer  ">
            {#if u["Notification"] }
                <span class=" absolute px-1  border-2 border-[#202225] rounded-full text-xs font-semibold  bg-red-500 text-white -bottom-2 -right-2 ">
                    {u["NumberOfNotification"]}
                </span>
            {/if}
        </div>

    </button>
        
    {/each}
    

    <!-- download button -->
    <div class="avatar  py-1 px-2">
        <div class=" bg-sky-600  h-14 w-14 hover:rounded-xl  rounded-3xl active:rounded-md hover:ring  transition-all duration-150 ease-linear cursor-pointer ">
            <svg class="w-10 h-10 m-2" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm3.293-7.707a1 1 0 011.414 0L9 10.586V3a1 1 0 112 0v7.586l1.293-1.293a1 1 0 111.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
        </div>
    </div>    
</div>

