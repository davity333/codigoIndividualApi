# 📚 Documentación de API - Módulo de Chat

Esta documentación cubre todos los endpoints disponibles para el funcionamiento del chat en la aplicación. Esto incluye el manejo de **Contactos**, el envío y lectura de **Mensajes** por HTTP, y la recepción en tiempo real vía **WebSocket**.

---

# 🫂 SECCIÓN 1: CONTACTOS (Contacts)

Esta sección gestiona la libreta de contactos de un usuario.

## 1.1 🔍 Buscar usuario por nombre (GET)
Busca a un usuario en el sistema por su nombre de usuario (para poder agregarlo después).
* **Ruta (ejemplo):** `GET /contacts/search/:username`
* **Path Parameters:**
  * `username` (Obligatorio): El string del nombre de usuario a buscar.
* **Respuestas Exitosas (200 OK):**
  ```json
  {
    "message": "user found",
    "data": { ...datos del usuario... }
  }
  ```
* **Errores Comunes:**
  * `400 Bad Request`: Nombre de usuario vacío.
  * `404 Not Found`: Usuario no encontrado.

## 1.2 ➕ Agregar Contacto (POST)
Agrega un usuario a tu lista de contactos.
* **Ruta (ejemplo):** `POST /contacts`
* **Body (JSON):**
  ```json
  {
    "userId": 1,
    "contactId": 2
  }
  ```
* **Respuestas Exitosas (201 Created):**
  ```json
  {
    "message": "Contacto agregado correctamente"
  }
  ```

## 1.3 📋 Obtener Lista de Contactos (GET)
Devuelve todos los contactos que un usuario tiene agregados.
* **Ruta (ejemplo):** `GET /contacts/:userId`
* **Path Parameters:**
  * `userId` (Obligatorio): El ID del dueño de la lista de contactos.
* **Respuestas Exitosas (200 OK):**
  ```json
  {
    "message": "Contactos obtenidos correctamente",
    "contacts": [
      {
        "id": 2,
        "name": "Juan Perez",
        "...": "..."
      }
    ]
  }
  ```
*(Nota: Si no hay contactos, la lista `contacts` viene vacía `[]` y el mensaje cambia a "No tienes ningún contacto agregado")*.

## 1.4 ❌ Eliminar Contacto (DELETE)
Elimina a una persona de tu lista de contactos.
* **Ruta (ejemplo):** `DELETE /contacts/:userId/:contactId`
* **Path Parameters:**
  * `userId`: Tu ID de usuario.
  * `contactId`: El ID del contacto que deseas eliminar.
* **Respuestas Exitosas (200 OK):**
  ```json
  {
    "message": "Contacto eliminado correctamente"
  }
  ```

---

# 💬 SECCIÓN 2: MENSAJERÍA HTTP (Messages)

Estos endpoints te permiten enviar, cargar historial y borrar mensajes individuales de forma estática (tipo REST).

## 2.1 📥 Cargar Historial de un Chat (GET)
Obtiene todos los mensajes intercambiados entre dos usuarios específicos (tú y tu contacto), ordenados cronológicamente.
* **Ruta (ejemplo):** `GET /messages?senderId=1&receiveId=2`
* **Query Parameters:**
  * `senderId` (Obligatorio): ID de uno de los participantes.
  * `receiveId` (Obligatorio): ID del otro participante.
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

## 2.2 📤 Enviar un Mensaje (POST)
Crea y envía un nuevo mensaje. Al hacerlo, el backend también empujará este mensaje por WebSocket a la persona receptora (si está en línea).
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
      ...
    }
  }
  ```

## 2.3 🗑️ Eliminar un Mensaje (DELETE)
Elimina un mensaje específico del historial para ambos usuarios (se borra de DB).
* **Ruta (ejemplo):** `DELETE /messages/:messageId`
* **Path Parameters:**
  * `messageId` (Obligatorio): El ID numérico único del mensaje a borrar.
* **Respuestas Exitosas (200 OK):**
  ```json
  {
    "message": "Mensaje eliminado exitosamente"
  }
  ```

---

# ⚡ SECCIÓN 3: TIEMPO REAL (WebSocket)

El endpoint WebSocket sirve **exclusivamente para escuchar** los mensajes que te envían en el momento exacto en que ocurren. No necesitas hacer *polling* o consultar la base de datos repetidamente.

## 3.1 🔌 Conexión WebSocket
* **Ruta de conexión:** `ws://localhost:8080/ws/messages` (Asegúrate de cambiar a `wss://` si el servidor usa HTTPS).
* **Autenticación (Obligatorio):**
  Debes indicar tu `userId` para que el servidor sepa qué mensajes debes recibir. Se puede de dos formas:
  * **Mediante Query String (Más común en frontend JS):** `ws://localhost:8080/ws/messages?userId=2`
  * **Mediante Header:** Enviar el header `X-User-ID: 2` al establecer conexión.

## 3.2 📨 Estructura del Evento (Diferencia con HTTP)
Cada vez que el *Usuario A* haga un `POST /messages` hacia ti, tú recibirás automáticamente un evento desde esta conexión WebSocket en un JSON **plano** sin envoltura (sin la palabra "data" inicial) y usando `snake_case`:

```json
{
  "id": 46,
  "sender_id": 1,
  "receiver_id": 2,
  "content": "¡Hola, cómo estás?",
  "created_at": "2026-02-26T15:30:00Z"
}
```
*(Compara este formato con la sección 2.1 y 2.2 para notar que HTTP devuelve camelCase y dentro de arreglos u objetos envolventes, mientras el Websocket escupe el objeto puro en snake_case).*
