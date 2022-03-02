const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {},
		// colors: {
		// 	"discord-dock": "#202225",
		// 	"discord-peoplelist": "#2f3136",
		// 	"discord-message":"#36393f",
		// },
	},

	// plugins: []
	plugins: [
		require("daisyui"),
	    require('tailwindcss-question-mark'),
		require('tailwind-scrollbar'),
		require('@tailwindcss/line-clamp'),
	],
};

module.exports = config;
