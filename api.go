/*
驾航服务 API 接口

////////////////////////////////////////////////////////////////////
注册登录
////////////////////////////////////////////////////////////////////
注册
http://www.jiahang0313.com/api/registeruser
       
send:
{
	name:""		   名字
	phone:""	   电话
	pass:""        md5加密
	serviceID:""   可为空
}
receive:
{
	status:0         0为失败，1为成功
	reason:""       失败原因，成功时为空
}

服务商注册
http://www.jiahang0313.com/api/registerservice
       
send:
{
	name:""		   用户名
	phone:""	   电话
	pass:""        md5加密
	checkid:""     不能为空
}
receive:
{
	status:0         0为失败，1为成功
	reason:""       失败原因，成功时为空
}
登录
http://www.jiahang0313.com/api/login​
send:
{
	phone:""		电话
	pass:""         md5加密
}

receive:
{
	id:int			user对应的id，userid为0登录失败。
	type:int		0为用户，1为商户，2为管理用户

}

/////////////////////////////////////////////////////////////
主页服务：
/////////////////////////////////////////////////////////////
服务类别
http://www.jiahang0313.com/api/servicelist
send:
{
	id:int
}

receive:
{
	services:[
		[	
			{name:"", serviceid:int },	服务名称, 服务id
			{name:"", serviceid:int }	服务名称, 服务id
		],
		[
			{name:"", serviceid:int },	服务名称, 服务id
			{name:"", serviceid:int }	服务名称, 服务id
		]
	]

}

服务详情
http://www.jiahang0313.com/api/servicedetail
send:
{
	id:int
	serviceid:int
}

receive:
{
	name:""			服务名
	image:""		介绍图片url
	price:int		价格
	context:""		服务描述
}
/////////////////////////////////////////////////////////////
用户​个人信息：
/////////////////////////////////////////////////////////////

意见反馈：
    http://www.jiahang0313.com/api/user/profile
send:
{
	id:int		 	用户id
	pass:""			密码
}
receive:
{
	name:""
	phone:""
	rank:""
}

意见反馈：
    http://www.jiahang0313.com/api/user/feedback
   
send:
{
	id:int		 	用户id
	feedback:""	     限制一下字数，不超过200个字。	
}
receive:
{
	status:0         0为失败，1为成功
	reason:""        失败原因，成功时为空
}
​

购买过的服务：
 http://www.jiahang0313.com/api/user/paid
send:
{
	id:int		id
	pass:""		用户密码
}
receive:
{
	services:[
		{name:"",buydate:timestamp, money:int},	服务名, 购买时间， 金额
		{name:"",buydate:timestamp, money:int}	服务名，购买时间， 金额
	]
}


我的消息：
 http://www.jiahang0313.com/api/user/message
send:
{
	id:int		id
	pass:""		用户密码
}
receive:
{
	services:[
		{title:"",body:""},	标题, 内容
		{title:"",body:""}	标题，内容
	]
}


添加积分：
 http://www.jiahang0313.com/api/user/addrank
send:
{
	id:int			用户id
	rank:int		增加的积分
}
receive:
{
	status:0         0为失败，1为成功
	reason:""        失败原因，成功时为空
}

/////////////////////////////////////////////////////////////
主后台：
/////////////////////////////////////////////////////////////
获取用户列表：
http://www.jiahang0313.com/api/admin/userlist
send:
{
	pass:""
}
receive:
{
	users:[
		{
			name:""				用户名
			phone:""			电话
			serviceID:""		客服专员id
			regdate:""			注册时间
			services:[
				{name:"",buydate:timestamp, money:int},	服务名, 购买时间， 金额
				{name:"",buydate:timestamp, money:int}	服务名，购买时间， 金额
			]
			feedback:""
		},

		{
			name:""				用户名
			phone:""			电话
			serviceID:""		客服专员id
			regdate:""			注册时间
			services:[
				{name:"",buydate:timestamp, money:int},	服务名, 购买时间， 金额
				{name:"",buydate:timestamp, money:int}	服务名，购买时间， 金额
			]
			feedback:""
		}
	]
}
查询单个用户：

http://www.jiahang0313.com/api/admin/seachuser
send:
{
	pass:""			管理员密码
	phone:""		要查询的用户手机号，仅提供根据手机号查询
}
receive:
{
	users:[
		{
			name:""				用户名
			phone:""			电话
			serviceID:""		客服专员id
			regdate:""			注册时间
			services:[
				{name:"",buydate:timestamp, money:int},	服务名, 购买时间， 金额
				{name:"",buydate:timestamp, money:int}	服务名，购买时间， 金额
			]
			feedback:""
		}
	]
}

查询单个服务：

http://www.jiahang0313.com/api/admin/seachservice
send:
{
	pass:""		管理员密码
	type:""		要查询服务类型，返回已经购买此服务的人员
}
receive:
{
	users:[
		{
			name:""				用户名
			phone:""			电话
			serviceID:""		客服专员id
			regdate:""			注册时间
			services:[
				{name:"",buydate:timestamp, money:int},	服务名, 购买时间， 金额
				{name:"",buydate:timestamp, money:int}	服务名，购买时间， 金额
			]
			feedback:""
		}，
		{
			name:""				用户名
			phone:""			电话
			serviceID:""		客服专员id
			regdate:""			注册时间
			services:[
				{name:"",buydate:timestamp, money:int},	服务名, 购买时间， 金额
				{name:"",buydate:timestamp, money:int}	服务名，购买时间， 金额
			]
			feedback:""
		}
	]
}


/////////////////////////////////////////////////////////////
子后台：
/////////////////////////////////////////////////////////////
待服务用户
http://www.jiahang0313.com/api/admin/waitservice
send:
{
	id:int		id
	pass:""		管理员密码
}
receive:
{
	users:[
		{
			name:""				用户名
			phone:""			电话
			services:[
				{name:"",buydate:timestamp, money:int},	服务名, 购买时间， 金额
				{name:"",buydate:timestamp, money:int}	服务名，购买时间， 金额
			]
		}，
		{
			name:""				用户名
			phone:""			电话
			services:[
				{name:"",buydate:timestamp, money:int},	服务名, 购买时间， 金额
				{name:"",buydate:timestamp, money:int}	服务名，购买时间， 金额
			]
		}
	]
}

已服务用户
http://www.jiahang0313.com/api/admin/alreadyserved
send:
{
	id:int		id
	pass:""		管理员密码
}
receive:
{
	users:[
		{
			name:""				用户名
			phone:""			电话
			services:[
				{name:"",buydate:timestamp, money:int},	服务名, 购买时间， 金额
				{name:"",buydate:timestamp, money:int}	服务名，购买时间， 金额
			]
		}，
		{
			name:""				用户名
			phone:""			电话
			services:[
				{name:"",buydate:timestamp, money:int},	服务名, 购买时间， 金额
				{name:"",buydate:timestamp, money:int}	服务名，购买时间， 金额
			]
		}
	]
}

服务某个用户
http://www.jiahang0313.com/api/admin/userservice
send:
{
	id:int		id
	pass:""		管理员密码
	userid:		待服务的用户id
}
receive:
{
	status:0         0为失败，1为成功
	reason:""        失败原因，成功时为空
}

*/