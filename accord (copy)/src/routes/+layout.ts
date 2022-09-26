import { redirect } from '@sveltejs/kit';

export async function load({ session, fetch }: any) {
	if (!session.user.authenticated) {
		throw redirect(302, '/login');
	}

	// FETCH USER PROFILE
	let x3 = {
		Name: "Sk Shahriar Ahmed Raka",
		Userid: 'skraka',
		ProfileImage:
			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
		CoverImage:"https://res.cloudinary.com/dqo0ssnti/image/upload/v1661613541/samples/Untitled_design_zmrybr.png"
	};

	throw new Error("@migration task: Migrate this return statement (https://github.com/sveltejs/kit/discussions/5774#discussioncomment-3292693)");
	return {
		props: {},
		stuff: {
			MyPro: x3
		}
	};
}
