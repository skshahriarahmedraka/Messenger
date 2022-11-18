<script lang="ts">

    import type { PageData } from './$types';
        import { UserProData } from '$lib/store2';
    
        export let data: PageData;
        
        let { Userdata } = data;
        console.log("🚀 ~ file: +page.svelte ~ line 10 ~ Userdata", Userdata)
        UserProData.update((d) => (d = Userdata));
    
        
        // import Twitter from '$lib/foot/icons/twitter.svelte';
        import ContractAdimin from '$lib/settingsCom/ContractAdimin.svelte';
        import ContractGameDev from '$lib/settingsCom/ContractGameDev.svelte';
        import General from '$lib/settingsCom/General.svelte';
        import PassSec from '$lib/settingsCom/Pass&Sec.svelte';
        import PaymentManagement from '$lib/settingsCom/PaymentManagement.svelte';
        import RequestToken from '$lib/settingsCom/RequestToken.svelte';
        import Wishlist from '$lib/settingsCom/Group List.svelte';
        import Coin from '$lib/icons/coin2.svelte';
        import List from '$lib/icons/list.svelte';
        import Message from '$lib/icons/message copy.svelte';
        import Paycard from '$lib/icons/paycard.svelte';
        import Security from '$lib/icons/security.svelte';
        import Transaction from '$lib/icons/transaction.svelte';
        import User from '$lib/icons/user.svelte';
        import {UserSettingSelect} from "$lib/store2";
        import Notification from '$lib/icons/notification.svelte'
        import Notifications from "$lib/settingsCom/FriendList.svelte"
	import FriendList from '$lib/settingsCom/FriendList.svelte';
	import Friends from '$lib/icons/friends.svelte';
	import Users from '$lib/icons/users.svelte';
	import GroupList from '$lib/settingsCom/Group List.svelte';
    
        let list = [
            { name: 'General', link: '', icon: User },
            { name: 'Friend List', link: '', icon: Friends },
            { name: 'Group List', link: '', icon: Users },
            { name: 'Payment Mangement', link: '', icon: Paycard },
            { name: 'Request Token', link: '', icon: Coin },
            { name: 'Password & Security', link: '', icon: Security },
            { name: 'Transaction', link: '', icon: Transaction },
            // { name: 'Wishlist', link: '', icon: List },
            { name: 'Contact Admin', link: '', icon: Message },
            // { name: 'Contact Game Developer', link: '', icon: Message }
        ];
        let active = $UserSettingSelect
    </script>
    
    <!-- markup (zero or more items) goes here -->
<div class=" min-h-screen w-full bg-[#36393f] flex flex-col gap-5 justify-between overflow-y-scroll">
    <div class="flex flex-row justify-center  gap-1 ">
        
        <div class="flex flex-col gap-1 xs:hidden sm:hidden md:hidden h-fit mt-5  ">
            {#each list as l}
                <div
                    on:click={() => {
                        UserSettingSelect.set(l.name)
                    }}
                    
                    on:keypress={()=>{}}
                    
                    class=" flex h-14 w-72 flex-row justify-center gap-2   rounded-md bg-gray-600 hover:cursor-pointer hover:bg-opacity-80  {active ===
                    l.name
                        ? ' border-x-8 bg-slate-500 border-blue-500 '
                        : ''} items-center "
                >
                    <svelte:component this={l.icon} class=" h-6 w-6 stroke-slate-200 " />
                    <p class=" font-Poppins text-slate-200 hover:text-white">{l.name}</p>
                </div>
            {/each}
        </div>
        <div
        class="  flex h-fit min-h-[300px] w-[880px] flex-col rounded-lg  bg-slate-700 bg-opacity-30 font-Poppins  "
        >
            {#if $UserSettingSelect === 'General'}
                <General />
                
            {:else if $UserSettingSelect === 'Payment Mangement'}
                <PaymentManagement />
            {:else if $UserSettingSelect === 'Request Token'}
                <RequestToken />
            {:else if $UserSettingSelect ==='Password & Security'}
                <PassSec />
            {:else if  $UserSettingSelect ==='Transaction'}
            <Transaction />
            {:else if $UserSettingSelect ==='Group List'}
                <GroupList />
            {:else if  $UserSettingSelect ==='Contact Admin'}
                <ContractAdimin />
            <!-- {:else if active ==='Contact Game Developer'}
                <ContractGameDev /> -->
                {:else if $UserSettingSelect ==='Friend List'}
                <FriendList />
            {:else}
                <!-- else content here -->
            {/if}
        </div>
    </div>
</div>
    
    <style>
        /* your styles go here */
    </style>
    