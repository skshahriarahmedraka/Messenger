<script lang="ts" context="module">
	export async function load({ session, fetch }: any) {
		if (!session.user.authenticated) {
			return {
				status: 302,
				redirect: '/login'
			};
		}

		// FETCH USER PROFILE
		let x3 = {
			Name: "Sk Shahriar Ahmed Raka",
			Userid: 'skraka',
			ProfileImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
			CoverImage:"https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png"
		};

		return {
			props: {},
			stuff: {
				MyPro: x3
			}
		};
	}
</script>

<script lang="ts">
	import '../app.css';
	// import peopleIcon from "./icons/users.svg"
	import Dock from '$lib/Dock/index.svelte';
	import Servers from '$lib/Servers/index.svelte';
	import NavBar from '$lib/Navbar/index.svelte';
	import Peoples from '$lib/Peoples/index.svelte';
	import Messages from '$lib/Messages/index.svelte';
	import { showPeopleList, ChatOrDock } from '$lib/store2';
	import Loading from '$lib/Loading/solarAccord.svelte';
	import { afterUpdate } from 'svelte';

	import Chat from '$lib/Chats/index.svelte';
	// export let MyPro:any

	let showPeopleListValue: number;
	// let ChatOrDockValue: number;

	showPeopleList.subscribe((val) => {
		showPeopleListValue = val;
	});

	// ChatOrDock.subscribe((val) => {
	// 	ChatOrDockValue = val;
	// });
	let loadingState: boolean = true;
	afterUpdate(() => {
		loadingState = false;
	});

	let x2: any = [
		{
			ServerURL: 'postgres',
			Name: 'Postgres',
			NewMessage: true,
			NumberOfNewMessage: 3,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'mongodb',
			Name: 'MongoDB',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: true,
			NumberOfNotification: 3,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'docker',
			Name: 'Docker',
			NewMessage: true,
			NumberOfNewMessage: 4,
			Notification: true,
			NumberOfNotification: 6,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'kubernetes',
			Name: 'Kubernetes',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: true,
			NumberOfNotification: 99,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'redis',
			Name: 'Redis',
			NewMessage: true,
			NumberOfNewMessage: 7,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'tensorflow',
			Name: 'Tensorflow',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'macity',
			Name: 'Mancity',
			NewMessage: true,
			NumberOfNewMessage: 2,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'cpp',
			Name: 'C++',
			NewMessage: true,
			NumberOfNewMessage: 4,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'ubuntu',
			Name: 'Ubuntu',
			NewMessage: true,
			NumberOfNewMessage: 7,
			Notification: true,
			NumberOfNotification: 2,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'realmadrid',
			Name: 'Real Madrid',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'bangladesh',
			Name: 'bangladesh',
			NewMessage: true,
			NumberOfNewMessage: 6,
			Notification: true,
			NumberOfNotification: 4,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'c',
			Name: 'C',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'coffee',
			Name: 'Coffee',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: true,
			NumberOfNotification: 45,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'chelsea',
			Name: 'Chelsea',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'go',
			Name: 'Go',
			NewMessage: true,
			NumberOfNewMessage: 9,
			Notification: true,
			NumberOfNotification: 65,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'mazda',
			Name: 'Mazda',
			NewMessage: false,
			NumberOfNewMessage: 2,
			Notification: true,
			NumberOfNotification: 5,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'pinterest',
			Name: 'Pinterest',
			NewMessage: false,
			NumberOfNewMessage: 0,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'svelte',
			Name: 'Svelte',
			NewMessage: true,
			NumberOfNewMessage: 3,
			Notification: false,
			NumberOfNotification: 0,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'tea',
			Name: 'Tea',
			NewMessage: true,
			NumberOfNewMessage: 345,
			Notification: true,
			NumberOfNotification: 999,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		},
		{
			ServerURL: 'linux',
			Name: 'Linux',
			NewMessage: true,
			NumberOfNewMessage: 3,
			Notification: true,
			NumberOfNotification: 90,
			ServerImage:
				'https://res.cloudinary.com/dqo0ssnti/image/upload/v1658404332/samples/1652821445210_umwhko.jpg'
		}
	];
	import { page } from '$app/stores';
	import Error1 from '$lib/Error/error1.svelte';
	import Error2 from '$lib/Error/error2.svelte';
	import { fade, fly } from 'svelte/transition';
</script>

{#if loadingState}
	<Loading class=" scale-150" />
{:else if $page.error}
	<!-- <Error1/> -->
	<Error2 />
{:else}
	<div class="h-screen w-full transition-all flex flex-row  duration-500 ease-linear ">
		<Dock ServerList={x2} />

		<slot />
	</div>
{/if}

<style>
</style>
