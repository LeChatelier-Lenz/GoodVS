import { Outlet } from "react-router-dom";
import Header from "../components/header.tsx";
import Footer from "../components/footer.tsx";
// import {Container} from "@mui/material";
// import Enter from "../components/enter.tsx";
// import MyIcon from "../components/Ilustration.tsx";
// import Box from "@mui/material/Box";

export default function Root() {

    return (
        <div id="root-element" style={
            {
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                // justifyContent: "center",
                height: "100%",
                width: "100%",
            }
        }>
            <Header />
            <Outlet />
            <Footer />
        </div>
    );
}