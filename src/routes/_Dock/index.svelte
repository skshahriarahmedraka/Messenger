<script lang="ts">
    // your script goes here
        import profilepic from "./_profilePic/law.png"
    import profilepic2 from "./_profilePic/ren.jpeg"
    import profilepic3 from "./_profilePic/e.jpg"
    import profilepic4 from "./_profilePic/s2.png"


    let show=false

    let UserList =[
        {
            "name": "ssar",
            "profilePic": profilepic,
            "active":true,
            "profieString":"Sk"
            
        },
         {
            "name": "ssar",
            "profilePic": profilepic2,
            "active":true,
            "profieString":"Sk"
        },
         {
            "name": "ssar",
            "profilePic": profilepic3,
            "active":true,
            "profieString":"Sk"
        },
         {
            "name": "ssar",
            "profilePic": profilepic4,
            "active":true,
            "profieString":"Sk"
        },
        {
            "name": "ssar",
            "profilePic": profilepic,
            "active":true,
            "profieString":"Sk"
        },
         {
            "name": "ssar",
            "profilePic": profilepic2,
            "active":true,
            "profieString":"Sk"
        },
         {
            "name": "ssar",
            "profilePic": profilepic3,
            "active":true,
            "profieString":"Sk"
        },
         {
            "name": "ssar",
            "profilePic": profilepic4,
            "active":true,
            "profieString":"Sk"
        },
        {
            "name": "ssar",
            "profilePic": profilepic,
            "active":true,
            "profieString":"Sk"
        },
         {
            "name": "ssar",
            "profilePic": profilepic2,
            "active":true,
            "profieString":"Sk"
        },
         {
            "name": "ssar",
            "profilePic": profilepic3,
            "active":true,
            "profieString":"Sk"
        },
         {
            "name": "ssar",
            "profilePic": profilepic4,
            "active":true,
            "profieString":"Sk"
        },
        {
            "name": "ssar",
            "profilePic": profilepic,
            "active":true,
            "profieString":"Sk"
        },
         {
            "name": "ssar",
            "profilePic": profilepic2,
            "active":true,
            "profieString":"Sk"
        },
         {
            "name": "ssar",
            "profilePic": profilepic3,
            "active":true,
            "profieString":"Sk"
        },
         {
            "name": "ssar",
            "profilePic": profilepic4,
            "active":true,
            "profieString":"Sk"
        },
        

    ]


</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div class=" flex flex-col  flex-none w-[72px] bg-[#202225] overflow-y-auto no-scrollbar  ">

        <!-- <div class="avatar online">
            <div class=" w-12 rounded-full">
                <img src="https://api.lorem.space/image/face?hash=28212">
            </div>
        </div> -->
         <div class="avatar  py-1 px-2">
            <div class=" w-14 h-14 hover:rounded-xl   rounded-3xl active:rounded-md hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer {false ? "ring  ring-offset-base-100  ring-blue-500" : "" }">
                <!-- <img src="https://api.lorem.space/image/face?hash=40361"> -->
                <!-- <img src="{}" alt="" class=""> -->
                <svg xmlns="http://www.w3.org/2000/svg" enable-background="new 0 0 100 100" viewBox="0 0 100 100"><path fill="#ce3056" d="M3.63,62.22c0.563,12.911,18.923,12.908,19.484,0c0,0,0-9.73,0-9.73h-9.742 C8.011,52.489,3.63,56.864,3.63,62.22z"/><path fill="#5bb381" d="M62.234,47.511c5.362,0,9.742-4.375,9.742-9.73V13.417c-0.56-12.905-18.92-12.914-19.484,0 c0,0,0,24.364,0,24.364C52.492,43.136,56.872,47.511,62.234,47.511z"/><path fill="#ce3056" d="M37.765,52.489c-5.362,0-9.742,4.375-9.742,9.73v24.364c0.56,12.905,18.92,12.914,19.484,0 c0,0,0-24.364,0-24.364C47.508,56.864,43.128,52.489,37.765,52.489z"/><path fill="#5bb381" d="M96.37,37.78c-0.563-12.911-18.923-12.908-19.484,0c0,0,0,9.73,0,9.73h9.742 C91.989,47.511,96.37,43.136,96.37,37.78z"/><path fill="#61c3ec" d="M13.372,47.511h24.393c12.921-0.559,12.93-18.897,0-19.461H13.372C0.45,28.61,0.445,46.948,13.372,47.511z"/><path fill="#e4b34b" d="M62.234 76.853h-9.742v9.73c0 5.355 4.38 9.73 9.742 9.73C75.161 95.751 75.158 77.414 62.234 76.853zM86.628 52.489H62.234c-12.923.56-12.928 18.898 0 19.461 0 0 24.393 0 24.393 0C99.548 71.391 99.557 53.052 86.628 52.489z"/><path fill="#61c3ec" d="M37.765,3.686c-12.926,0.562-12.924,18.9,0,19.461c0,0,9.742,0,9.742,0v-9.73 C47.508,8.061,43.128,3.686,37.765,3.686z"/><path d="M62.182 75.355h-9.701c-.765 0-1.385.62-1.385 1.383v9.689c0 6.105 4.973 11.073 11.086 11.073C76.887 96.891 76.885 75.962 62.182 75.355zM62.182 94.733c-4.585 0-8.316-3.726-8.316-8.305v-8.305h8.316C73.204 78.558 73.203 94.298 62.182 94.733zM86.472 51.095h-24.29c-14.702.607-14.707 21.537 0 22.145 0 0 24.29 0 24.29 0C101.172 72.635 101.18 51.704 86.472 51.095zM86.472 70.473h-24.29c-4.585 0-8.316-3.726-8.316-8.305s3.731-8.305 8.316-8.305h24.29C97.489 54.296 97.498 70.036 86.472 70.473zM76.771 48.905h9.701c14.704-.607 14.705-21.538 0-22.145-6.112 0-11.086 4.967-11.086 11.073v9.689C75.386 48.285 76.006 48.905 76.771 48.905zM78.156 37.832c0-4.579 3.731-8.305 8.316-8.305 11.023.436 11.021 16.176 0 16.611h-8.316V37.832zM37.818 2.5c-14.706.608-14.704 21.538 0 22.145 0 0 9.701 0 9.701 0 .765 0 1.385-.62 1.385-1.383v-9.689C48.903 7.467 43.93 2.5 37.818 2.5zM46.133 21.878h-8.316c-4.585 0-8.316-3.726-8.316-8.305.437-11.01 16.196-11.008 16.631 0C46.133 13.572 46.133 21.878 46.133 21.878zM23.229 51.095h-9.701c-14.704.607-14.705 21.538 0 22.145 6.112 0 11.086-4.967 11.086-11.072v-9.689C24.614 51.715 23.994 51.095 23.229 51.095zM21.844 62.168c-.436 11.009-16.196 11.008-16.631 0 0-4.579 3.731-8.305 8.316-8.305h8.316V62.168z"/><path d="M37.818 51.095c-6.112 0-11.086 4.967-11.086 11.072v24.26c.606 14.683 21.562 14.691 22.172 0 0 0 0-24.26 0-24.26C48.903 56.063 43.93 51.095 37.818 51.095zM46.133 86.427c-.434 11.004-16.193 11.013-16.631 0 0 0 0-24.26 0-24.26 0-4.579 3.731-8.305 8.316-8.305s8.316 3.726 8.316 8.305V86.427zM13.529 48.905c0 0 24.289 0 24.289 0 14.702-.607 14.707-21.536 0-22.145-1.827.036-1.816 2.735 0 2.767 11.017.433 11.026 16.174 0 16.611 0 0-24.289 0-24.289 0-11.017-.433-11.027-16.174 0-16.611 0 0 18.216 0 18.216 0 1.817-.027 1.823-2.738 0-2.767 0 0-18.216 0-18.216 0C-1.171 27.365-1.182 48.295 13.529 48.905z"/><path d="M62.182,48.905c6.112,0,11.086-4.967,11.086-11.072c-0.036-1.824-2.738-1.814-2.77,0 c-0.434,11.004-16.194,11.013-16.631,0c0,0,0-24.26,0-24.26c0.434-11.003,16.193-11.014,16.631,0c0,0,0,18.714,0,18.714 c0.027,1.814,2.741,1.821,2.77,0c0,0,0-18.714,0-18.714c-0.605-14.681-21.562-14.692-22.172,0c0,0,0,24.26,0,24.26 C51.096,43.937,56.069,48.905,62.182,48.905z"/></svg>
                <!-- <svg class="h-16 w-16  fill-indigo-700" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" id="Capa_1" x="0px" y="0px" viewBox="0 0 512 512" style="enable-background:new 0 0 512 512;" xml:space="preserve" width="512" height="512"> <g> <path d="M85.333,0h64c47.128,0,85.333,38.205,85.333,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   C38.205,234.667,0,196.462,0,149.333v-64C0,38.205,38.205,0,85.333,0z"/> <path d="M362.667,0h64C473.795,0,512,38.205,512,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   c-47.128,0-85.333-38.205-85.333-85.333v-64C277.333,38.205,315.538,0,362.667,0z"/> <path d="M85.333,277.333h64c47.128,0,85.333,38.205,85.333,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   C38.205,512,0,473.795,0,426.667v-64C0,315.538,38.205,277.333,85.333,277.333z"/> <path d="M362.667,277.333h64c47.128,0,85.333,38.205,85.333,85.333v64C512,473.795,473.795,512,426.667,512h-64   c-47.128,0-85.333-38.205-85.333-85.333v-64C277.333,315.538,315.538,277.333,362.667,277.333z"/> </g> </svg> -->
                <!-- <svg class="bg-green-600 " xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" id="Capa_1" x="0px" y="0px" viewBox="0 0 512.001 512.001" style="enable-background:new 0 0 512.001 512.001;" xml:space="preserve" width="512" height="512"> <g> <path d="M490.134,185.472L338.966,34.304c-45.855-45.737-120.076-45.737-165.931,0L21.867,185.472   C7.819,199.445-0.055,218.457,0,238.272v221.397C0.047,488.568,23.475,511.976,52.374,512h407.253   c28.899-0.023,52.326-23.432,52.373-52.331V238.272C512.056,218.457,504.182,199.445,490.134,185.472z M448,448H341.334v-67.883   c0-44.984-36.467-81.451-81.451-81.451c0,0,0,0,0,0h-7.765c-44.984,0-81.451,36.467-81.451,81.451l0,0V448H64V238.272   c0.007-2.829,1.125-5.541,3.115-7.552L218.283,79.552c20.825-20.831,54.594-20.835,75.425-0.01c0.003,0.003,0.007,0.007,0.01,0.01   L444.886,230.72c1.989,2.011,3.108,4.723,3.115,7.552V448z"/> </g> </svg> -->
            </div>
        </div>
        {#each UserList as u}
        <button class="my-1 mx-2 w-14 h-14 relative  ">
                <svg class="fill-white w-2 h-2 my-5 -left-3 place-content-center absolute " viewBox="0 0 512 512" xmlns="http://www.w3.org/2000/svg"><path d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"/></svg>

                <img src="{u.profilePic }" alt="" class=" rounded-3xl hover:rounded-xl active:rounded-md object-cover w-full h-full hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer  active:ring  active:ring-offset-base-50  active:ring-blue-600">
                <span class=" absolute px-1  border-2 border-[#202225] rounded-full text-xs font-semibold  bg-red-500 text-white -bottom-2 -right-2 ">
                    99+
                </span>
                <!-- <span class="absolute bg-red-600 px-2 py-1 text-xs font-bold rounded-full -top-3 -right-3">99+</span> -->

            </button>
            <!-- sample -->
            <!-- <button class="py-4 px-1 relative border-2 border-transparent text-gray-800 rounded-full hover:text-gray-400 focus:outline-none focus:text-gray-500 transition duration-150 ease-in-out" aria-label="Cart">
                <svg class="h-6 w-6" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" stroke="currentColor">
                    <path d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"></path>
                </svg>
                <span class="absolute inset-0 object-right-top -mr-6">
                    <div class="inline-flex items-center px-1.5 py-0.5 border-2 border-white rounded-full text-xs font-semibold leading-4 bg-red-500 text-white">
                    99
                    </div>
                </span>
            </button> -->
        {/each}
        

        <!-- download button -->
        <div class="avatar  py-1 px-2">
            <div class=" bg-sky-600  h-14 w-14 hover:rounded-xl  rounded-3xl active:rounded-md hover:ring  transition-all duration-150 ease-linear cursor-pointer ">
                <!-- <img src="https://api.lorem.space/image/face?hash=40361"> -->
                <!-- <img src="{}" alt="" class=""> -->
                <!-- <svg class="h-16 w-16  fill-indigo-700" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" id="Capa_1" x="0px" y="0px" viewBox="0 0 512 512" style="enable-background:new 0 0 512 512;" xml:space="preserve" width="512" height="512"> <g> <path d="M85.333,0h64c47.128,0,85.333,38.205,85.333,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   C38.205,234.667,0,196.462,0,149.333v-64C0,38.205,38.205,0,85.333,0z"/> <path d="M362.667,0h64C473.795,0,512,38.205,512,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   c-47.128,0-85.333-38.205-85.333-85.333v-64C277.333,38.205,315.538,0,362.667,0z"/> <path d="M85.333,277.333h64c47.128,0,85.333,38.205,85.333,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   C38.205,512,0,473.795,0,426.667v-64C0,315.538,38.205,277.333,85.333,277.333z"/> <path d="M362.667,277.333h64c47.128,0,85.333,38.205,85.333,85.333v64C512,473.795,473.795,512,426.667,512h-64   c-47.128,0-85.333-38.205-85.333-85.333v-64C277.333,315.538,315.538,277.333,362.667,277.333z"/> </g> </svg> -->
                <!-- <svg class="bg-green-600 " xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" id="Capa_1" x="0px" y="0px" viewBox="0 0 512.001 512.001" style="enable-background:new 0 0 512.001 512.001;" xml:space="preserve" width="512" height="512"> <g> <path d="M490.134,185.472L338.966,34.304c-45.855-45.737-120.076-45.737-165.931,0L21.867,185.472   C7.819,199.445-0.055,218.457,0,238.272v221.397C0.047,488.568,23.475,511.976,52.374,512h407.253   c28.899-0.023,52.326-23.432,52.373-52.331V238.272C512.056,218.457,504.182,199.445,490.134,185.472z M448,448H341.334v-67.883   c0-44.984-36.467-81.451-81.451-81.451c0,0,0,0,0,0h-7.765c-44.984,0-81.451,36.467-81.451,81.451l0,0V448H64V238.272   c0.007-2.829,1.125-5.541,3.115-7.552L218.283,79.552c20.825-20.831,54.594-20.835,75.425-0.01c0.003,0.003,0.007,0.007,0.01,0.01   L444.886,230.72c1.989,2.011,3.108,4.723,3.115,7.552V448z"/> </g> </svg> -->
                <svg class="w-10 h-10 m-2" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm3.293-7.707a1 1 0 011.414 0L9 10.586V3a1 1 0 112 0v7.586l1.293-1.293a1 1 0 111.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
            </div>
        </div>
        
        <!-- <div class="avatar offline">
            <div class=" w-16  active:mask active:mask-squircle rounded-md ring  ring-offset-base-100 ring-offset-2 ring-blue-600">
                <img src="https://api.lorem.space/image/face?hash=40361">
            </div>
        </div> -->
        
    </div>