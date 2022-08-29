<script lang="ts" context="module">
</script>

<script lang="ts">
	import Camera from '$lib/icons/camera.svelte';
	import Errorsign from '$lib/icons/Errorsign.svelte';
import RoundPlus from '$lib/icons/roundPlus.svelte';

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

	let AdminObj={
		input: '' as string,
		focus: false as boolean,
		list: [] as any,
		suggest: [] as any,
	};
	// AdminObj.input = '';
	// AdminObj.focus = false;

	AdminObj.list = [
		{ name: 'sk shahriar ahmed raka', UserId: 'skrajhkka', ProfileImage: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png' },
		{ name: 'sk  ahmed raka', UserId: 'skghjkhraka', ProfileImage: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png' },
		{ name: 'ahmed raka', UserId: 'skrajghka', ProfileImage: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png' },
		{ name: 'sk shahriar  raka', UserId: 'shjkghkraka', ProfileImage: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png' }
	];
	AdminObj.suggest = [
		{ name: 'what sk shahriar ahmed ', UserId: 'svghjkraka', ProfileImage: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png' },
		{ name: 'why shahriar ahmed raka', UserId: 'skhgvraka', ProfileImage: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png' },
		{ name: 'how shahriar ahmed raka', UserId: 'skhgjraka', ProfileImage: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png' },
		{ name: 'never shahriar ahmed raka', UserId: 'shgvjkraka', ProfileImage: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png' }
	];
	console.log("🚀 ~ file: index.svelte ~ line 43 ~ AdminObj", AdminObj)

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

<div
	class=" flex  h-screen w-full overflow-hidden bg-[#36393f] md:flex-col md:overflow-y-scroll lg:flex-row "
>
	<div class=" h-fit justify-center  md:w-full lg:w-1/2 lg:overflow-y-scroll ">
		<div class="flex w-full flex-col ">
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
			<!-- banner -->
			<div class="flex flex-row  justify-center ">
				{#if avatar}
					<img
						src={avatar}
						alt="ProfileImage"
						class=" mt-6 aspect-[16/9] w-[80%] self-center rounded-lg object-cover"
					/>
				{:else}
					<img
						src="https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png"
						alt="ProfileImage"
						class=" mt-6 aspect-[16/9] w-[80%]  self-center rounded-lg object-cover "
					/>
				{/if}
				<button class=" absolute   self-end   " on:click={() => FileInput.click()}>
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
						<img
							src="https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png"
							alt="ProfileImage"
							class=" aspect-[1/1] w-[80%] self-center rounded-lg object-cover  "
						/>
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
		<!-- admins -->
		<div class="flex flex-col justify-center  object-center ">
			<p class="self-center text-left  text-xl font-semibold text-gray-400">Admins:</p>
			<div class=" flex flex-row flex-wrap w-fit gap-2 self-center">
				{#each AdminObj.list as t}
					<li
						class=" bg-slate-600 rounded-lg h-10 w-fit flex flex-row "
					>
					<!-- class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200" -->
						<!-- {t} -->
						<img src={t.ProfileImage} alt="" class=" w-6 h-6 object-cover rounded-lg">
						<div class=" flex flex-col">
							<p class=" text-sm">{t.name}</p>
							<p class=" text-xs">{t.UserId}</p>
						</div>
					</li>
				{/each}
			</div>
			<!-- input admin name and suggession -->
			<div class="flex w-[65%] flex-col gap-1 self-center">
				<div class=" flex h-16 w-full flex-row">
					<RoundPlus/>
					<input
						on:keypress={onChangeAdminObj}
						maxlength="20"
						bind:value={AdminObj.input}
						on:focus={() => {
							AdminObj.focus = true;
						}}
						on:blur={() => {
							AdminObj.focus = false;
						}}
						type="text"
						class=" mx-2 mt-3 inline h-fit w-full  grow rounded-lg  border-2   bg-[#45494e] bg-inherit p-2  text-lg  font-medium text-[#98999e]  outline-none outline-0 focus:border-2 focus:shadow-none  focus:border-[#3ba55d] focus:outline-none active:border-none  active:outline-none "
					/>
				</div>
				<!-- render seggestions -->
				{#each AdminObj.suggest as i}
					<div
						class="ml-5 flex h-16 w-[85%] cursor-pointer   flex-row items-center gap-3 self-center rounded-lg bg-slate-600 pl-3 hover:bg-gray-500 active:bg-gray-400"
					>
						<img
							src="https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png"
							alt="ProfileImage"
							class=" h-12 w-12 rounded-lg object-cover"
						/>
						<div class="flex flex-col">
							<p class="">{i.name}</p>
							<p class="">@{i.UserId}</p>
						</div>
					</div>
				{/each}
			</div>
		</div>
	</div>
	<!-- middle bar -->
	<div class=" h-[80%] w-[1px] self-center bg-slate-500 md:hidden lg:block" />

	<div class=" flex h-fit flex-col md:w-full lg:w-1/2 lg:overflow-y-scroll ">
		<!--  -->
	</div>
</div>
