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
    import Datepicker from "$lib/datepicker/index.svelte"

    let datedrop:boolean=true
    let BirthDate:string="June 22, 2022"

</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->
<svelte:head>
  <title>Accord Register</title>
</svelte:head>

<div class=" w-screen h-screen  flex flex-row justify-center items-center bg-cover " style="background-image: url('{bg1}');">
  <a href="/login" class="absolute right-0 top-0     m-10">

        <button class="  h-12 w-60 bg-sky-700 active:bg-sky-500 hover:bg-sky-600 rounded-3xl text-xl font-bold text-gray-300"> Login </button>
    </a>

    <div class=" w-[485px] h-[640px] bg-[#36393f] flex flex-col  rounded-lg ">
        <div class=" self-center text-2xl font-semibold text-gray-300  mt-4 mb-2  ">
                Create a account
        </div>
        <!-- email -->
        <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">Email</div>
        <input type="email" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
          
        <!-- user name -->
         <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">User Name</div>
        <input type="text" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
         
        <!-- userID -->
         <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">UserID</div>
        <input type="text" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
         
        <!-- mobile -->
         <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">Mobile Number</div>
        <input type="text" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
         
        <!-- passowrd -->
         <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">Password</div>
        <input type="password" class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
         


        <!-- date of birth -->
        <div class=" text-left text-sm font-semibold text-gray-300 self-start ml-12 ">Birthday</div>
        <div on:click={()=>{datedrop=!datedrop}} type="text" class="cursor-pointer self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl ">
            {BirthDate}
        </div>
        {#if datedrop}
             
             <Datepicker bind:BirthDate={BirthDate} bind:datedrop={datedrop} />
        {/if}
        
        
        
        
        <!-- continue -->
        <button class="mt-4 self-center h-12 w-60 bg-sky-700 active:bg-sky-500 hover:bg-sky-600 rounded-xl text-xl font-bold text-gray-300"> Register </button>
        <div class="text-left text-sm font-bold text-sky-600 hover:text-sky-300 active:text-sky-500 cursor-pointer  ml-10 mt-2">Already have an account?</div>
        
        <div class="text-left text-xs mt-8 ml-4 font-semibold text-gray-500">By registering, you agree to Accord's <p class="inline text-sm font-bold text-sky-600 hover:text-sky-300 active:text-sky-500 cursor-pointer"> Terms of Service</p>  and <p class="inline text-sm font-bold text-sky-600 hover:text-sky-300 active:text-sky-500 cursor-pointer"> privacy Policy</p></div>
        
    </div>
    
</div>





