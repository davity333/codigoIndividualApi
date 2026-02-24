# SSE Mensajería en Tiempo Real

## Descripción
Se implementó **Server-Sent Events (SSE)** para notificaciones de mensajes en tiempo real. Cuando un usuario envía un mensaje, el receptor recibe una notificación instantánea sin necesidad de hacer polling.

## Componentes Implementados

### 1. **Broadcaster** (`Src/Core/sse/broadcaster.go`)
Servicio central que gestiona las conexiones SSE
- `Subscribe(userID int)` - Usuario se conecta al stream
- `Unsubscribe(userID int, ch)` - Usuario se desconecta
- `Broadcast(userID, message)` - Envía evento a un usuario
- `BroadcastToMultiple(userIDs[], message)` - Envía a múltiples usuarios

### 2. **Controlador SSE** (`Src/Endpoint/Message/Infrestructure/Controller/SubscribeMessage_controller.go`)
Endpoint `GET /api/v1/message/subscribe`
- Espera header `X-User-ID` con el ID del usuario
- Mantiene conexión abierta para recibir eventos
- Cierra cuando cliente desconecta o timeout

### 3. **SendMessage UseCase Actualizado**
Ahora emite eventos SSE cuando se envía un mensaje
- Llamadas a `broadcaster.Broadcast(receiverID, event)`

### 4. **Wire DI**
- `InitializeBroadcaster()` - Singleton global
- `ProvideSendMessageUseCase()` - Inyecta broadcaster en UseCase

---

## Uso de la API

### Cliente 1: Se conecta a SSE (para recibir mensajes)

```bash
curl -N -H "X-User-ID: 2" "http://localhost:8080/api/v1/message/subscribe"
```

**Respuesta en streaming:**
```
event: message
data: {"id":5,"sender_id":1,"receiver_id":2,"content":"Hola!","created_at":"2025-02-22T10:30:45Z"}

event: message
data: {"id":6,"sender_id":1,"receiver_id":2,"content":"¿Cómo estás?","created_at":"2025-02-22T10:31:00Z"}
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

**Respuesta:**
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

► **Cliente 1 recibe la notificación automáticamente en el stream SSE**

---

## Ejemplo JavaScript (Frontend)

```javascript
// Usuario se conecta al stream de mensajes
const eventSource = new EventSource(
  '/api/v1/message/subscribe',
  { headers: { 'X-User-ID': '2' } }
);

// Escucha mensajes entrantes
eventSource.addEventListener('message', (event) => {
  const msg = JSON.parse(event.data);
  console.log(`Nuevo mensaje de ${msg.sender_id}: ${msg.content}`);
  
  // Mostrar en UI
  addMessageToChat(msg);
});

// Manejar errores de conexión
eventSource.onerror = () => {
  console.log('Desconectado. Reconectando...');
  eventSource.close();
};

// Enviar mensaje
async function sendMessage(content, receiverId) {
  const response = await fetch('/api/v1/message/send', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      senderId: userId,
      receiverId: receiverId,
      content: content
    })
  });
  return response.json();
}
```

---

## Flujo en Tiempo Real

```
┌─────────────────────────────────────────────────────────────┐
│                      CLIENTE A (ID: 1)                       │
│                                                               │
│  POST /send {"content": "Hola!"}     [ENVÍA MENSAJE]         │
│                          │                                    │
│                          ▼                                    │
│                    SendMessageUseCase                        │
│                          │                                    │
│                          ├─► Guardarin BD                    │
│                          │                                    │
│                          └─► broadcaster.Broadcast(ID: 2)    │
│                                       │                       │
└──────────────────────────────┬────────┼──────────────────────┘
                               │        │
                               │        ▼
┌──────────────────────────────┤  ┌─────────────────────────┐
│   CLIENTE B (ID: 2)          │  │    SSE BROADCASTER      │
│                              │  │  (mantiene conexiones)  │
│ GET /subscribe ───►[ABIERTO] │  └─────────────────────────┘
│       │                      │         │
│       │                      └────────►│
│       │                                │
│   event: message◄───────────────[EVENTO EMITIDO]
│   data: {...}                         │
│       │                               │
│  [NOTIFICACIÓN MOSTRADA EN UI]        │
│                                       │
└───────────────────────────────────────┘
```

---

## Características

✅ **Real-time** - Notificaciones instantáneas sin polling
✅ **Eficiente** - Usa HTTP/1.1 Keep-Alive
✅ **Escalable** - Broadcaster gestiona múltiples conexiones
✅ **Integrado** - Se aplica automáticamente con Wire DI
✅ **Fallback** - Compatible con navegadores antiguos (con polyfill)
✅ **Seguro** - Requiere header X-User-ID (puede integrar JWT)

---

## Seguridad (Opcional)

Aplicar JWT Middleware al endpoint SSE:

```go
messageGroup.GET("/subscribe", 
  middleware.JWTMiddleware(),  // ◄── Agregar esta línea
  SubscribeMessage.Subscribe)
```

---

## Monitoreo

Ver usuarios conectados:
```go
connectedUsers := broadcaster.GetConnectedUsers()
```

---

## Próximos Pasos (Opcional)

1. **Agregar autenticación JWT** al endpoint `/subscribe`
2. **Persistir estado** - Guardar mensajes no entregados
3. **Historial de chats** - GET `/api/v1/message/getAll`
4. **Typing indicadores** - Evento diferente: `event: typing`
5. **Read receipts** - Confirmar cuando mensaje fue leído
