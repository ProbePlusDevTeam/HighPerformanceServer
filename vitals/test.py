import paho.mqtt.client as paho
import time,random,json
broker="127.0.0.1"
port=1883
def on_publish(client,userdata,result):             #create function for callback
    print("data published \n")
    pass
client1= paho.Client("control1")                           #create client object
client1.on_publish = on_publish                          #assign function to callback
client1.connect(broker,port)
client1.loop_start()
group_list=["5","6","7","8","9"]
while True:
    time.sleep(.01)
    #print(i)       
    #/vitals/group{str(i)[-1]}/{i}                    #establish connection
    
    ret= client1.publish(f"/vitals/group{random.choice(group_list)}",json.dumps({"time":time.time(),"cpu_pct":"{:.2f}".format(random.random()*100),"battery_pct":"{:.2f}".format(random.random()*100)}))
