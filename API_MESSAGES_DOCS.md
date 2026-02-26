# 📚 Documentación de API - Mensajería (Messages)

Esta documentación cubre todos los endpoints disponibles para el manejo de mensajes en el sistema: enviar, consultar historial, borrar mensajes y escuchar mensajes en tiempo real vía WebSocket.

---

## 1. 🔍 Obtener Historial de Chat (GET)
Obtiene todos los mensajes intercambiados entre dos usuarios específicos, ordenados cronológicamente.

* **Ruta (ejemplo):** `GET /messages?senderId=1&receiveId=2`
* **Query Parameters:**
  * `senderId` (Obligatorio): ID del usuario que consulta (o el otro participante).
  * `receiveId` (Obligatorio): ID del otro participante en el chat.
* **Respuestas Exitosas (200 OK):**
  ```json
  {
    "messages": [
      {
        "idMessage": 45,
        "senderId": 1,
        "receiveId": 2,
        "content": "¡Hola, cómo estás?",
        "timeMessage": "2026-02-26T15:30:00Z"
      }
    ]
  }
  ```
* **Errores Comunes (400 Bad Request):**
  * `{"error": "senderId es requerido"}`
  * `{"error": "receiveId debe ser un número entero"}`

---

## 2. 📝 Enviar un Mensaje (POST)
Crea y envía un nuevo mensaje. Al hacerlo con éxito, el mensaje también será transmitido en tiempo real a la persona que lo recibe a través de su conexión WebSocket (si está conectado).

* **Ruta (ejemplo):** `POST /messages`
* **Body (JSON):**
  ```json
  {
    "senderId": 1,
    "receiveId": 2,
    "content": "¡Hola, cómo estás?",
    "timeMessage": "2026-02-26T15:30:00Z"
  }
  ```
* **Respuestas Exitosas (200 OK):**
  ```json
  {
    "data": {
      "idMessage": 46,
      "senderId": 1,
      "receiveId": 2,
      "content": "¡Hola, cómo estás?",
      "timeMessage": "2026-02-26T15:30:00Z"
    }
  }
  ```
* **Errores Comunes (400 Bad Request):**
  * `{"error": "Error 404 - Solicitud incorrecta, el cuerpo de la solicitud no es válido", "Detail": "..."}`

---

## 3. 🗑️ Eliminar un Mensaje (DELETE)
Elimina un mensaje específico de la base de datos a partir de su ID.

* **Ruta (ejemplo):** `DELETE /messages/:messageId`
* **Path Parameters:**
  * `messageId` (Obligatorio): El ID numérico único del mensaje a eliminar.
* **Respuestas Exitosas (200 OK):**
  ```json
  {
    "message": "Mensaje eliminado exitosamente"
  }
  ```
* **Errores Comunes (400 Bad Request):**
  * `{"error": "Error 400 - Solicitud incorrecta, el messageId debe ser un número entero", "Detail": "..."}`

---

## 4. ⚡ Recibir Mensajes en Tiempo Real (WebSocket)
Este endpoint permite mantener una conexión abierta (`TCP`) para que el frontend reciba los mensajes en el instante exacto en que otra persona se los envíe.

* **Ruta de conexión:** `ws://localhost:8080/ws/messages` (Asegúrate de cambiar a `wss://` si usas HTTPS).
* **Parámetros de Autenticación (Obligatorio):**
  Debes indicar al servidor quién eres para que sepa qué mensajes pasarte.
  * **Opción A (Query String):** `ws://localhost:8080/ws/messages?userId=2`
  * **Opción B (Header):** Enviar el header `X-User-ID: 2` durante el handshake.
* **Formato del Evento Recibido:**
  Cada vez que el *Usuario A* haga un `POST` de un mensaje hacia ti, recibirás a través de este websocket el siguiente JSON plano (`snake_case`):
  ```json
  {
    "id": 46,
    "sender_id": 1,
    "receiver_id": 2,
    "content": "¡Hola, cómo estás?",
    "created_at": "2026-02-26T15:30:00Z"
  }
  ```

### **Diferencias Clave entre HTTP y WebSocket:**
* **WebSocket:** Emite un JSON plano, no envuelto en ningún `"data"`. Las propiedades usan formato de base de datos (`snake_case` como `sender_id`).
* **HTTP (POST / GET):** Devuelven el objeto envuelto dentro de una propiedad superior (`"data": {}` o `"messages": []`). Usan propiedades en formato de programación (`camelCase` como `senderId`).
