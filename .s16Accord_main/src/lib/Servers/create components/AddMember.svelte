<script lang="ts">
import Accord from "$lib/Dock/icons/accord.svelte"

</script>

<div class=" m-2 h-fit  w-[90%] rounded-lg  border-2 border-sky-400 ">
    
	<div
		class=" my-2 ml-1 mr-2 flex flex-grow items-center cursor-pointer flex-row transition-all duration-300 ease-linear "
	>
		<input
			placeholder="Search"
			autocomplete="off"
			type="text"
			class="     my-2   ml-4 mr-2  h-12 w-[50%] self-center rounded-2xl border-2 border-[#24262b]  bg-[#303338] p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
		<svg
			class=" h-8 w-8 place-content-center rounded-xl fill-slate-400 hover:fill-white"
			fill="currentColor"
			viewBox="0 0 20 20"
			xmlns="http://www.w3.org/2000/svg"
			><path
				fill-rule="evenodd"
				d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
				clip-rule="evenodd"
			/></svg
		>
        <div class=" grow"></div>
        <!--  -->
        <p class=" text-lg font-semibold mr-4 text-slate-400 "> {50234} Members</p>
	</div>
    <div class="">
        <div
            class="ml-5 flex h-16 w-fit  cursor-pointer   flex-row items-center gap-3 self-center rounded-lg bg-[#292b2f] pl-3 hover:bg-gray-700 active:bg-gray-600"
        >   
        {#if true}
                <Accord class=" h-10 w-10 rounded-lg object-cover  bg-slate-600 " />
        {:else}
             <!-- else content here -->
             <!-- <img
                 src={i.ProfileImage}
                 alt="ProfileImage"
                 class=" h-12 w-12 rounded-lg object-cover"
             /> -->
        {/if}
            <div class="flex flex-col mr-4 ">
                <p class=" overflow-hidden text-base text-slate-300 line-clamp-1 ">{"sk shahriar ahmed raka"}</p>
                <p class="cursor-pointer text-sm text-gray-500 line-clamp-1 hover:text-gray-400">
                    @{"skssar"}
                </p>
            </div>
        </div>
    </div>
</div>

<style>
	/* your styles go here */
</style>
