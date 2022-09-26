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
</script>

<svelte:head>
	<title>Accord</title>
</svelte:head>


{#if loadingState}
	<Loading class=" scale-150" />
{:else if $page.error}
	<!-- <Error1/> -->
	<Error2 />
{:else}
	<div class="flex flex-row h-screen w-full transition-all duration-500 ease-linear ">
		<slot />
	</div>
{/if}

<style>

</style>
