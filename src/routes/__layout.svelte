
<script lang="ts" context="module">
    
</script>

<script lang="ts">
	import '../app.css';
    
    // import peopleIcon from "./_icons/users.svg"
    import Dock from "$lib/_Dock/index.svelte"
    import Servers from "$lib/_Servers/index.svelte"
    import NavBar from "$lib/_Navbar/index.svelte"
    import Peoples from "$lib/_Peoples/index.svelte"
    import Messages from "$lib/_Messages/index.svelte"
    import {showPeopleList,ChatOrDock} from "$lib/store2"
    
    import Chat from "$lib/_Chats/index.svelte"


	
    
	
    let showPeopleListValue: number
    let ChatOrDockValue: number

    showPeopleList.subscribe(val=>{
        showPeopleListValue=val 
    }) 

    ChatOrDock.subscribe(val=>{
        ChatOrDockValue=val 
    })


	

</script>
<style>
	
</style>


<div class="flex flex-row h-screen w-full transition-all duration-500 ease-linear ">
    <!-- left avater bar -->
    
    <Dock/>
    
    <slot></slot>
    
</div>







