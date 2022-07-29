<script lang="ts">
    // import UserImg from "./_images/ot2.jpg"

    // let items1 ={
    //     "image":UserImg,
    //     "username":"Sk Shahriar Ahmed Raka",
    //     "mobile":"+8801838810189",
    //     "bio":"Lazy is such an ugly word, I prefer to selectively participate. I already want to take a nap tomorrow. ",
    //     "url":"shahriarraka",
    //     "notification":true ,
    //     "contents":[
    //         [213,"images"],
    //         [41,"videos"],
    //         [5,"files"],
    //         [6,"shared links"],
    //         [41,"audios"],
    //         [64,"gifs"],
    //         [23,"recomendedserver"],
    //         [67 ,"customizeChat"],
    //         [ 67,"ChangeTheme"],

    //         [4 ,"Changeimoji"],
    //         [675,"EditNickname"],
    //         [63 ,"shareThisContact"],
    //         [54 ,"DelteContact"],
    //         [436 ,"DeleteHistory"],
    //         [75 ,"DeleteChat"],
    //         [865 ,"BlockUser"]
    //     ],
    //     "recomendationlist":[
    //         []
    //     ],
    //     "cutomize":{
    //         "color":"blue",
    //         "emoji":"smile",
    //         "nickname":"ssar",
    //     },
        

    // }

    export let PeopleData:any
    
</script>

<style>
    /* your styles go here */
</style>


<div class=" w-1/6  bg-[#2f3136] transition-all duration-200 ease-linear flex flex-col overflow-y-scroll scrol"> 
     <div class=" w-full h-24  flex flex-row ">
                <img src="{PeopleData.image}" alt="" class=" w-20 h-20 object-cover rounded-3xl my-2 mx-2">
                <div class=" flex flex-col ">
                    <p class=" text-gray-300 text-lg my-2 mx-1 leading-5 line-clamp-2">
                        {PeopleData.username}
                    </p>
                <p class=" text-sm text-gray-400 my-2 mx-1 truncate">
                    {PeopleData.lastseen}
                </p>
            </div>
        </div>
    <div class=" text-gray-300">    
    <!-- mobile -->
        <div class=" w-full  h-10  bg-slate-700 rounded-lg flex flex-row place-content-center items-center">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z"></path></svg>
            <div class="">
                {PeopleData.mobile}
            </div>
        </div>
        <!-- Bio -->
        <div class=" w-full h-16 bg-slate-700 rounded-xl flex flex-row my-2 ">
            <svg class="w-6 h-6 fill-white m-1" width="24" height="24" xmlns="http://www.w3.org/2000/svg" fill-rule="evenodd" clip-rule="evenodd"><path d="M22 0h-20v24h14l6-6v-18zm-7 23h-12v-22h18v16h-6v6zm1-5h4.586l-4.586 4.586v-4.586zm-3 1h-8v1h8v-1zm0-3h-8v1h8v-1zm6-2v-1h-14v1h14zm0-4h-4v1h4v-1zm-6.006 1h-7.991l-.003-.789c-.003-.72-.006-1.615 1.314-1.92 1.483-.341 1.236-.418 1.158-.563-1.078-1.988-.71-3.173-.395-3.703.388-.651 1.089-1.025 1.923-1.025.827 0 1.523.368 1.91 1.011.545.904.409 2.222-.379 3.713-.105.196-.195.255 1.119.559 1.355.312 1.352 1.212 1.35 1.936l-.006.781zm-6.994-1h6c-.007-.547-.07-.626-.54-.734-.855-.198-1.629-.376-1.901-.972-.142-.311-.113-.66.087-1.039.61-1.151.758-2.146.407-2.729-.276-.458-.778-.526-1.053-.526-.48 0-.857.19-1.063.537-.352.59-.201 1.58.414 2.714.204.377.236.727.095 1.039-.269.598-1.036.774-1.847.962-.525.121-.593.202-.599.748zm13-2v-1h-4v1h4zm0-4h-4v1h4v-1z"/></svg>
            <div class=" line-clamp-2 ">
                {PeopleData.bio}
            </div>
        </div>
        <!-- user link -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207"></path></svg>
            <div class="">
                {PeopleData.url}
            </div>
        </div>
        <!-- notification -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.notification}
            </div>
        </div>
        <!-- photo -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[0][0]} {PeopleData.contents[0][1]}
            </div>
        </div>
        <!-- videos -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[1][0]} {PeopleData.contents[1][1]}

            </div>
        </div>
        <!-- files -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[2][0]} {PeopleData.contents[2][1]}

            </div>
        </div>
        <!-- shared links -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[3][0]} {PeopleData.contents[3][1]}

            </div>
        </div>
        <!-- Audio -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[4][0]} {PeopleData.contents[4][1]}

            </div>
        </div>
        <!-- Gifs -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[5][0]} {PeopleData.contents[5][1]}

            </div>
        </div>
        <!-- recomended Server -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[6][0]} {PeopleData.contents[6][1]}

            </div>
        </div>
        <!-- customize chat -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[7][0]} {PeopleData.contents[7][1]}

            </div>
        </div>
        <!-- change theme -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[8][0]} {PeopleData.contents[8][1]}

            </div>
        </div>
        <!-- change imoji -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[9][0]} {PeopleData.contents[9][1]}

            </div>
        </div>
        <!-- edit nickname -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[10][0]} {PeopleData.contents[10][1]}

            </div>
        </div>
        <!-- share this contact -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[11][0]} {PeopleData.contents[11][1]}

            </div>
        </div>
        <!-- edit this contact -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[12][0]} {PeopleData.contents[12][1]}

            </div>
        </div>
        <!-- delete this contact -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[13][0]} {PeopleData.contents[13][1]}

            </div>
        </div>
        <!-- clear history -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[14][0]} {PeopleData.contents[14][1]}

            </div>
        </div>
        <!-- delete chat -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[15][0]} {PeopleData.contents[15][1]}

            </div>
        </div>
        <!-- block user -->
        <div class="w-full h-10 flex flex-row bg-slate-700 rounded-lg my-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg>
            <div class="">
                {PeopleData.contents[0][0]} {PeopleData.contents[0][1]}

            </div>
        </div>
    </div>
</div>