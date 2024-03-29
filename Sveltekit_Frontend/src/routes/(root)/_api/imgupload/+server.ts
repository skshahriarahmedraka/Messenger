import { json } from '@sveltejs/kit';
// import axios from 'axios';
import minio from 'minio';
import * as fs from 'fs';

import { v4 as uuidv4 } from 'uuid';

const minioClient = new minio.Client({
	endPoint: '127.0.0.1',
	port: 9000,
	useSSL: false,
	accessKey: 'admin',
	secretKey: 'miniosecret'
});

import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ( {request} )=> {
	const data = await request.json();
	// console.log('🚀 ~ file: +server.ts ~ line 45 ~ POST ~ data', data);

	const file = data['img'];
	// const extention:string = data.name.substring(file.name.lastIndexOf('.'))
	const extention:string = "."+data["name"].split('.').pop()

	fs.writeFileSync(`static/avatar`, file, 'base64');

	// Generate UUID
	const myuuid = uuidv4();

	// UPLOAD in MINIO
	const f = `static/avatar`;
	const metaData = fs.statSync(f);
	const fileName= myuuid+extention
	minioClient.fPutObject('userimg', fileName, f, metaData, function (err, etag) {
		return console.log(err, etag); // err should be null
	});

	const mydata = {
		Link: '',
		Name: fileName,
		BucketName: 'userimg'
	};
	let ResData: {
		Link: string;
		Name: string;
		BucketName: string;
	}={} as  {
		Link: string;
		Name: string;
		BucketName: string;
	}

	ResData= {
		Link: `http://${process.env.MINIO_HOST}:${process.env.MINIO_PORT}/`+`${process.env.MINIO_IMGBUCKET}`+"/"+fileName,
		Name: fileName,
		BucketName: 'userimg'
	};
	// await fetch(`http://${process.env.GO_HOST}/api/getimglink`, {
	// 	method: 'POST',
	// 	mode: 'no-cors',

	// 	headers: {
	// 		'Content-Type': 'application/json',
	// 		Accept: 'application/json'
	// 	},
	// 	body: JSON.stringify(mydata)
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		console.log('🚀 ~ file: +server.ts ~ line 99 ~ .then ~ data', data);

	// 		ResData = data;
	// 	});

	return json(ResData);
}

// import * as fs from 'fs/promises';
// export async function POST({ request }) {
// 	const data = await request.formData();
// 	const file = data.get('img') as File;
// 	const blob = data.get('img') as Blob;
// 	console.log('🚀 ~ file: +server.ts ~ line 78 ~ POST ~ file.name', file.name);
// 	console.log('🚀 ~ file: +server.ts ~ line 78 ~ POST ~ file.size', file.size);

// 	// await fs.writeFile(`static/avatar.png/${file.name}`, file.stream());

// 	///////////////////
// 	// const data = await request.json();
// 	// // console.log('🚀 ~ file: +server.ts ~ line 45 ~ POST ~ data', data);

// 	// const file = data['img'];

// 	// fs.writeFileSync(`static/avatar.png`, file, 'base64');

// 	// Generate UUID
// 	let myuuid = uuidv4();

// 	// UPLOAD in MINIO
// 	// const f = `static/avatar.png`;
// 	// const metaData = fs.statSync(f);

// 	// interface FileWithPath extends File {
// 	// 	path: string
// 	//   }

// 	//   const path = (file as FileWithPath).path

// 	// const metaData = {
// 	// 	'Content-Type': file.type,
// 	//   }

// 	let path= URL.createObjectURL(blob)


// 	fs.stat(path, (error, stats) => {
// 		if (error) {
// 			console.log('fs stat error : ', error);
// 		} else {
// 			console.log('fs stat successfull : ', stats);

// 			minioClient.fPutObject(
// 				'myimg',
// 				file.name,
// 				path,
// 				stats,
// 				function (err, etag) {
// 					return console.log('fPutObject error :  ', err, etag); // err should be null
// 				}
// 			);
// 		}
// 	});

// 	const mydata = {
// 		Link: '',
// 		Name: file.name,
// 		BucketName: 'myimg'
// 	};
// 	let ResData: {
// 		Link: string;
// 		Name: string;
// 		BucketName: string;
// 	};

// 	await fetch('http://localhost:8001/imglink', {
// 		method: 'POST',
// 		mode: 'no-cors',

// 		headers: {
// 			'Content-Type': 'application/json',
// 			Accept: 'application/json'
// 		},
// 		body: JSON.stringify(mydata)
// 	})
// 		.then((res) => {
// 			return res.json();
// 		})
// 		.then((data) => {
// 			console.log('🚀 ~ file: +server.ts ~ line 99 ~ .then ~ data', data);

// 			ResData = data;
// 		});

// 	return json(ResData);
// }
