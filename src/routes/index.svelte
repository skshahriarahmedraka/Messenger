<script context=module>
     
</script>
<script lang="ts">
    
    // import peopleIcon from "./_icons/users.svg"
    import Dock from "./_Dock/index.svelte"
    import Servers from "./_Servers/index.svelte"
    import NavBar from "./_Navbar/index.svelte"
    import Peoples from "./_Peoples/index.svelte"
    import Messages from "./_Messages/index.svelte"
    import {showPeopleList,ChatOrDock} from "$lib/store"
    import Dock2 from "./_Dock2/index.svelte"
    import Chat from "./_Chats/index.svelte"
   
    
      
    let showPeopleListValue
    let ChatOrDockValue

    showPeopleList.subscribe(val=>{
        showPeopleListValue=val 
    }) 

    ChatOrDock.subscribe(val=>{
        ChatOrDockValue=val 
    })


   



</script>
<style>

</style>


<div class="flex flex-row h-screen transition-all duration-500 ease-linear ">
    <!-- left avater bar -->
    
    {#if ChatOrDockValue}
         <!-- content here -->
         <Dock/>
     
         <!-- hash bar -->
         <Servers/>
    {:else}
        <Chat/>
         <!-- else content here -->
    {/if}
    

    <!-- <Dock2/> -->
    
    <!-- everything -->
    <div class=" w-full  bg-stone-600 flex flex-col flex-nowrap ">    
        <!-- nav bar -->
        <NavBar/>
       
        <!-- message and people -->
        <div class="  w-full h-full  bg-[#2f3136] flex flex-row overflow-hidden">
            <!-- messaging 830px -->
            <Messages/>
            <!-- people -->
            {#if showPeopleListValue!=0 }
                 <Peoples/>
            {/if}
            
        </div>
    </div>
    
</div>