import { json } from "@sveltejs/kit";



import type { RequestHandler } from './$types';

export const GET: RequestHandler = async (  ) =>{
	// const data = await request.json();
    let resData
    await fetch(`http://${process.env.GO_HOST}/admin/moneymanagement`, {
			// credentials: 'same-origin',
			method: 'GET',
			mode: 'no-cors',
			headers: new Headers({ 'Content-Type': 'application/json', Accept: 'application/json' }),

			// body: JSON.stringify(data)
		})
			.then((response) => response.json())
			.then((data) => {
                resData=data
				console.log('🚀 ~ file: +page.svelte ~ line 312 ~ .then ~ REsponse', data);
		});

    return json(resData)
}

