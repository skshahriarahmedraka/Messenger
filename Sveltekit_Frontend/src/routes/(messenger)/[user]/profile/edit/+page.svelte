<script lang="ts" context="module">
</script>

<script lang="ts">
	import Camera from '$lib/icons/camera.svelte';
	import Errorsign from '$lib/icons/Errorsign.svelte';
	import RoundPlus from '$lib/icons/roundPlus.svelte';
import AddAdmin from '$lib/Servers/create components/AddAdmin.svelte';
import Addmoderator from '$lib/Servers/create components/Addmoderator.svelte';
import Accord_Default from "$lib/icons/accord default.svelte"
import Accord from "$lib/Dock/icons/accord.svelte"
import AeroDown from "$lib/icons/aeroDown.svelte"
import Bd from "$lib/icons/bd.svelte"
import AddMember from  "$lib/Servers/create components/AddMember.svelte"
import Message from "$lib/icons/message.svelte"
import CatagoryChannel from "$lib/catagoryChannel/index.svelte"
	import WalpaperLogo from '$lib/icons/walpaperLogo.svelte';

	let avatar: any;
	let FileInput: any;
	const onFileSelected = (e: any) => {
		let image = e.target.files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		reader.onload = (e) => {
			avatar = e.target?.result;
		};
	};

	let UserAvatar: any;
	let UserFileInput: any;
	const onFileSelectedUser = (e: any) => {
		let image = e.target.files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		reader.onload = (e) => {
			UserAvatar = e.target?.result;
		};
	};

	let ServerData = {
		Name: 'Offensive Security',
		UserId: 'offsec'
	};

	let ErrorMsg = {
		Name: [true, "User Name is't valid"],
		UserId: [true, 'User Id is not Available']
	};

	function OnSubmit() {}

	
const imgUrl = new URL('Accord', import.meta.url).href



</script>

<div
	class=" flex  h-screen w-full overflow-hidden bg-[#36393f] md:flex-col md:overflow-y-scroll lg:flex-row "
>
	<div class=" no-scrollbar h-fit flex flex-col items-center md:w-full lg:w-1/2 lg:overflow-y-scroll mb-6 ">
		<!-- save & discard button -->
		<div class=" absolute right-0">
			<button
				class=" m-4 h-12 w-40 self-center rounded-xl bg-[#0077d3] text-xl font-bold text-gray-300 hover:bg-sky-600 active:bg-sky-500"
			>
				save
			</button>

			<button
				class=" m-4 h-12 w-40 self-center rounded-xl border-[1px] border-red-500 bg-transparent text-xl font-bold text-red-500 hover:bg-red-300 hover:bg-opacity-50 active:bg-red-200 active:bg-opacity-30"
			>
				discard
			</button>
		</div>
		<div class="flex w-[80%] flex-col   justify-center items-center  ">
			<!-- banner -->
			<div class=" flex  flex-row justify-center w-full relative ">
				{#if avatar}
					<!-- banner image -->
					<img
						src={avatar}
						alt="ProfileImage"
						class=" mt-6 aspect-[16/9] w-[80%] self-center rounded-lg object-cover"
					/>
				{:else}
				
				<div
					class="mt-6 w-full aspect-[16/9]   self-center rounded-lg object-fill bg-transparent flex  items-center justify-center  border-2 border-slate-400 border-opacity-30 "
				>
					<WalpaperLogo class=" h-20 w-20 stroke-slate-400 " />
				</div>
			
					<!-- <Accord class="mt-6 w-[80%] aspect-[16/9] h-80 self-center rounded-lg object-fill bg-slate-600 " /> -->
					
				{/if}
				<button class=" absolute   self-end h-full w-full  " on:click={() => FileInput.click()}>
					<!-- <img src="https://static.thenounproject.com/png/625182-200.png" alt="uploadImage" class=" h-28 w-28"> -->
					<!-- banner camera logo -->
					<!-- <Camera class="h-10 w-10 stroke-slate-100 stroke-1 opacity-50  hover:opacity-100" /> -->
					<input
						bind:this={FileInput}
						on:change={(e) => onFileSelected(e)}
						type="file"
						class=" hidden h-full w-full "
						accept=".jpg, .jpeg, .png, .svg"
					/>
				</button>
			</div>
			<!-- profile image , name , id -->
			<div class=" flex w-full flex-row  justify-center  ">
				<!-- profile image -->
				<div class="  relative mt-6 flex  w-1/4  justify-center">
					{#if UserAvatar}
						<img
							src={UserAvatar}
							alt="ProfileImage"
							class=" aspect-[1/1]   w-[60%] self-center rounded-lg object-cover "
						/>
					{:else}
					<Accord class=" aspect-[1/1] w-44 self-center rounded-lg object-cover  bg-gray-600 " />

						<!-- <img
							src="https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png"
							alt="ProfileImage"
							class=" aspect-[1/1] w-[80%] self-center rounded-lg object-cover  "
						/> -->
					{/if}
					<button class=" absolute  self-end   " on:click={() => UserFileInput.click()}>
						<!-- <img src="https://static.thenounproject.com/png/625182-200.png" alt="uploadImage" class=" h-28 w-28"> -->
						<Camera class="h-10 w-10 stroke-slate-100 stroke-1 opacity-50  hover:opacity-100" />
						<input
							bind:this={UserFileInput}
							on:change={(e) => onFileSelectedUser(e)}
							type="file"
							class=" hidden h-10 w-10 "
							accept=".jpg, .jpeg, .png, .svg"
						/>
					</button>
				</div>
				<!-- name , id -->
				<div class=" mt-6  w-2/4 ">
					<!--  -->
					<!-- Name  -->
					<div class=" flex flex-row self-start text-left text-sm font-semibold text-gray-300 ">
						<div class=" ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
							Name :
						</div>
						{#if ErrorMsg.Name[0]}
							<div class=" mt-2 ml-3 inline-flex ">
								<Errorsign class="h-6 w-6   fill-red-500" />
								<p class="text-red-300">{ErrorMsg.Name[1]}</p>
							</div>
						{/if}
					</div>
					<input
						on:keyup={(e) => (e.key === 'Enter' ? OnSubmit() : '')}
						bind:value={ServerData.Name}
						type="text"
						class=" mx-4 my-2 h-12 w-[415px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
					/>
					<!-- User id -->
					<div class=" flex flex-row self-start text-left text-sm font-semibold text-gray-300 ">
						<div class=" ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
							User Id :
						</div>
						{#if ErrorMsg.UserId[0]}
							<!-- content here -->
							<div class=" mt-2 ml-3 inline-flex ">
								<Errorsign class="h-6 w-6   fill-red-500" />
								<p class="text-red-300">{ErrorMsg.UserId[1]}</p>
							</div>
						{/if}
					</div>
					<input
						on:keyup={(e) => (e.key === 'Enter' ? OnSubmit() : '')}
						bind:value={ServerData.UserId}
						type="text"
						class=" mx-4 my-2 h-12 w-[415px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
					/>
				</div>
			</div>
		</div>
		<AddAdmin />
		<Addmoderator />
		<!-- ADD MEMBER -->
		<AddMember />
	</div>
	<!-- middle bar -->
	<div class=" h-[80%] w-[1px] self-center bg-slate-500 md:hidden lg:block " />
	
	<div class="no-scrollbar h-fit  justify-center md:w-full lg:w-1/2 lg:overflow-y-scroll mb-6">
		<!--  -->
		<CatagoryChannel />
	</div>
</div>

