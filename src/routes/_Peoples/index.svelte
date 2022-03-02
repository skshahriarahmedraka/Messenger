<script lang="ts">
    import UserImg from "./_images/ot2.jpg"

    let items1 ={
        "image":UserImg,
        "username":"Sk Shahriar Ahmed Raka",
        "mobile":"+8801838810189",
        "bio":"Lazy is such an ugly word, I prefer to selectively participate. I already want to take a nap tomorrow. ",
        "url":"shahriarraka",
        "notification":true ,
        "contents":[
            [213,"images"],
            [41,"videos"],
            [5,"files"],
            [6,"shared links"],
            [41,"audios"],
            [64,"gifs"]
        ],
        "recomendationlist":[
            []
        ],
        "cutomize":{
            "color":"blue",
            "emoji":"smile",
            "nickname":"ssar",
        },
        

    }
    
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div class=" w-1/6  bg-[#2f3136] transition-all duration-200 ease-linear flex flex-col overflow-y-scroll scrol">
     <!-- icon 80*80 -->
     <!-- icon name div 380*100 -->
     <!-- comonent 44 width -->
     {#each Array.from(Array(20+1).keys()).slice(1) as i }
     <div class=" w-full h-24  flex flex-row ">
                <img src="{UserImg}" alt="" class=" w-20 h-20 object-cover rounded-3xl my-2 mx-2">
                
                <div class=" flex flex-col ">
                    <p class=" text-gray-300 text-lg my-2 mx-1 leading-5 line-clamp-2">
                        Sk Shahriar Ahmed Raka
                </p>
                <p class=" text-sm text-gray-400 my-2 mx-1 truncate">
                    last seen recently
                </p>
            </div>
        </div>
        {/each}
        <div class=" w-full  bg-slate-500 rounded-lg ">
            
            
        </div>
</div>