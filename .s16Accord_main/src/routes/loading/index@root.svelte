<script lang="ts" context="module">
	export async function load({ session }: any) {
		if (!session.user.authenticated) {
			return {
				status: 302,
				redirect: '/login'
			};
		}

		return {
			props: {}
		};
	}
</script>

<script lang="ts">
	import Loading from '$lib/Loading/solarAccord.svelte';
</script>

<!-- markup (zero or more items) goes here -->

<Loading class=" scale-150" />

<style>
	/* your styles go here */
</style>
