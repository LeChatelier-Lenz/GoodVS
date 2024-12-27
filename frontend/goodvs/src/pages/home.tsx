import {Container} from "@mui/material";
import Box from "@mui/material/Box";
import Enter from "../components/enter.tsx";
import MyIcon from "../components/Ilustration.tsx";
import "../style/home.css";

export default function Home() {
    return (
        <Container
            className="container"
            sx = {{
                display: "flex",
                flexDirection: "row",
                // alignItems: "center",
                alignItems:"flex-start",
                verticalAlign: "middle",
                justifyContent: "center",
                // height: "80%",
                minHeight: "800px",
                maxHeight: "1000px",
                gap: { xs: 4, sm: 8 },
                py: { xs: 8, sm: 20 },
                // border:"solid"
            }}>
            <Box sx={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                verticalAlign: "middle",
                justifyContent: "center",
                gap: { xs: 4, sm: 8 },
                py: { xs: 8, sm: 10 },
                textAlign: { sm: 'center', md: 'left' },
                // height: "100%",
                margin: "auto",
                width: "40%",
                // border:"solid",
            }}>
                <Enter />
            </Box>
            <Box
                sx={{
                    display: "flex",
                    flexDirection: "column",
                    alignItems: "center",
                    // gap: { xs: 4, sm: 8 },
                    py: { sm: 10 },
                    textAlign: { sm: 'center', md: 'left' },
                    height: "100%",
                    width: "60%",
                }}
            >
                <MyIcon />
            </Box>
        </Container>
    );
}