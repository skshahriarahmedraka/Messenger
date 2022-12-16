import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import * as jwt from 'jsonwebtoken';

export const load: PageServerLoad = async ({ cookies, locals }) => {
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
	let Userdata:	{
		UUID: string;
		UserID: string;

		Email: string;
		Password: string;
		UserName: string;
		Mobile: string;
		BirthDate: string;

		ProfileImg: string;
		BannerImg: string;
		Accounttype: string;
		Coin: number;
		TransactionHistory: string[];
		ContactAdminMsg: string[];
	
		UserBio: string;
		FriendList: { UserID: string; CollectionID: string }[];
		GroupList: { GroupID: string; CollectionID: string }[];
		City: string;
		Address: string;
		Country: string;
		ZipCode: string;
	} = {
		UUID: '' as string,
		UserID: '' as string,
	
		Email: '' as string,
		Password: '' as string,
		UserName: '' as string,
		Mobile: '' as string,
		BirthDate: '' as string,
		UserBio: '' as string,
	
		ProfileImg: '' as string,
		BannerImg: '' as string,
	
		Coin: 0 as number,
		Accounttype: '' as string,
		TransactionHistory: [] as string[],
	
		City: '' as string,
		Address: '' as string,
		ZipCode: '' as string,
		Country: '' as string,
	
		FriendList: [] as { UserID: string; CollectionID: string }[],
		GroupList: [] as { GroupID: string; CollectionID: string }[],
		ContactAdminMsg: [] as string[]
	};
	if (MyCookie != '') {
		interface tokeninterface {
			UserName: string;
			Email: string;
			UserID: string;
			UUID: string;
			exp: number;
		}
		const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
		console.log('decoded: ', decoded);
		//   let resdata
		console.log(`http://${process.env.GO_HOST}/user/${(decoded as tokeninterface).UUID}`);
		await fetch(`http://${process.env.GO_HOST}/user/${(decoded as tokeninterface).UUID}`, {
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


	
	console.log("🚀 ~ file: +page.server.ts ~ line 99 ~ constload:PageServerLoad= ~ Userdata", Userdata)
	return {
		Userdata
	};
};
