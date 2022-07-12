const config = {
	mode: 'jit',
	content: ['./src/**/*.{html,js,svelte,ts}'],
	

	theme: {
		extend: {}
	},

	plugins: [
		require('tailwindcss-question-mark'),
		require('@tailwindcss/line-clamp'),
		require('@tailwindcss/typography'),
		require('@tailwindcss/aspect-ratio'),
		require('@tailwindcss/forms'),
	]
};

module.exports = config;
