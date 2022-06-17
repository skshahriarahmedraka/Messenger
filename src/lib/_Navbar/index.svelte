<script lang="ts">
    
    import {showPeopleList, MyPro} from "$lib/store2"
    import Taw from "./_profileImg/t.jpg"

    let showPeopleList1:number  
    showPeopleList.subscribe(val=>{
        showPeopleList1=val 
    }) 

    function showPeopleListFunc(){
        if (showPeopleList1 != 0){

            showPeopleList.update(n=> n=0)
        }else{
            showPeopleList.update(n=> n=1)
        }
    }
    let inputValue:string=""
    let searchIconEvent:boolean=false


</script>

<style>
    /* your styles go here */
</style>


 <div class=" bg-[#36393f] h-14 w-full flex flex-row border-b-2 border-solid border-[#32353a] ">
     <div class="flex flex-row hover:fill-white fill-slate-400 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer">
         <svg class="w-8 h-8 m-2 " fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M14.243 5.757a6 6 0 10-.986 9.284 1 1 0 111.087 1.678A8 8 0 1118 10a3 3 0 01-4.8 2.401A4 4 0 1114 10a1 1 0 102 0c0-1.537-.586-3.07-1.757-4.243zM12 10a2 2 0 10-4 0 2 2 0 004 0z" clip-rule="evenodd"></path></svg>
         <p class="my-3 text-xl font-semibold ">Gaming Chat</p>
     </div>
     <!-- gap -->
     <!-- search -->
     {#if searchIconEvent}
          
          <div class=" flex-grow my-2 ml-1 mr-2 flex flex-row transition-all duration-300 ease-linear cursor-pointer">
                    <input placeholder="Search" autocomplete="off"  on:blur="{()=>{ inputValue==="" ? searchIconEvent=!searchIconEvent : searchIconEvent=searchIconEvent }}"  bind:value="{inputValue}"  type="text" class="     w-full    self-center h-12  p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 ml-4 mr-2 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
                     <svg on:click={()=>{searchIconEvent=!searchIconEvent ; inputValue="" }} class="w-8 h-8 mt-1 place-content-center rounded-xl hover:fill-white fill-slate-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
         </div>
    {:else}
        <div class=" flex-grow"></div>
        <div class="">
            <svg on:click="{()=>searchIconEvent=!searchIconEvent}" class="w-7 h-7 my-4 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
        </div>
        
    {/if}
    <!-- audio call -->
    <svg class="w-7 h-7 mt-4 ml-2 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path></svg>
    <!-- video call -->
    <svg class="w-8 h-8 my-3 ml-2 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path></svg>
    <!-- hash -->
     <svg class="w-8 h-8 my-3 mx-2 hover:text-white text-slate-400  transition-all duration-200 ease-linear cursor-pointer" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14"></path></svg>
     <!-- notification -->
    <div class=" flex flex-row ">

        <svg class="w-7 h-7 my-4 mr-4 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
        <div class=" absolute mt-3 ml-2 h-4  px-1  border-1 border-[#202225] rounded-xl text-xs font-semibold  bg-red-500 text-white  ">
                99+
        </div>
    </div>

     <!-- show people pannel -->
     <div class=" w-8 h-8 my-4 mx-0" on:click={showPeopleListFunc}> 
        <svg class="w-6 h-6 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"></path></svg>
     </div>
     <div class=" w-1"></div>

     <div class="flex flex-col ">
         <div class="mr-2 mt-1 w-60 text-right font-medium text-lg truncate text-gray-300 ">
            Sk Shahriar Ahmed Raka
         </div>
         {#if false} 
            <div class=" w-56 truncate text-right text-sm text-[#6e6a5c] mr-2 mt-1 ">
                Active 1 hour ago
            </div>
         {:else}
            <div class=" content-end place-content-end  flex flex-row mr-2  mt-1">
                <svg class=" fill-green-400 w-2 h-2 m-2 text-right   place-content-center " viewBox="0 0 512 512" xmlns="http://www.w3.org/2000/svg"><path d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"/></svg>
                <div class=" text-right  text-gray-400">
                    Active
                </div>
            </div>             
         {/if}
     </div>
     <div class=" h-12 w-12 my-1 mx-2">
          <img src="{ $MyPro["ProfileImage"]}" alt="" class=" rounded-2xl hover:rounded-xl active:rounded-md object-cover w-full h-full hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer  active:ring  active:ring-offset-base-50  active:ring-blue-600">
         
     </div>
     <div class="my-2 mr-0 ml-0">
         <svg class="w-7 h-7 mt-1 hover:text-white text-slate-400 transition-all duration-200 ease-linear cursor-pointer" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M10 6a2 2 0 110-4 2 2 0 010 4zM10 12a2 2 0 110-4 2 2 0 010 4zM10 18a2 2 0 110-4 2 2 0 010 4z"></path></svg>
     </div>


            
</div>