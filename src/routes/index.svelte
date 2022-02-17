<script context=module>
     
</script>
<script lang="ts">
    
    import peopleIcon from "./_icons/users.svg"
    import Dock from "./_Dock/index.svelte"
    import Servers from "./_Servers/index.svelte"
    import NavBar from "./_Navbar/index.svelte"
    import Peoples from "./_Peoples/index.svelte"
    import Messages from "./_Messages/index.svelte"
    import {showPeopleList} from "$lib/store"
   
    
      
    let showPeopleListValue

    showPeopleList.subscribe(val=>{
        showPeopleListValue=val 
    }) 


   



</script>
<style>

</style>


<div class="flex flex-row h-screen">
    <!-- left avater bar -->
    <Dock/>
    <!--  -->

    <!-- hash bar -->
    <!-- 132px -->
    <Servers/>
    
    <!-- everything -->
    <div class=" w-full bg-stone-600 flex flex-col">
        <!-- nav bar -->
        <NavBar/>
       
        <!-- message and people -->
        <div class=" w-full h-full bg-[#2f3136] flex flex-row">
            <!-- messaging 830px -->
            <Messages/>
            <!-- people -->
            {#if showPeopleListValue!=0 }
                 <Peoples/>
            {/if}
            
        </div>
    </div>
    
</div>