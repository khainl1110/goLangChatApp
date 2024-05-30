import { useState } from "react"

export default function SendMessage() {

    let [message, setMessage] = useState('dfdf')

    let sendMessage = () => {
        fetch('http://127.0.0.1:8080/sendMsg', {
            method: 'POST',
            body: JSON.stringify({"message": message})
        }).then(res => res.json())
        .then(data => console.log(data))
    }
    return (
        <div>
            <h2>Send message</h2>
            <textarea value = {message} onChange={(e) => setMessage(e.target.value)}/>
            <div></div>
            <button onClick={sendMessage}>Send message</button>

        </div>
    )
}