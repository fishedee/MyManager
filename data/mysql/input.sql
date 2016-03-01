#创建数据库
drop database if exists FishMoney;
create database FishMoney;
use FishMoney;

#创建用户表
create table t_user(
	userId integer not null auto_increment,
	name char(32) not null,
	password char(48) not null,
	type integer not null,
	createTime timestamp not null default CURRENT_TIMESTAMP,
	modifyTime timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP, 
	primary key( userId )
)engine=innodb default charset=utf8 auto_increment = 10001;

alter table t_user add index nameIndex(name,password);

#创建类目表
create table t_category(
	categoryId integer not null auto_increment,	
	userId integer not null,
	name char(32) not null,
	remark varchar(128) not null,
	createTime timestamp not null default CURRENT_TIMESTAMP,
	modifyTime timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP, 
	primary key( categoryId )
)engine=innodb default charset=utf8 auto_increment = 10001;

alter table t_category add index userIdIndex(userId);

#创建银行卡表
create table t_card(
	cardId integer not null auto_increment,	
	userId integer not null,
	name char(32) not null,
	bank char(32) not null,
	card varchar(32) not null,
	money integer not null,
	remark varchar(128) not null,
	createTime timestamp not null default CURRENT_TIMESTAMP,
	modifyTime timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP, 
	primary key( cardId )
)engine=innodb default charset=utf8 auto_increment = 10001;

alter table t_card add index userIdIndex(userId);

#创建记账表
create table t_account(
	accountId integer not null auto_increment,
	userId integer not null,
	name char(32) not null,
	money integer not null,
	remark varchar(128) not null,
	categoryId integer not null,
	cardId integer not null,
	type integer not null,
	createTime timestamp not null default CURRENT_TIMESTAMP,
	modifyTime timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP, 
	primary key( accountId )
)engine=innodb default charset=utf8 auto_increment = 10001;

alter table t_account add index userIdIndex(userId);

#创建博客同步任务表
create table t_blog_sync(
	blogSyncId integer not null auto_increment,
	userId integer not null,
	accessToken varchar(1024) not null,
	gitUrl varchar(1024) not null,
	syncType integer not null,
	state integer not null,
	stateMessage varchar(10240) not null,
	createTime timestamp not null default CURRENT_TIMESTAMP,
	modifyTime timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP, 
	primary key( blogSyncId )
)engine=innodb default charset=utf8 auto_increment = 10001;

alter table t_blog_sync add index stateIndex(userId,state);
alter table t_blog_sync add index syncTypeIndex(userId,syncType);

#创建博客自动同步表
create table t_blog_sync_auto(
	blogSyncAutoId integer not null auto_increment,
	userId integer not null,
	accessToken varchar(1024) not null,
	gitUrl varchar(1024) not null,
	createTime timestamp not null default CURRENT_TIMESTAMP,
	modifyTime timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP, 
	primary key( blogSyncAutoId )
)engine=innodb default charset=utf8 auto_increment = 10001;

alter table t_blog_sync_auto add index userIdIndex(userId);

#创建session表
CREATE TABLE IF NOT EXISTS `ci_sessions` (
    `id` varchar(40) NOT NULL,
    `ip_address` varchar(45) NOT NULL,
    `timestamp` int(10) unsigned DEFAULT 0 NOT NULL,
    `data` blob NOT NULL,
    PRIMARY KEY (id),
    KEY `ci_sessions_timestamp` (`timestamp`)
);

#创建初始数据
insert into t_user(userId,name,password,type) values
(10001,"fish",SHA1("123"),1);

insert into t_category(categoryId,userId,name,remark) values
(10001,10001,"日常收支",''),
(10002,10001,"衣着服装",''),
(10003,10001,"理财投资",''),
(10004,10001,"薪酬工资",'');

insert into t_card(cardId,userId,name,bank,card,money,remark) values
(10001,10001,'工资卡',"农业银行卡",'',0,''),
(10002,10001,'消费卡',"工商银行卡",'',0,''),
(10003,10001,'理财卡',"工商银行卡",'',0,'');

insert into t_account(accountId,userId,name,money,remark,categoryId,cardId,type,createTime,modifyTime) values
(10001,10001,"日常支出",100,'',10001,10002,1,now(),now()),
(10002,10001,"日常收入",100,'',10001,10002,2,'2014-11-10 12:0:0','2014-11-10 12:0:0'),
(10003,10001,"日常收入",100,'',10001,10002,3,'2014-10-10 12:0:0','2014-10-10 12:0:0');

#显示一下所有数据
select * from t_user;
select * from t_category;
select * from t_card;
select * from t_account;
