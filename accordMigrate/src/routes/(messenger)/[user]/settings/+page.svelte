<script lang="ts">
	import CameraPlus from '$lib/icons/cameraPlus.svelte';
    import Datepicker from '$lib/datepicker/index.svelte';
	import Errorsign from '$lib/icons/Errorsign.svelte';
    import Eye from '$lib/icons/eye.svelte';
	import EyeClosed from '$lib/icons/eyeClosed.svelte';
	let myData = {
		UserID: '',
		UserName: '',
		ProfileImage: '',
		BannerImage: '',
		Email: '',
		Password: '',
		NumOfFriends: '',
		NumOfConnectedServer: '',
		Mobile: '',
		Birthday: 'June 22, 2022'
	};

	let ThisBannerImage: any;
	let ThisBannerImgInput: any;

	let ThisProfileImage: any;
	let ThisProfileImageInput: any;
    let ShowPass: boolean = true;

	let datedrop: boolean = false;
	// let BirthDate: string = 'June 22, 2022';
    let ErrorMsg = {
		UserID: [false, ''],
		UserName: [false, ''],
		ProfileImage: [false, ''],
		BannerImage: [false, ''],
		Email: [false, ''],
		Password: [false, ''],
		NumOfFriends: [false, ''],
		NumOfConnectedServer: [false, ''],
		Mobile: [false, ''],
		Birthday: [false, '']
	};
	// let ErrorMsg2 = {
	// 	Email: [false, ''],
	// 	Username: [false, ''],
	// 	Userid: [false, ''],
	// 	Mobile: [false, ''],
	// 	Password: [false, ''],
	// 	Birth: [false, '']
	// };
	// let RegData2 = {
	// 	Email: '',
	// 	Username: '',
	// 	Userid: '',
	// 	Mobile: '',
	// 	Password: '',
	// 	Birth: ''
	// };
	const ValidateEmail = (E: string) => {
		return String(E)
			.toLowerCase()
			.match(
				/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
			);
	};
	const ValidateUsername = (E: string) => {
		console.log(
			'🚀 ~ file: index@root.svelte ~ line 44 ~ Username ~ E.match',
			E.match('[a-z A-Z]*')
		);
		return E.match('[a-z A-Z]*');
	};
	const ValidateMobile = (M: string) => {
		// console.log("🚀 ~ file: index@root.svelte ~ line 48 ~ ValidateMobile ~ String(M).toLowerCase().match('^[0-9]+$')",
		// 	String(M).toLowerCase().match('^[0-9]+$')
		// );
		return String(M).toLowerCase().match('[0-9]*');
	};
	const ValidateUserid = (M: string) => {
		return String(M).toLowerCase().match('[a-zA-Z0-9]*');
	};
	const ValidateBirth = (M: string) => {
		return String(M).toLowerCase().match('[a-zA-Z0-9]*');
	};
	function ErrorChecking() {
		// Email Validation check
		if (myData.Email.trim() === '') {
			ErrorMsg.Email[1] = 'Email is Empty';
			ErrorMsg.Email[0] = true;
		} else if (!ValidateEmail(myData.Email)) {
			ErrorMsg.Email[1] = 'invalid/wrong email or password';
			ErrorMsg.Email[0] = true;
		} else {
			ErrorMsg.Email[0] = false;
		}
		// Username Validation Check
		if (myData.UserName.trim() === '') {
			ErrorMsg.UserName[1] = 'Name is Empty';
			ErrorMsg.UserName[0] = true;
		} else if (!ValidateUsername(myData.UserName)) {
			ErrorMsg.UserName[1] = 'invalid Username';
			ErrorMsg.UserName[0] = true;
		} else {
			ErrorMsg.UserName[0] = false;
		}
		// Userid Valication check
		if (myData.UserID.trim() === '') {
			ErrorMsg.UserID[1] = 'Userid is Empty';
			ErrorMsg.UserID[0] = true;
		} else if (!ValidateUserid(myData.UserID)) {
			ErrorMsg.UserID[1] = 'invalid Userid';
			ErrorMsg.UserID[0] = true;
		} else {
			ErrorMsg.UserID[0] = false;
		}

		// Password validation check
		if (myData.Password.trim() === '') {
			ErrorMsg.Password[1] = 'Password is Empty';
			ErrorMsg.Password[0] = true;
		} else if (myData.Password.length < 8 || myData.Password.length > 30) {
			ErrorMsg.Password[1] = 'Password is Too short';
			ErrorMsg.Password[0] = true;
		} else {
			ErrorMsg.Password[0] = false;
		}
		//  mobile Validation check

		if (myData.Mobile.trim() === '') {
			ErrorMsg.Mobile[1] = 'Mobile is Empty';
			ErrorMsg.Mobile[0] = true;
		} else if (!ValidateMobile(myData.Mobile)) {
			ErrorMsg.Mobile[1] = 'invalid mobile';
			ErrorMsg.Mobile[0] = true;
		} else {
			ErrorMsg.Mobile[0] = false;
		}
		// birthday validation check
		if (myData.Birthday.trim() === '') {
			ErrorMsg.Birthday[1] = 'Birth is Empty';
			ErrorMsg.Birthday[0] = true;
		} else if (!ValidateBirth(myData.Birthday)) {
			ErrorMsg.Birthday[1] = 'invalid birthday';
			ErrorMsg.Birthday[0] = true;
		} else {
			ErrorMsg.Birthday[0] = false;
		}

		return (
			!ErrorMsg.Email[0] &&
			!ErrorMsg.Password[0] &&
			!ErrorMsg.UserName[0] &&
			!ErrorMsg.UserID[0] &&
			!ErrorMsg.Mobile[0] &&
			!ErrorMsg.Birthday[0]
		);
	}

	const onBannerImgUpload = (e: any) => {
		let image = e.target.files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		let name = e.target.files[0].name;
		reader.onload = (e) => {
			myData.BannerImage = String(e.target?.result);

			UploadBannerImage(String(e.target?.result), name);
		};
	};

	async function UploadBannerImage(img: string, name: string) {
		const data = { img: '', name: '' };
		const imgData = img.split(',');
		data['img'] = imgData[1];
		data['name'] = name;

		// console.log(data);
		const res = await fetch('/api/imgupload', {
			method: 'POST',
			// mode: 'no-cors',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			},
			body: JSON.stringify(data)
		});
		const imgLink = await res.json();
		console.log('🚀 ~ file: inputGameProfile.svelte ~ line 56 ~ UploadImage ~ imgLink', imgLink);
		myData.BannerImage = imgLink.Link;
	}
	const onProfileImgUpload = (e: any) => {
		let image = e.target.files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		let name = e.target.files[0].name;
		reader.onload = (e) => {
			myData.BannerImage = String(e.target?.result);

			UploadBannerImage(String(e.target?.result), name);
		};
	};
	async function UploadProfileImage(img: string, name: string) {
		const data = { img: '', name: '' };
		const imgData = img.split(',');
		data['img'] = imgData[1];
		data['name'] = name;

		// console.log(data);
		const res = await fetch('/api/imgupload', {
			method: 'POST',
			// mode: 'no-cors',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			},
			body: JSON.stringify(data)
		});
		const imgLink = await res.json();
		console.log('🚀 ~ file: inputGameProfile.svelte ~ line 56 ~ UploadImage ~ imgLink', imgLink);
		myData.BannerImage = imgLink.Link;
	}
    async function OnSubmit(){

    }
</script>

<!-- markup (zero or more items) goes here -->

<div
	class=" flex  h-screen w-full overflow-hidden bg-[#36393f] font-Poppins md:flex-col md:overflow-y-scroll lg:flex-row "
>
	<div class=" flex w-1/2 flex-col  items-center overflow-y-scroll ">
		<p class="text-xl font-semibold text-slate-300">Banner Image :</p>
        <!-- banner image -->
		<div class=" relative  flex w-[800px] h-[500px]  aspect-[16/9] flex-row justify-center  ">
			{#if myData.BannerImage != ''}
				<img
					src={myData.BannerImage}
					bind:this={ThisBannerImage}
					on:error={() => {
						ThisBannerImage.src = myData.BannerImage;
					}}
					alt="ProfileImage"
					class=" aspect-[16/9] h-full w-full self-center rounded-lg object-cover "
				/>
			{:else}
				<CameraPlus
					class="  w-full h-full aspect-[16/9] rounded-xl border-2  border-slate-200 stroke-slate-100 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
				/>

			{/if}
			<button
				class=" absolute   h-full   w-full self-end  "
				on:click={() => ThisBannerImgInput.click()}
			>
				<input
					bind:this={ThisBannerImgInput}
					on:change={(e) => onBannerImgUpload(e)}
					type="file"
					class=" hidden h-10 w-10 "
					accept=".jpg, .jpeg, .png, .svg"
				/>
			</button>
		</div>
		<p class="text-xl font-semibold text-slate-300 ">Profile Image :</p>
		<!-- profile image -->
        <div class=" relative  flex h-72 w-96 flex-row justify-center  ">
			{#if myData.ProfileImage != ''}
				<img
					src={myData.ProfileImage}
					bind:this={ThisProfileImage}
					on:error={() => {
						ThisProfileImage.src = myData.ProfileImage;
					}}
					alt="ProfileImage"
					class=" aspect-[1/1]  w-full self-center rounded-lg object-cover "
				/>
			{:else}
				<CameraPlus
					class=" h-full w-full aspect-[1/1] rounded-xl border-2  border-slate-200 stroke-slate-100 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
				/>
			{/if}
			<button
				class=" absolute   h-full   w-full self-end  "
				on:click={() => ThisProfileImageInput.click()}
			>
				<input
					bind:this={ThisProfileImageInput}
					on:change={(e) => onProfileImgUpload(e)}
					type="file"
					class=" hidden h-10 w-10  "
					accept=".jpg, .jpeg, .png, .svg"
				/>
			</button>
		</div>
        <!-- USER NAME -->
		<div class=" flex flex-row text-left w-full text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				User Name :
			</div>
			{#if ErrorMsg.UserName[0]}
				
				<div class=" mt-2 ml-3 flex flex-row w-fit   ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.UserName[1]}</p>
				</div>
			{/if}
		</div>
		<input
			autocomplete="off"
			bind:value={myData.UserName}
			type="text"
			class=" self-center h-12 w-[415px] xs:w-[90%] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
		/>

         <!-- USER ID -->
		<div class=" flex flex-row text-left w-full text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				User ID :
			</div>
			{#if ErrorMsg.UserID[0]}
				
				<div class=" mt-2 ml-3 flex flex-row w-fit   ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.UserID[1]}</p>
				</div>
			{/if}
		</div>
		<input
			autocomplete="off"
			bind:value={myData.UserID}
			type="text"
			class=" self-center h-12 w-[415px] xs:w-[90%] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
		/>

         <!-- USERID -->
		<!-- <div class=" flex flex-row text-left w-full text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				User ID :
			</div>
			{#if ErrorMsg.UserName[0]}
				
				<div class=" mt-2 ml-3 flex flex-row w-fit   ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.UserName[1]}</p>
				</div>
			{/if}
		</div>
		<input
			autocomplete="off"
			bind:value={myData.UserName}
			type="text"
			class=" self-center h-12 w-[415px] xs:w-[90%] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
		/> -->

        <!-- EMAIL -->
        <div class=" flex flex-row text-left w-full text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Email :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 flex flex-row w-fit   ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<input
			autocomplete="off"
			bind:value={myData.Email}
			type="text"
			class=" self-center h-12 w-[415px] xs:w-[90%] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
		/>
        <!-- PASSOWRD -->
		<div class=" flex flex-row text-left w-full text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Password :
			</div>
			{#if ErrorMsg.Password[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 flex flex-row w-fit  ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.Password[1]}</p>
				</div>
			{/if}
		</div>
		<div class="relative h-fit w-full flex justify-center  ">
			<button
				on:click={() => {
					ShowPass = !ShowPass;
				}}
				class=" absolute right-14 top-5  "
			>
				{#if ShowPass}
					<EyeClosed class="h-6 w-6 stroke-gray-500 " />
				{:else}
					<Eye class="h-6 w-6 stroke-gray-500 " />
				{/if}
			</button>
			<input
			id="foo"
				value={myData.Password}
				
				autocomplete="off"
				type={ShowPass ? 'password' : 'text'}
				class=" self-center h-12 xs:w-[90%] w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
			/>
		</div>

        <!-- Mobile Number -->

        <div class=" flex flex-row text-left w-full text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Mobile :
			</div>
			{#if ErrorMsg.Mobile[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 flex flex-row w-fit   ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.Mobile[1]}</p>
				</div>
			{/if}
		</div>
		<input
			autocomplete="off"
			bind:value={myData.Mobile}
			type="text"
			class=" self-center h-12 w-[415px] xs:w-[90%] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
		/>

        <!-- date of birth -->
		<div class=" flex flex-row text-left w-full text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				BirthDay :
			</div>
			{#if ErrorMsg.Birthday[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 flex flex-row w-fit  ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.Birthday[1]}</p>
				</div>
			{/if}
		</div>
		<div
			autocomplete="off"
			on:click={() => {
				datedrop = !datedrop;
			}}
			type="text"
			class=" self-center cursor-pointer xs:w-[90%] h-12   w-[415px]  text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
		>
			{myData.Birthday} 
		</div>

		{#if datedrop}
			<Datepicker class=" top-44 self-center " bind:BirthDate={myData.Birthday} bind:datedrop />
		{/if}
	</div>
	<div class="h-full w-1/2 border-2 border-gray-500 " />
</div>

<style>
	input:-webkit-autofill,
	input:-webkit-autofill:hover,
	input:-webkit-autofill:focus,
	textarea:-webkit-autofill,
	textarea:-webkit-autofill:hover,
	textarea:-webkit-autofill:focus,
	select:-webkit-autofill,
	select:-webkit-autofill:hover,
	select:-webkit-autofill:focus {
		border: 2px solid rgb(14 165 233 / var(--tw-border-opacity));
		-webkit-text-fill-color: #98999e;
		-webkit-box-shadow: 0 0 0px 1000px #303338 inset;
		box-shadow: 0 0 0px 1000px #303338 inset;
		transition: background-color 5000s ease-in-out 0s;
	}
</style>
