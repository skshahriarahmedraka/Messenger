
        {#each Array.from(Array(10+1).keys()).slice(1) as i }
         
   
        
        <div class=" w-full  bg-[#2f4136] flex flex-col ">
            <!-- Sub Field list -->
            <div class=" bg-gray-400 flex flex-row">
                <button class="" on:click={()=>{demoField= !demoField}}>

                    <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
                </button>
                <p class="">{subFieldName}</p>
            </div>
           {#if demoField}
                {#each listOfField as field}
                        <div class="flex flex-row ml-5 ">
                        {#if field.type==="announcement" }
                            <Announcement/>
                            <div class=" truncate">
                                {field.name}
                            </div>
                        {:else if  field.type==="message"}
                            <Messaging/>
                            <div class=" truncate">
                                {field.name}
                            </div>
                        {:else if  field.type==="audio"}
                            <Audio/>
                            <div class=" truncate">
                                {field.name}
                            </div>
                        {:else if  field.type==="video"}
                            <Video/>
                            <div class=" truncate">
                                {field.name}
                            </div>
                        
                        {/if}
                    </div>
                {/each}
           {/if}

        </div>
        {/each}