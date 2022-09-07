<script lang='ts'>
	import RoundPlus from '$lib/icons/roundPlus.svelte';
import Accord from "$lib/Dock/icons/accord.svelte"
    
    let ModeratorObj = {
		input: '' as string,
		focus: false as boolean,
		list: [] as any,
		suggest: [] as any
	};
	// ModeratorObj.input = '';
	// ModeratorObj.focus = false;

	ModeratorObj.list = [
		{
			name: 'sk shahriar ahmed raka',
			UserId: 'skrajhkka',
			ProfileImage:
				''
		},
		{
			name: 'sk  ahmed raka',
			UserId: 'skghjkhraka',
			ProfileImage:
				''
		},
		{
			name: 'ahmed raka',
			UserId: 'skrajghka',
			ProfileImage:
				''
		},
		{
			name: 'sk shahriar  raka',
			UserId: 'shjkghkraka',
			ProfileImage:
				''
		},
		{
			name: 'sk shahriar ahmed raka',
			UserId: 'skrajhkka',
			ProfileImage:
				''
		},
		{
			name: 'sk  ahmed raka',
			UserId: 'skghjkhraka',
			ProfileImage:
				''
		},
		{
			name: 'ahmed raka',
			UserId: 'skrajghka',
			ProfileImage:
				''
		},
		{
			name: 'sk shahriar  raka',
			UserId: 'shjkghkraka',
			ProfileImage:
				''
		}
	];
	ModeratorObj.suggest = [
		{
			name: 'what sk shahriar ahmed ',
			UserId: 'svghjkraka',
			ProfileImage:
				''
		},
		{
			name: 'why shahriar ahmed raka',
			UserId: 'skhgvraka',
			ProfileImage:
				''
		},
		{
			name: 'how shahriar ahmed raka',
			UserId: 'skhgjraka',
			ProfileImage:
				''
		},
		{
			name: 'never shahriar ahmed raka',
			UserId: 'shgvjkraka',
			ProfileImage:
				''
		}
	];
	console.log('🚀 ~ file: index.svelte ~ line 43 ~ ModeratorObj', ModeratorObj);

	function onChangeModeratorObj(e: { keyCode: any }) {
		switch (e.keyCode) {
			case 32:
			case 13:
				onKeyTag();
				break;
		}
	}
	function onKeyTag() {
		if (ModeratorObj.input.trim() != ('' as String)) {
			ModeratorObj.list = [...ModeratorObj.list, ModeratorObj.input.trim()];
			ModeratorObj.input = '' as string;
		}
	}
</script>

<style>
    /* your styles go here */
</style>


<!-- admins -->
<div class="flex flex-col justify-center  object-center ">
    <p class="self-center text-left  text-xl font-semibold text-[#3ba55d] my-3">Moderator:</p>
    <div class=" mx-4 flex w-fit flex-row flex-wrap justify-center gap-2 self-center">
        {#each ModeratorObj.list as t}
            <li
                class="  flex h-10 w-fit cursor-pointer   flex-row items-center gap-2 self-center rounded-lg bg-[#292b2f] pl-3 hover:bg-gray-700 active:bg-gray-600 "
            >
                <!-- class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200" -->
                <!-- {t} -->
				{#if t.ProfileImage.trim()===""}
					<Accord class=" h-6 w-6 rounded-lg object-cover  bg-slate-600 " />

				{:else}
					 <img src={t.ProfileImage} alt="" class=" h-6 w-6 rounded-lg object-cover " />
				{/if}
                <div class=" mr-3 flex flex-col ">
                    <p class=" overflow-hidden text-sm text-slate-300 line-clamp-1  ">{t.name}</p>
                    <p class=" cursor-pointer text-xs text-gray-500 line-clamp-1 hover:text-gray-400">
                        @{t.UserId}
                    </p>
                </div>
                <!-- delete admin -->
                <svg
                    on:click={() => {}}
                    class="mr-2 h-6 w-6 stroke-red-400 hover:stroke-red-600"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                    ><path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
                    /></svg
                >
            </li>
        {/each}
    </div>
    <!-- input admin name and suggession -->
    <div class="flex w-[65%] flex-col gap-1 self-center">
        <div class=" flex h-16 w-full flex-row">
            <RoundPlus />
            <input
                on:keypress={onChangeModeratorObj}
                maxlength="20"
                bind:value={ModeratorObj.input}
                on:focus={() => {
                    ModeratorObj.focus = true;
                }}
                on:blur={() => {
                    ModeratorObj.focus = false;
                }}
                type="text"
                class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
            />
            <!-- class=" mx-2 mt-3 inline h-fit w-full  grow rounded-lg  border-2   bg-[#45494e] bg-inherit p-2  text-lg  font-medium text-[#98999e]  outline-none outline-0 focus:border-2	 focus:shadow-none  focus:border-[#3ba55d] focus:outline-none active:border-none  active:outline-none " -->
        </div>
        <!-- render seggestions -->
        {#each ModeratorObj.suggest as i}
            <div
                class="ml-5 flex h-16 w-[85%] cursor-pointer   flex-row items-center gap-3 self-center rounded-lg bg-[#292b2f] pl-3 hover:bg-gray-700 active:bg-gray-600"
            >   
            {#if i.ProfileImage.trim()===""}
					<Accord class=" h-12 w-12 rounded-lg object-cover  bg-slate-600 " />
            {:else}
                 <!-- else content here -->
                 <img
                     src={i.ProfileImage}
                     alt="ProfileImage"
                     class=" h-12 w-12 rounded-lg object-cover"
                 />
            {/if}
                <div class="flex flex-col">
                    <p class=" overflow-hidden text-base text-slate-300 line-clamp-1 ">{i.name}</p>
                    <p class="cursor-pointer text-sm text-gray-500 line-clamp-1 hover:text-gray-400">
                        @{i.UserId}
                    </p>
                </div>
            </div>
        {/each}
    </div>
</div>