

<script lang="ts">
	import bg1 from './bg/Colored Shapes.svg';
	// import bg2 from './_bg/Icon Grid.svg';
	// import bg3 from './_bg/Sprinkle.svg';
	import QrCode from './bg/qrcode.svelte';
	import Datepicker from '$lib/datepicker/index.svelte';
	import Errorsign from '$lib/icons/Errorsign.svelte';
	import Eye from '$lib/icons/eye.svelte';
	import EyeClosed from '$lib/icons/eyeClosed.svelte';
	import { goto } from '$app/navigation';

	let ShowPass: boolean = true;

	let datedrop: boolean = false;
	// let BirthDate: string = 'June 22, 2022';
	let ErrorMsg = {
		Email: [false, ''],
		UserName: [false, ''],
		Userid: [false, ''],
		Mobile: [false, ''],
		Password: [false, ''],
		BirthDate: [false, '']
	};
	let RegData = {
		// UUID:"",
		Email: '',
		UserName: '',
		Userid: '',
		Mobile: '', 
		Password: '',
		BirthDate: '',
    // Accounttype : "",
	// ProfileImg:"",
	// BannerImg:"",
	// Coin:0,
	// FrinedListID:[]


	};
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
		if (RegData.Email.trim() === '') {
			ErrorMsg.Email[1] = 'Email is Empty';
			ErrorMsg.Email[0] = true;
		} else if (!ValidateEmail(RegData.Email)) {
			ErrorMsg.Email[1] = 'invalid/wrong email or password';
			ErrorMsg.Email[0] = true;
		} else {
			ErrorMsg.Email[0] = false;
		}
		// Username Validation Check
		if (RegData.UserName.trim() === '') {
			ErrorMsg.UserName[1] = 'Name is Empty';
			ErrorMsg.UserName[0] = true;
		} else if (!ValidateUsername(RegData.UserName)) {
			ErrorMsg.UserName[1] = 'invalid UserName';
			ErrorMsg.UserName[0] = true;
		} else {
			ErrorMsg.UserName[0] = false;
		}
		// Userid Valication check
		if (RegData.Userid.trim() === '') {
			ErrorMsg.Userid[1] = 'Userid is Empty';
			ErrorMsg.Userid[0] = true;
		} else if (!ValidateUserid(RegData.Userid)) {
			ErrorMsg.Userid[1] = 'invalid Userid';
			ErrorMsg.Userid[0] = true;
		} else {
			ErrorMsg.Userid[0] = false;
		}

		// Password validation check
		if (RegData.Password.trim() === '') {
			ErrorMsg.Password[1] = 'Password is Empty';
			ErrorMsg.Password[0] = true;
		} else if (RegData.Password.length < 8 || RegData.Password.length > 30) {
			ErrorMsg.Password[1] = 'Password is Too short';
			ErrorMsg.Password[0] = true;
		} else {
			ErrorMsg.Password[0] = false;
		}
		//  mobile Validation check

		if (RegData.Mobile.trim() === '') {
			ErrorMsg.Mobile[1] = 'Mobile is Empty';
			ErrorMsg.Mobile[0] = true;
		} else if (!ValidateMobile(RegData.Mobile)) {
			ErrorMsg.Mobile[1] = 'invalid mobile';
			ErrorMsg.Mobile[0] = true;
		} else {
			ErrorMsg.Mobile[0] = false;
		}
		// birthday validation check
		if (RegData.BirthDate.trim() === '') {
			ErrorMsg.BirthDate[1] = 'BirthDate is Empty';
			ErrorMsg.BirthDate[0] = true;
		} else if (!ValidateBirth(RegData.BirthDate)) {
			ErrorMsg.BirthDate[1] = 'invalid BirthDateday';
			ErrorMsg.BirthDate[0] = true;
		} else {
			ErrorMsg.BirthDate[0] = false;
		}

		return (
			!ErrorMsg.Email[0] &&
			!ErrorMsg.Password[0] &&
			!ErrorMsg.UserName[0] &&
			!ErrorMsg.Userid[0] &&
			!ErrorMsg.Mobile[0] &&
			!ErrorMsg.BirthDate[0]
		);
	}
	async function OnSubmit() {
		// console.log('🚀 ~ file: index@root.svelte ~ line 45 ~ OnSubmit ~ OnSubmit');
		// console.log('🚀 ~ file: index@root.svelte ~ line 166 ~ OnSubmit ~ RegData', RegData);
		// console.log('🚀 ~ file: index@root.svelte ~ line 28 ~ ErrorMsg', ErrorMsg);

		// console.log(
		// 	'🚀 ~ file: index@root.svelte ~ line 141 ~ OnSubmit ~ !ErrorChecking()',
		// 	ErrorChecking()
		// );
		if (ErrorChecking()) {
			let res = await fetch('/api/register', {
				// credentials: 'same-origin',

				method: 'POST',
				// mode: 'no-cors',
				credentials: 'include',
				body: JSON.stringify(RegData)
			})
				.then((res) => {
                    return res.json()
					// console.log('🚀 ~ file: index@root.svelte ~ line 150 ~ OnSubmit ~ res', res);
				}).then((d) => {
				    console.log("🚀 ~ file: +page.svelte ~ line 156 ~ .then ~ REgister response  data", d)
				    
                    
					goto('/');
                })
				.catch((err) => {
					console.log('🚀 ~ file: index@root.svelte ~ line 93 ~ OnSubmit ~ err', err);
					ErrorMsg.Email[1] = `${err}`;
					ErrorMsg.Email[0] = true;
				});
			// console.log('🚀 ~ file: index@blank.svelte ~ line 45 ~ OnSubmit ~ res', res);
			// if (res != null) {
			// 	// let value= await res.json()
			// 	console.log('🚀 ~ file: index.svelte ~ line 49 ~ OnSubmit ~ value : ');
			// 	// LoginData = {
			// 	// 	Email: '',
			// 	// 	Password: ''
			// 	// };
			// 	// let s:string="/"+String(value.ID)

			// 	goto('/');
			// }
		}
	}

	async function SetToDefault() {
		RegData = {
			Email: '',
			UserName: '',
			Userid: '',
			Mobile: '',
			Password: '',
			BirthDate: '',
            // Accounttype: ""
		};
		ErrorMsg = {
			Email: [false, ''],
			UserName: [false, ''],
			Userid: [false, ''],
			Mobile: [false, ''],
			Password: [false, ''],
			BirthDate: [false, '']
		};
		datedrop = false;
	}
</script>

<!-- markup (zero or more items) goes here -->
<svelte:head>
	<title>Accord Register</title>
</svelte:head>

<div
	class=" w-full h-screen font-Poppins  flex flex-row justify-center items-center bg-cover "
	style="background-image: url('{bg1}');"
>
	<a href="/login" class="absolute right-0 top-0 m-10 xs:mt-2 xs:mr-2">
		<button
			class="  h-12 w-60 bg-sky-700 active:bg-sky-500 hover:bg-sky-600 rounded-3xl text-xl font-bold text-gray-300"
		>
			Login
		</button>
	</a>

	<div class=" w-[485px] xs:w-full xs:mt-28 h-fit bg-[#36393f] flex flex-col  rounded-lg ">
		<div class=" self-center text-2xl font-semibold text-gray-300  mt-4 mb-2  ">
			Create a account
		</div>
		<!-- email -->
		<div class=" flex flex-row w-full  text-left text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Email :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 flex flex-row w-fit  ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<input
			autocomplete="off"
			bind:value={RegData.Email}
			type="email"
			class=" self-center h-12 w-[415px] xs:w-[90%] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
		/>

		<!-- passowrd -->
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
				on:keyup={(e) => (e.key === 'Enter' ? OnSubmit() : '')}
				value={RegData.Password}
				on:input={(e) => {
					//  var element = e.target as HTMLElement
					RegData.Password = e.target?.value
				}}
				autocomplete="off"
				type={ShowPass ? 'password' : 'text'}
				class=" self-center h-12 xs:w-[90%] w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
			/>
		</div>
		<!-- user name -->
		<div class=" flex flex-row text-left w-full text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				User Name :
			</div>
			{#if ErrorMsg.UserName[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 flex flex-row w-fit   ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.UserName[1]}</p>
				</div>
			{/if}
		</div>
		<input
			autocomplete="off"
			bind:value={RegData.UserName}
			type="text"
			class=" self-center h-12 w-[415px] xs:w-[90%] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
		/>
		<!-- <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">User Name</div>
        <input type="text" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
          -->
		<!-- userID -->
		<div class=" flex flex-row text-left w-full text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				UserID :
			</div>
			{#if ErrorMsg.Userid[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 flex flex-row w-fit  ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.Userid[1]}</p>
				</div>
			{/if}
		</div>
		<input
			autocomplete="off"
			bind:value={RegData.Userid}
			type="text"
			class=" self-center h-12 w-[415px] xs:w-[90%] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
		/>
		<!-- <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">UserID</div>
        <input type="text" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
          -->
		<!-- mobile -->
		<div class=" flex flex-row text-left w-full text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Mobile Number :
			</div>
			{#if ErrorMsg.Mobile[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 flex flex-row w-fit  ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.Mobile[1]}</p>
				</div>
			{/if}
		</div>
		<input
			autocomplete="off"
			bind:value={RegData.Mobile}
			type="text"
			class=" self-center h-12 w-[415px] p-2 xs:w-[90%] text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
		/>
		
		
		<!-- <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">Password</div>
        <input type="password" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
          -->

		<!-- date of birth -->
		<div class=" flex flex-row text-left w-full text-sm font-semibold text-gray-300 self-start ">
			<div class=" ml-12 xs:ml-6 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				BirthDay :
			</div>
			{#if ErrorMsg.BirthDate[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 flex flex-row w-fit  ">
					<Errorsign class="h-6 w-6   fill-red-500" />
					<p class="text-red-300">{ErrorMsg.BirthDate[1]}</p>
				</div>
			{/if}
		</div>
		<div
			autocomplete="off"
			on:click={() => {
				datedrop = !datedrop;
			}}
			type="text"
			class=" self-center cursor-pointer xs:w-[90%] h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
		>
			<!--  -->
			{RegData.BirthDate}
		</div>

		{#if datedrop}
			<Datepicker class=" top-44 self-center " bind:BirthDate={RegData.BirthDate} bind:datedrop />
		{/if}

		<!-- <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">Birthday</div>
        <div on:click={()=>{datedrop=!datedrop}} type="text" class="cursor-pointer self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
            {BirthDate}
        </div>
         -->

		<!-- continue -->
		<button
			on:click={() => {
				OnSubmit();
			}}
			class="mt-4 self-center h-12 w-60  bg-sky-700 active:bg-sky-500 hover:bg-sky-600 rounded-xl text-xl font-bold text-gray-300"
		>
			Register
		</button>
		<div
			class="text-left text-sm font-bold text-sky-600 hover:text-sky-300 active:text-sky-500 cursor-pointer  ml-10 mt-2"
		>
			Already have an account?
		</div>

		<div class="text-left text-xs mt-8 ml-4 font-semibold text-gray-500">
			By registering, you agree to Accord's <p
				class="inline text-sm font-bold text-sky-600 hover:text-sky-300 active:text-sky-500 cursor-pointer"
			>
				Terms of Service
			</p>
			and
			<p
				class="inline text-sm font-bold text-sky-600 hover:text-sky-300 active:text-sky-500 cursor-pointer"
			>
				privacy Policy
			</p>
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
		box-shadow: 0 0 0px 1000px #303338 inset;
		transition: background-color 5000s ease-in-out 0s;
	}
</style>
