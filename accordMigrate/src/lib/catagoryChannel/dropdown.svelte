<script lang="ts">
	let serverTypes = ['chat', 'video', 'voice', 'annauncement', 'post'];
	let setType = 'sk shahriargsfsfdg';
	let show = false;
</script>

<!-- markup (zero or more items) goes here -->

<!-- This example requires Tailwind CSS v2.0+ -->
<div class="relative inline-block w-1/3 text-left">
	
		<button
			type="button"
			class=" h-8 w-fit flex  flex-row self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]   text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800"
		>
			<p class="ml-2">

                {setType}
            </p>
			<svg
				class="mx-2 mt-1 h-5 w-5"
				xmlns="http://www.w3.org/2000/svg"
				viewBox="0 0 20 20"
				fill="currentColor"
				aria-hidden="true"
			>
				<path
					fill-rule="evenodd"
					d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z"
					clip-rule="evenodd"
				/>
			</svg>
		</button>
	

	{#if show}
		<!-- content here -->
		<div
			class="absolute left-0 z-10 mt-0 -ml-28  flex w-56 origin-top-right  flex-col gap-1 rounded-md bg-transparent shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
			role="menu"
			aria-orientation="vertical"
			aria-labelledby="menu-button"
			tabindex="-1"
		>
			<!-- Active: "bg-gray-100 text-gray-900", Not Active: "text-gray-700" -->
			{#each serverTypes as s}
				<!-- content here -->
				<button
					type="button"
					class="inline-flex h-8 w-full  justify-center  self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800"
					id="menu-button"
					aria-expanded="true"
					aria-haspopup="true"
				>
					{s}
				</button>
			{/each}
		</div>
	{:else}
		<!-- else content here -->
	{/if}
</div>

<style>
	/* your styles go here */
</style>
