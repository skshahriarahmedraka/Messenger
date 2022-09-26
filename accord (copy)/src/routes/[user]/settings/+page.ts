import { redirect } from '@sveltejs/kit';

export async function load({ session }: any) {
	if (!session.user.authenticated) {
		throw redirect(302, '/login');
	}

	return {};
}
