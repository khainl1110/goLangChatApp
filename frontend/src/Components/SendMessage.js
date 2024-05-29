import { useState } from "react"

export default function SendMessage() {

    let [message, setMessage] = useState('dfdf')

    let sendMessage = () => {
        fetch('http://127.0.0.1:8080/sendMsg', {
            method: 'POST',
            // headers: {'Content-Type': 'application/json',
            // },
            body: JSON.stringify({"message": "send from front end"})
        }).then(res => res.json())
        .then(data => console.log(data))
    }
    return (
        <div>
            <h1>Send message</h1>
            <textarea value = {message} onChange={(e) => setMessage(e.target.value)}/>
            <button onClick={sendMessage}>Send</button>

        </div>
    )
}