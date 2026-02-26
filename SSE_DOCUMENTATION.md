# WebSockets Mensajería en Tiempo Real

## Descripción
Se implementó **WebSockets** para notificaciones de mensajes en tiempo real. Cuando un usuario envía un mensaje, el receptor recibe una notificación instantánea a través de una conexión bidireccional continua de baja latencia.

## Componentes Implementados

### 1. **Broadcaster** (`Src/Core/sse/broadcaster.go`)
Servicio central que gestiona las conexiones WebSocket. (Mantuvo el nombre de paquete `sse` temporalmente por compatibilidad estructural, pero usa WebSockets por debajo).
- `Subscribe(userID int, conn *websocket.Conn)` - Registra una nueva conexión WebSocket para un usuario.
- `Unsubscribe(userID int, conn *websocket.Conn)` - Elimina una conexión WebSocket cuando el usuario se desconecta.
- `Broadcast(userID, message)` - Escribe un evento JSON directamente en el socket del usuario.
- `BroadcastToMultiple(userIDs[], message)` - Envía a múltiples usuarios.

### 2. **Controlador WebSocket** (`Src/Endpoint/Message/Infrestructure/Controller/SubscribeMessage_controller.go`)
Endpoint `GET /api/v1/message/subscribe`
- Espera header HTTP `X-User-ID` o parámetro de query `?userId=X` con el ID del usuario.
- Promueve (Upgrade) la conexión HTTP tradicional a una conexión WebSocket completa usando `gorilla/websocket`.
- Mantiene la conexión abierta en un loop de lectura infinito escuchando eventos de desconexión del cliente.

### 3. **SendMessage UseCase Actualizado**
Emite eventos JSON al WebSocket cuando se envía un mensaje.
- Llamadas a `broadcaster.Broadcast(receiverID, event)`

### 4. **Wire DI**
- `InitializeBroadcaster()` - Singleton global.
- `ProvideSendMessageUseCase()` - Inyecta broadcaster en UseCase.

---

## Uso de la API

### Cliente 1: Se conecta al WebSocket (para recibir mensajes)

Dependiendo de tu cliente (App Móvil o Web), la conexión se hace mediante el protocolo `ws://`.

Ejemplo de conexión de red cruda (o utilizando un cliente WS como Postman):
```
ws://localhost:8080/api/v1/message/subscribe?userId=2
```

**Respuesta recibida en el socket (formato JSON crudo):**
```json
{"id":5,"sender_id":1,"receiver_id":2,"content":"Hola!","created_at":"2025-02-22T10:30:45Z"}
{"id":6,"sender_id":1,"receiver_id":2,"content":"¿Cómo estás?","created_at":"2025-02-22T10:31:00Z"}
```

### Cliente 2: Envía un mensaje

```bash
curl -X POST http://localhost:8080/api/v1/message/send \
  -H "Content-Type: application/json" \
  -d '{
    "senderId": 1,
    "receiverId": 2,
    "content": "Hola!"
  }'
```

**Respuesta HTTP Clásica:**
```json
{
  "data": {
    "idMessage": 5,
    "senderId": 1,
    "receiveId": 2,
    "content": "Hola!",
    "timeMessage": 1708591845000
  }
}
```

► **Cliente 1 recibe la notificación automáticamente empujada a su WebSocket abierto de manera instantánea**

---

## Ejemplo JavaScript / Frontend (WebSockets)

```javascript
const userId = 2; // ID del usuario actual

// 1. Conectar al WebSocket (nota el protocolo ws:// en lugar de http://)
const socket = new WebSocket(`ws://localhost:8080/api/v1/message/subscribe?userId=${userId}`);

// 2. Escuchar conexión exitosa
socket.onopen = function(e) {
  console.log("[open] Conexión establecida con el servidor WebSocket");
};

// 3. Escuchar mensajes entrantes
socket.onmessage = function(event) {
  // Los datos llegan como un string JSON
  const msg = JSON.parse(event.data);
  console.log(`[message] Nuevo mensaje de ${msg.sender_id}: ${msg.content}`);
  
  // Mostrar en la UI
  addMessageToChat(msg);
};

// 4. Manejar desconexiones o cierres
socket.onclose = function(event) {
  if (event.wasClean) {
    console.log(`[close] Conexión cerrada limpiamente, código=${event.code} motivo=${event.reason}`);
  } else {
    // ej. el servidor se cayó o la red se perdió
    console.log('[close] Conexión interrumpida. Sugerencia: Programar reconexión automática.');
  }
};

// 5. Manejar errores
socket.onerror = function(error) {
  console.error(`[error]`, error);
};

// Para ENVIAR mensajes, actualmente sigues usando tu API REST POST normal, 
// pero en un futuro podrías enviar mensajes escribiendo al socket mediante:
// socket.send(JSON.stringify({ content: "Hola", receiverId: 1 }));
```

---

## Flujo Dúplex en Tiempo Real (Actual)

```
┌─────────────────────────────────────────────────────────────┐
│                      CLIENTE A (ID: 1)                       │
│                                                               │
│  POST /send {"content": "Hola!"}     [ENVÍA MENSAJE]         │
│                          │                                    │
│                          ▼                                    │
│                    SendMessageUseCase                        │
│                          │                                    │
│                          ├─► Guardar en BD                   │
│                          │                                    │
│                          └─► broadcaster.Broadcast(ID: 2)    │
│                                       │                       │
└──────────────────────────────┬────────┼──────────────────────┘
                               │        │
                               │        ▼
┌──────────────────────────────┤  ┌─────────────────────────┐
│   CLIENTE B (ID: 2)          │  │     WEBSOCKET HUB       │
│                              │  │  (mantiene conexiones)  │
│ WS UPGRADE ────────►[ABIERTO]│  └─────────────────────────┘
│       │                      │         │
│       │                      └────────►│
│       │                                │
│ {"content": "Hola!"}◄───────────[JSON EMITIDO]
│       │                               │
│  [NOTIFICACIÓN MOSTRADA EN UI]        │
│                                       │
└───────────────────────────────────────┘
```

---

## Características de WebSockets respecto a SSE

✅ **Full-Duplex** - Bidireccional. Actualmente solo el servidor habla, pero con WS tienes la infraestructura lista para que los clientes envíen mensajes directamente por el socket sin hacer peticiones HTTP POST (latencia mínima absoluta).
✅ **Estándar de Chat** - Usado por WhatsApp, Discord, Slack.
✅ **Manejo de estados nativo** - Facilita rastrear si un usuario está "escribiendo..." sin sobrecargar la red.
✅ **Baja Latencia** - No hay sobrecarga de headers de HTTP cada vez que se quiere emitir un mensaje.

## Seguridad (Opcional)

Puedes proteger la inicialización del WebSocket comprobando tokens de acceso justo en el momento donde se evalúa el header (antes de llamar a la función de upgrade `upgrader.Upgrade()`).

## Futuras mejoras
*   **Aceptar lectura del socket**: Puedes crear un caso de uso donde si el servidor lee del socket `msg := ws.ReadMessage()`, invoque al `SendMessageUseCase` automáticamente, y así eliminas por completo la necesidad del endpoint REST POST `/send`.
