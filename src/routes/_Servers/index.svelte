<script lang="ts">
    // your script goes here
        import userBanner from "./_images/whiteBeard.jpg"
        import Messaging from "./_icons/message.svelte"
        import Audio from "./_icons/microphone.svelte"
        import Video from "./_icons/video.svelte"
        import Announcement from "./_icons/announcement.svelte"


        let subFieldName:string="| Python |"
        let listOfField=[
            {type:"announcement",name:"General"},
            {type:"message",name:"messages"},
            {type:"audio",name:"speaking"},
            {type:"video",name:"Conference"},
            {type:"message",name:"Competitive Programming"},
            
        ]
        let demoField:boolean=false

        let showRoomList =false 
        

</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->


<div class=" w-52 bg-[#2f3136] flex flex-col overflow-y-auto no-scrollbar ">
        <!-- banner image -->
        <div class=" w-full aspect-video  bg-cover  " style="background-image: url({userBanner});">
            <!-- <img src="{userBanner}" alt="" class=" bg-cover "> -->
        </div>
        <!-- sub Field -->
        <div class="px-1 pt-4 ">
        <ul class="flex flex-col space-y-2">
          <!-- dash board -->
          <li class="relative text-gray-400 hover:text-white focus-within:text-white">
            <div class="absolute inset-y-0 left-0 flex items-center pl-2 pointer-events-none" >
              <svg class="w-5 h-5 stroke-current" fill="none" viewBox="0 0 24 24">                 <path                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                   d="M4.75 6.75C4.75 5.64543 5.64543 4.75 6.75 4.75H17.25C18.3546 4.75 19.25 5.64543 19.25 6.75V17.25C19.25 18.3546 18.3546 19.25 17.25 19.25H6.75C5.64543 19.25 4.75 18.3546 4.75 17.25V6.75Z"                 ></path>                 <path                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                   d="M9.75 8.75V19"                 ></path>                 <path                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                   d="M5 8.25H19"                 ></path>               </svg>
            </div>
            <a href="#" class="inline-block w-full py-2 pl-8 pr-4 text-xs rounded hover:bg-gray-800 focus:outline-none focus:ring-1 focus:ring-gray-500 focus:bg-gray-800"               >Dashboard</a             >
          </li>
          <!-- content -->
          <li class="">
            <div class="relative  flex justify-between  hover:text-white focus-within:text-white" >
              <div class="flex items-center w-full">
                <div class="absolute inset-y-0 left-0 flex items-center pl-2 pointer-events-none" >
                    <!-- room icon -->
                  <svg class=" w-5 h-5 stroke-current" fill="none" viewBox="0 0 24 24">                     <path                       stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M7.75 19.25H16.25C17.3546 19.25 18.25 18.3546 18.25 17.25V9L14 4.75H7.75C6.64543 4.75 5.75 5.64543 5.75 6.75V17.25C5.75 18.3546 6.64543 19.25 7.75 19.25Z"                     ></path>                     <path                       stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M18 9.25H13.75V5"                     ></path>                     <path                       stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M9.75 15.25H14.25"                     ></path>                     <path                       stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M9.75 12.25H14.25"                     ></path>                   </svg>
                </div>
                <a href="#" class="inline-block w-full py-2 pl-8 pr-4 text-xs rounded hover:bg-gray-800 focus:outline-none focus:ring-1 focus:ring-gray-500 focus:bg-gray-800"                   >
                    Content</a>
              </div>
              <button class="absolute right-0 flex items-center p-1 " tabindex="-1" on:click={()=> showRoomList=!showRoomList}>
                <svg class="w-5 h-5 stroke-current" fill="none" viewBox="0 0 24 24">                   <path                     stroke="currentColor"                     stroke-linecap="round"                     stroke-linejoin="round"                     stroke-width="1.5"                     d="M15.25 10.75L12 14.25L8.75 10.75"                   ></path>                 </svg>
              </button>
            </div>
            {#if showRoomList}
                <div class="pt-2 pl-4">
                    <ul class="flex flex-col pl-2 text-gray-500 border-l border-gray-700">
                        <li>
                        <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white" >Courses</a >
                        </li>
                        <li>
                        <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white" >Categories</a >
                        </li>
                        <li>
                        <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white" >Instructors</a >
                        </li>
                        <li>
                        <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white" >Video Library</a >
                        </li>
                    </ul>
                </div>
            {/if}
           
          </li>
          <!-- design -->
          <li class="relative  text-gray-500 hover:text-white focus-within:text-white">
            <div class="absolute inset-y-0 left-0 flex items-center pl-2 pointer-events-none" >
              <svg class="w-5 h-5 stroke-current" fill="none" viewBox="0 0 24 24">                 <path                   d="M12 4.75L19.25 9L12 13.25L4.75 9L12 4.75Z"                   stroke="currentColor"                   stroke-width="1.5"                   stroke-linecap="round"                   stroke-linejoin="round"                 ></path>                 <path                   d="M9.25 12L4.75 15L12 19.25L19.25 15L14.6722 12"                   stroke="currentColor"                   stroke-width="1.5"                   stroke-linecap="round"                   stroke-linejoin="round"                 ></path>               </svg>
            </div>
            <a href="#" class="inline-block w-full py-2 pl-8 pr-4 text-xs rounded hover:bg-gray-800 focus:outline-none focus:ring-1 focus:ring-gray-500 focus:bg-gray-800" >Design</a >
          </li>
          <!-- marketing -->
          <li class="relative  text-gray-500 hover:text-white focus-within:text-white">
            <div class="absolute inset-y-0 left-0 flex items-center pl-2 pointer-events-none" >
              <svg class="w-5 h-5 stroke-current" fill="none" viewBox="0 0 24 24">                 <circle cx="15" cy="9" r="1" fill="currentColor"></circle>                 <path                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                   d="M12 4.75H19.25V12L12.5535 18.6708C11.7544 19.4668 10.4556 19.445 9.68369 18.6226L5.28993 13.941C4.54041 13.1424 4.57265 11.8895 5.36226 11.1305L12 4.75Z"                 ></path>               </svg>
            </div>
            <a href="#" class="inline-block w-full py-2 pl-8 pr-4 text-xs rounded hover:bg-gray-800 focus:outline-none focus:ring-1 focus:ring-gray-500 focus:bg-gray-800" >Market & sell</a >
          </li>
          <!-- reporting -->
          <li class="relative  text-gray-500 hover:text-white focus-within:text-white">
            <div class="absolute inset-y-0 left-0 flex items-center pl-2 pointer-events-none" >
              <svg class="w-5 h-5 stroke-current" fill="none" viewBox="0 0 24 24">                 <path                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                   d="M5.75 19.2502H6.25C6.80229 19.2502 7.25 18.8025 7.25 18.2502V15.75C7.25 15.1977 6.80229 14.75 6.25 14.75H5.75C5.19772 14.75 4.75 15.1977 4.75 15.75V18.2502C4.75 18.8025 5.19772 19.2502 5.75 19.2502Z"                 ></path>                 <path                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                   d="M11.75 19.2502H12.25C12.8023 19.2502 13.25 18.8025 13.25 18.2502V12.75C13.25 12.1977 12.8023 11.75 12.25 11.75H11.75C11.1977 11.75 10.75 12.1977 10.75 12.75V18.2502C10.75 18.8025 11.1977 19.2502 11.75 19.2502Z"                 ></path>                 <path                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                   d="M17.75 19.2502H18.25C18.8023 19.2502 19.25 18.8025 19.25 18.2502V5.75C19.25 5.19772 18.8023 4.75 18.25 4.75H17.75C17.1977 4.75 16.75 5.19772 16.75 5.75V18.2502C16.75 18.8025 17.1977 19.2502 17.75 19.2502Z"                 ></path>               </svg>
            </div>
            <a href="#" class="inline-block w-full py-2 pl-8 pr-4 text-xs rounded hover:bg-gray-800 focus:outline-none focus:ring-1 focus:ring-gray-500 focus:bg-gray-800" >Reporting</a >
          </li>
          <!-- support -->
          <li class="relative  text-gray-500 hover:text-white focus-within:text-white">
            <div class="absolute inset-y-0 left-0 flex items-center pl-2 pointer-events-none" >
              <svg class="w-5 h-5 stroke-current" fill="none" viewBox="0 0 24 24">                 <circle                   cx="12"                   cy="12"                   r="7.25"                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                 ></circle>                 <circle                   cx="12"                   cy="12"                   r="3.25"                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                 ></circle>                 <path                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                   d="M7 17L9.5 14.5"                 ></path>                 <path                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                   d="M17 17L14.5 14.5"                 ></path>                 <path                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                   d="M9.5 9.5L7 7"                 ></path>                 <path                   stroke="currentColor"                   stroke-linecap="round"                   stroke-linejoin="round"                   stroke-width="1.5"                   d="M14.5 9.5L17 7"                 ></path>               </svg> </div>
            <a href="#" class="inline-block w-full py-2 pl-8 pr-4 text-xs rounded hover:bg-gray-800 focus:outline-none focus:ring-1 focus:ring-gray-500 focus:bg-gray-800" >Support</a >
          </li>
        </ul>
      </div>
        <!-- mysubfield -->


    </div>