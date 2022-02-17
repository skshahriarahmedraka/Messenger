<script lang="ts">
    // your script goes here
        import userBanner from "./_images/whiteBeard.jpg"
        let subFieldName:string="| Python |"
        let listOfField=[
            "General",
            "object",
            "Competitive Problem",
            ""
        ]
        

</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->


<div class=" w-52 bg-[#2f3136] flex flex-col no-scrollbar ">
        <!-- banner image -->
        <div class=" w-full aspect-video bg-cover bg-center " style="background-image: url({userBanner});">
            <!-- <img src="{userBanner}" alt="" class=" bg-cover "> -->
        </div>
        <!-- sub Field -->
        <div tabindex="0" class="collapse collapse-plus">
            <input type="checkbox"> 
            <div class="collapse-title text-sm font-medium">
                {subFieldName}
            </div>
            <div class="collapse-content"> 
                <!-- <p>tabindex="0" attribute is necessary to make the div focusable</p> -->
                <ul class="">
                    <li class="">raka</li>
                    <li class="">ssar</li>
                    <li class="">ssar</li>
                </ul>
            </div>
        </div>
        <div class=" w-full h-96 bg-[#2f4136] flex flex-col ">
            <!-- Sub Field list -->
            <div class=" bg-gray-400 flex flex-row">
                <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
                <p class="">{subFieldName}</p>
            </div>
            <div class="">
               {#each listOfField as field}
                    {field}
               {/each}
            </div>

        </div>
    </div>