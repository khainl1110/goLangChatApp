import ReceiveMessage from "./ReceiveMessage";
import SendMessage from "./SendMessage";

export default function App() {
    return (
        <div>
            <h1>Hello to new app</h1>
            <SendMessage />
            <ReceiveMessage />
        </div>
    )
}