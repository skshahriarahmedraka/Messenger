<script lang="ts">
    import Ek2 from "./_profilePic/e2.jpg"
    export let message:string

</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->


     <!-- content here -->
    <div class=" flex flex-row max-w-full mr-20 mt-4 ml-5   text-left text-base text-white overflow-ellipsis " >
          
        <img src="{Ek2}" alt="" class=" w-14 h-14 m-2 rounded-2xl">
         <div class="flex flex-col">
             <div class="flex flex-row">
                 <div class="">
                     sk shahriar ahmed raka
                 </div>
                 <div class=" text-xs mt-2 ml-2">
                     Yesterday at 9:40 PM
                 </div>

             </div>
             <div class="">
                 {message}
             </div>
         </div>
     </div>

