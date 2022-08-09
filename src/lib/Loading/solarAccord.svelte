<script>
    import { fade, fly } from 'svelte/transition';
    
</script>

<style>
   
.portfolio-loader{
        background-color:#2c3e50;
        height:100%;
        width:100%;
        margin-bottom:200px;
        position: fixed;
    }
    .sun{
        /* background:radial-gradient(#ff0,#f90); */
        height:50px;
        width:50px;
        /* border-radius:50%; */
        position:absolute;
        left:0;
        right:0;
        top:0;
        bottom:0;
        margin:auto;
    }
    .planetX{
        position:absolute;
        z-index:100;
        border-radius:50%;
    }
    .planet1{
        left:20px;
        height:13px;
        width:13px;
        background-color:#ff8;
    }
    .planet2{
        left:23px;
        height:20px;
        width:20px;
        background:linear-gradient(#00ff00,#09f,#09f);
        -webkit-animation: rotation 1s infinite linear;
        animation: rotation 1s infinite linear;
    }
    .planet3{
        left:49px;
        height:17px;
        width:17px;
        background:radial-gradient(#ff9900,#ff4400);
    }
    .orbit{
        background:transparent;
        border-radius:50%;
        border:1px solid #fff;
        position:absolute;
        left:0;
        right:0;
        top:0;
        bottom:0;
        margin:auto;
    }
    .orbit1{
        height:100px;
        width:100px;
        -webkit-animation: rotation 2s infinite linear;
        -moz-animation: rotation 2s infinite linear;
        animation: rotation 2s infinite linear;
    }
    .orbit2{
        height:150px;
        width:150px;
        -webkit-animation: rotation 3s infinite linear;
        -moz-animation: rotation 3s infinite linear;
        animation: rotation 3s infinite linear;


    }
    .orbit3{
        height:200px;
        width:200px;
        -moz-animation: rotation 6s infinite linear;
        -webkit-animation: rotation 6s infinite linear;
        animation: rotation 6s infinite linear;
    }

    @-webkit-keyframes rotation {
		from {
				-webkit-transform: rotate(0deg);
		}
		to {
				-webkit-transform: rotate(359deg);
		}
    }
    @keyframes rotation {
		from {
				transform: rotate(0deg);
		}
		to {
				transform: rotate(359deg);
		}
    }
    @-moz-keyframes rotation {
		from {
				-moz-transform: rotate(0deg);
		}
		to {
				-moz-transform: rotate(359deg);
		}
    }
</style>

<div out:fade class="portfolio-loader {$$props.class}">
    <svg class="h-16 w-16 sun" width="50" height="50" viewBox="0 0 96 96" xmlns="http://www.w3.org/2000/svg">
        <defs>
            <linearGradient id="e399c19f-b68f-429d-b176-18c2117ff73c" x1="-1032.172" x2="-1059.213" y1="145.312" y2="65.426" gradientTransform="matrix(1 0 0 -1 1075 158)" gradientUnits="userSpaceOnUse">
                <stop offset="0" stop-color="#114a8b"/>
                <stop offset="1" stop-color="#0669bc"/>
            </linearGradient>
            <linearGradient id="ac2a6fc2-ca48-4327-9a3c-d4dcc3256e15" x1="-1023.725" x2="-1029.98" y1="108.083" y2="105.968" gradientTransform="matrix(1 0 0 -1 1075 158)" gradientUnits="userSpaceOnUse">
                <stop offset="0" stop-opacity=".3"/>
                <stop offset=".071" stop-opacity=".2"/>
                <stop offset=".321" stop-opacity=".1"/>
                <stop offset=".623" stop-opacity=".05"/>
                <stop offset="1" stop-opacity="0"/>
            </linearGradient>
            <linearGradient id="a7fee970-a784-4bb1-af8d-63d18e5f7db9" x1="-1027.165" x2="-997.482" y1="147.642" y2="68.561" gradientTransform="matrix(1 0 0 -1 1075 158)" gradientUnits="userSpaceOnUse">
                <stop offset="0" stop-color="#3ccbf4"/>
                <stop offset="1" stop-color="#2892df"/>
            </linearGradient>
        </defs>
        <path fill="url(#e399c19f-b68f-429d-b176-18c2117ff73c)" d="M33.338 6.544h26.038l-27.03 80.087a4.152 4.152 0 0 1-3.933 2.824H8.149a4.145 4.145 0 0 1-3.928-5.47L29.404 9.368a4.152 4.152 0 0 1 3.934-2.825z"/>
        <path fill="#0078d4" d="M71.175 60.261h-41.29a1.911 1.911 0 0 0-1.305 3.309l26.532 24.764a4.171 4.171 0 0 0 2.846 1.121h23.38z"/>
        <path fill="url(#ac2a6fc2-ca48-4327-9a3c-d4dcc3256e15)" d="M33.338 6.544a4.118 4.118 0 0 0-3.943 2.879L4.252 83.917a4.14 4.14 0 0 0 3.908 5.538h20.787a4.443 4.443 0 0 0 3.41-2.9l5.014-14.777 17.91 16.705a4.237 4.237 0 0 0 2.666.972H81.24L71.024 60.261l-29.781.007L59.47 6.544z"/>
        <path fill="url(#a7fee970-a784-4bb1-af8d-63d18e5f7db9)" d="M66.595 9.364a4.145 4.145 0 0 0-3.928-2.82H33.648a4.146 4.146 0 0 1 3.928 2.82l25.184 74.62a4.146 4.146 0 0 1-3.928 5.472h29.02a4.146 4.146 0 0 0 3.927-5.472z"/>
    </svg>
    <div class="orbit orbit1"><div class="planetX planet1"></div></div>
    <div class="orbit orbit2"><div class="planetX planet2"></div></div>
    <div class="orbit orbit3"><div class="planetX planet3"></div></div>
  
</div>