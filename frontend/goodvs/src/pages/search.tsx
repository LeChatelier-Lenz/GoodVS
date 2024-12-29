import CustomizedInputBase from "../components/searchbar.tsx";
import Box from "@mui/material/Box";
import ProductCard from "../components/productItem.tsx";
import {Card, CardMedia, Container, Grid2, Typography} from "@mui/material";
import  {useState} from "react";
import "../style/loading.css"


export default function Search() {
    interface product{
        "id":string,
        "name":string,
        "url":string,
        "img_url":string,
        "price":number,
        "title":string,
        "category":string,
        "platform":string
    }
    const [productList, setProductList] = useState<Map<string,product[]>>(new Map<string,product[]>());

    const logoList = [
        "https://img10.360buyimg.com/img/jfs/t1/265960/38/828/10287/676565f6Fcdb37884/072d830437959819.png",
        "https://image5.suning.cn/uimg/cms/img/171818139083441716.png",
        "https://ts4.cn.mm.bing.net/th?id=ODLS.5f7a98c3-fd79-4885-9046-5068ac784fd7&w=32&h=32&qlt=90&pcl=fffffa&o=6&pid=1.2"
    ]

    const platformSheet = [
        "京东",
        "苏宁",
        "淘宝"
    ]

    const changeProductList = (data:any[]) => {
        if (data.length === 0){
            alert("No result found");
            return;
        }
        let itemDict = new Map<string,product[]>();
        // itemDict.clear();
        itemDict.set("京东",[]);
        itemDict.set("苏宁",[]);
        for (let i = 0; i < data.length; i++) {
            if (itemDict.has(data[i].platform)){
                itemDict.get(data[i].platform).push({
                    "id":data[i].id,
                    "name":data[i].name,
                    "url":data[i].url,
                    "img_url":data[i].img_url,
                    "price":data[i].price,
                    "title":data[i].title,
                    "category":data[i].category,
                    "platform":data[i].platform
                });
            } else {
                itemDict.set(data[i].platform,[{
                    "id":data[i].id,
                    "name":data[i].name,
                    "url":data[i].url,
                    "img_url":data[i].img_url,
                    "price":data[i].price,
                    "title":data[i].title,
                    "category":data[i].category,
                    "platform":data[i].platform
                }]);
            }
        }
        setProductList(itemDict);
        console.log(productList);
    }


    return (
        <Container id="search"
             sx = {{
            display: "flex",
            flexDirection: "column",

            // justifyContent: "center",
            // height: "100%",
            // minHeight: "600px",
             maxHeight: "1200px",

            width: "80%",
            gap: { xs: 4, sm: 8 },
            py: { xs: 8, sm: 20 },
             my : 4,
             borderBottom: " 1px solid #e0e0e0",
             }}
        >
            {
                (productList.size === 0) ?
                    <Box
                        sx={{

                            display: "flex",
                            flexDirection: "row",
                            alignItems: "center",
                            justifyContent: "center",
                            width: "100%",
                            height: "100%",
                            minHeight: "100px",
                            // border:"solid"
                        }}
                    >
                    <Typography
                        variant="h3"
                        sx={{
                            display: "flex",
                            flexDirection: "row",
                            alignItems: "center",
                            justifyContent: "center",
                            width: "800px",
                            height: "50px",
                            minHeight: "100px",
                            // border:"solid"
                        }}
                    >开始搜索吧！
                    </Typography>
                    <div className="loader">
                        <div className="justify-content-center jimu-primary-loading"></div>
                    </div>
                    </Box>
                : <></>
            }
            <CustomizedInputBase handleResult={changeProductList}/>
            {
                (productList.size === 0) ? <></> :
            <Box id="product-list" sx = {{
                display: "flex",
                flexDirection: "row",
                // justifyContent: "

                // alignItems: "center",
                width : "100%",

                // minWidth: "1200px",
                maxHeight: "90%",
                border: "1px solid #e0e0e0",
                borderRadius: "5px",
            }}>
                {platformSheet.map((item_p, index_p) => (
                    <Box
                        sx={{
                            display: "flex",
                            flexDirection: "column",
                            // alignItems: "flex-start",
                            // justifyContent: "start",
                            width: "100%",
                            // border:"solid"
                        }}
                    >
                        <Box
                            component={Card}
                            sx={{
                                height:"60px",
                                display: "flex",
                                flexDirection: "row",
                                alignItems: "center",
                                justifyContent: "center",
                                width: "100%",
                                minHeight: "60px",
                                py:1,
                                // border:"solid"
                            }}
                        >
                            <CardMedia
                                component="img"
                                height="50px"
                                image={logoList[index_p]}
                                alt={item_p}
                                sx={{ objectFit: 'contain'}}
                            />
                            <Typography
                                variant="h6"
                                sx = {{
                                    display: "flex",
                                    flexDirection: "row",
                                    alignItems: "center",
                                    justifyContent: "center",
                                    width: "100%",
                                    height: "50px",
                                    minHeight: "100px",
                                    // border:"solid"
                                }}
                            >
                                {item_p}
                            </Typography>
                        </Box>
                    <Grid2 key={index_p} component={Box} spacing={2}
                           container={true} size={12} gap={1}
                        sx = {{
                            marginTop: "10px",
                            display: "flex",
                            flexDirection: "row",
                            // alignItems: "center",
                            justifyContent: "center",
                            width: "100%",
                            overflowY:"auto",
                        }}
                    >
                        {productList.get(item_p) === undefined ? <></> :
                            productList.get(item_p).map((item, index) => (
                            <Grid2  key={index} component={Box} size={10}
                            >
                                <ProductCard
                                    id={item.id}
                                    imageUrl={item.img_url}
                                    productName={item.name}
                                    platform={item.platform}
                                    url={item.url}
                                    title={item.title}
                                    price={item.price}
                                    category={item.category}
                                    favorited={false}
                                 />
                            </Grid2>
                        ))}
                    </Grid2>
                </Box>
                ))}
            </Box>
            }
        </Container>
    );
}