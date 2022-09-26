throw new Error("@migration task: Migrate the load function input (https://github.com/sveltejs/kit/discussions/5774#discussioncomment-3292693)");
export async function load({ stuff, fetch }) {
	let FriendData = {
		Name: 'Portgas D. Ace',
		ProfileUrl: 'ace',
		ProfileImage:
			'https://res.cloudinary.com/dqo0ssnti/image/upload/v1642489744/samples/1007550_t0uscy.jpg'
	};

	let FriendMsg = [
		{
			Conversation: {
				id: 123,
				members: ['skraka', 'ace']
			}
		},
		[
			{
				id: 0,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',
				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 1,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',
				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 2,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 3,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 4,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 5,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 6,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 7,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 8,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 9,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 10,
				time: '20 jun22',
				writer: 'skraka',
				writerName: 'Sk Shahriar Ahmed Raka',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			// {id:10.1, time:"20 jun22", writer: 'Sk Shahriar Ahmed Raka', message: "hello brother , what's you doing" },
			{
				id: 11,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 12,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 13,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message:
					"hello brother , what's you doingSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed RakaSk Shahriar Ahmed Raka",

				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 14,
				time: '20 jun22',
				writer: 'skraka',
				writerName: 'Sk Shahriar Ahmed Raka',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 15,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 16,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },

				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 17,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 18,
				time: '20 jun22',
				writer: 'skraka',
				writerName: 'Sk Shahriar Ahmed Raka',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 19,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 1, skraka: 2 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 20,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 7, skraka: 1 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 21,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 5, skraka: 4 },
				numberOfReact: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 22,
				time: '20 jun22',
				writer: 'skraka',
				writerName: 'Sk Shahriar Ahmed Raka',

				message:
					"hello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doinghello brother , what's you doing",
				react: { ace: 2, skraka: 9, fr: 3, dfz: 2, sg: 3 },
				numberOfReact: [1, 1, 2, 3, 4, 5, 6, 7, 8, 9]
			},
			{
				id: 23,
				time: '20 jun22',
				writer: 'ace',
				writerName: 'Portgas D. Ace',

				message: "hello brother , what's you doing",
				react: { ace: 0, skraka: 1, sk: 2, f: 3, g: 4, p: 5, o: 6, i: 7, u: 8, w: 9 },
				numberOfReact: [1, 1, 2, 3, 4, 5, 67435, 7, 8, 99]
			}
		]
	];

	let FrinedStat = {};

	return {
		FriendData: FriendData,
		FriendMsg: FriendMsg,
		MyPro: stuff.MyPro
	};
}
