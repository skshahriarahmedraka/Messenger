<script>
    
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
        background:radial-gradient(#ff0,#f90);
        height:50px;
        width:50px;
        border-radius:50%;
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

<div class="portfolio-loader {$$props.class}">
    <div class="sun"></div>
    <div class="orbit orbit1"><div class="planetX planet1"></div></div>
    <div class="orbit orbit2"><div class="planetX planet2"></div></div>
    <div class="orbit orbit3"><div class="planetX planet3"></div></div>
</div>