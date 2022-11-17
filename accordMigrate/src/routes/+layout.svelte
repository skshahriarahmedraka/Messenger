

<script lang="ts">
	import '../app.css';
	import Error2 from '$lib/Error/error2.svelte';
	import { page } from '$app/stores';
	import Loading from '$lib/Loading/solarAccord.svelte';
	import { afterUpdate } from 'svelte';

	let loadingState: boolean = true;
	afterUpdate(() => {
		loadingState = false;
	});

const AccordImg = new URL('../lib/icons/accord.svg', import.meta.url).href

</script>

<!-- <svelte:head>
	<link rel="icon" href="{AccordImg}">
</svelte:head> -->


{#if loadingState}
	<Loading class=" " />
{:else if $page.error}
	<Error2 />
{:else}
	<div class="flex flex-row min-h-screen max-h-screen w-screen transition-all duration-500 ease-linear ">
		<slot />
	</div>
{/if}

<style>

</style>
