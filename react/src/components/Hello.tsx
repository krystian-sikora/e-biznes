import axios from "axios";

export function Hello() {

    const hello = () => {
        axios.get("http://localhost:8080/hello").then(r => console.log(r.data)).catch(e => console.error("Error fetching hello:", e));
    };

    return (
        <div>
            <button onClick={hello}>Hello</button>
        </div>
    );
}