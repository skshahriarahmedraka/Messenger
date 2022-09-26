<script lang="ts" context="module">
	export async function load({ session }: any) {
		if (!session.user.authenticated) {
			return {
				status: 302,
				redirect: '/login'
			};
		}

		return {
			props: {},
			
		};
	}
</script>

<script lang="ts">
</script>

<!-- markup (zero or more items) goes here -->

<svelte:head>
	<title>Settings</title>
</svelte:head>

<div class=" w-full h-screen overflow-hidden bg-stone-600 flex flex-col flex-nowrap  " >
	<h1 class="">User Settings</h1>
</div>

<style>
	/* your styles go here */
</style>
