<script lang="ts">
	// your script goes here
	import userBanner from './images/whiteBeard.jpg';
	// import Moji2 from "./_images/moji2.jpg"
	// import { MyPro} from "$lib/store2"
	import Messaging from './icons/message.svelte';
	import Microphone from './icons/microphone.svelte';
	import Video from './icons/video.svelte';
	import Announcement from './icons/announcement.svelte';
	import Hash from './icons/hash.svelte' 
	import Rule from './icons/rule.svelte'
	import InvitePeople from './icons/invitePeople.svelte';
	import AeroUp from './icons/aeroUp.svelte';
	import AeroDown from './icons/aeroDown.svelte';
	import AudioCall from '$lib/Navbar/profileImg/audioCall.svelte';
	import RedOff from './icons/redOff.svelte';
	import Speaker from './icons/Speaker.svelte';
	import Notification from './icons/notification.svelte';
	import Settings from './icons/settings.svelte';

  import UserDashPannel from "$lib/UserDashPanel/index.svelte"

	export let MyPro: any;
	let subFieldName: string = '| Python |';
	let listOfField = [
		{ type: 'announcement', name: 'General' },
		{ type: 'message', name: 'messages' },
		{ type: 'audio', name: 'speaking' },
		{ type: 'video', name: 'Conference' },
		{ type: 'message', name: 'Competitive Programming' }
	];
	let demoField: boolean = false;
  // let SearchStyle:string="sm"

	let showRoomList = false;
	let listOfField2 = [
		{
			name: 'Dave',
			notification: true,
			icon: Hash,
			dropdownIcon: true,
			content: [
				{ name: 'rules', notification: true, icon: Rule, notificationNumber: 3 },
				{ name: 'social-media', notification: false, icon: Announcement, notificationNumber: 0 },
				{ name: 'annauncement', notification: true, icon: Announcement, notificationNumber: 10 }
			]
		},
		{
			name: 'TechnologyTechnologyTechnologyTechnology',
			notification: false,
			icon: Hash,
			dropdownIcon: false,
			content: [
				{ name: 'tech-news', notification: false, icon: Announcement, notificationNumber: 0 },
				{ name: 'pc-Build', notification: false, icon: Hash, notificationNumber: 0 },
				{
					name: 'pc-newsTechnologyTechnologyTechnologyTechnology',
					notification: false,
					icon: Hash,
					notificationNumber: 0
				},
				{ name: 'annauncement', notification: false, icon: Announcement, notificationNumber: 10 }
			]
		},
		{
			name: 'Voice Chat',
			notification: false,
			icon: Hash,
			dropdownIcon: true,
			content: [
				{ name: 'Discussion', notification: false, icon: Announcement, notificationNumber: 0 },
				{ name: 'New-topic', notification: false, icon: Hash, notificationNumber: 0 }
			]
		},
		{
			name: 'Tech-chat',
			notification: true,
			icon: Hash,
			dropdownIcon: true,
			content: [
				{ name: 'tech-news', notification: false, icon: Announcement, notificationNumber: 0 },
				{ name: 'pc-news', notification: false, icon: Hash, notificationNumber: 0 },
				{ name: 'annauncement', notification: true, icon: Announcement, notificationNumber: 10 }
			]
		},
		{
			name: 'Random',
			notification: true,
			icon: Hash,
			dropdownIcon: true,
			content: [{ name: 'general', notification: false, icon: Announcement, notificationNumber: 0 }]
		},
		{
			name: 'Other Topic',
			notification: true,
			icon: Hash,
			dropdownIcon: true,
			content: [
				{ name: 'meeting', notification: false, icon: Announcement, notificationNumber: 0 },
				{ name: 'new-tech', notification: false, icon: Hash, notificationNumber: 0 }
			]
		},
		{
			name: 'Dave',
			notification: true,
			icon: Hash,
			dropdownIcon: true,
			content: [
				{ name: 'rules', notification: true, icon: Rule, notificationNumber: 3 },
				{ name: 'social-media', notification: false, icon: Announcement, notificationNumber: 0 },
				{ name: 'annauncement', notification: true, icon: Announcement, notificationNumber: 10 }
			]
		},
		{
			name: 'TechnologyTechnologyTechnologyTechnology',
			notification: false,
			icon: Hash,
			dropdownIcon: false,
			content: [
				{ name: 'tech-news', notification: false, icon: Announcement, notificationNumber: 0 },
				{ name: 'pc-Build', notification: false, icon: Hash, notificationNumber: 0 },
				{
					name: 'pc-newsTechnologyTechnologyTechnologyTechnology',
					notification: false,
					icon: Hash,
					notificationNumber: 0
				},
				{ name: 'annauncement', notification: false, icon: Announcement, notificationNumber: 10 }
			]
		},
		{
			name: 'Voice Chat',
			notification: false,
			icon: Hash,
			dropdownIcon: true,
			content: [
				{ name: 'Discussion', notification: false, icon: Announcement, notificationNumber: 0 },
				{ name: 'New-topic', notification: false, icon: Hash, notificationNumber: 0 }
			]
		},
		{
			name: 'Tech-chat',
			notification: true,
			icon: Hash,
			dropdownIcon: true,
			content: [
				{ name: 'tech-news', notification: false, icon: Announcement, notificationNumber: 0 },
				{ name: 'pc-news', notification: false, icon: Hash, notificationNumber: 0 },
				{ name: 'annauncement', notification: true, icon: Announcement, notificationNumber: 10 }
			]
		},
		{
			name: 'Random',
			notification: true,
			icon: Hash,
			dropdownIcon: true,
			content: [{ name: 'general', notification: false, icon: Announcement, notificationNumber: 0 }]
		},
		{
			name: 'Other Topic',
			notification: true,
			icon: Hash,
			dropdownIcon: true,
			content: [
				{ name: 'meeting', notification: false, icon: Announcement, notificationNumber: 0 },
				{ name: 'new-tech', notification: false, icon: Hash, notificationNumber: 0 }
			]
		}
	];

	let userOptions = { video: false, microphone: false, notificationSound: false, sound: false };
  let showUsername:boolean=false
</script>

<!-- markup (zero or more items) goes here -->

<div class="w-60 h-screen flex flex-col ">
	<div
		class=" w-60 bg-[#2f3136] flex flex-col overflow-y-auto no-scrollbar transition-all duration-500 ease-linear  "
	>
		<!-- banner image -->
		<div class=" w-full  " style="">
			<img src={userBanner} alt="" class="aspect-video w-full object-cover " />
		</div>
		<!-- cutttt note -->
		<!-- mysubfield -->
		<!-- mysubfield -->

		<div class="px-1 pt-4 ">
			<ul class="flex flex-col space-y-2">
				<!-- content -->
				{#each listOfField2 as li}
					<!-- content here -->
					<li class="">
						<div
							class="relative  flex justify-between {li.notification && !li.dropdownIcon
								? ' text-white '
								: 'text-gray-500'}  hover:text-white focus-within:text-white"
						>
							<div
								class="flex items-center w-full"
								on:click={() => (li.dropdownIcon = !li.dropdownIcon)}
							>
								<div class="absolute inset-y-0 left-0 flex items-center pl-2 pointer-events-none">
									<!-- room icon -->
									<!-- {li.icon} -->
									<svg
										class="w-5 h-5 "
										fill="currentColor"
										viewBox="0 0 20 20"
										xmlns="http://www.w3.org/2000/svg"
										><path
											fill-rule="evenodd"
											d="M14.243 5.757a6 6 0 10-.986 9.284 1 1 0 111.087 1.678A8 8 0 1118 10a3 3 0 01-4.8 2.401A4 4 0 1114 10a1 1 0 102 0c0-1.537-.586-3.07-1.757-4.243zM12 10a2 2 0 10-4 0 2 2 0 004 0z"
											clip-rule="evenodd"
										/></svg
									>
									<!-- <svg class=" w-5 h-5 stroke-current" fill="none" viewBox="0 0 24 24"><path stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M7.75 19.25H16.25C17.3546 19.25 18.25 18.3546 18.25 17.25V9L14 4.75H7.75C6.64543 4.75 5.75 5.64543 5.75 6.75V17.25C5.75 18.3546 6.64543 19.25 7.75 19.25Z"                     ></path>                     <path                       stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M18 9.25H13.75V5"                     ></path>                     <path                       stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M9.75 15.25H14.25"                     ></path>                     <path                       stroke="currentColor"                       stroke-linecap="round"                       stroke-linejoin="round"                       stroke-width="1.5"                       d="M9.75 12.25H14.25"                     ></path>                   </svg> -->
								</div>
								<a
									href="/"
									class="truncate inline-block w-full py-2 pl-8 pr-4 text-xs rounded hover:bg-gray-800 focus:outline-none focus:ring-1 focus:ring-gray-500 focus:bg-gray-800"
								>
									{li.name}
								</a>
							</div>
							<button
								class="absolute right-0 flex items-center p-1 "
								tabindex="-1"
								on:click={() => (li.dropdownIcon = !li.dropdownIcon)}
							>
								{#if li.dropdownIcon}
									<!-- aero up -->
									<svg
										class="w-5 h-5"
										fill="currentColor"
										viewBox="0 0 20 20"
										xmlns="http://www.w3.org/2000/svg"
										><path
											fill-rule="evenodd"
											d="M14.707 12.707a1 1 0 01-1.414 0L10 9.414l-3.293 3.293a1 1 0 01-1.414-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 010 1.414z"
											clip-rule="evenodd"
										/></svg
									>
								{:else}
									<!-- aero down -->
									<svg
										class="w-5 h-5"
										fill="currentColor"
										viewBox="0 0 20 20"
										xmlns="http://www.w3.org/2000/svg"
										><path
											fill-rule="evenodd"
											d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
											clip-rule="evenodd"
										/></svg
									>
								{/if}
								<!-- <svg class="w-5 h-5 stroke-current" fill="none" viewBox="0 0 24 24">                   <path                     stroke="currentColor"                     stroke-linecap="round"                     stroke-linejoin="round"                     stroke-width="1.5"                     d="M15.25 10.75L12 14.25L8.75 10.75"                   ></path>                 </svg> -->
							</button>
						</div>
						{#if li.dropdownIcon}
							{#each li.content as con}
								<!-- content here -->
								<div class=" pl-2 text-left">
									<ul
										class="flex flex-col {con.notification
											? 'text-white'
											: 'text-gray-500'}   border-l border-gray-700"
									>
										<li class="flex flex-row">
											{#if con.notification}
												<!-- content here -->
												<svg
													class="fill-white w-2 h-2 mt-3  -left-4 relative "
													viewBox="0 0 512 512"
													xmlns="http://www.w3.org/2000/svg"
													><path
														d="M512 256c0 141.4-114.6 256-256 256s-256-114.6-256-256s114.6-256 256-256S512 114.6 512 256z"
													/></svg
												>
											{/if}

											<a
												href="/"
												class="truncate inline-block w-full {con.notification
													? 'px-2'
													: 'px-4'}  py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white"
											>
												{con.name}
											</a>
										</li>
										<!-- <li>
                          <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white" >Categories</a >
                          </li>
                          <li>
                            <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white" >Instructors</a >
                            </li>
                            <li>
                              <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-gray-800 hover:text-white focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white" >Video Library</a >
                              </li>
                            </ul> -->
									</ul>
								</div>
							{/each}
						{/if}
					</li>
				{/each}
			</ul>
		</div>

		<!-- <div class=" absolute bg-slate-500 h-14 w-60 align-bottom">

      </div> -->
	</div>
	<div class="">
    <UserDashPannel  {showUsername} />
  </div>
</div>

<style>
	/* your styles go here */
</style>
