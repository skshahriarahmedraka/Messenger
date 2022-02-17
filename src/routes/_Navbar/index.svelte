<script lang="ts">
    
    import {showPeopleList} from "$lib/store"

    let showPeopleList1 
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


</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

 <div class=" h-14 w-full ">
            <div class="navbar bg-base-100">
            <div class="flex-1">
                <a class="btn btn-ghost normal-case text-xl">daisyUI</a>
                <button class=" w-8 h-8 " on:click={showPeopleListFunc}> 
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"></path></svg>
                </button>
            </div>
            <div class="flex-none gap-2">
                <div class="form-control">
                <input type="text" placeholder="Search" class="input input-bordered">
                </div>
                <div class="dropdown dropdown-end">
                <label tabindex="0" class="btn btn-ghost btn-circle avatar">
                    <div class="w-10 rounded-full">
                    <img src="https://api.lorem.space/image/face?hash=33791">
                    </div>
                </label>
                <ul tabindex="0" class="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52">
                    <li>
                    <a class="justify-between">
                        Profile
                        <span class="badge">New</span>
                    </a>
                    </li>
                    <li><a>Settings</a></li>
                    <li><a>Logout</a></li>
                </ul>
                </div>
            </div>
            </div>
        </div>