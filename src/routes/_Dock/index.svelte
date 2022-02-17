<script lang="ts">
    // your script goes here
        import profilepic from "./_profilePic/law.png"
    import profilepic2 from "./_profilePic/ren.jpeg"
    import profilepic3 from "./_profilePic/e.jpg"
    import profilepic4 from "./_profilePic/s2.png"


  

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

<div class=" flex flex-col  flex-none w-20 bg-[#202225] overflow-y-auto no-scrollbar  ">

        <!-- <div class="avatar online">
            <div class=" w-12 rounded-full">
                <img src="https://api.lorem.space/image/face?hash=28212">
            </div>
        </div> -->
        {#each UserList as u}
             <!-- content here -->
        <div class="avatar online py-1 px-2">
            <div class="  w-16 hover:rounded-3xl rounded-md active:rounded-lg hover:ring hover:ring-cyan-500 transition-all duration-150 ease-linear cursor-pointer {false ? "ring  ring-offset-base-100  ring-blue-500" : "" }">
                <!-- <img src="https://api.lorem.space/image/face?hash=40361"> -->
                <img src="{u.profilePic }" alt="" class="">
            </div>
        </div>
        {/each}
        
        <!-- <div class="avatar offline">
            <div class=" w-16  active:mask active:mask-squircle rounded-md ring  ring-offset-base-100 ring-offset-2 ring-blue-600">
                <img src="https://api.lorem.space/image/face?hash=40361">
            </div>
        </div> -->
        
    </div>