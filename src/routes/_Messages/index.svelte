<script lang="ts">
  
    import {showPeopleList} from "$lib/store"

    let showPeopleListValue

    showPeopleList.subscribe(val=>{
        showPeopleListValue=val 
    }) 

</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div class=" {showPeopleListValue!=0 ? "w-5/6": "w-full"} transition-all duration-200 ease-linear bg-[#36393f] flex flex-col">
                <!-- all messaging -->
                <div class=" flex-grow ">
                    
                </div>
                <!-- write message -->
                <div class=" flex-none min-w-full py-2 px-1 ">
                    <input type="text" placeholder="Type here" class="input min-w-full ">
                </div>

</div>