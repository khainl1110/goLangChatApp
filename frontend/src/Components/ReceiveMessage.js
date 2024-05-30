import { useState } from "react"

export default function ReceiveMessage() {
    let [message, setMessage] = useState('')

    let pullMsg = () => {
        fetch('http://127.0.0.1:8080/getMsg')
        .then(res => res.json())
        .then(data => setMessage(data.message))
    }
    return (
        <div>
            <h2>Receive message</h2>
            <p>Wait a few seconds because pub/sub is slow</p>
            <h3>{message}</h3>
            <button onClick={pullMsg}>Pull message</button>
        </div>
    )
}