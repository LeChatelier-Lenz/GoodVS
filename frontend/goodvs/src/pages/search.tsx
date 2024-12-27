import CustomizedInputBase from "../components/searchbar.tsx";
import Box from "@mui/material/Box";
import ProductCard from "../components/productItem.tsx";
import {Container, Grid2} from "@mui/material";

export default function Search() {

    const products = [
        {
            imageUrl: 'https://via.placeholder.com/150',
            productName: '商品1',
            platform: '平台A',
            productLink: 'https://example.com/product1'
        },
        {
            imageUrl: 'https://via.placeholder.com/150',
            productName: '商品2',
            platform: '平台B',
            productLink: 'https://example.com/product2'
        },
        {
            imageUrl: 'https://via.placeholder.com/150',
            productName: '商品3',
            platform: '平台C',
            productLink: 'https://example.com/product3'
        },
        {
            imageUrl: 'https://via.placeholder.com/150',
            productName: '商品4',
            platform: '平台D',
            productLink: 'https://example.com/product4'
        },
        {
            imageUrl: 'https://via.placeholder.com/150',
            productName: '商品5',
            platform: '平台E',
            productLink: 'https://example.com/product5'
        }
        // 更多商品数据...
    ];

    return (
        <Container id="search"
             sx = {{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            // justifyContent: "center",
            // height: "100%",
            minHeight: "800px",
             maxHeight: "1400px",
            width: "90%",
            gap: { xs: 4, sm: 8 },
            py: { xs: 8, sm: 20 },
             my : 4,
             }}
        >

            <CustomizedInputBase />
            <Box id="product-list" sx = {{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                width : "100%",
                // border:"solid"
            }}>
                <Grid2 container spacing={1} justifyContent="flex-start">
                    {products.map((product, index) => (
                        <Grid2  key={index} component={Box} size={4}
                        >
                            <ProductCard {...product} />
                        </Grid2>
                    ))}
                </Grid2>
            </Box>
        </Container>
    );
}