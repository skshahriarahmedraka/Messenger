import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import * as jwt from 'jsonwebtoken';

export const load: PageServerLoad = async ({ cookies,locals }) => {
	// redirect user if not logged in
	if (!locals.user.Authenticated) {
		console.log(
			'🚀 ~ file: +page.server.ts ~ line 7 ~ constload:PageServerLoad= ~ locals.user',
			locals.user
		);
		throw redirect(302, '/login');
	}

	console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies);
	console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies.get('Auth1'));
	const MyCookie = cookies.get('Auth1') || '';
	const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;
	let Userdata: any;
	if (MyCookie != '') {
		const decoded:any = jwt.verify(MyCookie, JWT_Auth_KEY);
		console.log('decoded: ', decoded);
		//   let resdata
		console.log(`http://${process.env.GO_HOST}/user/${decoded.UUID}`);
		await fetch(`http://${process.env.GO_HOST}/user/${decoded.UUID}`, {
			mode: 'no-cors'
		})
			.then((res) => {
				return res.json();
			})
			.then((d) => {
				Userdata = d;
				//   console.log("🚀 ~ file: +page.server.ts ~ line 34 ~ constload:PageServerLoad= ~ resdata", resdata)
				//   console.log("🚀 ~ file: +layout.server.ts ~ line 24 ~ constload:PageServerLoad= ~ resdata", resdata)
				//   UserProData.set(resdata);
			});

		// const Userdata=GetUserData(decoded)
		// console.log("🚀 ~ file: +layout.server.ts ~ line 17 ~ Userdata", Userdata)
	}
	return {
		Userdata
	};
};
