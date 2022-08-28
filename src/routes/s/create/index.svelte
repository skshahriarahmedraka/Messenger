<script lang="ts" context="module">
</script>

<script lang="ts">
	import Camera from '$lib/icons/camera.svelte';
	import Errorsign from '$lib/icons/Errorsign.svelte';

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
		UserId: [true, 'Password is too short']
	};

	function OnSubmit() {}
</script>

<div class=" flex  h-screen w-full flex-row overflow-hidden bg-[#36393f] ">
	<div class=" h-full w-1/2 justify-center overflow-y-scroll ">
		<div class="flex flex-col w-full ">
			<!-- save & discard button -->
			<div class=" absolute right-0">
				<button
					class=" m-4 h-12 w-40 self-center rounded-xl bg-[#0077d3] text-xl font-bold text-gray-300 hover:bg-sky-600 active:bg-sky-500"
				>
					save
				</button>

				<button
					class=" m-4 h-12 w-40 self-center rounded-xl bg-transparent border-[1px] border-red-500 text-xl font-bold text-red-500 hover:bg-red-300 hover:bg-opacity-50 active:bg-opacity-30 active:bg-red-200"
				>
					discard
				</button>
			</div>
			<!-- banner -->
			<div class="flex flex-row  justify-center">
				{#if avatar}
					<img
						src={avatar}
						alt="ProfileImage"
						class=" aspect-video mt-6 w-[80%] self-center rounded-lg object-cover"
					/>
				{:else}
					<img
						src="https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png"
						alt="ProfileImage"
						class=" aspect-video mt-6 w-[80%] self-center rounded-lg object-cover "
					/>
				{/if}
				<button class=" -mt-20 -ml-16 self-end hover:contents  " on:click={() => FileInput.click()}>
					<!-- <img src="https://static.thenounproject.com/png/625182-200.png" alt="uploadImage" class=" h-28 w-28"> -->
					<Camera class="h-10 w-10 stroke-slate-100 stroke-1 opacity-50  hover:opacity-100" />
					<input
						bind:this={FileInput}
						on:change={(e) => onFileSelected(e)}
						type="file"
						class=" hidden h-10 w-10 "
						accept=".jpg, .jpeg, .png, .svg"
					/>
				</button>
			</div>
			<!-- profile image , name , id -->
			<div class=" flex flex-row w-full  ">
				<!-- profile image -->
				<div class="  mt-6 flex w-1/3   justify-center">
					{#if UserAvatar}
						<img
							src={UserAvatar}
							alt="ProfileImage"
							class=" aspect-square mt-6 w-[40%] self-center rounded-lg object-cover"
						/>
					{:else}
						<img
							src="https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png"
							alt="ProfileImage"
							class=" aspect-[1/1]   w-[60%] self-center rounded-lg object-cover  "
						/>
					{/if}
					<button class=" -mt-0 -ml-10  self-end   " on:click={() => FileInput.click()}>
						<!-- <img src="https://static.thenounproject.com/png/625182-200.png" alt="uploadImage" class=" h-28 w-28"> -->
						<Camera class="h-10 w-10 stroke-slate-100 stroke-1 opacity-50  hover:opacity-100" />
						<input
							bind:this={UserFileInput}
							on:change={(e) => onFileSelected(e)}
							type="file"
							class=" hidden h-10 w-10 "
							accept=".jpg, .jpeg, .png, .svg"
						/>
					</button>
				</div>
				<!-- name , id -->
				<div class=" mt-6  self-start w-2/3 ">
					<!--  -->
					<!-- email or  mobile -->
					<div class=" flex flex-row self-start text-left text-sm font-semibold text-gray-300 ">
						<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
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
					<!-- email or  mobile -->
					<div class=" flex flex-row self-start text-left text-sm font-semibold text-gray-300 ">
						<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
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
	</div>
	<!-- middle bar -->
	<div class=" h-[80%] w-[1px] self-center bg-white" />

	<div class=" h-full w-1/2 overflow-y-scroll" />
</div>
