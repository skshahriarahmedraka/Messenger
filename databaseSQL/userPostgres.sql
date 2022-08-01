
DROP TABLE userdata ;



CREATE TABLE userdata (
	id BIGSERIAL  NOT NULL  PRIMARY KEY  ,
	Email VARCHAR (255) NOT NULL,
	Password VARCHAR (255) NOT NULL,
	Userid VARCHAR (255) NOT NULL,
	Username VARCHAR (255) NOT NULL,
	Mobile VARCHAR (255) NOT NULL,
	birthdate DATE NOT NULL,
	Accounttype VARCHAR (255) NOT NULL
);

INSERT INTO userdata (Email, Password, Userid, Username, Mobile, birthdate, Accounttype) VALUES ( 'raka@gmail.com', 'whatsupraka','skraka' ,'sk shahriar ahmed raka' , '01838810189','1998-7-8','premium' );
INSERT INTO userdata (Email, Password, Userid, Username, Mobile, birthdate, Accounttype) VALUES ( 'ssar@gmail.com', 'whatsupssar','skssar' ,'sk shahriar ahmed ssar' , '01838810189','1998-7-8','premium' );
INSERT INTO userdata (Email, Password, Userid, Username, Mobile, birthdate, Accounttype) VALUES ( 'ahmed@gmail.com', 'whatsupahmed','skahmed' ,'sk shahriar ahmed' , '01838810189','1998-7-8','anonymus' );
INSERT INTO userdata (Email, Password, Userid, Username, Mobile, birthdate, Accounttype) VALUES ( 'shahriar@gmail.com', 'whatsupshahriar','skshahriar' ,'sk shahriar ahmed shahriar' , '01838810189','1998-7-8','premium' );
INSERT INTO userdata (Email, Password, Userid, Username, Mobile, birthdate, Accounttype) VALUES ( 'ekbal@gmail.com', 'whatsupekbal','skekbal' ,'sk shahriar ahmed ekbal' , '01838810189','1998-7-8','free' );
INSERT INTO userdata (Email, Password, Userid, Username, Mobile, birthdate, Accounttype) VALUES ( 'rakib@gmail.com', 'whatsuprakib','skrakib' ,'sk shahriar ahmed rakib' , '01838810189','1998-7-8','premium' );


SELECT * from userdata;


