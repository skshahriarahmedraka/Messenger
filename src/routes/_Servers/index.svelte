<script lang="ts">
    // your script goes here
        import userBanner from "./_images/whiteBeard.jpg"
        // import Moji2 from "./_images/moji2.jpg"
        import {Moji} from "$lib/store"
        import Messaging from "./_icons/message.svelte"
        import Audio from "./_icons/microphone.svelte"
        import Video from "./_icons/video.svelte"
        import Announcement from "./_icons/announcement.svelte"
        import Hash from "./_icons/hash.svelte"
        import Rule from "./_icons/rule.svelte"
        import InvitePeople from "./_icons/invitePeople.svelte"
        import AeroUp from "./_icons/aeroUp.svelte"
        import AeroDown from "./_icons/aeroDown.svelte"


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
        let listOfField2=[
            {"name":"Dave","notification":true,"icon":Hash,"dropdownIcon":true,
              "content":[
                {"name":"rules","notification":true,"icon":Rule,"notificationNumber":3},
                {"name":"social-media","notification":false,"icon":Announcement,"notificationNumber":0},
                {"name":"annauncement","notification":true,"icon":Announcement,"notificationNumber":10},
              ]
            },
          {"name":"TechnologyTechnologyTechnologyTechnology","notification":false,"icon":Hash,"dropdownIcon":false,
              "content":[
                {"name":"tech-news","notification":false,"icon":Announcement,"notificationNumber":0},
                {"name":"pc-Build","notification":false,"icon":Hash,"notificationNumber":0},
                {"name":"pc-newsTechnologyTechnologyTechnologyTechnology","notification":false,"icon":Hash,"notificationNumber":0},
                {"name":"annauncement","notification":false,"icon":Announcement,"notificationNumber":10},
              ]
            },
            {"name":"Voice Chat","notification":false,"icon":Hash,"dropdownIcon":true,
              "content":[
                {"name":"Discussion","notification":false,"icon":Announcement,"notificationNumber":0},
                {"name":"New-topic","notification":false,"icon":Hash,"notificationNumber":0},
              ]
            },
            {"name":"Tech-chat","notification":true,"icon":Hash,"dropdownIcon":true,
              "content":[
                {"name":"tech-news","notification":false,"icon":Announcement,"notificationNumber":0},
                {"name":"pc-news","notification":false,"icon":Hash,"notificationNumber":0},
                {"name":"annauncement","notification":true,"icon":Announcement,"notificationNumber":10},
              ]
            },
            {"name":"Random","notification":true,"icon":Hash,"dropdownIcon":true,
              "content":[
                {"name":"general","notification":false,"icon":Announcement,"notificationNumber":0},
              ]
            },
            {"name":"Other Topic","notification":true,"icon":Hash,"dropdownIcon":true,
              "content":[
                {"name":"meeting","notification":false,"icon":Announcement,"notificationNumber":0},
                {"name":"new-tech","notification":false,"icon":Hash,"notificationNumber":0},
              ]
            },
            {"name":"Dave","notification":true,"icon":Hash,"dropdownIcon":true,
              "content":[
                {"name":"rules","notification":true,"icon":Rule,"notificationNumber":3},
                {"name":"social-media","notification":false,"icon":Announcement,"notificationNumber":0},
                {"name":"annauncement","notification":true,"icon":Announcement,"notificationNumber":10},
              ]
            },
          {"name":"TechnologyTechnologyTechnologyTechnology","notification":false,"icon":Hash,"dropdownIcon":false,
              "content":[
                {"name":"tech-news","notification":false,"icon":Announcement,"notificationNumber":0},
                {"name":"pc-Build","notification":false,"icon":Hash,"notificationNumber":0},
                {"name":"pc-newsTechnologyTechnologyTechnologyTechnology","notification":false,"icon":Hash,"notificationNumber":0},
                {"name":"annauncement","notification":false,"icon":Announcement,"notificationNumber":10},
              ]
            },
            {"name":"Voice Chat","notification":false,"icon":Hash,"dropdownIcon":true,
              "content":[
                {"name":"Discussion","notification":false,"icon":Announcement,"notificationNumber":0},
                {"name":"New-topic","notification":false,"icon":Hash,"notificationNumber":0},
              ]
            },
            {"name":"Tech-chat","notification":true,"icon":Hash,"dropdownIcon":true,
              "content":[
                {"name":"tech-news","notification":false,"icon":Announcement,"notificationNumber":0},
                {"name":"pc-news","notification":false,"icon":Hash,"notificationNumber":0},
                {"name":"annauncement","notification":true,"icon":Announcement,"notificationNumber":10},
              ]
            },
            {"name":"Random","notification":true,"icon":Hash,"dropdownIcon":true,
              "content":[
                {"name":"general","notification":false,"icon":Announcement,"notificationNumber":0},
              ]
            },
            {"name":"Other Topic","notification":true,"icon":Hash,"dropdownIcon":true,
              "content":[
                {"name":"meeting","notification":false,"icon":Announcement,"notificationNumber":0},
                {"name":"new-tech","notification":false,"icon":Hash,"notificationNumber":0},
              ]
            },
        ]
        
        let userOptions = {"video":false ,"microphone":false ,"notificationSound":false,"sound":false}
        
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div class="w-60 h-screen flex flex-col ">

  <div class=" w-60 bg-[#2f3136] flex flex-col overflow-y-auto no-scrollbar transition-all duration-500 ease-linear ">
    <!-- banner image -->
    <div class=" w-full  " style="">
      <img src="{userBanner}" alt="" class="aspect-video w-full object-cover ">
    </div>
       <!-- cutttt note -->
       <!-- mysubfield -->
        <!-- mysubfield -->
       
      <div class="px-1 pt-4 ">
        <ul class="flex flex-col space-y-2">
          
          <!-- content -->
          {#each listOfField2 as li}
             <!-- content here -->
             <li class="">
              <div class="relative  flex justify-between {li.notification && !li.dropdownIcon ? " text-white ": "text-gray-500" }  hover:text-white focus-within:text-white" >
                <div class="flex items-center w-full" on:click={()=> li.dropdownIcon=!li.dropdownIcon}>
                  <div class="absolute inset-y-0 left-0 flex items-center pl-2 pointer-events-none" >
                    <!-- room icon -->
                    <!-- {li.icon} -->
                    <svg class="w-5 h-5 " fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M14.243 5.757a6 6 0 10-.986 9.284 1 1 0 111.087 1.678A8 8 0 1118 10a3 3 0 01-4.8 2.401A4 4 0 1114 10a1 1 0 102 0c0-1.537-.586-3.07-1.757-4.243zM12 10a2 2 0 10-4 0 2 2 0 004 0z" clip-rule="evenodd"></path></svg>
                    <!-- <svg class=" w-5 h-5 stroke-current" fill="none" viewBox="0 0 24 24"><path stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M7.75 19.25H16.25C17.3546 19.25 18.25 18.3546 18.25 17.25V9L14 4.75H7.75C6.64543 4.75 5.75 5.64543 5.75 6.75V17.25C5.75 18.3546 6.64543 19.25 7.75 19.25Z"                     ></path>                     <path                       stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M18 9.25H13.75V5"                     ></path>                     <path                       stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M9.75 15.25H14.25"                     ></path>                     <path                       stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M9.75 12.25H14.25"                     ></path>                   </svg> -->
                  </div>
                <a href="#" class="truncate inline-block w-full py-2 pl-8 pr-4 text-xs rounded hover:bg-gray-800 focus:outline-none focus:ring-1 focus:ring-gray-500 focus:bg-gray-800"                   >
                {li.name}    
                </a>
              </div>
              <button class="absolute right-0 flex items-center p-1 " tabindex="-1" on:click={()=> li.dropdownIcon=!li.dropdownIcon}>
                {#if li.dropdownIcon}
                  <!-- aero up -->
                  <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M14.707 12.707a1 1 0 01-1.414 0L10 9.414l-3.293 3.293a1 1 0 01-1.414-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 010 1.414z" clip-rule="evenodd"></path></svg>
               {:else}
                  <!-- aero down -->
                  <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
               {/if}
                <!-- <svg class="w-5 h-5 stroke-current" fill="none" viewBox="0 0 24 24">                   <path                     stroke="currentColor"                     stroke-linecap="round"                     stroke-linejoin="round"                     stroke-width="1.5"                     d="M15.25 10.75L12 14.25L8.75 10.75"                   ></path>                 </svg> -->
              </button>
              </div>
              {#if li.dropdownIcon}
                  {#each li.content as con}
                    <!-- content here -->
                    <div class=" pl-2 text-left">
                      <ul class="flex flex-col { con.notification ? 'text-white' : 'text-gray-500' }   border-l border-gray-700">
                        <li class="flex flex-row">
                          {#if con.notification}
                            <!-- content here -->
                            <svg class="fill-white w-2 h-2 mt-3  -left-4 relative " viewBox="0 0 512 512" xmlns="http://www.w3.org/2000/svg"><path d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"/></svg>
                          {/if}
                          
                          <a href="#" class="truncate inline-block w-full {con.notification ? "px-2" : "px-4" }  py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white" >
                            {con.name}
                            </a >
                        </li>
                          <!-- <li>
                          <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white" >Categories</a >
                          </li>
                          <li>
                            <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white" >Instructors</a >
                            </li>
                            <li>
                              <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white" >Video Library</a >
                              </li>
                            </ul> -->
                    </div>
                  {/each}
              {/if}
           
            </li>
          
          {/each}
          
        </ul>
      </div>
      
      <!-- <div class=" absolute bg-slate-500 h-14 w-60 align-bottom">

      </div> -->
</div>
<div class="w-full h-14 bg-[#292b2f] flex flex-row ">
     <div class=" my-1 ml-2 h-12 w-12 transition-all duration-300 ease-linear cursor-pointer">
               <img src="{ Moji}" alt="" class=" rounded-2xl hover:rounded-xl active:rounded-md object-cover w-full h-full hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer  active:ring  active:ring-offset-base-50  active:ring-blue-600">
      </div>
      <!-- <div class="flex flex-col">
        <div class=" text-sm"> 
          h18dl3dnef
        </div>
      </div> -->
      <div class=" flex-grow"></div>
      <!-- video -->
      <div class=" hover:bg-gray-700 rounded-lg my-3" on:click="{()=>{userOptions["video"]=!userOptions["video"]}}">
        {#if !userOptions["video"] }
           <svg class=" absolute  fill-[#a4393c]  h-6 w-6  mt-1 mb-2 ml-1 " version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 1000 1000" enable-background="new 0 0 1000 1000" xml:space="preserve"> <metadata>  </metadata> <g><path d="M364.3,926.9l-2.4,2.4c0,0.8-0.2,1.6-0.6,2.4c-0.4,0.8-0.6,1.6-0.6,2.4c-0.8,1.6-1.2,2.8-1.2,3.6v-3.6c-11.3,18.5-27.5,31.8-48.8,39.8c-21.3,8-42.4,12.9-63.3,14.5c-24.9,2.4-51.1,2-78.4-1.2L635.7,73.1l2.4-2.4c0-2.4,0.4-4,1.2-4.8c0-0.8,0.4-1.6,1.2-2.4v2.4C651.8,47.3,668,34.1,689.3,26c21.3-8,42.4-12.9,63.3-14.5c24.9-2.4,51.1-2,78.4,1.2L364.3,926.9z"/></g> </svg>
        {/if}
        <!-- <svg class=" absolute  fill-[#a4393c] stroke-2 h-4 w-4 my-4" viewBox="0 0 640 512" xmlns="http://www.w3.org/2000/svg"><path d="M24.03 0c5.156 0 10.37 1.672 14.78 5.109l591.1 463.1c10.44 8.172 12.25 23.27 4.062 33.7c-8.125 10.41-23.19 12.28-33.69 4.078L9.189 42.89c-10.44-8.172-12.25-23.27-4.062-33.7C9.845 3.156 16.91 0 24.03 0z"/></svg> -->
        <svg class="w-6 h-6 m-1 fill-slate-400 hover:fill-white transition-all duration-100 ease-linear cursor-pointer" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M2 6a2 2 0 012-2h6a2 2 0 012 2v8a2 2 0 01-2 2H4a2 2 0 01-2-2V6zM14.553 7.106A1 1 0 0014 8v4a1 1 0 00.553.894l2 1A1 1 0 0018 13V7a1 1 0 00-1.447-.894l-2 1z"></path></svg>
      </div>
      <!-- microphone -->
      <div class="hover:bg-gray-700 rounded-lg my-3" on:click={()=>{userOptions["microphone"]=!userOptions["microphone"]}}>
        {#if !userOptions["microphone"] }
           <svg class=" absolute  fill-[#a4393c]  h-6 w-6  mt-1 mb-2 ml-1 " version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 1000 1000" enable-background="new 0 0 1000 1000" xml:space="preserve"> <metadata>  </metadata> <g><path d="M364.3,926.9l-2.4,2.4c0,0.8-0.2,1.6-0.6,2.4c-0.4,0.8-0.6,1.6-0.6,2.4c-0.8,1.6-1.2,2.8-1.2,3.6v-3.6c-11.3,18.5-27.5,31.8-48.8,39.8c-21.3,8-42.4,12.9-63.3,14.5c-24.9,2.4-51.1,2-78.4-1.2L635.7,73.1l2.4-2.4c0-2.4,0.4-4,1.2-4.8c0-0.8,0.4-1.6,1.2-2.4v2.4C651.8,47.3,668,34.1,689.3,26c21.3-8,42.4-12.9,63.3-14.5c24.9-2.4,51.1-2,78.4,1.2L364.3,926.9z"/></g> </svg>
        {/if}
        <svg class="w-6 h-6 m-1 fill-slate-400 hover:fill-white transition-all duration-100 ease-linear cursor-pointer" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M7 4a3 3 0 016 0v4a3 3 0 11-6 0V4zm4 10.93A7.001 7.001 0 0017 8a1 1 0 10-2 0A5 5 0 015 8a1 1 0 00-2 0 7.001 7.001 0 006 6.93V17H6a1 1 0 100 2h8a1 1 0 100-2h-3v-2.07z" clip-rule="evenodd"></path></svg>
      </div>
      <!-- sound -->
      <div class="hover:bg-gray-700 rounded-lg my-3" on:click={()=>{userOptions["sound"]=!userOptions["sound"]}}>
        {#if !userOptions["sound"] }
           <svg class=" absolute  fill-[#a4393c]  h-6 w-6  mt-1 mb-2 ml-1" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 1000 1000" enable-background="new 0 0 1000 1000" xml:space="preserve"> <metadata>  </metadata> <g><path d="M364.3,926.9l-2.4,2.4c0,0.8-0.2,1.6-0.6,2.4c-0.4,0.8-0.6,1.6-0.6,2.4c-0.8,1.6-1.2,2.8-1.2,3.6v-3.6c-11.3,18.5-27.5,31.8-48.8,39.8c-21.3,8-42.4,12.9-63.3,14.5c-24.9,2.4-51.1,2-78.4-1.2L635.7,73.1l2.4-2.4c0-2.4,0.4-4,1.2-4.8c0-0.8,0.4-1.6,1.2-2.4v2.4C651.8,47.3,668,34.1,689.3,26c21.3-8,42.4-12.9,63.3-14.5c24.9-2.4,51.1-2,78.4,1.2L364.3,926.9z"/></g> </svg>
        {/if}
        <svg class="w-6 h-6 m-1 fill-slate-400 hover:fill-white transition-all duration-100 ease-linear cursor-pointer" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M9.383 3.076A1 1 0 0110 4v12a1 1 0 01-1.707.707L4.586 13H2a1 1 0 01-1-1V8a1 1 0 011-1h2.586l3.707-3.707a1 1 0 011.09-.217zM14.657 2.929a1 1 0 011.414 0A9.972 9.972 0 0119 10a9.972 9.972 0 01-2.929 7.071 1 1 0 01-1.414-1.414A7.971 7.971 0 0017 10c0-2.21-.894-4.208-2.343-5.657a1 1 0 010-1.414zm-2.829 2.828a1 1 0 011.415 0A5.983 5.983 0 0115 10a5.984 5.984 0 01-1.757 4.243 1 1 0 01-1.415-1.415A3.984 3.984 0 0013 10a3.983 3.983 0 00-1.172-2.828 1 1 0 010-1.415z" clip-rule="evenodd"></path></svg>
      </div>
      <!-- notification sound -->
      <div class="hover:bg-gray-700 rounded-lg my-3" on:click={()=>{userOptions["notificationSound"]=!userOptions["notificationSound"]}}>
        {#if !userOptions["notificationSound"] }
          <svg class=" absolute  fill-[#a4393c]  h-6 w-6 mt-1 mb-2 ml-1 " version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 1000 1000" enable-background="new 0 0 1000 1000" xml:space="preserve"> <metadata>  </metadata> <g><path d="M364.3,926.9l-2.4,2.4c0,0.8-0.2,1.6-0.6,2.4c-0.4,0.8-0.6,1.6-0.6,2.4c-0.8,1.6-1.2,2.8-1.2,3.6v-3.6c-11.3,18.5-27.5,31.8-48.8,39.8c-21.3,8-42.4,12.9-63.3,14.5c-24.9,2.4-51.1,2-78.4-1.2L635.7,73.1l2.4-2.4c0-2.4,0.4-4,1.2-4.8c0-0.8,0.4-1.6,1.2-2.4v2.4C651.8,47.3,668,34.1,689.3,26c21.3-8,42.4-12.9,63.3-14.5c24.9-2.4,51.1-2,78.4,1.2L364.3,926.9z"/></g> </svg>
        {/if}
        <svg class="w-6 h-6 m-1 fill-slate-400 hover:fill-white transition-all duration-100 ease-linear cursor-pointer" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M10 2a6 6 0 00-6 6v3.586l-.707.707A1 1 0 004 14h12a1 1 0 00.707-1.707L16 11.586V8a6 6 0 00-6-6zM10 18a3 3 0 01-3-3h6a3 3 0 01-3 3z"></path></svg>
        <!-- <svg class="w-8 h-8 my-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path></svg> -->
      </div>
      <!-- settings -->
      <div class=" hover:bg-gray-700 rounded-lg my-3 mr-1">
        <svg class="w-7 h-7  m-1 mr-1 fill-slate-400 hover:fill-white transition-all duration-100 ease-linear cursor-pointer" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z" clip-rule="evenodd"></path></svg>
      </div>
</div>
</div>