import { Outlet } from "react-router-dom";
import Header from "../components/header.tsx";
import Footer from "../components/footer.tsx";

export default function Root() {
    return (
        <div id="root" style={
            {
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                justifyContent: "center",
                height: "100vh",
                width: "100%"
            }
        }>
            <Header />
            <h1>GoodVS</h1>
            <Outlet />
            <Footer />
        </div>
    );
}