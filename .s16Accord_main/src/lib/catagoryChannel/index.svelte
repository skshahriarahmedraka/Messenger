<script lang="ts">
	import AeroDown from '$lib/icons/aeroDown.svelte';
	import Message from '$lib/icons/message.svelte';
	import serverIconMessage from '$lib/icons/serverCatagory/message.svelte';
	import RoundPlus from '$lib/icons/roundPlus.svelte';
	import Dropdown from './dropdown.svelte';
	let catagoryChannelList = [
		{
			name: 'offensive ',
			icon: serverIconMessage,
			type: 'chat',
			channels: [
				{ cName: 'certification', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'classroom', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'voiceChat', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'Video Chat', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'gossip', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'future goal', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'acchivement', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' }
			]
		},
		{
			name: 'Security ',
			icon: serverIconMessage,
			type: 'chat',
			channels: [
				{ cName: 'certification', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'classroom', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'voiceChat', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'Video Chat', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'gossip', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'future goal', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' },
				{ cName: 'acchivement', cIcon: serverIconMessage, cType: 'chat', cLink: 'https://' }
			]
		}
	];
	let li = { 1: serverIconMessage };
	let AdminObj = {
		input: '' as string,
		focus: false as boolean,
		list: [] as any,
		suggest: [] as any
	};
	function onChangeAdminObj(e: { keyCode: any }) {
		switch (e.keyCode) {
			case 32:
			case 13:
				onKeyTag();
				break;
		}
	}
	function onKeyTag() {
		if (AdminObj.input.trim() != ('' as String)) {
			AdminObj.list = [...AdminObj.list, AdminObj.input.trim()];
			AdminObj.input = '' as string;
		}
	}
</script>

<p class=" mt-4 ml-4 overflow-hidden text-4xl text-gray-300 ">Create Catagoris & Channel</p>
{#each catagoryChannelList as i}
	<div class="flex  flex-col justify-center mt-5">
		<!-- catagoris -->
		<div class="-ml-28 flex flex-row justify-center text-2xl text-gray-200">
			<svelte:component this={i.icon} class="h-10 w-10 fill-slate-400" />
			<p class="">{i.name}</p>
			<p class="">{'(' + i.type + ')'}</p>
		</div>
		<!-- channel -->

		{#each i.channels as j}
			<div
				class="ml-6 flex h-8 w-64  flex-row justify-start self-center text-left text-lg text-gray-300  "
			>
				<svelte:component this={j.cIcon} class="h-7 w-7  fill-slate-400" />
				<p class="">{j.cName}</p>
				<p class="">{'(' + j.cType + ')'}</p>
			</div>
		{/each}
         
            <div class="  ml-6 flex h-8 w-fit  flex-row justify-start self-center text-left text-lg ">
                <p class=" text-base text-slate-200 w-36 mt-1">Add channel :</p>
            <RoundPlus on:click={()=>{}} class="hover:cursor-pointer  hover:stroke-green-400 active:stroke-green-600" />
                
                <input
                    
                    type="text"
                    class=" mx-2 my-2 h-8 w-64  self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
                />
                <Dropdown />
                <!-- class=" mx-2 mt-3 inline h-fit w-full  grow rounded-lg  border-2   bg-[#45494e] bg-inherit p-2  text-lg  font-medium text-[#98999e]  outline-none outline-0 focus:border-2	 focus:shadow-none  focus:border-[#3ba55d] focus:outline-none active:border-none  active:outline-none " -->
            </div>
        

	</div>
    {/each}
    <div class="flex flex-col  mt-5 ml-4">
        <p class="text-2xl text-slate-200">Create Catagory :</p>
        <div class=" flex h-16 w-full flex-row ">
            <RoundPlus on:click={()=>{}} class="hover:cursor-pointer hover:stroke-green-400 active:stroke-green-600  " />
            <input
                
                type="text"
                class=" mx-4 my-2 h-10 w-2/3  self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
            />
            <Dropdown />
            <!-- class=" mx-2 mt-3 inline h-fit w-full  grow rounded-lg  border-2   bg-[#45494e] bg-inherit p-2  text-lg  font-medium text-[#98999e]  outline-none outline-0 focus:border-2	 focus:shadow-none  focus:border-[#3ba55d] focus:outline-none active:border-none  active:outline-none " -->
        </div>
    </div>

<!-- <div class="mt-4 ml-10 flex h-28 w-full flex-row justify-start gap-2 ">
	<Message class="mt-2 h-6 w-6 fill-slate-400" />
	<div class=" text-2xl text-slate-200">offensive Chats</div>
	<AeroDown class="mt-2 h-6 w-6 fill-slate-400" />
</div>
<div class="ml-4  h-28 ">
	<svelte:component this={li[1]} />
</div> -->
<style>
	/* your styles go here */
</style>
