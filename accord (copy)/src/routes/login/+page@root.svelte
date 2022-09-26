<script lang="ts">
	import bg1 from './_bg/Colored Shapes.svg';
	import bg2 from './_bg/Icon Grid.svg';
	import bg3 from './_bg/Sprinkle.svg';
	import QrCode from './_bg/qrcode.svelte';
	import Errorsign from '$lib/icons/Errorsign.svelte';

	// import Qrcode from "./_bg/qrcode.svelte"
	// your script goes here

	import { createEventDispatcher } from 'svelte';
	import { goto } from '$app/navigation';
	import Eye from '$lib/icons/eye.svelte';
	import EyeClosed from '$lib/icons/eyeClosed.svelte';

	let ShowPass: boolean = true;
	const dispatch = createEventDispatcher();
	let ErrorMsg = {
		Email: [false, 'invalid/wrong email'],
		Password: [false, 'Password is too short']
	};
	let LoginData = {
		Email: '',
		Password: ''
	};

	// let EmailOrMobileOrUserID: string = '';
	// let Mobile:number
	// let Password: string = '';
	// let InputError: boolean = false;
	const ValidateEmail = (E: string) => {
		return String(E)
			.toLowerCase()
			.match(
				/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
			);
	};
	const ValidateMobile = (M: string) => {
		return String(M).toLowerCase().match('^[0-9]+$');
	};

	function ErrorChecking() {
		if (LoginData.Email.trim() === '') {
			ErrorMsg.Email[1] = 'Email is Empty';
			ErrorMsg.Email[0] = true;
		} else if (!ValidateEmail(LoginData.Email)) {
			ErrorMsg.Email[1] = 'invalid/wrong email or password';
			ErrorMsg.Email[0] = true;
		} else {
			ErrorMsg.Email[0] = false;
		}
		if (LoginData.Password.trim() === '') {
			ErrorMsg.Password[1] = 'Password is Empty';
			ErrorMsg.Password[0] = true;
		} else if (LoginData.Password.length < 8 || LoginData.Password.length > 30) {
			ErrorMsg.Password[1] = 'Password is Too short';
			ErrorMsg.Password[0] = true;
		} else {
			ErrorMsg.Password[0] = false;
		}
		return ErrorMsg.Email[0] && ErrorMsg.Password[0];
	}
	async function SetToDefault() {
		LoginData = {
			Email: '',
			Password: ''
		};
		ErrorMsg = {
			Email: [false, 'invalid/wrong email'],
			Password: [false, 'Password is too short']
		};
		ShowPass = true;
	}
	async function OnSubmit() {
		

		console.log('🚀 ~ file: index@root.svelte ~ line 45 ~ OnSubmit ~ OnSubmit');
		if (!ErrorChecking()) {
			let res = await fetch('http://localhost:8888/login', {
				// credentials: 'same-origin',

				method: 'POST',
				mode: 'no-cors',
				credentials: 'include',
				body: JSON.stringify(LoginData)
			}).then((res)=>{
            	console.log("🚀 ~ file: index@root.svelte ~ line 93 ~ OnSubmit ~ res", res)
				location.replace("/")
				
			}).catch((err)=>{
            console.log("🚀 ~ file: login index@root.svelte ~ line 93 ~ OnSubmit ~ err", err)
			ErrorMsg.Email[1] = `${err}`;
			ErrorMsg.Email[0] = true;
			
		})
		// console.log('🚀 ~ file: login index@blank.svelte ~ line 45 ~ OnSubmit ~ res', res);
			// if (res!=null) {
			// // let value= await res.json()
			// console.log('🚀 ~ file: login index.svelte ~ line 49 ~ OnSubmit ~ value : ');
			// // LoginData = {
			// // 	Email: '',
			// // 	Password: ''
			// // };
			// // let s:string="/"+String(value.ID)

			// 	goto('/');
			// }
		}
	}
</script>

<!-- markup (zero or more items) goes here -->
<svelte:head>
	<title>Accord Login</title>
</svelte:head>

<div class=" w-screen h-screen  grid place-items-center" style="background-image: url('{bg3}');">
	<!-- register button -->
	<a href="/register" class="absolute right-0 top-0  m-10">
		<button
			class="     h-12 w-60 bg-sky-700 active:bg-sky-500 hover:bg-sky-600 rounded-3xl text-xl font-bold text-gray-300"
		>
			Register
		</button>
	</a>

	<div class=" w-[785px] h-[410px] bg-[#36393f] flex flex-row rounded-lg ">
		<!-- login -->
		<div class=" w-[490px] h-full flex flex-col  ">
			<div class=" text-2xl font-semibold text-gray-300  mt-4 mb-2  self-center">
				{"✨ Welcome back 🔥"}
			</div>
			<div class=" text-sm font-medium text-gray-400  mb-4  self-center">
				We're so excited to see you again!
			</div>
			<!-- email or  mobile -->
			<div class=" flex flex-row text-left text-sm font-semibold text-gray-300 self-start ">
				<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
					Email :
				</div>
				{#if ErrorMsg.Email[0]}
					<!-- content here -->
					<div class=" mt-2 ml-3 inline-flex ">
						<Errorsign class="h-6 w-6   fill-red-500" />
						<p class="text-red-300">{ErrorMsg.Email[1]}</p>
					</div>
				{/if}
			</div>
			<input
			on:keyup={e=>e.key==='Enter' ? OnSubmit(): ""}

				bind:value={LoginData.Email}
				type="text"
				class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
			/>
			<!-- password -->
			<div class=" flex flex-row  text-left text-sm font-semibold text-gray-300 self-start  mt-2 ">
				<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
					Password :
				</div>
				{#if ErrorMsg.Password[0]}
					<!-- content here -->
					<div class=" mt-2 ml-3 inline-flex ">
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
				on:keyup={e=>e.key==='Enter' ? OnSubmit(): ""}

					value={LoginData.Password}
					on:input={(e)=>{LoginData.Password = e.target.value }}
					type={ ShowPass ? "password" : "text"}
					class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 mt-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
				/>
			</div>
			<!-- forget password -->
			<div
				class="text-left text-xs text-sky-600 hover:text-sky-300 active:text-sky-500 cursor-pointer hover:underline hover:underline-offset-2 ml-10 mt-1"
			>
				Forget password?
			</div>
			<!-- submit -->
			<button
				on:click={() => OnSubmit()}
				
				class=" m-4 self-center h-12 w-60 bg-sky-700 active:bg-sky-500 hover:bg-sky-600 rounded-xl text-xl font-bold text-gray-300"
			>
				Login
			</button>
			<!-- register -->
			<div class="ml-10 mt-4 text-xs text-slate-500 ">
				Need a account? <p
					class="inline text-left text-xs text-sky-600 hover:text-sky-300 active:text-sky-500 cursor-pointer hover:underline hover:underline-offset-2 "
				>
					Register
				</p>
			</div>
		</div>
		<!-- qrcode -->
		<div class="w-[295px] h-full   flex flex-col ">
			<!-- <div class=" w-[180px] h-[180px]" style="background-image: url('{QrCode}');"></div> -->

			<QrCode />

			<div class=" text-2xl font-semibold text-gray-300 mt-4 grid justify-items-center">
				Login with QR Code
			</div>
			<div class=" text-sm text-[#8a8c90] m-6 grid justify-items-center">
				Scan with <p class="font-semibold inline text-gray-400">Accord Mobile app</p>
				to login instantly
			</div>
		</div>
	</div>
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
		transition: background-color 5000s ease-in-out 0s;
	}
</style>
