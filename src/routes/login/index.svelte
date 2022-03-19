<script lang="ts">
    import bg1 from "./_bg/Colored Shapes.svg"
    import bg2 from "./_bg/Icon Grid.svg"
    import bg3 from "./_bg/Sprinkle.svg"
    import QrCode from "./_bg/qrcode.svelte"
// import Qrcode from "./_bg/qrcode.svelte"
    // your script goes here
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
    <a href="/register" class="absolute justify-self-end place-self-start m-10">

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
            <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">Email <p class="text-slate-500 inline">or</p> Phone Number <p class="text-slate-500 inline">or</p> UserID</div>
            <input type="text" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-0 focus:border-sky-600  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
            <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 mt-2 ">Password</div>
            <input type="password" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-0 focus:border-sky-600  bg-[#303338] border-2 mx-4 mt-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
            
            <!-- forget password -->
            <div class="text-left text-xs text-sky-600 hover:text-sky-300 active:text-sky-500 cursor-pointer hover:underline hover:underline-offset-2 ml-10 mt-1">Forget password?</div>
            <!-- submit -->
            <button class="m-4 self-center h-12 w-60 bg-sky-700 active:bg-sky-500 hover:bg-sky-600 rounded-xl text-xl font-bold text-gray-300"> Login </button>
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