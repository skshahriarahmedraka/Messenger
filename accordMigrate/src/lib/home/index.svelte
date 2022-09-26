<script lang="ts">
	let suggestedFrineds = [
		{
			name: 'md abdullah ',
			bio: 'work at Google ',
			img: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1660151599/samples/about_dmmlnj.jpg'
		},
		{
			name: '',
			bio: '',
			img: 'https://res.cloudinary.com/dqo0ssnti/image/upload/v1660151599/samples/about_dmmlnj.jpg'
		}
	];
</script>

<div class="m-3 flex flex-row flex-wrap justify-center gap-3 ">
	{#each Array.from(Array(10 + 1).keys()).slice(1) as i}
		<div class=" flex flex-col h-80  w-48    rounded-lg bg-gray-600 bg-opacity-80 ">
			<!--  -->
			<img
				src="https://res.cloudinary.com/dqo0ssnti/image/upload/v1660151599/samples/about_dmmlnj.jpg"
				alt="UserName"
				class=" h-[60%]  w-fit rounded-t-lg object-cover"
			/>
            <p class="text-base font-sans font-semibold mx-1 line-clamp-2">Shahin Alam joyShahin Alam joy Shahin Alam joy</p>
            <p class=" text-sm text-gray-300 font-sans mx-1 font-light
			lamp-2">Work at Google INCWork at Google INCWork at Google INC </p>
            <div class="flex flex-row text-base gap-2 justify-center     my-2 items-center">
                <button class="px-2 py-1 rounded-md bg-[#3982e4]">Connect</button>
                <button class="px-2  py-1 rounded-md border-0 bg-[#3a3b3c] ">Remove</button>
            </div>
		</div>
	{/each}
</div>

<style>
	/* your styles go here */
</style>
