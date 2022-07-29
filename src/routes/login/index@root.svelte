<script lang="ts" context="module">
    export async function load({}){

        return {
            props:{
                
            }
        }
    }
</script>
<script lang="ts">
    import bg1 from "./_bg/Colored Shapes.svg"
    import bg2 from "./_bg/Icon Grid.svg"
    import bg3 from "./_bg/Sprinkle.svg"
    import QrCode from "./_bg/qrcode.svelte"
// import Qrcode from "./_bg/qrcode.svelte"
    // your script goes here
    import {createEventDispatcher} from 'svelte'
    const dispatch=createEventDispatcher()

    let EmailOrMobileOrUserID:string=""
    // let Mobile:number
    let Password:string=""
    let InputError:boolean=false
    const ValidateEmail = (E: string) => {
        return String(E)
            .toLowerCase()
            .match(
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
            );
    };
    const ValidateMobile =(M: string)=>{
           return String(M)
            .toLowerCase()
            .match('^[0-9]+$');
    }
   
    function OnSubmit(){
        if (ValidateEmail(EmailOrMobileOrUserID) ){
            InputError=false
            SubmitRequest()
        }
        else if( ValidateMobile(EmailOrMobileOrUserID)){

        }
        else{
            InputError=true
        }

    }
    async function SubmitRequest(){
        try{
            console.log("successfull")
            // const res= await fetch("/auth/login",{
            //     method:"POST",
            //     body:JSON.stringify({
            //         EmailOrMobile,
            //         Password
            //     }),
            //     headers:{
            //         "Content-Type":"application/json"
            //     }
            // })
            // if(res.ok) {
            //     dispatch("success")
            // }
        }catch(err){
            console.log(err)
        }
    }
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->
<svelte:head>
  <title>Accord Login</title>
</svelte:head>

<div class=" w-screen h-screen  grid place-items-center" style="background-image: url('{bg3}');">
     <!-- register button -->
    <a href="/register" class="absolute right-0 top-0  m-10">

        <button class="     h-12 w-60 bg-sky-700 active:bg-sky-500 hover:bg-sky-600 rounded-3xl text-xl font-bold text-gray-300"> Register </button>
    </a>

    <div class=" w-[785px] h-[410px] bg-[#36393f] flex flex-row rounded-lg ">
        <!-- login -->
        <div class=" w-[490px] h-full flex flex-col  ">
            <div class=" text-2xl font-semibold text-gray-300  mt-4 mb-2  self-center">
                Welcome back :-)
            </div>
            <div class=" text-sm font-medium text-gray-400  mb-4  self-center">
                We're so excited to see you again!
            </div>
            <!-- email or  mobile -->
            <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">Email <p class="text-slate-500 inline">or</p> Phone Number <p class="text-slate-500 inline">or</p> UserID
                {#if InputError}
                    <svg class=" inline-flex ml-1  fill-red-500  w-6 h-6" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd">  </path></svg>
                    <p class=" inline-flex text-red-600">
                        Invalid Email/Mobile
                    </p> 
                     <!-- content here -->
                {/if}
            </div>
            <input bind:value={EmailOrMobileOrUserID} type="text" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
            <!-- password -->
            <div  class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 mt-2 ">Password</div>
            <input bind:value={Password} type="password" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 mt-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
            <!-- forget password -->
            <div class="text-left text-xs text-sky-600 hover:text-sky-300 active:text-sky-500 cursor-pointer hover:underline hover:underline-offset-2 ml-10 mt-1">Forget password?</div>
            <!-- submit -->
            <button on:click="{OnSubmit}" class=" m-4 self-center h-12 w-60 bg-sky-700 active:bg-sky-500 hover:bg-sky-600 rounded-xl text-xl font-bold text-gray-300"> Login </button>
            <!-- register -->
            <div class="ml-10 mt-4 text-xs text-slate-500 ">Need a account? <p class="inline text-left text-xs text-sky-600 hover:text-sky-300 active:text-sky-500 cursor-pointer hover:underline hover:underline-offset-2 ">Register</p></div>
        </div>
        <!-- qrcode -->
        <div class="w-[295px] h-full   flex flex-col ">
            <!-- <div class=" w-[180px] h-[180px]" style="background-image: url('{QrCode}');"></div> -->

            
            <QrCode/>

            <div class=" text-2xl font-semibold text-gray-300 mt-4 grid justify-items-center">
                Login with QR Code
            </div>            
            <div class=" text-sm text-[#8a8c90] m-6 grid justify-items-center">
                Scan with <p class="font-semibold inline text-gray-400">Accord Mobile app</p> to login instantly
            </div>
            
        </div>
    </div>

</div>