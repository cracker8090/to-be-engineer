# coding:utf-8
from twilio.rest import Client  # 先导入
 
# sid和token都是在twilio网站的个人设置中看到的
account_sid ='AC4d79e18fd3a75ab644598348e510375e'
auth_token ='6893b0ce47570760803ad5db4cc210ac'
# 实例化
client = Client(account_sid, auth_token)
 
# 开始发短信
def send_msg(message):
    u'自定义短信内容message'
    msg = client.messages.create(
        to='+8615651797525',  # 要给谁发短信, 必须带区号, 中国要加上+86
        from_='+12013351008', # 你自己twilio网站申请的手机号码, 必须带上+号
        body=message  # 你的短信内容
    )
 
# 开始打电话
def call_num(number):
    u'自定义打电话的号码'
    call = client.calls.create(
        to='+86'+number,  # 要给谁打电话, 必须带区号, 中国要加上+86
        from_='+12013351008', # 你自己twilio网站申请的手机号码, 必须带上+号
        url="http://demo.twilio.com/docs/voice.xml" # 要播放的mp3
    )
 
if __name__ == '__main__':
    send_msg('伤心')