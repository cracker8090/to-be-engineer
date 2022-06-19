from twilio.rest import Client
 
# Your Account Sid and Auth Token from twilio.com/console
account_sid = 'AC9a6e51be2xxxxxxxxxxxxx'#注册，创建项目后获得
auth_token = 'you auth token'
client = Client(account_sid, auth_token)
call = client.messages.create(
                        to='+86133xxxxxx',#86是区号后面是你注册时用的手机
                        from_='+16162xxxxxx',#创建项目时，给你的号码
                        body='test info',#发送的短信信息
                    )


