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






        <!-- date of birth -->

        
            <!-- <div class="flex justify-center">
                <div>
                    <div class="dropdown relative">
                    <button on:click="{()=>{datedrop=!datedrop}}" class=" dropdown-toggle px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg active:text-white transition duration-150 ease-in-out flex items-center whitespace-nowrap " type="button" id="dropdownMenuButton1" data-bs-toggle="dropdown" aria-expanded="false" >
                        Dropdown button
                        <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="caret-down" class="w-2 ml-2" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512" > <path fill="currentColor" d="M31.3 192h257.3c17.8 0 26.7 21.5 14.1 34.1L174.1 354.8c-7.8 7.8-20.5 7.8-28.3 0L17.2 226.1C4.6 213.5 13.5 192 31.3 192z" ></path> </svg>
                    </button>
                    <ul class=" dropdown-menu min-w-max absolute { datedrop ?'hidden' : ''} bg-white text-base z-50 float-left py-2 list-none text-left rounded-lg shadow-lg mt-1  m-0 bg-clip-padding border-none " aria-labelledby="dropdownMenuButton1" >
                        <li>
                            <a class=" dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-gray-700 hover:bg-gray-100 " href="#" >Action</a >
                        </li>
                        <li>
                            <a class=" dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-gray-700 hover:bg-gray-100 " href="#" >Another action</a >
                        </li>
                        <li>
                            <a class=" dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-gray-700 hover:bg-gray-100 " href="#" >Something else here</a >
                        </li>
                    </ul>
                    </div>
                </div>
            </div> -->