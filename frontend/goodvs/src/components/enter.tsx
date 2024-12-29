import "../style/enter.css";

export default function Enter() {

    const handleClick = () => {
        console.log("jump into '/search'");
        window.location.href = "/search";
    }

    return (
        <button className ="button"
            onClick={handleClick}
        >
            Try GoodVS | 开始使用
        </button>
    );
}