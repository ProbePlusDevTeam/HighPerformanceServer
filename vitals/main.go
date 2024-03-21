package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"sync"
	"syscall"
	"time"
	"vitals/dbinsertmodule"

	//"vitals/dbinsertmodule"
	//"vitals/tlvparser"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	wg               sync.WaitGroup
	clientOptions    = options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err = mongo.Connect(context.Background(), clientOptions)
	groupMap         = make(map[string]context.Context)
)
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var x = msg.Payload()
	fmt.Printf("Received message: %T from topic: %s,\n", x, msg.Topic())
	fmt.Println(reflect.TypeOf(x))

}
var onMessageReceived mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
	parts := strings.Split(message.Topic(), "/")
	groupId := parts[2]
	var doc bson.D
	if err := bson.UnmarshalExtJSON(message.Payload(), true, &doc); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(groupId, time.Now(), "number of go routine :", runtime.NumGoroutine())
	if _, ok := groupMap[groupId]; !ok {
		dbName := "vitals"
		duration := 5 * time.Second // Insert data every 5 seconds
		groupMap[groupId] = context.Background()
		wg.Add(1)
		go dbinsertmodule.StartInserting(mongoClient, groupMap, dbName, groupId, duration, &wg)
		fmt.Println("part is not in list ", parts[2])

	}
	dbinsertmodule.AddData(doc, groupId)

}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	var broker = "127.0.0.1"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		panic(token.Error())

	}

	if token := client.Subscribe("/vitals/#", 0, onMessageReceived); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error subscribing to topic:", token.Error()))
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

}
