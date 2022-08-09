<script lang="ts">
	import Angryx from './reactions/angryx.svelte';

	import Carex from './reactions/carex.svelte';
	import Cry from './reactions/cry.svelte';
	import Hahax from './reactions/hahax.svelte';

	import Happy from './reactions/happy.svelte';

	import Likex from './reactions/likex.svelte';
	import Lovex from './reactions/lovex.svelte';
	import Sadx from './reactions/sadx.svelte';

	import Unlike from './reactions/unlike.svelte';
	import Wowx from './reactions/wowx.svelte';

	import Ek2 from './_profilePic/e2.jpg';
	export let message: any;
	export let MyPro: any;

	// var person = { fname: 'Nick', lname: 'Jonas', age: 26 };
	// for (let x in person) {
	// 	console.log(x + ': ' + person[x]);
	// }

	function ReactCount(react: { [key: string]: number }) {
		let r = [
			[0, 0],
			[1, 0],
			[2, 0],
			[3, 0],
			[4, 0],
			[5, 0],
			[6, 0],
			[7, 0],
			[8, 0],
			[9, 0]
		];
		for (let x in react) {
			// console.log(x + ': ' + react[x]);
			r[react[x]][1] += 1;
		}

		r.sort(compareSecondColumn);

		function compareSecondColumn(a, b) {
			if (a[1] === b[1]) {
				return 0;
			} else {
				return a[1] > b[1] ? -1 : 1;
			}
		}

		// r = r.sort(function (a, b) {
		// 	return a[1] - b[1];
		// });

		// r.sort();
		// r.reverse();
		return r;
	}

	let reactions = [Likex, Lovex, Happy, Carex, Wowx, Hahax, Unlike, Sadx, Cry, Angryx];

	// like 1 , love 2 ,happy 3 , care 4 ,wow 5,haha 6,Unlike 7 ,sad 8, cry 9,angry 10
</script>

<!-- content here -->
<!-- { message.writer=== MyPro.Name  ? "flex-row-reverse" : "flex-row" }  -->
<!-- {#if true} -->
<div class=" flex flex-col">
	<div
		class=" flex  {message.writer === MyPro.Userid
			? 'flex-row-reverse'
			: 'flex-row'}    mr-20 mt-4 ml-5 max-w-full   overflow-ellipsis text-base text-white "
	>
		<!-- <div class="w-[80%] ?"> -->
		<img src={Ek2} alt="" class=" m-2 h-10 w-10 rounded-2xl " />
		<div class="flex max-w-[80%] flex-col flex-wrap ">
			<div
				class="flex flex-row  {message.writer != MyPro.Userid ? 'self-start' : 'self-end'} gap-1"
			>
				<button class=" h-6 text-[#04b1aa] line-clamp-1 hover:text-[#47faf4]">
					{message.writerName}
				</button>
				<div class=" mt-2 ml-2 text-xs text-[#6e6a5c]">
					{message.time}
				</div>
			</div>
			<!-- MESSAGES -->
			<div
				class=" text-[15px] {message.writer === MyPro.Userid
					? 'bg-blue-400 bg-opacity-20 p-2 '
					: 'bg-gray-700'}  rounded-lg p-2    "
			>
				{message.message}
			</div>
			<!-- REACTIONS -->
			<div class="no-scrollbar flex h-16 w-full flex-row overflow-x-scroll ">
				{#each ReactCount(message.react) as x}
					<!-- content here -->
					<!-- reactions[{x[1]}] -->

					<!-- {console.log(reactions[x[0]])} -->
					{#if x[1] != 0}
						<!-- content here -->
						<svelte:component this={reactions[x[0]]} class="h-16 w-16" />
						{/if}
						{:else}
						<!-- empty list -->
						{/each}
						<svg class="h-16 w-16" aria-hidden="false" width="24" height="24" viewBox="0 0 24 24"><path fill="currentColor" fill-rule="evenodd" clip-rule="evenodd" d="M12.2512 2.00309C12.1677 2.00104 12.084 2 12 2C6.477 2 2 6.477 2 12C2 17.522 6.477 22 12 22C17.523 22 22 17.522 22 12C22 11.916 21.999 11.8323 21.9969 11.7488C21.3586 11.9128 20.6895 12 20 12C15.5817 12 12 8.41828 12 4C12 3.31052 12.0872 2.6414 12.2512 2.00309ZM10 8C10 6.896 9.104 6 8 6C6.896 6 6 6.896 6 8C6 9.105 6.896 10 8 10C9.104 10 10 9.105 10 8ZM12 19C15.14 19 18 16.617 18 14V13H6V14C6 16.617 8.86 19 12 19Z"></path><path d="M21 3V0H19V3H16V5H19V8H21V5H24V3H21Z" fill="currentColor"></path></svg>
			</div>
		</div>
	</div>
</div>

<!-- {/if} -->
<style>
	/* your styles go here */
</style>
