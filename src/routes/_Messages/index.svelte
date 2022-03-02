<script lang="ts">
  
    import {showPeopleList , MessageList} from "$lib/store"
    import Message from "./message.svelte"


    let showPeopleListValue

    showPeopleList.subscribe(val=>{
        showPeopleListValue=val 
    }) 
    let messagelist= [
        {"user":"sk shahriar ahmed raka","msg":"what's up homie . what are you doing bro"},
        {"user":"ekbal hasan","msg":"what's up homie . what are you doing bro"},
        {"user":"Portgas D. Ace","msg":"what's up homie . Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corporis impedit totam optio quasi sequi fuga magnam non repellendus libero, pariatur excepturi exercitationem laboriosam recusandae illum a doloribus eius ab."},
        {"user":"monkey D. luffy","msg":"what are you doing bro . Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corporis impedit totam optio quasi sequi fuga magnam non repellendus libero, pariatur excepturi exercitationem laboriosam recusandae illum a doloribus eius ab."},
        {"user":"Gol D. Roger","msg":"Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corporis impedit totam optio quasi sequi fuga magnam non repellendus libero, pariatur excepturi exercitationem laboriosam recusandae illum a doloribus eius ab."},
        {"user":"Adward Newgate","msg":"good bye . Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corporis impedit totam optio quasi sequi fuga magnam non repellendus libero, pariatur excepturi exercitationem laboriosam recusandae illum a doloribus eius ab."},
        {"user":"marco the phinix","msg":"that's a good point .Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corporis impedit totam optio quasi sequi fuga magnam non repellendus libero, pariatur excepturi exercitationem laboriosam recusandae illum a doloribus eius ab."},
        {"user":"inosuke hashibira","msg":" i am hungry Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corporis impedit totam optio quasi sequi fuga magnam non repellendus libero, pariatur excepturi exercitationem laboriosam recusandae illum a doloribus eius ab. . what are you doing bro"},
        {"user":"Miss tomayo","msg":"Now way home .Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi corporis impedit totam optio quasi sequi fuga magnam non repellendus libero, pariatur excepturi exercitationem laboriosam recusandae illum a doloribus eius ab. . what are you doing bro"},
        {"user":"Kibutsuji Mujan","msg":" there is a long way . what's up homie . what are you doing bro"},
    ]

    let writeFocus:boolean=false

    let messengerValue:string=""

    // let sampleMessaelist:string[] =["ssar","raka","sk"]
    // function onSubmitMessage(e){
    //     let  formData = new FormData(e.target);
    //     const data = {};
    //     for (let field of formData) {
    //     const [key, value] = field;
    //     data[key] = value;
    //     }
    //     $MessageList=[...$MessageList,data["string"]]    
    //     console.log(data)
    //     // sampleMessaelist.push(data["string"])
    //     console.log($MessageList)
    // }
    function SendMessage(){
        $MessageList=[...$MessageList,messengerValue]
        console.log($MessageList)
        messengerValue=""

    }
    function keyHandler(e){
        if (e.key ==="Enter"){
            SendMessage()
        }
    }

</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div class=" {showPeopleListValue!=0 ? "w-5/6": "w-full"} h-full  transition-all duration-200 ease-linear bg-[#36393f] flex flex-col">
                <!-- all messaging -->
    <div class=" flex-grow bg-gray-600 z-10 flex flex-col overflow-y-scroll scrollbar-thin  scrollbar-thumb-gray-500 scrollbar-track-gray-700 ">
        <!-- scrollbar-thin  scrollbar-thumb-gray-500 scrollbar-track-gray-700 -->
    
    {#each $MessageList as i}
            <!-- content here -->
            <Message message={i}/>
    {/each}

    </div>
    <!-- write message -->
    <!-- <form action="" class=" " on:submit|preventDefault={onSubmitMessage}> -->
    <div class="  w-full  py-2 px-1 flex flex-row h-14 place-items-center  ">
        <!-- all -->
        <button class=" hover:bg-gray-800 hover:rounded-xl">
        <svg class=" w-7 h-7 m-2  fill-gray-400" version="1.1" id="Capa_1" x="0px" y="0px" viewBox="0 0 512 512" style="enable-background:new 0 0 512 512;" xml:space="preserve" width="512" height="512"><g>	<path d="M85.333,0h64c47.128,0,85.333,38.205,85.333,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   C38.205,234.667,0,196.462,0,149.333v-64C0,38.205,38.205,0,85.333,0z"/>	<path d="M362.667,0h64C473.795,0,512,38.205,512,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   c-47.128,0-85.333-38.205-85.333-85.333v-64C277.333,38.205,315.538,0,362.667,0z"/>	<path d="M85.333,277.333h64c47.128,0,85.333,38.205,85.333,85.333v64c0,47.128-38.205,85.333-85.333,85.333h-64   C38.205,512,0,473.795,0,426.667v-64C0,315.538,38.205,277.333,85.333,277.333z"/><path d="M362.667,277.333h64c47.128,0,85.333,38.205,85.333,85.333v64C512,473.795,473.795,512,426.667,512h-64   c-47.128,0-85.333-38.205-85.333-85.333v-64C277.333,315.538,315.538,277.333,362.667,277.333z"/></g></svg>
        </button>
        <!-- camera -->
        <button class=" hover:bg-gray-800 hover:rounded-xl">
        <svg class="w-7 h-7 m-2  fill-gray-400" xmlns="http://www.w3.org/2000/svg" id="Filled" viewBox="0 0 24 24" width="512" height="512"><path d="M17.721,3,16.308,1.168A3.023,3.023,0,0,0,13.932,0H10.068A3.023,3.023,0,0,0,7.692,1.168L6.279,3Z"/><circle cx="12" cy="14" r="4"/><path d="M19,5H5a5.006,5.006,0,0,0-5,5v9a5.006,5.006,0,0,0,5,5H19a5.006,5.006,0,0,0,5-5V10A5.006,5.006,0,0,0,19,5ZM12,20a6,6,0,1,1,6-6A6.006,6.006,0,0,1,12,20Z"/></svg>

        </button>
        <!-- images -->
        <button class=" hover:bg-gray-800 hover:rounded-xl">
        <svg class="w-6 h-6 m-2  fill-gray-400" xmlns="http://www.w3.org/2000/svg" id="Filled" viewBox="0 0 24 24" width="512" height="512"><path d="M11.122,12.536a3,3,0,0,0-4.244,0l-6.84,6.84A4.991,4.991,0,0,0,5,24H19a4.969,4.969,0,0,0,2.753-.833Z"/><circle cx="18" cy="6" r="2"/><path d="M19,0H5A5.006,5.006,0,0,0,0,5V16.586l5.464-5.464a5,5,0,0,1,7.072,0L23.167,21.753A4.969,4.969,0,0,0,24,19V5A5.006,5.006,0,0,0,19,0ZM18,10a4,4,0,1,1,4-4A4,4,0,0,1,18,10Z"/></svg>

        </button>
        <!-- voice -->
        <button class=" hover:bg-gray-800 hover:rounded-xl ">
        <svg class="w-7 h-7 m-2  fill-gray-400" xmlns="http://www.w3.org/2000/svg" id="Layer_1" data-name="Layer 1" viewBox="0 0 24 24" width="512" height="512"><path d="M22,13a9.01,9.01,0,0,1-9,9H11a9.011,9.011,0,0,1-9-9H0A11.013,11.013,0,0,0,11,24h2A11.013,11.013,0,0,0,24,13Z"/><path d="M9,13H5.071a7,7,0,0,0,13.858,0H15V11h4V8H15V6h3.929A7,7,0,0,0,5.071,6H9V8H5v3H9Z"/></svg>

        </button>
        <!-- <svg class="w-7 h-7 m-2" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"   viewBox="0 0 410.414 410.414" style="enable-background:new 0 0 410.414 410.414;" xml:space="preserve"> <g>  <path d="M323.14,148.583c-5.523,0-10,4.477-10,10v53.789c0,63.615-55.755,86.282-107.934,86.282S97.273,275.987,97.273,212.372   v-53.789c0-5.523-4.477-10-10-10s-10,4.477-10,10v53.789c0,33.888,13.334,61.751,38.561,80.58   c20.407,15.231,47.703,23.932,79.373,25.441v72.022h-61.52c-5.523,0-10,4.477-10,10s4.477,10,10,10h143.039c5.523,0,10-4.477,10-10   s-4.477-10-10-10h-61.52v-72.022c31.67-1.508,58.966-10.21,79.373-25.441c25.227-18.828,38.561-46.692,38.561-80.58v-53.789   C333.14,153.06,328.663,148.583,323.14,148.583z"/>  <path d="M138.642,267.736c16.35,10.689,38.746,16.109,66.565,16.109s50.215-5.42,66.565-16.109   c19.517-12.76,29.833-33.036,29.833-58.637V85.573C301.606,35.189,261.965,0,205.207,0c-56.758,0-96.398,35.189-96.398,85.573   v123.526C108.808,234.7,119.125,254.976,138.642,267.736z M205.207,20c38.028,0,76.399,20.276,76.399,65.573v9.207h-40.312   c-5.523,0-10,4.477-10,10s4.477,10,10,10h40.312v23.305h-40.312c-5.523,0-10,4.477-10,10s4.477,10,10,10h40.312v23.305h-40.312   c-5.523,0-10,4.477-10,10s4.477,10,10,10h40.312v7.71c0,19.052-6.796,32.756-20.778,41.897   c-12.857,8.406-32.091,12.849-55.621,12.849s-42.763-4.443-55.621-12.849c-13.981-9.141-20.777-22.845-20.777-41.897v-7.71h40.312   c5.523,0,10-4.477,10-10s-4.477-10-10-10h-40.312v-23.305h40.312c5.523,0,10-4.477,10-10s-4.477-10-10-10h-40.312V114.78h40.312   c5.523,0,10-4.477,10-10s-4.477-10-10-10h-40.312v-9.207C128.808,40.276,167.179,20,205.207,20z"/></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g></svg>  -->

            <input bind:value="{messengerValue}" on:keypress="{keyHandler}" type="text" name="string" placeholder="Type Anything" class="input  h-10 min-w-fit grow rounded-2xl transition-all duration-250 ease-linear " on:focus="{()=>{writeFocus=true}}" on:blur="{()=>{writeFocus = false}}">
            <button  class="w-10 h-10 ml-2  hover:bg-gray-800 hover:rounded-xl">
                <svg class="w-10 h-10" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM7 9a1 1 0 100-2 1 1 0 000 2zm7-1a1 1 0 11-2 0 1 1 0 012 0zm-.464 5.535a1 1 0 10-1.415-1.414 3 3 0 01-4.242 0 1 1 0 00-1.415 1.414 5 5 0 007.072 0z" clip-rule="evenodd"></path></svg>
                <!-- <svg class=" w-12 h-12" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M17.75 3C19.5449 3 21 4.45507 21 6.25V12.999L16.2512 13L16.038 13.0069C14.8404 13.0848 13.8195 13.8111 13.3236 14.8393C12.9039 14.9527 12.4636 15.0095 12.0009 15.0095C10.9561 15.0095 10.0214 14.7199 9.17654 14.1356C8.83587 13.9 8.36869 13.9851 8.13307 14.3258C7.89744 14.6665 7.98259 15.1336 8.32326 15.3693C9.41979 16.1277 10.6524 16.5095 12.0009 16.5095C12.2279 16.5095 12.4518 16.4987 12.6724 16.477L13.0009 16.4364L13 21H6.25C4.45507 21 3 19.5449 3 17.75V6.25C3 4.45507 4.45507 3 6.25 3H17.75ZM20.341 14.7197L14.7197 20.341C14.6508 20.4098 14.578 20.4737 14.5017 20.5325L14.5008 16.2501L14.5066 16.1066C14.5759 15.2573 15.2516 14.5794 16.0999 14.5065L16.2514 14.5L20.5332 14.5007C20.4743 14.5774 20.4101 14.6505 20.341 14.7197ZM9.00045 7.75116C8.31048 7.75116 7.75116 8.31048 7.75116 9.00045C7.75116 9.69041 8.31048 10.2497 9.00045 10.2497C9.69041 10.2497 10.2497 9.69041 10.2497 9.00045C10.2497 8.31048 9.69041 7.75116 9.00045 7.75116ZM15.0004 7.75116C14.3105 7.75116 13.7512 8.31048 13.7512 9.00045C13.7512 9.69041 14.3105 10.2497 15.0004 10.2497C15.6904 10.2497 16.2497 9.69041 16.2497 9.00045C16.2497 8.31048 15.6904 7.75116 15.0004 7.75116Z" fill="#212121"/></svg> -->
            </button>
            {#if writeFocus || messengerValue.length>0 || true }
            
            <button on:click="{SendMessage}" type="submit" class=" w-10  h-10   "> 
                <svg class=" fill-cyan-600" style="enable-background:new 0 0 24 24;" version="1.1" viewBox="0 0 24 24" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><g id="info"/><g id="icons"><path d="M21.5,11.1l-17.9-9C2.7,1.7,1.7,2.5,2.1,3.4l2.5,6.7L16,12L4.6,13.9l-2.5,6.7c-0.3,0.9,0.6,1.7,1.5,1.2l17.9-9   C22.2,12.5,22.2,11.5,21.5,11.1z" id="send"/></g></svg>
            </button>
            {/if}
        </div>
                <!-- </form> -->


</div>