<script lang="ts">
	let day = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
	let months = [
		'January',
		'February',
		'March',
		'April',
		'May',
		'June',
		'July',
		'August',
		'September',
		'October',
		'November',
		'December'
	];
	let setDate = [1, 0, 6, 2022];
	// let setDate=[date,month,week,year]
	function FindDay() {
		// const d = new Date("July 22, 1983");
        let s=`${months[setDate[1]]} ${setDate[0]}, ${setDate[3]}`
        // let s=`${months[setDate[3]]}-${setDate[0]}-${setDate[1]}`
        // console.log("🚀 ~ file: index.svelte ~ line 23 ~ FindDay ~ s", s)
		let d = new Date(s);
		setDate[2] = d.getDay()
		
	}
	// program to check leap year
	function checkLeapYear(year: number) {
		//three conditions to find out the leap year
		if ((0 === (year % 4) && 0 != year % 100) || 0 === year % 400) {
			return true;
		} else {
			return false;
		}
	}
	function NumOfDayinMonth(i: number,j:number) {
        if (checkLeapYear(i) && months[j]==="February" ){
            return 29
        }else if (months[j]==="February"){
            return 28
        }else if (j<7){
            if (j%2===0){
                return 31
            }else{
                return 30
            }
        }
            if (j%2===0){
                return 30
            }else{
                return 31
            }
        
    }

	function SetTime(){
		BirthDate=`${months[setDate[1]]} ${setDate[0]}, ${setDate[3]}`
	}
	export let BirthDate
	export let datedrop
</script>

<div class="flex bg-[#374151] absolute  {$$props.class} text-[#fcfcfc] font-semibold shadow-lg rounded-xl">
	<div class="flex flex-col">
		<div class="flex divide-x">
			<div class="flex flex-col px-6 pt-5 pb-6 border-b border-gray-100">
				<div class="flex items-center justify-between">
					<button
						on:click={() => {
							setDate[3] === 1850 ? (setDate[3] = 1850) : (setDate[3] -= 1);
                            FindDay()
						}}
						class="flex items-center justify-center p-1 m-2 rounded-xl border-[2px]  border-[#fcfcfc]    hover:bg-gray-500"
					>
						<svg class="w-6 h-6 fill-[#fcfcfc] stroke-current" fill="none">
							<path
								d="M13.25 8.75L9.75 12l3.5 3.25"
								stroke-width="1.5"
								stroke-linecap="round"
								stroke-linejoin="round"
							/>
						</svg>
					</button>
					<div class="text-sm font-semibold">{setDate[3]}</div>
					<button
						on:click={() => {
							setDate[3] === 2023 ? (setDate[3] = 2023) : (setDate[3] += 1);
                            FindDay()
						}}
						class="flex items-center justify-center p-1 m-2 rounded-xl border-[2px]  border-[#fcfcfc]    hover:bg-gray-500"
					>
						<svg class="w-6 h-6 fill-[#fcfcfc]   stroke-current" fill="none">
							<path
								d="M10.75 8.75l3.5 3.25-3.5 3.25"
								stroke-width="1.5"
								stroke-linecap="round"
								stroke-linejoin="round"
							/>
						</svg>
					</button>
				</div>
				<div class="flex items-center justify-between">
					<button
						on:click={() => {
							setDate[1] === 0 ? (setDate[1] = 11) : (setDate[1] -= 1);
                            FindDay()

						}}
						class="flex items-center justify-center p-1 m-2 rounded-xl border-[2px]  border-[#fcfcfc]    hover:bg-gray-500"
					>
						<svg class="w-6 h-6 fill-[#fcfcfc] stroke-current" fill="none">
							<path
								d="M13.25 8.75L9.75 12l3.5 3.25"
								stroke-width="1.5"
								stroke-linecap="round"
								stroke-linejoin="round"
							/>
						</svg>
					</button>

					<div class="text-sm font-semibold">{months[setDate[1]]}</div>
					<button
						on:click={() => {
							setDate[1] === 11 ? (setDate[1] = 0) : (setDate[1] += 1);
                            FindDay()

						}}
						class="flex items-center justify-center p-1 m-2 rounded-xl border-[2px]  border-[#fcfcfc]    hover:bg-gray-500"
					>
						<svg class="w-6 h-6 fill-[#fcfcfc]   stroke-current" fill="none">
							<path
								d="M10.75 8.75l3.5 3.25-3.5 3.25"
								stroke-width="1.5"
								stroke-linecap="round"
								stroke-linejoin="round"
							/>
						</svg>
					</button>
				</div>
				<div class="grid grid-cols-7 text-xs text-center text-[#fcfcfc]">
					{#each day as i}
						<span class="flex items-center justify-center w-10 h-10 font-semibold rounded-lg">
							{i}
						</span>
					{/each}
					
					{#each Array.from(Array(setDate[2]+1).keys()).slice(1) as i}
						<span class="flex items-center  justify-center w-10 h-10 text-gray-700 rounded-lg" >
                        <!--  -->
                        <!-- {i} -->
                        </span>

					{/each}
					{#each Array.from(Array(NumOfDayinMonth(setDate[3],setDate[1])+1).keys()).slice(1) as i}
						<!-- content here -->
						<span on:click="{()=>{setDate[0]=i}}"
							class="flex items-center justify-center hover:border-2 border-blue-600 w-10 h-10 text-white hover:cursor-pointer transition-all ease-in duration-100 { setDate[0]===i ? " bg-blue-600" : "" } rounded-l-lg"
						>
							{i}
						</span>
					{/each}

					
				</div>
			</div>
		</div>
		<div class="flex items-center justify-between px-6 py-4">
			<div class="flex items-center space-x-2">
				<button on:click="{()=>{datedrop=false}}"
					class="px-4 py-2 text-sm rounded-lg bg-gray-500 hover:text-gray-600 focus:outline-none focus:ring-1 focus:ring-blue-600 hover:bg-gray-100"
				>
					Cancel
				</button>
				<button on:click="{()=>{SetTime() ; datedrop=false}}"
					class="px-4 py-2 text-sm text-white bg-blue-600 rounded-lg focus:outline-none focus:ring-1 focus:ring-blue-600 hover:bg-blue-700"
				>
					Set Date
				</button>
			</div>
		</div>
	</div>
</div>

<style>
	/* your styles go here */
</style>
