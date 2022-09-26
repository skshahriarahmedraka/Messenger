throw new Error("@migration task: Migrate the load function input (https://github.com/sveltejs/kit/discussions/5774#discussioncomment-3292693)");
export async function load({ stuff, session }: any) {
	console.log('🚀 ~ file: index.svelte ~ line 3 ~ load ~ session.user', session.user);

	return {
		MyPro: stuff.MyPro
	};
}
