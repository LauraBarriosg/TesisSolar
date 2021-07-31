package mqtt

import (
	"github.com/LauraBarriosg/PRUEBA/models"
    mqtt "github.com/eclipse/paho.mqtt.golang"
	"fmt"
	"time"
	//"github.com/LauraBarriosg/PRUEBA/libs"
)

// Varibles para el manejador de Mensajes
var X int
var Valor_Mandado mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    X++
    fmt.Printf("Mensajess %d\n", X)
    fmt.Printf("Mensaje: %s desde el topic: %s\n", msg.Payload(), msg.Topic())
	Json := fmt.Sprintf("%s",msg.Payload())
	models.ExtraerDatos(Json)
}
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    fmt.Println("Conectado")
}
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Printf("Connect lost: %v", err)
}

// Coneccion y cliente MQTT
func Mqtt(){
	//var broker = "localhost"
	var broker = "mosquitto"
    var port = 1883
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
    opts.SetDefaultPublishHandler(Valor_Mandado)
    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler
    client := mqtt.NewClient(opts)
    
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
	defer client.Disconnect(0)
	sub(client)
	time.Sleep(time.Second * 10000)
}

// Suscriptor mqtt
func sub(client mqtt.Client) {
    topic := "Si/#"
    token := client.Subscribe(topic, 1, nil)
    token.Wait()
}
